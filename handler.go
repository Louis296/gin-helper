package gin_helper

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"time"
)

// MainHandler will return a gin handler that could router request to a func
// with name like "Action"+"Version", Action and Version needed to be
// given in context as query params. And you can use "errHandler" to
// handle error which happened when cannot route. If it is nil, gin-helper
// will use a default error handler.
// Param h is a function receiver, which actually is a struct.
func MainHandler(errHandler func(c *gin.Context, err error), h interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		if errHandler == nil {
			errHandler = doErrResp
		}
		action, ok := c.GetQuery("Action")
		if !ok {
			errHandler(c, errors.New("no action"))
			c.Abort()
			return
		}
		version, ok := c.GetQuery("Version")
		if !ok {
			errHandler(c, errors.New("no version"))
			c.Abort()
			return
		}
		hv := reflect.ValueOf(h)
		f := hv.MethodByName(action + version)
		if !f.IsValid() {
			errHandler(c, errors.New("no such api"))
			c.Abort()
			return
		} else {
			f.Call([]reflect.Value{reflect.ValueOf(c)})
		}
	}
}

// response is a default resp
type response struct {
	Status string
	Time   time.Time
	Data   interface{}
}

func doErrResp(c *gin.Context, err error) {
	resp := response{
		Status: "Error",
		Time:   time.Now(),
	}
	resp.Data = struct {
		Message string
	}{Message: err.Error()}
	c.JSON(200, resp)
}

type ParserHandler interface {
	Handler(c *gin.Context) (interface{}, error)
}

type ProtectHandler interface {
	CheckPermission(c *gin.Context) (bool, error)
}

type LoggableHandler interface {
	OperatingLog(c *gin.Context) (bool, error)
}

func HandleParserHandler(c *gin.Context, h interface{}, respHandler func(c *gin.Context, data interface{}, err error)) {
	if respHandler == nil {
		respHandler = doResp
	}
	if handler, ok := h.(ParserHandler); ok {
		err := c.Bind(handler)
		if err != nil {
			fmt.Println("[gin-helper-error] " + err.Error())
		}
		data, err := handler.Handler(c)
		if err != nil {
			fmt.Println("[gin-helper-error] " + err.Error())
		}
		respHandler(c, data, err)
	}
}

func doResp(c *gin.Context, data interface{}, err error) {
	resp := response{Time: time.Now()}
	if err != nil {
		resp.Status = "Error"
		resp.Data = struct {
			Message string
		}{Message: err.Error()}
	} else {
		resp.Status = "Success"
		resp.Data = data
	}
	c.JSON(200, resp)
}
