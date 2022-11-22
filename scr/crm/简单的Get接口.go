package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func apiReequest() {
	var r = gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "成功了---哈哈")
	})
	r.Run(":8080")
}

//在终端输入http://localhost:8181/,会显示成功了，这代码就是监听端口号8181
