package main

import "github.com/gin-gonic/gin"

type errString string

func (e errString) Error() string {
	return string(e)
}

func main() {
	r := gin.Default()
	r.GET("/get", Get)

}

func Get(c *gin.Context) {
}
