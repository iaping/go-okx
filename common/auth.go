package common

type Auth struct {
	ApiKey     string
	SecretKey  string
	Passphrase string
	Simulated  bool
}

func NewAuth(apiKey, secretKey, passphrase string, simulated bool) Auth {
	return Auth{
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		Passphrase: passphrase,
		Simulated:  simulated,
	}
}

func (a Auth) Signature(method, path, body string, isUnix bool) *Signature {
	return &Signature{
		Key:    a.SecretKey,
		Method: method,
		Path:   path,
		Body:   body,
		IsUnix: isUnix,
	}
}
