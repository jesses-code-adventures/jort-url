package passwords

import (
	"github.com/matthewhartstonge/argon2"
)

type PasswordHandler struct {
	argon argon2.Config
}

func NewPasswordHandler() *PasswordHandler {
	config := argon2.DefaultConfig()
	return &PasswordHandler{config}
}

func (p *PasswordHandler) Hash(password []byte) ([]byte, error) {
	encoded, err := p.argon.HashEncoded(password)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

func (p *PasswordHandler) Verify(password []byte, existing_encoded []byte) (bool, error) {
	ok, err := argon2.VerifyEncoded(password, existing_encoded)
	if err != nil {
		return false, err
	}
	if ok {
		return true, nil
	}
	return false, nil
}
