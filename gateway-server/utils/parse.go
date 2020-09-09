package Utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//获取结构体中字段的名称
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

//获取结构体中Tag的值，如果没有tag则返回字段值
func GetTagName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result, tagName)
	}
	return result
}
func GetStructValue(obj interface{}, name string) int64 {
	immutable := reflect.ValueOf(obj)
	val := immutable.FieldByName(name).Int()
	return val
}
func GetStructValueString(obj interface{}, name string) string {
	immutable := reflect.ValueOf(obj)
	val := immutable.FieldByName(name).String()
	return val
}

func GetJsonBody(c *gin.Context) map[string]interface{} {
	var (
		reqInfo map[string]interface{}
		body    []byte
		err     error
	)
	if body, err = ioutil.ReadAll(c.Request.Body); err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &reqInfo)
	return reqInfo
}

func GenerateUUID() string {
	withLineUUID := uuid.NewV4().String()
	uuidList := strings.Split(withLineUUID, "-")
	return strings.Join(uuidList, "")
}
