package objects

import "github.com/gin-gonic/gin"

func Handlers(r *gin.Engine) {

	group := r.Group("/objects/")
	{
		group.GET("/getObject/:name", get)
		group.PUT("/putObject/:name", put)
	}

}
