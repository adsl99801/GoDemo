package router

import (
	"fmt"
	"log"
	"net/http"

	Model "GoDemo/model"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Post r
func Post(r *gin.Engine) {
	r.POST("/post/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, "/home/keith/original/upload/"+file.Filename)
		//curl -X POST http://localhost:8081/photo/upload  -F "file=@/home/keith/92.jpg" -H " Content-Type: multipart/form-data"
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.POST("/post/new", func(c *gin.Context) {
		var response Model.BaseResponse
		c.BindJSON(&response)
		c.String(http.StatusOK, fmt.Sprintf("done:"+response.Msg))
	})
}
