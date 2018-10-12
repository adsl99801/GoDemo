package router

import (
	Dao "GoDemo/dao"
	Model "GoDemo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Home o
func Home(r *gin.Engine) {

	r.POST("/home/login", func(c *gin.Context) {
		base := Model.BaseResponse{Sus: true, Status: 0, Msg: ""}
		response := Model.LoginResponse{BaseResponse: base, Token: ""}
		c.JSON(http.StatusOK, response)
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user": "u1", "value": "v1"})
	})
	r.GET("/data", func(c *gin.Context) {
		dao := new(Dao.MembersDao)
		members := dao.All()
		if len(members) <= 0 {
			dao.Insert()
			c.String(http.StatusOK, "members<=0")
			return
		}
		c.String(http.StatusOK, "data:"+members[0].Name)
	})
}
