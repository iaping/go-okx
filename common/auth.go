package common

type Auth struct {
	ObjectID   string
	ApiKey     string
	SecretKey  string
	Passphrase string
	Simulated  bool
}

func NewAuth(objectID, apiKey, secretKey, passphrase string, simulated bool) Auth {
	return Auth{
		ObjectID:   objectID,
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
