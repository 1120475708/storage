package objects

import (
	"fmt"
	"github.com/1120475708/common/utils"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func put(c *gin.Context) {
	path := utils.GetPrefixPath()
	fmt.Println(path)
	f, err := os.Create(path + c.Param("name"))
	if err != nil {
		log.Println(err)
		fmt.Fprint(c.Writer, http.StatusInternalServerError)
	}
	defer f.Close()
	io.Copy(f, c.Request.Body)

}
