package client

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/kafka/v5/client/spec"
	"github.com/stretchr/testify/require"
)

type certificateFiles struct {
	certPath string
	keyPath  string
}

func writeSelfSignedCertificate(t *testing.T) certificateFiles {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
	}
	certDER, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	require.NoError(t, err)
	keyDER, err := x509.MarshalECPrivateKey(key)
	require.NoError(t, err)

	dir := t.TempDir()
	files := certificateFiles{certPath: filepath.Join(dir, "cert.pem"), keyPath: filepath.Join(dir, "key.pem")}
	require.NoError(t, os.WriteFile(files.certPath, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER}), 0o600))
	require.NoError(t, os.WriteFile(files.keyPath, pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER}), 0o600))
	return files
}

func startTLSServer(t *testing.T, server certificateFiles) string {
	t.Helper()
	cert, err := tls.LoadX509KeyPair(server.certPath, server.keyPath)
	require.NoError(t, err)
	listener, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{Certificates: []tls.Certificate{cert}})
	require.NoError(t, err)
	t.Cleanup(func() { listener.Close() })
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go func() {
				defer conn.Close()
				_ = conn.(*tls.Conn).Handshake()
			}()
		}
	}()
	return listener.Addr().String()
}

func handshake(t *testing.T, address string, userSpec *spec.Spec) error {
	t.Helper()
	tlsConfig, err := createTLSConfiguration(userSpec)
	require.NoError(t, err)
	conn, err := tls.Dial("tcp", address, tlsConfig)
	if err != nil {
		return err
	}
	return conn.Close()
}

func specWithCA(client certificateFiles, caPath string) *spec.Spec {
	return &spec.Spec{TlsDetails: &spec.TlsDetails{CertFile: &client.certPath, KeyFile: &client.keyPath, CaFile: &caPath}}
}

func TestTLSConfigurationRejectsUntrustedBrokerByDefault(t *testing.T) {
	broker := writeSelfSignedCertificate(t)
	other := writeSelfSignedCertificate(t)
	address := startTLSServer(t, broker)

	err := handshake(t, address, specWithCA(other, other.certPath))

	var unknownAuthority x509.UnknownAuthorityError
	require.ErrorAs(t, err, &unknownAuthority)
}

func TestTLSConfigurationAcceptsBrokerSignedByConfiguredCA(t *testing.T) {
	broker := writeSelfSignedCertificate(t)
	address := startTLSServer(t, broker)

	require.NoError(t, handshake(t, address, specWithCA(broker, broker.certPath)))
}

func TestTLSConfigurationAcceptsUntrustedBrokerWhenVerificationIsSkipped(t *testing.T) {
	broker := writeSelfSignedCertificate(t)
	other := writeSelfSignedCertificate(t)
	address := startTLSServer(t, broker)
	userSpec := specWithCA(other, other.certPath)
	userSpec.InsecureSkipTLSVerify = true

	require.NoError(t, handshake(t, address, userSpec))
}

func TestTLSConfigurationFailsOnInvalidCAFile(t *testing.T) {
	client := writeSelfSignedCertificate(t)
	caPath := filepath.Join(t.TempDir(), "ca.pem")
	require.NoError(t, os.WriteFile(caPath, []byte("not a certificate"), 0o600))

	_, err := createTLSConfiguration(specWithCA(client, caPath))

	require.ErrorContains(t, err, "failed to parse CA file")
}
