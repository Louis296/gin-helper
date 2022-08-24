package gin_helper

import (
	"github.com/gin-gonic/gin"
	"testing"
)

type Handler struct {
}

func (r Handler) Test20220101(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Test001"})
}

func TestMainHandler(t *testing.T) {
	r := gin.Default()
	//r.GET("/v1", MainHandler(func(c *gin.Context, err error) {
	//	c.JSON(200, gin.H{"err": err.Error()})
	//},reflect.ValueOf(Handler{})))
	r.GET("/v1", MainHandler(nil, Handler{}))
	err := r.Run(":8081")
	if err != nil {
		return
	}
}

// controller layer
func (r Handler) ParserHandlerTest20220101(c *gin.Context) {
	HandleParserHandler(c, &SampleReq{}, nil)
}

// service layer
type SampleReq struct {
	Message string
}

type SampleResp struct {
	Message string
}

func (r *SampleReq) Handler(c *gin.Context) (interface{}, error) {
	return SampleResp{Message: "hello " + r.Message}, nil
}

func TestParserHandler(t *testing.T) {
	r := gin.Default()
	r.GET("/v1", MainHandler(nil, Handler{}))
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
