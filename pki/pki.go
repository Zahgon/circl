package pki

import (
	"crypto/x509/pkix"
	"encoding/asn1"

	"github.com/cloudflare/circl/sign"
	"github.com/cloudflare/circl/sign/schemes"
)

var (
	allSchemesByOID map[string]sign.Scheme
	allSchemesByTLS map[uint]sign.Scheme
)

type pkixPrivKey struct {
	Version    int
	Algorithm  pkix.AlgorithmIdentifier
	PrivateKey []byte
}

func init() {
	allSchemesByOID = make(map[string]sign.Scheme)
	allSchemesByTLS = make(map[uint]sign.Scheme)
	for _, scheme := range schemes.All() {
		if cert, ok := scheme.(CertificateScheme); ok {
			allSchemesByOID[cert.Oid().String()] = scheme
		}
		if tlsScheme, ok := scheme.(TLSScheme); ok {
			allSchemesByTLS[tlsScheme.TLSIdentifier()] = scheme
		}
	}
}

func SchemeByOid(oid asn1.ObjectIdentifier) sign.Scheme {
	_ = "STUB: not implemented"
	return *new(sign.Scheme)
}

func SchemeByTLSID(id uint) sign.Scheme {
	_ = "STUB: not implemented"
	return *

	// Additional methods when the signature scheme is supported in X509.
	new(sign.Scheme)
}

type CertificateScheme interface {
	// Return the appropriate OIDs for this instance.  It is implicitly
	// assumed that the encoding is simple: e.g. uses the same OID for
	// signature and public key like Ed25519.
	Oid() asn1.ObjectIdentifier
}

// Additional methods when the signature scheme is supported in TLS.
type TLSScheme interface {
	TLSIdentifier() uint
}

func UnmarshalPEMPublicKey(data []byte) (sign.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PublicKey), nil
}

func UnmarshalPKIXPublicKey(data []byte) (sign.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PublicKey), nil
}

func UnmarshalPEMPrivateKey(data []byte) (sign.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PrivateKey), nil
}

func isMLDSA(scheme sign.Scheme) bool { _ = "STUB: not implemented"; return false }

func UnmarshalPKIXPrivateKey(data []byte) (sign.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(sign.PrivateKey), nil
}

// ML-DSA unfortunately has a complex private key format, which we
// handle here separately. If future schemes require custom parsing
// as well, we can introduce an interface for that.

// Handle case of seed-only private key

// We don't support expanded-only private keys, so the only remaining
// option is a SEQUENCE of both seed and expanded private key.

func MarshalPEMPublicKey(pk sign.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MarshalPKIXPublicKey(pk sign.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MarshalPEMPrivateKey(sk sign.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MarshalPKIXPrivateKey(sk sign.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ML-DSA is special. See comment in UnmarshalPKIXPrivateKey().
