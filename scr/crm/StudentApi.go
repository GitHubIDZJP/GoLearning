package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 接口转为 JSON 格式
// type Student map[string]interface{}

// type Address map[string]interface{ data }

type resultInfo struct {
	stat string

	arrc []zjp_UserInfo //定义数据类型结构体
}

type zjp_UserInfo struct {
	Name     string `json:"name"`
	Age      string `json:"age"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Url      string `json:"url"`
}

func mainStudentApi() {
	//var arrCount int = 20
	//var vv = [4]zjp_UserInfo{}
	//var arr = [...]zjp_UserInfo{}
	//var name = [4]string{"王强", "刘树", "罗强", "画画"}
	//var age string = [4]string{"王强", "刘树", "罗强", "画画"}
	//var title string = [4]string{"王强", "刘树", "罗强", "画画"}
	//var Category string = [4]string{"王强", "刘树", "罗强", "画画"}
	//var url string = [4]string{"王强", "刘树", "罗强", "画画"}
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"}) //设置代理，必须加上，不然下面会提示代理提示权限不安全
	r.GET("/", func(obj *gin.Context) {

		var uData resultInfo
		uData.stat = "sdsd"

		ss := zjp_UserInfo{
			Name:     "王强",
			Age:      "23",
			Title:    "一条倡议 一分钟征来10辆车",
			Category: "头条",
			Url:      "https://mini.eastday.com/mobile/221115053930984942512.html",
		}
		ss1 := zjp_UserInfo{
			Name:     "习近平",
			Age:      "20",
			Title:    "甘肃农信系统切换停业公告",
			Category: "小红书",
			Url:      "https://mini.eastday.com/mobile/221115052848888128599.html",
		}
		//ss := zjp_UserInfo{
		//	Name:     name,
		//	Age:      "23",
		//	Title:    "一条倡议 一分钟征来10辆车",
		//	Category: "头条",
		//	Url:      "https://mini.eastday.com/mobile/221115053930984942512.html",
		//}

		//for i := 0; i < len(name); i++ {
		//	//fmt.Println("打印i:\n", i)
		//	//append(vv, ss1)
		//	//ss.Name = "sd" //string(i)
		//	//ss.Name = "名字" + string(i)
		//	var sc string = string(i)
		//	sc += fmt.Sprintf("%s%d", name, i)
		//	fmt.Println("醋酸钠:\n", sc)
		//	ss := zjp_UserInfo{
		//		Name:     string(name[i]),
		//		Age:     "23",
		//		Title:    string(name[i]),
		//		Category: string(name[i]),
		//		Url:      string(name[i]),
		//	}
		//	//fmt.Println("打印数组sdsd:\n", ss.Name)
		//	vv = [4]zjp_UserInfo{ss}
		//	arr.apped(vv, ss)
		//}

		vv := []zjp_UserInfo{ss, ss1}
		//fmt.Println("打印数组:\n", uData.arrc)

		Data := map[string]interface{}{
			"stat": "1",
			"data": vv,
		}

		obj.JSON(http.StatusOK, gin.H{

			"status":   200,
			"message":  "响应成功",
			"page":     "1",
			"pageSize": len(vv),
			"result":   Data,
		})

	})
	r.Run(":8080")

}
