package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gateway/pkg/app"
	"gateway/pkg/e"
)

// @Tags  字典项模块
// @Summary 服务类型
// @Description 查询服务类型
// @Accept  json
// @Produce  json
// @Success 200 {string} string	Result 成功后返回值
// @Router /uiApi/v1/eumn/serverTypeList [get]
func GetServerType(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"serverTypeList": [1]map[string]interface{}{
			{
				"label": "http",
				"value": "http",
			},
		},
	})
}
