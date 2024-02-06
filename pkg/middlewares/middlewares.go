package middlewares

import (
	"context"
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/atlas/resource"
	auth_helper "github.com/sergey23144V/BotanyBackend/pkg/auth-helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

const (
	KeyToken  = "token"
	KeyUserId = "userId"
)

// AuthInterceptor - промежуточный слой для проверки токена в gRPC
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	if info.FullMethod == "/auth.AuthService/SignUpUser" || info.FullMethod == "/auth.AuthService/SignInUser" {
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

	userId, err := authorization.(auth_helper.TokenAuth).GetUserIdFromToken()
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, KeyToken, token)

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
	id := ctx.Value(KeyUserId).(string)

	return &resource.Identifier{
		ResourceId: id,
	}
}
func GetTokenFromContext(ctx context.Context) *string {
	raw := ctx.Value(KeyToken).(string)

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
