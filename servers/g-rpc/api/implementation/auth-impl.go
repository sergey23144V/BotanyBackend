package implementation

import (
	context "context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	resource "github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"gorm.io/gorm"

	"github.com/sergey23144V/BotanyBackend/pkg"
	auth_helper "github.com/sergey23144V/BotanyBackend/pkg/auth-helper"

	"time"
	"unicode"
)

type AuthServerImpl struct {
	db *gorm.DB
}

func (a AuthServerImpl) SignUpSuperUser(ctx context.Context, input *api.SignUpUserInput) (*api.SignInUserResponse, error) {
	_, err := a.CreatSuperUser(ctx, input)
	if err != nil {
		return nil, err
	}

	token, err := a.GenerateToken(ctx, &api.SignInUserInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	return &api.SignInUserResponse{
		Status:       "200",
		AccessToken:  token,
		RefreshToken: "",
	}, nil
}

func NewAuthServer(db *gorm.DB) AuthServerImpl {
	return AuthServerImpl{db: db}
}

func (a AuthServerImpl) SignUpUser(ctx context.Context, input *api.SignUpUserInput) (*api.SignInUserResponse, error) {
	_, err := a.CreatUser(ctx, input)
	if err != nil {
		return nil, err
	}

	token, err := a.GenerateToken(ctx, &api.SignInUserInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	return &api.SignInUserResponse{
		Status:       "200",
		AccessToken:  token,
		RefreshToken: "",
	}, nil
}

func (a AuthServerImpl) SignInUser(ctx context.Context, input *api.SignInUserInput) (*api.SignInUserResponse, error) {
	token, err := a.GenerateToken(ctx, input)
	if err != nil {
		return nil, err
	}

	return &api.SignInUserResponse{
		Status:       "200",
		AccessToken:  token,
		RefreshToken: "",
	}, nil
}

func (a AuthServerImpl) MustEmbedUnimplementedAuthServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *AuthServerImpl) GenerateToken(ctx context.Context, input *api.SignInUserInput) (string, error) {
	userResult, err := api.ReadUserByEmailAndPassword(ctx, &api.User{Email: input.Email, Password: generatePasswordHash(input.Password)}, s.db)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth_helper.TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(auth_helper.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userResult.Id.ResourceId,
		userResult.Role,
	})

	return token.SignedString([]byte(auth_helper.SigningKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(auth_helper.Salt)))
}

func (s *AuthServerImpl) CreatUser(ctx context.Context, input *api.SignUpUserInput) (*api.User, error) {
	dubl, err := api.CheckingForDuplicateEmails(ctx, &api.User{Email: input.Email}, s.db)
	if err != nil {
		return nil, err
	} else if !dubl {
		return nil, errors.New("duplicate email")
	}

	err = ValidatePassword(input.Password)
	if err != nil {
		return nil, err
	}

	userInput := &api.User{
		Id:        &resource.Identifier{ResourceId: pkg.GenerateUUID()},
		Name:      input.Name,
		Email:     input.Email,
		Role:      api.RoleType_NormalUser,
		Password:  generatePasswordHash(input.Password),
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	return api.DefaultCreateUser(ctx, userInput, s.db)
}

func (s *AuthServerImpl) CreatSuperUser(ctx context.Context, input *api.SignUpUserInput) (*api.User, error) {
	dubl, err := api.CheckingForDuplicateEmails(ctx, &api.User{Email: input.Email}, s.db)
	if err != nil {
		return nil, err
	} else if !dubl {
		return nil, errors.New("duplicate email")
	}

	err = ValidatePassword(input.Password)
	if err != nil {
		return nil, err
	}

	userInput := &api.User{
		Id:        &resource.Identifier{ResourceId: pkg.GenerateUUID()},
		Name:      input.Name,
		Email:     input.Email,
		Role:      api.RoleType_SuperUser,
		Password:  generatePasswordHash(input.Password),
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	return api.DefaultCreateUser(ctx, userInput, s.db)
}

func ValidatePassword(password string) error {
	// Правила валидации
	minLength := 8
	hasUpperCase := false
	hasLowerCase := false
	hasDigit := false

	// Проверка минимальной длины
	if len(password) < minLength {
		return fmt.Errorf("Пароль должен содержать не менее %d символов", minLength)
	}

	// Проверка наличия букв верхнего, нижнего регистра, цифр и специальных символов
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpperCase = true
		case unicode.IsLower(char):
			hasLowerCase = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	// Проверка выполнения всех правил
	if !hasUpperCase {
		return fmt.Errorf("Пароль должен содержать хотя бы одну букву верхнего регистра")
	}
	if !hasLowerCase {
		return fmt.Errorf("Пароль должен содержать хотя бы одну букву нижнего регистра")
	}
	if !hasDigit {
		return fmt.Errorf("Пароль должен содержать хотя бы одну цифру")
	}

	// Пароль прошел валидацию
	return nil
}
