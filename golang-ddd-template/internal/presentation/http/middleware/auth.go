package middleware

import (
	//"fmt"
	"github.com/andriykutsevol/WeatherServer/internal/domain/auth"
	"github.com/andriykutsevol/WeatherServer/internal/domain/contextx"
	"github.com/andriykutsevol/WeatherServer/internal/domain/errors"
	"github.com/andriykutsevol/WeatherServer/internal/presentation/http"
	"github.com/gin-gonic/gin"
)


func wrapUserAuthContext(c *gin.Context, userID string) {
	http.SetUserID(c, userID)
	ctx := contextx.NewUserID(c.Request.Context(), userID)
	// TODO. Setpu logging
	//ctx = logger.NewUserIDContext(ctx, userID)
	c.Request = c.Request.WithContext(ctx)
}



func UserAuthMiddleware(a auth.Repository, skippers ...SkipperFunc) gin.HandlerFunc {

	// if !configs.C.JWTAuth.Enable {
	// 	return func(c *gin.Context) {
	// 		wrapUserAuthContext(c, configs.C.Root.UserName)
	// 		c.Next()
	// 	}
	// }

	// TODO. Look at the original files. Handle Authentication.

	return func(c *gin.Context) {

		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}		
		//wrapUserAuthContext(c, configs.C.Root.UserName)

		userID, err := a.ParseUserID(c.Request.Context(), http.GetToken(c))
		if err != nil {
			if err == auth.ErrInvalidToken {
				// if configs.C.IsDebugMode() {
				// 	wrapUserAuthContext(c, configs.C.Root.UserName)
				// 	c.Next()
				// 	return
				// }
				http.ResError(c, errors.ErrInvalidToken)
				return
			}
			http.ResError(c, errors.WithStack(err))
			return
		}

		//wrapUserAuthContext(c, "root")
		wrapUserAuthContext(c, userID)
		c.Next()
	}

}















