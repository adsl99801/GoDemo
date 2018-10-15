package router

import (
	"fmt"
	"log"
	"net/http"

	Model "GoDemo/model"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Upload o
func Upload(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	c.SaveUploadedFile(file, "/home/keith/original/upload/"+file.Filename)
	//curl -X POST http://localhost:8081/photo/upload  -F "file=@/home/keith/92.jpg" -H " Content-Type: multipart/form-data"
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

//NewOrder o
func NewOrder(c *gin.Context) {
	var response Model.BaseResponse
	c.BindJSON(&response)
	c.String(http.StatusOK, fmt.Sprintf("done:"+response.Msg))
}

//Test r
func Test(c *gin.Context) {
	// curl -H "token:t1" http://localhost:8081/post/test
	userID := c.MustGet("userID").(string)
	c.JSON(http.StatusOK, gin.H{"post": "u1", "userID": "" + userID})
}
