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

func get(c *gin.Context) {
	path := utils.GetPrefixPath()
	fmt.Println(path)
	f, err := os.Open(path + c.Param("name"))
	if err != nil {
		log.Println(err)
		fmt.Fprint(c.Writer, http.StatusInternalServerError)
	}
	defer f.Close()
	io.Copy(c.Writer, f)

}
