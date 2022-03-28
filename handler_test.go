package gin_helper

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func (r Handler) Test001(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Test001"})
}

func TestMainHandler(t *testing.T) {
	r := gin.Default()
	r.GET("/v1", MainHandler(func(c *gin.Context, err error) {
		c.JSON(200, gin.H{"err": err.Error()})
	}))
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
