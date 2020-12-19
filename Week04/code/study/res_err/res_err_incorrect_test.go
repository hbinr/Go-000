package res_err

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

const (
	CodeUserNotExit = 10001
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func TestResErr(t *testing.T) {
	res := Response{
		Code: CodeUserNotExit, // 用户不存在业务状态码
		Msg:  "用户不存在",
		Data: nil,
	}

	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, res) // 正常http响应+res
	})
	t.Logf("res:%v\n", res)

	r.Run()

}
