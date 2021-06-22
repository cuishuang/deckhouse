package certificate

import (
	"bytes"
	"log"
	"os"

	"github.com/cloudflare/cfssl/cli/genkey"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/helpers"
	"github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/signer/local"
	"github.com/sirupsen/logrus"
)

type Certificate struct {
	Key  string `json:"key"`
	Cert string `json:"cert"`
	CA   string `json:"ca"`
}

type SigningOption func(signing *config.Signing)

func GenerateSelfSignedCert(logger *logrus.Entry, cn string, hosts []string, ca Authority, options ...interface{}) (Certificate, error) {
	logger.Debugf("Generate self-signed cert for %s %v", cn, hosts)
	request := &csr.CertificateRequest{
		CN: cn,
		KeyRequest: &csr.KeyRequest{
			A: "ecdsa",
			S: 256,
		},
	}

	for _, option := range options {
		if f, ok := option.(Option); ok {
			f(request)
		}
	}

	// Catch cfssl logs message
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	g := &csr.Generator{Validator: genkey.Validator}
	csrBytes, key, err := g.ProcessRequest(request)
	if err != nil {
		return Certificate{}, err
	}

	req := signer.SignRequest{
		Hosts:   hosts,
		Request: string(csrBytes),
	}

	parsedCa, err := helpers.ParseCertificatePEM([]byte(ca.Cert))
	if err != nil {
		return Certificate{}, err
	}

	priv, err := helpers.ParsePrivateKeyPEM([]byte(ca.Key))
	if err != nil {
		return Certificate{}, err
	}

	signingConfig := &config.Signing{
		Default: config.DefaultConfig(),
	}

	for _, option := range options {
		if f, ok := option.(SigningOption); ok {
			f(signingConfig)
		}
	}

	s, err := local.NewSigner(priv, parsedCa, signer.DefaultSigAlgo(priv), signingConfig)
	if err != nil {
		return Certificate{}, err
	}

	cert, err := s.Sign(req)
	if err != nil {
		return Certificate{}, err
	}

	return Certificate{CA: ca.Cert, Key: string(key), Cert: string(cert)}, nil
}
