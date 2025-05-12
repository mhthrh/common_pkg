package cryptox

type ICrypto interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
	Md5Sum(string) (string, error)
	Sha256(string) string
}
