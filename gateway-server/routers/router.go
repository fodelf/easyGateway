package routers

import (
	"github.com/gin-gonic/gin"

	_ "gateway/docs"

	proxy "gateway/middleware/proxy"
	v1 "gateway/routers/api/v1"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(proxy.ReverseProxy())
	// proxyMiddle := r.Group("/*")
	// proxyMiddle.Use(proxy.ReverseProxy())
	r.LoadHTMLGlob("web/*.html")          // 添加入口index.html
	r.LoadHTMLFiles("web/static/*/*")     // 添加资源路径
	r.Static("/static", "web/static")     // 添加资源路径
	r.StaticFile("/ui", "web/index.html") // 前端接口
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiRoot := r.Group("/uiApi/v1")
	apiRoot.Use()
	// 总览部分
	indexApi := apiRoot.Group("/index")
	{
		//获取首页汇总
		indexApi.GET("/sum", v1.GetSum)
		//获取图表信息
		indexApi.GET("/charts/:id", v1.GetCharts)
		//实时状态查询
		indexApi.GET("/actualTime", v1.GetActualTime)
		//实时状态查询
		indexApi.GET("/warningList", v1.GetWarningList)
	}
	eumnApi := apiRoot.Group("/eumn")
	{
		//服务类型
		eumnApi.GET("/serverTypeList", v1.GetServerType)
	}
	// 服务部分
	serviceApi := apiRoot.Group("/service")
	{
		//获取服务汇总
		serviceApi.GET("/serviceSum", v1.GetServerSum)
		//获取图表信息
		serviceApi.GET("/serviceList", v1.GetServerList)
		//新增服务
		serviceApi.POST("/addService", v1.ImportService)
		//查询服务明细
		serviceApi.GET("/serviceDetail", v1.GetServerDetail)
		//编辑服务
		serviceApi.POST("/editService", v1.EditService)
		//编辑服务
		serviceApi.POST("/deleteService", v1.DeleteService)
	}
	// 系统部分
	systemApi := apiRoot.Group("/system")
	{
		//修改mq
		systemApi.POST("/editRabbitMq", v1.EditRabbitMq)
		//修改consul
		systemApi.POST("/editConsul", v1.EditConsul)
		// 查询配置详情
		systemApi.GET("/systemDetail", v1.GetSystemDetail)
	}
	return r
}
