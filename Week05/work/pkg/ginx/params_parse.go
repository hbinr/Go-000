package ginx

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func QueryInt(param string, c *gin.Context) (intVar int64, err error) {
	intStr := c.Query(param)
	if intVar, err = strconv.ParseInt(intStr, 10, 64); err != nil {
		logrus.Error("strconv.ParseInt(intStr) failed,err:", err)
		return
	}
	return
}
