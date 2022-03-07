package gin_helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
)

// Handler when a func need to be a handle, assign it as Handler's
// public method
type Handler struct {
}

// MainHandler is a handler that could router request to a func
// with name like "Action+Version", Action and Version needed be
// given in context as query params. And you can use "doResp" to
// decide how to process the response
func MainHandler(c *gin.Context, doResp func(c *gin.Context, data interface{}, err error)) {
	action, ok := c.GetQuery("Action")
	if !ok {
		doResp(c, nil, errors.New("no action"))
		c.Abort()
	}
	version, ok := c.GetQuery("Version")
	if !ok {
		doResp(c, nil, errors.New("no version"))
		c.Abort()
	}
	h := Handler{}
	hv := reflect.ValueOf(h)
	f := hv.MethodByName(action + version)
	if !f.IsValid() {
		doResp(c, nil, errors.New("no such api"))
		c.Abort()
	}
	f.Call([]reflect.Value{reflect.ValueOf(c)})
}
