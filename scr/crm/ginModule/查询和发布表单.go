package ginModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
可以查看go函数
1.postRequest()
2.queryParams()
查询和发布表单

	query 查询
	DefaultQuery
	PostForm发布表单
*/
func QueryAndPublishForms() {
	router := gin.Default()
	//GET请求 url ?后面是querystring参数,key = value格式,多个key-value用&连接
	router.POST("/post", func(c *gin.Context) {
		///queryParams?name=randySun&age=18
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.SecureJSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "回传成功",
			//这下面是为了方便查看传回来的参数值，实际开发了前端传回来后直接显示回传成功code：200就行
			"id":   id,
			"page": page,
			"name": name,
		})
	})
	router.Run(":8080")

}

/*
post接口:
http://localhost:8080/post?id=1234&page=0&name=TestName&message=meg

回传成功的参数:
在终端可以看到
/post?id=1234&page=0&name=TestName&message=msg"



*/
