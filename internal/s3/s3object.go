package s3

import (
	"context"
	"io"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Session(bucket, region, roleARN string) (*s3.S3, error) {
	if region == "" {
		if reg, err := s3manager.GetBucketRegion(
			context.Background(),
			session.Must(session.NewSession()),
			bucket,
			"us-east-1",
		); err != nil {
			return nil, err
		} else {
			region = reg
		}
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	awsCfg := &aws.Config{}
	if roleARN != "" {
		parsedArn, err := arn.Parse(roleARN)
		if err != nil {
			return nil, err
		}
		creds := stscreds.NewCredentials(sess, parsedArn.String())
		awsCfg.Credentials = creds
	}

	return s3.New(sess, awsCfg), nil
}

// Glob will try to resolve a (star and double-star) glob expression given in pattern on the bucket and return results.
// If the pattern doesn't contain any globs it will be returned as-is in an array, without checking for object existence
func Glob(svc s3iface.S3API, bucket, pattern string) ([]string, error) {
	const (
		star       = "*"
		doubleStar = "**"
	)

	if !strings.Contains(pattern, star) {
		return []string{pattern}, nil
	}

	var prefix, rest string

	// Decide if first glob char we encounter is * or **, and parse it into prefix & rest
	{
		starLoc := strings.Index(pattern, star)
		doubleStarLoc := strings.Index(pattern, doubleStar)
		if doubleStarLoc == -1 || starLoc < doubleStarLoc {
			prefix = pattern[:starLoc]
			rest = pattern[starLoc+1:]
		} else {
			prefix = pattern[:doubleStarLoc]
			rest = pattern[doubleStarLoc+2:]
		}
	}

	var r *regexp.Regexp // if pattern ends with * no regexp will be built, just include all results below
	if rest != "" {
		filterRegex := regexp.QuoteMeta(pattern)
		filterRegex = strings.ReplaceAll(filterRegex, regexp.QuoteMeta(doubleStar), ".*?")
		filterRegex = strings.ReplaceAll(filterRegex, regexp.QuoteMeta(star), "[^/]+?")
		var err error
		r, err = regexp.Compile("^" + filterRegex + "$")
		if err != nil {
			return nil, err
		}
	}

	objs := make(map[string]struct{})

	if err := svc.ListObjectsV2Pages(&s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Delimiter: nil,
		Prefix:    aws.String(prefix),
	}, func(o *s3.ListObjectsV2Output, lastPage bool) bool {
		for _, obj := range o.Contents {
			k := *obj.Key
			if r == nil || r.MatchString(k) {
				objs[k] = struct{}{}
			}
		}

		return true
	}); err != nil {
		return nil, err
	}

	ret := make([]string, len(objs))
	i := 0
	for k := range objs {
		ret[i] = k
		i++
	}

	return ret, nil
}

func GetObject(svc *s3.S3, bucket, key string) (io.ReadCloser, error) {
	out, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return out.Body, nil
}
