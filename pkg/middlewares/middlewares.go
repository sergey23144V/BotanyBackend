package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/infobloxopen/atlas-app-toolkit/v2/rpc/resource"
	auth_helper "github.com/sergey23144V/BotanyBackend/pkg/auth-helper"
	"github.com/sergey23144V/BotanyBackend/servers/g-rpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const (
	KeyToken    = "token"
	KeyUserId   = "userId"
	KeyUserRole = "userRole"
)

func AuthInterceptorGraphQL() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path != "/api" {
				next.ServeHTTP(w, r)
				return
			}

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
				return
			}
			// Восстанавливаем тело запроса
			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			var reqBody map[string]interface{}
			if err := json.Unmarshal(body, &reqBody); err != nil {
				http.Error(w, "Ошибка разбора JSON", http.StatusInternalServerError)
				return
			}

			if findAuth(string(body)) || reqBody["operationName"].(string) == "IntrospectionQuery" {
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("authorization")

			if token == "" {
				http.Error(w, "Отсутствует токен авторизации", http.StatusUnauthorized)
				return
			}
			authorization := ParseAuthorization(token)

			userId, role, err := authorization.(auth_helper.TokenAuth).GetUserFromToken()
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			ctx := r.Context()

			ctx = context.WithValue(ctx, KeyToken, authorization.String())

			ctx = context.WithValue(ctx, KeyUserRole, role)

			ctx = context.WithValue(ctx, KeyUserId, userId)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return

		})
	}
}

func AuthInterceptorREST() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := r.Header.Get("authorization")

			if token == "" {
				next.ServeHTTP(w, r)
				return
			}
			authorization := ParseAuthorization(token)

			userId, role, err := authorization.(auth_helper.TokenAuth).GetUserFromToken()
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			ctx := r.Context()

			ctx = context.WithValue(ctx, KeyToken, authorization.String())

			ctx = context.WithValue(ctx, KeyUserRole, role)

			ctx = context.WithValue(ctx, KeyUserId, userId)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return

		})
	}
}

func findAuth(text string) bool {
	// Создаем регулярное выражение для поиска строки "auth"
	re := regexp.MustCompile(`auth \{`)

	// Ищем все совпадения в тексте
	matches := re.FindAllString(text, -1)

	// Если найдено хотя бы одно совпадение, возвращаем первое
	if len(matches) > 0 {
		return true
	}

	// Если совпадений не найдено, возвращаем пустую строку
	return false
}

func findPort(text string) bool {
	// Создаем регулярное выражение для поиска строки "auth"
	re := regexp.MustCompile(`8080`)

	// Ищем все совпадения в тексте
	matches := re.FindAllString(text, -1)

	// Если найдено хотя бы одно совпадение, возвращаем первое
	if len(matches) > 0 {
		return true
	}

	// Если совпадений не найдено, возвращаем пустую строку
	return false
}

// AuthInterceptor - промежуточный слой для проверки токена в gRPC
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	if info.FullMethod == "/auth.AuthService/SignUpUser" || info.FullMethod == "/auth.AuthService/SignInUser" ||
		info.FullMethod == "/auth.AuthService/SignUpSuperUser" || info.FullMethod == "/auth.AuthService/RefreshToken" {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "не удалось получить метаданные")
	}

	token := getToken(md)
	if token == "" {
		return nil, status.Errorf(codes.Unauthenticated, "требуется токен аутентификации")
	}

	authorization := ParseAuthorization(token)

	userId, role, err := authorization.(auth_helper.TokenAuth).GetUserFromToken()
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, KeyToken, authorization.String())

	ctx = context.WithValue(ctx, KeyUserRole, role)

	ctx = context.WithValue(ctx, KeyUserId, userId)

	return handler(ctx, req)
}

// getToken - вспомогательная функция для извлечения токена из метаданных
func getToken(md metadata.MD) string {
	if val, ok := md["authorization"]; ok && len(val) > 0 {
		return val[0]
	}
	return ""
}

func GetUserIdFromContext(ctx context.Context) *resource.Identifier {
	id, ok := ctx.Value(KeyUserId).(string)
	if !ok {
		return nil
	}
	return &resource.Identifier{
		ResourceId: id,
	}
}
func GetUserRoleFromContext(ctx context.Context) *api.RoleType {
	role, ok := ctx.Value(KeyUserRole).(*api.RoleType)
	if !ok {
		return nil
	}
	return role
}

func GetTokenFromContext(ctx context.Context) *string {
	raw, ok := ctx.Value(KeyToken).(string)
	if !ok {
		return nil
	}
	return &raw
}

func ParseAuthorization(authToken string) auth_helper.Authorization {
	// Разбиваем строку заголовка на части по пробелу
	parts := strings.SplitN(authToken, " ", 2)

	// Проверяем, что есть две части и первая часть не пустая
	if len(parts) != 2 || parts[0] == "" {
		fmt.Errorf("неверный формат заголовка аутентификации")
	}

	// Получаем тип аутентификации (Bearer или Basic)
	authType := parts[0]

	// Получаем значение аутентификации после пробела
	authValue := parts[1]

	var authorization auth_helper.Authorization

	var err error

	if authType == "Basic" {
		authorization, err = auth_helper.ParseBase(authValue)
		if err != nil {
			fmt.Errorf("Incorrect authorization")
		}
	} else if authType == "Bearer" {
		authorization = auth_helper.NewTokenAuth(authValue)
	} else {
		fmt.Errorf("неподдерживаемый тип аутентификации: %s", authType)
	}
	return authorization
}

func ValidToken(ctx context.Context) error {

	token := GetTokenFromContext(ctx)
	if token == nil {
		return errors.New("request does not have a token token")
	}
	result, err := jwt.ParseWithClaims(*token, &auth_helper.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(auth_helper.SigningKey), nil
	})
	if err != nil {
		return errors.New("error parsing the token :" + err.Error())
	}
	if result == nil {
		return errors.New("the token is invalid or expired")
	}
	return nil
}
