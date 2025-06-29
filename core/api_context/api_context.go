package api_context

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vfyuliawan/my-golang-sdk/core/api_constant"
	"github.com/vfyuliawan/my-golang-sdk/core/audit"
	"github.com/vfyuliawan/my-golang-sdk/core/base_response"
)

const CacheStatefulSession = "CACHE_STATEFUL_SESSION"
const ApiContextKey = "apicontext"

type Metadata struct {
	api_constant.HeaderList
	ClientIp string `json:"client_ip"`
}

type User struct {
	UserId      int64
	Name        string
	Type        string
	PhoneNumber string
	Email       string
	Token       string
}

type Context struct {
	Metadata Metadata
	User     User
	Error    base_response.Error
	Audit    audit.Payload
}

func GetApiContext(c context.Context) *Context {
	if ctxValue, ok := c.Value(ApiContextKey).(*Context); ok {
		return ctxValue
	}
	ctx := checkApiContext(c)
	return GetApiContext(ctx)
}

func UpdateApiContext(c context.Context, apiContextPayload *Context) {
	if ginContext, ok := c.(*gin.Context); ok {
		ctx := ginContext.Request.Context()
		ctx = context.WithValue(ctx, ApiContextKey, apiContextPayload)
		ginContext.Request = ginContext.Request.WithContext(ctx)
		return
	}

	c = context.WithValue(c, ApiContextKey, apiContextPayload)

}

func checkApiContext(c context.Context) context.Context {
	apiContext := new(Context)

	// Case 1: If context is a Gin context
	if ginContext, ok := c.(*gin.Context); ok {
		if existingCtx, exists := ginContext.Value(ApiContextKey).(*Context); exists && existingCtx != nil {
			return ginContext.Request.Context() // Return existing context
		}

		// Create new context and update Gin's request
		ctx := ginContext.Request.Context()
		ctx = context.WithValue(ctx, ApiContextKey, apiContext)
		ginContext.Request = ginContext.Request.WithContext(ctx)
		ginContext.Set(ApiContextKey, apiContext) // Optional, for convenience
		return ctx
	}

	// Case 2: If context is a standard context
	if existingCtx, exists := c.Value(ApiContextKey).(*Context); exists && existingCtx != nil {
		return c // Return existing context
	}

	// Create new context with value
	newCtx := context.WithValue(c, ApiContextKey, apiContext)
	return newCtx
}
