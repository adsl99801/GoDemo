package router

import (
	Base "GoDemo/dao"
	Members "GoDemo/dao/members"
	Model "GoDemo/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Home o
func Home(r *gin.Engine) {

	// r.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// })
	r.GET("/json", func(c *gin.Context) {
		base := Model.BaseResponse{Sus: true, Status: 0, Msg: ""}
		response := Model.LoginResponse{BaseResponse: base, Token: ""}
		c.JSON(http.StatusOK, response)
	})
	//curl localhost:8081/from --request POST --data '{"Sus":true,"Status":0,"Msg":"msg1","Token":""}'
	r.POST("/from", func(c *gin.Context) {
		var response Model.BaseResponse
		c.BindJSON(&response)
		log.Println(response)
		c.String(http.StatusOK, fmt.Sprintf("done:"+response.Msg))
	})
	r.GET("/members", func(c *gin.Context) {
		var dao = new(Members.MembersDao)
		db := Base.Open()
		defer db.Close()
		members := dao.All(db)
		if len(members) <= 0 {
			// dao.Insert()
			c.String(http.StatusOK, "members<=0")
			return
		}
		c.JSON(http.StatusOK, members)
	})

}
