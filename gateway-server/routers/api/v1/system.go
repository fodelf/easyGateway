package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
)

// @Tags  系统模块
// @Summary 编辑consul数据
// @Description 编辑consul数据
// @Accept  json
// @Produce  json
// @Param address path string true "Address"  consul地址
// @Param port path string false "Port"  consul端口
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/system/editConsul [post]
func EditConsul(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
}

// @Tags  系统模块
// @Summary 编辑rabbitMq数据
// @Description 编辑rabbitMq数据
// @Accept  json
// @Produce  json
// @Param address path string true "Address"  abbitMq地址
// @Param port path string false "Port"  abbitMq端口
// @Param userName path string false "UserName"  用户名
// @Param psssword path string false "Psssword"  密码
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/system/editRabbitMq [post]
func EditRabbitMq(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
}

// @Tags  系统模块
// @Summary 系统详情
// @Description 查询系统配置详情
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/system/systemDetail [get]
func GetSystemDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"rabbitMq": map[string]interface{}{
			"address":  "187:10:10;10",
			"port":     "8500",
			"userName": "xxx",
			"psssword": "xxxxxxx",
		},
		"consul": map[string]interface{}{
			"address": "187:10:10;10",
			"port":    "8500",
		},
	})
}
