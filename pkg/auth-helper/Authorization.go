package auth_helper

import (
	"encoding/json"
	"strings"
)

// Структура Authorization (пустая)
type Authorization interface {
	String() string
	ToHttpHeaderValue() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(bytes []byte) error
}

// Ваши структуры для CredentialsAuth и TokenAuth
type CredentialsAuth struct {
	Login    string
	Password string
}

func (ca *CredentialsAuth) UnmarshalJSON(bytes []byte) error {
	var jsonData string
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		return err
	}
	credentials := strings.Split(jsonData, ":")

	// Устанавливаем логин и пароль в структуру
	ca.Login = credentials[0]
	ca.Password = credentials[1]

	return nil
}

func (ca *CredentialsAuth) MarshalJSON() ([]byte, error) {
	data := ca.Login + ":" + ca.Password

	return json.Marshal(data)
}

type TokenAuth struct {
	Token string
}

func (ta TokenAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(ta.Token)
}

func (ta TokenAuth) UnmarshalJSON(bytes []byte) error {
	var token string
	err := json.Unmarshal(bytes, &token)
	ta.Token = token
	return err
}
