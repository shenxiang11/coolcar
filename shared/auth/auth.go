package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/shenxiang11/coolcar/shared/auth/token"
	"github.com/shenxiang11/coolcar/shared/id"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"strings"
)

const ImpersonateAccountHeader = "impersonate-account-id"
const authorizationHeader = "authorization"
const bearerPrefix = "Bearer "

func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {

	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open public key file: %v", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key: %v", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key: %v", err)
	}

	i := &interceptor{verifier: &token.JWTTokenVerifier{PublicKey: pubKey}}

	return i.HandleReq, nil
}

type tokenVerifer interface {
	Verify(token string) (string, error)
}

type interceptor struct {
	verifier tokenVerifer
}

type accountIDKey struct{}

func (i *interceptor) HandleReq(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	aid, err := i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)
	}

	return handler(context.WithValue(ctx, accountIDKey{}, aid), req)
}

func impersonationFromContext(c context.Context) string {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return ""
	}

	imp := m[ImpersonateAccountHeader]
	if len(imp) == 0 {
		return ""
	}

	return imp[0]
}

func tokenFromContext(c context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", unauthenticated
	}

	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn == "" {
		return "", unauthenticated
	}

	return tkn, nil
}

func AccountIDFromContext(c context.Context) (id.AccountID, error) {
	v := c.Value(accountIDKey{})
	aid, ok := v.(string) // FIXME: 直接断言到 AccountID 会报错
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return id.AccountID(aid), nil
}

func ContextWithAccount(c context.Context, aid id.AccountID) context.Context {
	return context.WithValue(c, accountIDKey{}, aid)
}
