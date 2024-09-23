package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/yatoyun/todo-app/api/config"
)

func JWTMiddleware(cfg *config.Config) gin.HandlerFunc {
	keyFunc := func(ctx context.Context) (interface{}, error) {
		// Auth0の公開鍵を取得するための設定が必要です。
		// ここでは簡略化のため、署名キーを直接返していますが、
		// 実際にはAuth0のJWKエンドポイントからキーを取得する必要があります。
		// 例えば、https://YOUR_AUTH0_DOMAIN/.well-known/jwks.json
		return []byte("YOUR_AUTH0_SECRET"), nil
	}

	jwtValidator, err := validator.New(
		keyFunc,
		validator.RS256,
		cfg.Auth0Domain,
		[]string{cfg.Auth0Audience},
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		log.Fatalf("JWT Validatorの設定に失敗しました: %v", err)
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
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
