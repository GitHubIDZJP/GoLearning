package ginModule

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func MultipartUrlBind() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 可以使用显式绑定声明绑定多部分表单:
		//	 c.ShouldBindWith(&form, binding.Form)
		//或者您可以简单地使用ShouldBind方法自动绑定:
		var form LoginForm
		// 在这种情况下，将自动选择适当的绑定
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "登录了"})
			} else {
				c.JSON(401, gin.H{"status": "未经许可"})
			}
		}
	})
	router.Run(":8080")
}

/*
404 page not found


*/
