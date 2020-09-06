package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
)
// @Tags  服务模块
// @Summary 服务数据汇总
// @Description 查询服务数据汇总
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/serviceSum [get]
func GetServerSum(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"serverList": [1]map[string]interface{}{
		 {
		  "label":"http",
		  "count":1,
		  "value":"http",
		 },
		},
		"sum":1,
	})
}
type Rule struct {
    url string
    pathReWrite string
}
// @Tags  服务模块
// @Summary 服务列表
// @Description 查询服务列表
// @Accept  json
// @Produce  json
// @Param type path string false "Type"  查询服务类型	
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/serviceList [get]
func GetServerList(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"serverList": [1]map[string]interface{}{
		 {
		  "serverId":"xxxxxx",
		  "serviceName":"XX",
		  "serviceType":"http",
		  "serviceAddress":"127.0.0.1",
		  "servicePort":"300",
		  "serviceLmit":200,
		  "serviceBreak":30,
		  "serviceRules":[1]map[string]interface{}{
			{
				"url":"/api",
				"pathReWrite":`{/api:''}`,
			},
		   },
		   "useConsulId":"xxxxx",
		   "useConsulTag":"ss",
		   "useConsulCheckPath":"/checkHealth",
		   "useConsulPort":9990,
		   "useConsulInterval":10,
		   "useConsulTimeout":3,
		   "dingdingAccessToken":"xxxxxxxxxx",
		   "dingdingSercet":"xxxxxxxxxx",
		   "dingdingList":[]int{18651892475},
		 },
		},
		"total":1,
	})
}
// @Tags  服务模块
// @Summary 服务详情
// @Description 查询服务详情
// @Accept  json
// @Produce  json
// @Param type path string false "Type"  查询服务类型	
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/serviceDetail/{id} [get]
func GetServerDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"serverId":"xxxxxx",
		"serviceName":"XX",
		"serviceType":"http",
		"serviceAddress":"127.0.0.1",
		"servicePort":"300",
		"serviceLimt":200,
		"serviceBreak":30,
		"serviceRules":[1]map[string]interface{}{
		{
			"url":"/api",
			"pathReWrite":`{/api:''}`,
		},
		},
		"useConsulId":"xxxxx",
		"useConsulTag":"ss",
		"useConsulCheckPath":"/checkHealth",
		"useConsulPort":9990,
		"useConsulInterval":10,
		"useConsulTimeout":3,
		"dingdingAccessToken":"xxxxxxxxxx",
		"dingdingSercet":"xxxxxxxxxx",
		"dingdingList":[]int{18651892475},
	})
}
// @Tags  服务模块
// @Summary 新增服务
// @Description 新增服务
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/addService [post]
func ImportService(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
	})
}
// @Tags  服务模块
// @Summary 编辑服务
// @Description 编辑服务
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/editService [post]
func EditService(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
	})
}
// @Tags  服务模块
// @Summary 删除服务
// @Description 删除服务
// @Accept  json
// @Produce  json
// @Param serviceId path string true "ServiceId"  删除服务的id
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/service/deleteService [post]
func DeleteService(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
	})
}
