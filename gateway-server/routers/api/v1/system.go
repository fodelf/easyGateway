package v1

import (
	InterfaceEntity "gateway/models/InterfaceEntity"
	"gateway/pkg/e"
	"net/http"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "strconv"
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
type ConsulPost struct {
	ConsulAddress string `json:"address"`
	ConsulPort    int    `json:"port"`
}

func EditConsul(c *gin.Context) {
	appG := app.Gin{C: c}
	var consulPost ConsulPost
	if err := c.ShouldBind(&consulPost); err != nil {
		return
	}
	consulInfo := InterfaceEntity.ConsulInfo{
		ConsulAddress: consulPost.ConsulAddress,
		ConsulPort:    consulPost.ConsulPort,
	}
	// fmt.Println(consulInfo)
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc")
	tx := DB.Begin()
	if err := tx.Model(&InterfaceEntity.ConsulInfo{}).Update(&consulInfo).Error; err != nil {
		tx.Rollback()
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		tx.Commit()
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
	}
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
type RabbitMqPost struct {
	RabbitMqAddress  string `json:"address"`
	RabbitMqPort     int    `json:"port"`
	RabbitMqUserName string `json:"userName"`
	RabbitMqPassWord string `json:"password"`
}

func EditRabbitMq(c *gin.Context) {
	appG := app.Gin{C: c}
	var rabbitMqPost RabbitMqPost
	if err := c.ShouldBind(&rabbitMqPost); err != nil {
		return
	}
	rabbitMq := InterfaceEntity.RabbitMQInfo{
		RabbitMQAddress:  rabbitMqPost.RabbitMqAddress,
		RabbitMQPort:     rabbitMqPost.RabbitMqPort,
		RabbitMQUserName: rabbitMqPost.RabbitMqUserName,
		RabbitMQPassword: rabbitMqPost.RabbitMqPassWord,
	}
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc")
	tx := DB.Begin()
	if err := tx.Model(&InterfaceEntity.RabbitMQInfo{}).Update(&rabbitMq).Error; err != nil {
		tx.Rollback()
		appG.Response(http.StatusOK, e.ERROR, map[string]interface{}{})
	} else {
		tx.Commit()
		appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{})
	}
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
	var (
		consulInfo   InterfaceEntity.ConsulInfo
		rabbitMQInfo InterfaceEntity.RabbitMQInfo
	)
	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc")
	if err := DB.First(&consulInfo).Error; err != nil {
	} else {

	}
	if err := DB.First(&rabbitMQInfo).Error; err != nil {
	} else {

	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"rabbitMq": rabbitMQInfo,
		"consul":   consulInfo,
	})
	DB.Close()
}
