package database

import (
	"log"

	"github.com/xormplus/xorm" // 记得 go get 获取哦
)

// ORM xorm引擎的实例，供其他模块可以直接使用，注意**首字母大写**，因为`Go`语音的隐藏和公开的规则，大写为公开，小写为隐藏。
var ORM *xorm.Engine

func init() { // 使用init来自动连接数据库，并创建ORM实例
	var err error
	ORM, err = xorm.NewEngine("sqlite3", "./database/test.db")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = ORM.Ping() // 测试能操作数据库
	if err != nil {
		log.Fatalln(err)
		return
	}
	ORM.ShowSQL(true) // 测试环境，显示每次执行的sql语句长什么样子
}
