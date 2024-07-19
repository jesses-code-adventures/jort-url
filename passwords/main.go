package passwords

import (
	"fmt"
	"github.com/matthewhartstonge/argon2"
	"log/slog"
)

type PasswordEncoder struct {
	argon argon2.Config
}

func NewPasswordEncoder() *PasswordEncoder {
	config := argon2.DefaultConfig()
	return &PasswordEncoder{config}
}

func (p *PasswordEncoder) EncodePassword(password string) ([]byte, error) {
	encoded, err := p.argon.HashEncoded([]byte(password))
	if err != nil {
		return nil, err
	}
	fmt.Println(string(encoded))
	return encoded, nil
}

func (p *PasswordEncoder) VerifyPassword(password string, existing_encoded []byte) (bool, error) {
	ok, err := argon2.VerifyEncoded([]byte(password), existing_encoded)
	if err != nil {
		slog.Error("error verifying password")
		return false, err
	}
	if ok {
		slog.Info("correct password")
		return true, nil
	}
	slog.Warn("incorrect password")
	return false, nil
}
