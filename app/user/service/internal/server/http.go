package server

import (
	"context"
	userv1 "kratos-base/api/user/service/v1"
	"kratos-base/app/user/service/internal/conf"
	"kratos-base/app/user/service/internal/service"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/go-kratos/kratos/v2/log"
	// "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwt4 "github.com/golang-jwt/jwt/v4"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			// jwt.Server(jwtKeyFunc(c), jwt.WithSigningMethod(jwt4.SigningMethodRS256)),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	userv1.RegisterUserHTTPServer(srv, user)

	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	return srv
}

func jwtKeyFunc(c *conf.Server) func(token *jwt4.Token) (interface{}, error) {
	options := keyfunc.Options{
		Ctx: context.Background(),
		RefreshErrorHandler: func(err error) {
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(c.Auth.Jwk, options)
	if err != nil {
		log.Fatalf("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
	}
	return func(token *jwt4.Token) (interface{}, error) {
		checkAud := token.Claims.(jwt4.MapClaims).VerifyAudience(c.Auth.Audience, false)
		if !checkAud {
			vErr := new(jwt4.ValidationError)
			vErr.Errors |= jwt4.ValidationErrorAudience
			return token, vErr
		}

		checkIss := token.Claims.(jwt4.MapClaims).VerifyIssuer(c.Auth.Issuer, false)
		if !checkIss {
			vErr := new(jwt4.ValidationError)
			vErr.Errors |= jwt4.ValidationErrorIssuer
			return token, vErr
		}
		return jwks.Keyfunc(token)
	}
}
