package gin_helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
)

// Handler is a func receiver. When a func need to be a handler, assign it as Handler's
// public method
type Handler struct {
}

// MainHandler will return a handler that could router request to a func
// with name like "Action"+"Version", Action and Version needed be
// given in context as query params. And you can use "doResp" to
// decide how to process the response when a error happened
func MainHandler(doErrResp func(c *gin.Context, err error)) func(*gin.Context) {
	return func(c *gin.Context) {
		action, ok := c.GetQuery("Action")
		if !ok {
			doErrResp(c, errors.New("no action"))
			c.Abort()
		}
		version, ok := c.GetQuery("Version")
		if !ok {
			doErrResp(c, errors.New("no version"))
			c.Abort()
		}
		h := Handler{}
		hv := reflect.ValueOf(h)
		f := hv.MethodByName(action + version)
		if !f.IsValid() {
			doErrResp(c, errors.New("no such api"))
			c.Abort()
		} else {
			f.Call([]reflect.Value{reflect.ValueOf(c)})
		}
	}
}
