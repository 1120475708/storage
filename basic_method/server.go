package main

import (
	"basic_method/objects"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	objects.Handlers(r)
	r.Run(":8888")

}
