package main //声明main包，表明当前是一个可执行程序
import (
	//导入外部包
	"crm/function1"
	"crm/function2"
	"crm/ginModule"
	"fmt"
)

//"GO学习/scr/crm/test/common/message"

// main函数 程序执行的入口
func main() {
	//s1 := "字符串"
	//s2 := "拼接"
	//s3 := s1 + s2
	//fmt.Println(s3) //s3 = "打印字符串"

	TestApiRequest()
	//CEMTERComMain_sd()
	//调用外部的函数
	function1.Function1000()
	function2.Function2()
	//ginModule.Model_binding_and_validation() //模型绑定和验证
	//ginModule.MultipartUrlBind() //URL编码绑定
	//ginModule.RenderingMain() //渲染
	//ginModule.ginNew() //中间件
	//ginModule.CEMTERComMain()
	//ginModule.A_file_upload() //单文件上传
	//ginModule.Multiple_file_upload() //多个文件上传
	//ginModule.HttpEncrypt() //加密,搞不懂
	//ginModule.SetOrFetchCookie() //设置并获取cookie
	//ginModule.ProvideStaticFile() //加载模版文件夹里的HTML
	//ginModule.FeedDataFromTheReader() //从阅读器提供数据
	//ginModule.SecureJSON() //安全json---->写了个Api

	//ginModule.RunMultipteServer() //运行多个服务

	//ginModule.Redirection() //重定向
	//ginModule.PostSendHttpRedirection() //从 POST 发出 HTTP 重定向
	//	ginModule.SenderRouterRedirection() //发出路由器重定向，使用HandleContext
	//ginModule.SearchStrParameter() //查询字符串参数

	//ginModule.QueryAndPublishForms() //查询和发布表单

	//ginModule.PureJSON() //纯JSON
	//ginModule.ParamtersPath() //路径中的参数//--没搞定有啥用
	//ginModule.BindOnlyInQueryString() //只绑定查询字符串
	//ginModule.BindOnlyInQueryStringTest() //只绑定查询字符串
	//ginModule.MultipleTemplatePlate() //模版
	ginModule.MultipleTemplateAdvanced() //多模版进阶
	//ginModule. //HTML
	//ginModule.
	//ginModule.
	//ginModule.
	//ginModule.
	//ginModule.
	//ginModule.
	//ginModule.
	//ginModule.

	//mainsd()

	//aa_main()
	//moduleMain()
	////fmt.Printin("叫爸爸123")
	fmt.Println("sdsdds")
	//get01()

	//init_fetch()
	//init_variable() //变量
	//init_constant() //常量
	//mainApi()
	//mainStudentApi() //学生接口
	fmt.Println("\n")
	//mainLogin()

	/**
	注意事项1：执行顺序只从上到下，同端口时 下面的会覆盖上面的
	前提是:每新建一个类来写一个接口
	网络Api,但是一旦只要端口号都是8080，回调方法按顺序来，哪个在下面哪个就执行，会直接覆盖上面的
	*/
	/*
	 注意事项2：执行顺序只从下到上，同端口时 上面的会覆盖下面的(run时只运行最上面的接口)
	 前提是：接口写一个类里才行
	*/
	//AsciiJSON_Api()
	//apiReequest()

	//api01()
	//requestSimpleApi() //请求最简单Api例子
	//queryGet()
	//queryParams()
	//DefaultQuery()
	//bashRequestFramework() //基础网络框架
	//post 请求Api

	//postRequest()
	//gorm1() //gorm快速入门--增删改查
	//gorm2() //gorm声明model
	//gorm3()//gorm模型使用
	//gorm4() //gorm模型加强
	//gorm5() //gorm模型加强1
	//gorm_batch_insert()
	//gorm_Creating_Record() //gorm样板
	//gorm_association_relationship_between_tables() //表之间关联关系
	//gorm_association_gl() //表之间关联关系1

	//官网:

	//mainKoyeb() //Koyeb
	//rootMain()

	//gin官网:
	//Model_binding_and_validation() //模型绑定和验证
	//MultipartUrlBind() //多部分/URL 编码绑定
	//t1 := time.Time{}
	//t2 := time.Now()

	//create_model_and_generate_table() //创建模型和生成表
	//PartialFieldPropertiesAreInserted() //部分字段属性插入
	//t1Second := t1.Unix()
	//t2Second := t2.Unix()
	//fmt.Printf("时间1:%v, 时间1Second:%v \n", t1, t1Second) // UTC时间1970-01-01 08:00:50的秒数为0,所以在此时间之前的为负值,且在int32范围之外
	//fmt.Printf("时间2:%v, 时间2Second:%v \n", t2, t2Second)
	//fmt.Println("t2Second-t1Second=", t2Second-t1Second)

}
