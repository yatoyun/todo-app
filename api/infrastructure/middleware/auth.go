package middleware

import (
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/yatoyun/todo-app/api/config"
)

func JWTMiddleware(cfg *config.Config) gin.HandlerFunc {
	issuerURL, err := url.Parse("https://" + cfg.Auth0Domain + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{cfg.Auth0Audience},
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		log.Fatalf("JWT Validatorの設定に失敗しました: %v", err)
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		slog.Error("request header: %v", r.Header)
		log.Printf("JWT検証エラー: %v", err)
	}

	jwtMiddleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(c *gin.Context) {
		encounteredError := true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			c.Request = r
			c.Next()
		}

		jwtMiddleware.CheckJWT(handler).ServeHTTP(c.Writer, c.Request)

		if encounteredError {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"message": "JWTが無効です"},
			)
			return
		}
	}
}
