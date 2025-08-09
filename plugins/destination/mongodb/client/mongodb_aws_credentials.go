// Copyright (C) MongoDB, Inc. 2025-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package client

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsv4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
	"go.mongodb.org/mongo-driver/x/mongo/driver/auth"
)

const (
	sourceExternal      = "$external"
	responceNonceLength = 64
	amzDateFormat       = "20060102T150405Z"
	awsSessionToken     = "AWS_SESSION_TOKEN"
	defaultRegion       = "us-east-1"
	maxHostLength       = 255
)

const MongoDBCQAWS = "MONGODB-CQ-AWS"

// Authenticator is an authenticator that uses the AWS SDK rather than the
// lightweight AWS package used internally by the driver.
type Authenticator struct {
	userCred *auth.Cred    // MongoDB TLS credentials with AWS keys
	awsCfg   aws.Config    // AWS SDK config
	signer   *awsv4.Signer // SigV4 signer
}

var _ driver.Authenticator = (*Authenticator)(nil)

func mapToCredentials(credProps map[string]string) (spec.Credentials, error) {
	var awsCred spec.Credentials

	if val, ok := credProps["LocalProfile"]; ok {
		awsCred.LocalProfile = val
	}
	if val, ok := credProps["RoleARN"]; ok {
		awsCred.RoleARN = val
	}
	if val, ok := credProps["RoleSessionName"]; ok {
		awsCred.RoleSessionName = val
	}
	if val, ok := credProps["ExternalID"]; ok {
		awsCred.ExternalID = val
	}
	if val, ok := credProps["Default"]; ok {
		boolValue, err := strconv.ParseBool(val)
		if err != nil {
			return awsCred, fmt.Errorf("failed to parse Default value: %w", err)
		}
		awsCred.Default = boolValue
	}
	return awsCred, nil
}

func getAWSConfig(awsCred spec.Credentials) (*aws.Config, error) {
	ctx := context.Background()
	var cfg aws.Config
	var err error
	configFns := []func(*config.LoadOptions) error{}

	if awsCred.Default {
		// Use default AWS credentials
		cfg, err = config.LoadDefaultConfig(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		configFns = append(configFns, config.WithDefaultRegion("us-east-1"))
		if awsCred.LocalProfile != "" {
			configFns = append(configFns, config.WithSharedConfigProfile(awsCred.LocalProfile))
		}

		cfg, err = config.LoadDefaultConfig(ctx, configFns...)
		if err != nil {
			return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
		}

		if awsCred.RoleARN != "" {
			opts := make([]func(*stscreds.AssumeRoleOptions), 0, 1)
			if awsCred.ExternalID != "" {
				opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
					opts.ExternalID = &awsCred.ExternalID
				})
			}
			if awsCred.RoleSessionName != "" {
				opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
					opts.RoleSessionName = awsCred.RoleSessionName
				})
			}
			stsClient := sts.NewFromConfig(cfg)
			provider := stscreds.NewAssumeRoleProvider(stsClient, awsCred.RoleARN, opts...)

			cfg.Credentials = aws.NewCredentialsCache(provider)
		}
	}

	return &cfg, nil
}

// NewAuthenticator creates a new AWS SDK authenticator. It loads the AWS
// SDK config (honoring AWS_STS_REGIONAL_ENDPOINTS & AWS_REGION) and returns an
// Authenticator that uses it
func NewAuthenticator(cred *auth.Cred, _ *http.Client) (driver.Authenticator, error) {
	// details are in the cred.Props
	awsCred, err := mapToCredentials(cred.Props)
	if err != nil {
		return nil, fmt.Errorf("failed to map credentials: %w", err)
	}
	cfg, err := getAWSConfig(awsCred)
	if err != nil {
		return nil, fmt.Errorf("failed to get AWS config: %w", err)
	}

	return &Authenticator{
		userCred: cred,
		awsCfg:   *cfg,
		signer:   awsv4.NewSigner(),
	}, nil
}

var _ auth.AuthenticatorFactory = NewAuthenticator

// Auth starts the SASL conversation by constructing a custom SASL adapter that
// uses the AWS SDK for singing.
func (a *Authenticator) Auth(ctx context.Context, cfg *driver.AuthConfig) error {
	// Build a SASL adapter that uses AWS SDK for signing.
	adapter := &awsSdkSaslClient{
		userCred: a.userCred,
		awsCfg:   a.awsCfg,
		signer:   a.signer,
	}

	return auth.ConductSaslConversation(ctx, cfg, sourceExternal, adapter)
}

// REauth is not supported for AWS SDK authentication.
func (*Authenticator) Reauth(context.Context, *driver.AuthConfig) error {
	return errors.New("AWS reauthentication not supported")
}

type conversationState uint8

const (
	conversationStateStart       conversationState = 1 // before sending anything.
	conversationStateServerFirst conversationState = 2 // after sending client-first, awaiting server reply.
	conversationStateDone        conversationState = 3 // after sending client-final.
)

// awsSdkSaslClient is a SASL client that uses the AWS SDK for signing.
type awsSdkSaslClient struct {
	state    conversationState // handshake state machine
	nonce    []byte            // client nonce
	userCred *auth.Cred        // MongoDB TLS credentials with AWS keys
	awsCfg   aws.Config        // AWS SDK config
	signer   *awsv4.Signer     // SigV4 signer
}

var _ auth.SaslClient = (*awsSdkSaslClient)(nil)

// Start will create the client-first SASL message.
// { p: 110, r: <32-byte nonce>}; per the current Go Driver behavior.
func (client *awsSdkSaslClient) Start() (string, []byte, error) {
	client.state = conversationStateServerFirst
	client.nonce = make([]byte, 32)
	_, _ = rand.Read(client.nonce)

	idx, msg := bsoncore.AppendDocumentStart(nil)
	msg = bsoncore.AppendInt32Element(msg, "p", 110)
	msg = bsoncore.AppendBinaryElement(msg, "r", 0x00, client.nonce)
	msg, _ = bsoncore.AppendDocumentEnd(msg, idx)

	return auth.MongoDBAWS, msg, nil
}

func getRegion(host string) (string, error) {
	region := defaultRegion

	if len(host) == 0 {
		return "", errors.New("invalid STS host: empty")
	}
	if len(host) > maxHostLength {
		return "", errors.New("invalid STS host: too large")
	}
	// The implicit region for sts.amazonaws.com is us-east-1
	if host == "sts.amazonaws.com" {
		return region, nil
	}
	if strings.HasPrefix(host, ".") || strings.HasSuffix(host, ".") || strings.Contains(host, "..") {
		return "", errors.New("invalid STS host: empty part")
	}

	// If the host has multiple parts, the second part is the region
	parts := strings.Split(host, ".")
	if len(parts) >= 2 {
		region = parts[1]
	}

	return region, nil
}

// Next handles the server's "server-first" message, then builds and returns the
// "client-final" payload containing the SigV4-signed STS GetCallerIdentity
// request.
func (client *awsSdkSaslClient) Next(ctx context.Context, challenge []byte) ([]byte, error) {
	if client.state != conversationStateServerFirst {
		return nil, fmt.Errorf("invalid state: %v", client.state)
	}
	client.state = conversationStateDone

	// Unmarhal the server's BSON: { s: <server nonce>, h: "<sts host>"}
	var sm struct {
		Nonce primitive.Binary `bson:"s"`
		Host  string           `bson:"h"`
	}

	if err := bson.Unmarshal(challenge, &sm); err != nil {
		return nil, err
	}

	// Check nonce prefix
	if sm.Nonce.Subtype != 0x00 {
		return nil, errors.New("server reply contained unexpected binary subtype")
	}

	if len(sm.Nonce.Data) != responceNonceLength {
		return nil, fmt.Errorf("server reply nonce was not %v bytes", responceNonceLength)
	}

	if !bytes.HasPrefix(sm.Nonce.Data, client.nonce) {
		return nil, errors.New("server nonce did not extend client nonce")
	}

	currentTime := time.Now().UTC()
	// body := "Action=GetCallerIdentity&Version=2011-06-15"
	body := strings.NewReader("Action=GetCallerIdentity&Version=2011-06-15")

	// Create http.Request
	req, _ := http.NewRequestWithContext(ctx, "POST", "/", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", "43")

	req.Host = sm.Host
	req.Header.Set("X-Amz-Date", currentTime.Format(amzDateFormat))

	// Include session token if present.
	if tok := client.userCred.Props[awsSessionToken]; tok != "" {
		req.Header.Set("X-Amz-Security-Token", tok)
	}

	req.Header.Set("X-MongoDB-Server-Nonce", base64.StdEncoding.EncodeToString(sm.Nonce.Data))
	req.Header.Set("X-MongoDB-GS2-CB-Flag", "n")

	region, err := getRegion(sm.Host)
	if err != nil {
		return nil, fmt.Errorf("failed to get AWS region: %w", err)
	}

	// Retrieve AWS creds and sign the request using AWS SDK v4.
	creds, err := client.awsCfg.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve AWS credentials: %w", err)
	}

	h := sha256.New()
	_, _ = io.Copy(h, body)
	payloadHash := hex.EncodeToString(h.Sum(nil))

	// Create signer with credentials
	err = client.signer.SignHTTP(ctx, creds, req, payloadHash, "sts", region, currentTime)
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	// create message
	// { a: Authorization, d: X-Amz-Date, t: X-Amz-Security-Token }
	idx, msg := bsoncore.AppendDocumentStart(nil)
	msg = bsoncore.AppendStringElement(msg, "a", req.Header.Get("Authorization"))
	msg = bsoncore.AppendStringElement(msg, "d", req.Header.Get("X-Amz-Date"))
	if tok := req.Header.Get("X-Amz-Security-Token"); tok != "" {
		msg = bsoncore.AppendStringElement(msg, "t", tok)
	}
	msg, _ = bsoncore.AppendDocumentEnd(msg, idx)

	return msg, nil
}

// complete signals that the SASL conversation is done.
func (client *awsSdkSaslClient) Completed() bool {
	return client.state == conversationStateDone
}
