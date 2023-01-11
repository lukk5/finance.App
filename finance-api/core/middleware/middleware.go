package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

const (
	XRequestIdKey = "X-Request-ID" // request id header key
)

func RequestMiddlewareID() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := context.Request.Header.Get(XRequestIdKey)

		if requestId == "" {
			requestId = uuid.New().String()
		}

		// context.Request = context.Request.WithContext(trace.Wit)
		//
		//context.Request = context.Request.WithContext()

		context.Writer.Header().Set(XRequestIdKey, requestId)
	}
}

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.AbortWithStatus(http.StatusGatewayTimeout)
			}
			cancel()
		}()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
