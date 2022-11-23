package main //声明main包，表明当前是一个可执行程序
import (
	"fmt"
)

//"GO学习/scr/crm/test/common/message"

// main函数 程序执行的入口
func main() {
	aa_main()
	//moduleMain()
	////fmt.Printin("叫爸爸123")
	fmt.Println("sdsdds")
	get01()

	init_fetch()
	init_variable() //变量
	init_constant() //常量
	mainApi()
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
