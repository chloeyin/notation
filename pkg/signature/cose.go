package signature

import (
	"crypto/x509"
	"fmt"
	"os"

	"github.com/notaryproject/notation-go"
	"github.com/notaryproject/notation-go/signature/cose"
)

// NewSignerFromFiles creates a verifier from certificate files
// copy from jws.go
func NewCoseVerifierFromFiles(certPaths []string) (notation.Verifier, error) {
	verifier := cose.NewVerifier()
	verifier.VerifyOptions.Roots = x509.NewCertPool()
	for _, path := range certPaths {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		if !verifier.VerifyOptions.Roots.AppendCertsFromPEM(data) {
			return nil, fmt.Errorf("failed to parse PEM certificate: %q", path)
		}
	}
	return verifier, nil
}
