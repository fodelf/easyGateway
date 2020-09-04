package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/EDDYCJY/go-gin-example/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	v1 "gateway/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("web/*.html")          // 添加入口index.html
	r.LoadHTMLFiles("web/static/*/*")     // 添加资源路径
	r.Static("/static", "web/static")     // 添加资源路径
	r.StaticFile("/ui", "web/index.html") // 前端接口

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/uiApi/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// //新建标签
		// apiv1.POST("/tags", v1.AddTag)
		// //更新指定标签
		// apiv1.PUT("/tags/:id", v1.EditTag)
		// //删除指定标签
		// apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// //导出标签
		// r.POST("/tags/export", v1.ExportTag)
		// //导入标签
		// r.POST("/tags/import", v1.ImportTag)

		// //获取文章列表
		// apiv1.GET("/articles", v1.GetArticles)
		// //获取指定文章
		// apiv1.GET("/articles/:id", v1.GetArticle)
		// //新建文章
		// apiv1.POST("/articles", v1.AddArticle)
		// //更新指定文章
		// apiv1.PUT("/articles/:id", v1.EditArticle)
		// //删除指定文章
		// apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		// //生成文章海报
		// apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
