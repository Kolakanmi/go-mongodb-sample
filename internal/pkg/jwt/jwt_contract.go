package jwt

type (
	Signer interface {
		Sign(claims Claims) (string, error)
	}
	Verifier interface {
		Verify(token string) (*Claims, error)
	}
	SignerVerifier interface {
		Signer
		Verifier	}
)
