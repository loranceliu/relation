package svc

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// ServiceContext 实现 context.Context 接口
type ServiceContext struct {
	GinContext *gin.Context
}

func NewServiceContext(c *gin.Context) *ServiceContext {
	return &ServiceContext{
		GinContext: c,
	}
}

func (c *ServiceContext) Deadline() (deadline time.Time, ok bool) {
	return c.GinContext.Deadline()
}

func (c *ServiceContext) Done() <-chan struct{} {
	return c.GinContext.Done()
}

func (c *ServiceContext) Err() error {
	return c.GinContext.Err()
}

func (c *ServiceContext) Value(key any) string {
	obj := c.GinContext.Value(key)
	var value string
	if floatValue, ok := obj.(float64); ok {
		value = strconv.FormatFloat(floatValue, 'f', -1, 64)
	}

	return value
}
