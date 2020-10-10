package proxy

import (
	// "bytes"
	// "compress/gzip"

	"errors"
	"gateway/hystrix"
	"gateway/models/InterfaceEntity"
	"time"

	// "io/ioutil"

	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	ProxyConfig    []map[string]interface{}
	RequestSum     int = 0 // 请求成功汇总
	Total          int = 0
	Conn           interface{}
	MyFirstChannel = make(chan string)
)

// var mqConn *amqp.Connection

// func init() {
// 	mqConn, _ := amqp.Dial("amqp://admin:admin@111.229.133.9:5672/")
// }

func grepProxy(url string) map[string]interface{} {
	var (
		ServerConfig map[string]interface{} = map[string]interface{}{
			"flag": true,
		}
		urlkey string
	)
	// conn, _ := amqp.Dial("amqp://admin:admin@111.229.133.9:5672/")
	// ch, _ := conn.Channel()
	// // q, _ := ch.QueueDeclare(
	// // 	"hello", // name
	// // 	false,   // durable
	// // 	false,   // delete when unused
	// // 	false,   // exclusive
	// // 	false,   // no-wait
	// // 	nil,     // arguments
	// // )
	// body := "Hello World!"
	// ch.Publish(
	// 	"",      // exchange
	// 	"hello", // routing key
	// 	false,   // mandatory
	// 	false,   // immediate
	// 	amqp.Publishing{
	// 		ContentType: "text/plain",
	// 		Body:        []byte(body),
	// 	})
	// ch.Close()
	// conn.Close()
	// ProxyConfig = append(ProxyConfig, map[string]interface{}{
	// 	"serviceAddress": "127.0.0.1",
	// 	"prot":           3000,
	// 	"serviceRules": []map[string]interface{}{
	// 		{
	// 			"url":         "/api",
	// 			"pathReWrite": `{/api:''}`,
	// 		},
	// 	},
	// })
	for i := 0; i < len(ProxyConfig); i++ {
		var child = ProxyConfig[i]
		var rules = child["serviceRules"].([]map[string]interface{})
		for k := 0; k < len(rules); k++ {
			var rule = rules[k]
			urlkey = rule["url"].(string)
			if strings.HasPrefix(url, urlkey) {
				var pathReWrite = rule["pathReWrite"].(map[string]interface{})
				ServerConfig = map[string]interface{}{
					"flag":           false,
					"serviceAddress": child["serviceAddress"].(string),
					"servicePort":    child["servicePort"].(int),
					"serviceRule": map[string]interface{}{
						"url": urlkey,
						"pathReWrite": map[string]interface{}{
							"oldPath": pathReWrite["oldPath"].(string),
							"newPath": pathReWrite["newPath"].(string),
						},
					},
					"index":    i,
					"serverId": child["serverId"].(string),
				}
				break
			}
		}
	}
	return ServerConfig
}

func ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
		// serviceInfoCount InterfaceEntity.ServiceInfo
		// sum              int
		// sumInfo InterfaceEntity.SumInfo
		// chartInfo  InterfaceEntity.ChartInfo
		// chartInfo InterfaceEntity.ChartInfo
		)
		urlPath := c.Request.URL.String()
		var proxyObj = grepProxy(urlPath)
		if proxyObj["flag"].(bool) || strings.HasPrefix(urlPath, "/uiApi") {
			c.Next()
			return
		}
		// fmt.Println("前面")

		// MyFirstChannel <- "hello" // Send
		// // myVariable := <-MyFirstChannel
		// // fmt.Println(myVariable)
		// fmt.Println("后面")
		// payload := Payload{}
		// work := Job{Payload: payload}
		// var complete chan int = make(chan int)
		// // Push the work onto the queue.
		// fmt.Println("前面")
		// // JobQueue <- work
		// go func() {
		// 	complete <- 0
		// }()

		// <-complete
		// fmt.Println("后面")
		// ch := make(chan int)
		// // c := 0
		// // stopCh := make(chan bool)
		// ch <- 10
		// // go Chann(ch, stopCh)

		// for {
		// 	select {
		// 	case c := <-ch:
		// 		fmt.Println("Recvice", c)
		// 		fmt.Println("channel")
		// 		// case s := <-ch:
		// 		// 	fmt.Println("Receive", s)
		// 	}
		// }
		// fmt.Println("后面")
		// fmt.Println(<-pkg.JobQueue)
		// f1 := func() error {
		// 	fmt.Println("处理业务逻辑")
		// 	return nil
		// }
		// //第二步：
		// //回调函数，只有 err不为空，才会执行回调函数(如果发生了超时，熔断，
		// //限流，超时之后也会回调)
		// fallBack := func(err error) error {
		// 	fmt.Println("熔断了")
		// 	return err
		// }
		// //第三步：
		// if err := hystrix.Do(proxyObj["serverId"].(string), f1, fallBack); err != nil {
		// 	c.Next()
		// 	fmt.Println("不执行后面的逻辑")
		// 	return
		// }
		// DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc&_journal_mode=WAL")
		// if err := DB.First(&sumInfo).Update("request_sum", gorm.Expr("request_sum + ?", 1)).Error; err != nil {
		// }
		// if err := DB.First(&chartInfo).Where("time = ? AND server_id = ?", time.Now().Format("2006/01/02"), "all").Update("total", gorm.Expr("total + ?", 1)).Error; err != nil {
		// }
		// if err := DB.First(&chartInfo).Where("time = ? AND server_id = ?", time.Now().Format("2006/01/02"), proxyObj["serverId"]).Update("total", gorm.Expr("total + ?", 1)).Error; err != nil {
		// }
		// DB.Close()
		err := hystrix.Do(proxyObj["serverId"].(string), func() error {
			var serviceAddress = proxyObj["serviceAddress"].(string)
			var prot = strconv.Itoa(proxyObj["servicePort"].(int))
			var serviceRule = proxyObj["serviceRule"].(map[string]interface{})
			var pathReWrite = serviceRule["pathReWrite"].(map[string]interface{})
			var oldPath = pathReWrite["oldPath"].(string)
			var newPath = pathReWrite["newPath"].(string)
			var Host = serviceAddress + ":" + prot
			director := func(req *http.Request) {
				req.URL.Scheme = "http"
				req.URL.Host = Host
				req.URL.Path = strings.Replace(urlPath, oldPath, newPath, 1)
			}
			var proxyError error = nil
			errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
				var chartInfo InterfaceEntity.ChartInfo
				DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc&_journal_mode=WAL")
				if err := DB.First(&chartInfo).Where("time = ? AND server_id = ?", time.Now().Format("2006/01/02"), proxyObj["serverId"]).Update("fail", gorm.Expr("fail + ?", 1)).Error; err != nil {
				}
				DB.Close()
				proxyError = err
			}
			proxy := &httputil.ReverseProxy{Director: director, ErrorHandler: errFunc}
			proxy.ServeHTTP(c.Writer, c.Request)
			return proxyError
		}, func(err error) error {
			return errors.New("子系统熔断")
		})
		if err != nil {
			// c.Abort()
			c.JSON(http.StatusUnauthorized, "子系统熔断")
			// c.Next()
		} else {
			// c.Next()
		}
		// err1 := hystrix.Do(proxyObj["serverId"].(string), func() error {
		// 	if flag == true {
		// 		fmt.Println("子系统异常")
		// 		return errors.New("子系统异常")
		// 	}
		// 	return nil
		// }, nil)
		// if err1 != nil {
		// 	log.Println("hystrix breaker err: ", err1)
		// }

		// return
		// c.Next()
		// fmt.Println("继续")
		// DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc&_journal_mode=WAL")
		// if err := DB.First(&sumInfo).Update("request_sum", gorm.Expr("request_sum + ?", 1)).Error; err != nil {
		// }
		// if err := DB.First(&chartInfo).Where("time = ? AND server_id = ?", time.Now().Format("2006/01/02"), "all").Update("total", gorm.Expr("total + ?", 1)).Error; err != nil {
		// }
		// if err := DB.First(&chartInfo).Where("time = ? AND server_id = ?", time.Now().Format("2006/01/02"), proxyObj["serverId"]).Update("total", gorm.Expr("total + ?", 1)).Error; err != nil {
		// }
		// var serviceAddress = proxyObj["serviceAddress"].(string)
		// var prot = strconv.Itoa(proxyObj["servicePort"].(int))
		// var serviceRule = proxyObj["serviceRule"].(map[string]interface{})
		// var pathReWrite = serviceRule["pathReWrite"].(map[string]interface{})
		// var oldPath = pathReWrite["oldPath"].(string)
		// var newPath = pathReWrite["newPath"].(string)
		// var Host = serviceAddress + ":" + prot
		// director := func(req *http.Request) {
		// 	req.URL.Scheme = "http"
		// 	req.URL.Host = Host
		// 	req.URL.Path = strings.Replace(urlPath, oldPath, newPath, 1)
		// }
		//更改内容
		// modifyFunc := func(resp *http.Response) error {
		// 	//todo 部分章节功能补充2
		// 	//todo 兼容websocket
		// 	if strings.Contains(resp.Header.Get("Connection"), "Upgrade") {
		// 		return nil
		// 	}
		// 	var payload []byte
		// 	var readErr error

		// 	//todo 部分章节功能补充3
		// 	//todo 兼容gzip压缩
		// 	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		// 		gr, err := gzip.NewReader(resp.Body)
		// 		if err != nil {
		// 			return err
		// 		}
		// 		payload, readErr = ioutil.ReadAll(gr)
		// 		resp.Header.Del("Content-Encoding")
		// 	} else {
		// 		payload, readErr = ioutil.ReadAll(resp.Body)
		// 	}
		// 	if readErr != nil {
		// 		return readErr
		// 	}

		// 	//异常请求时设置StatusCode
		// 	if resp.StatusCode != 200 {
		// 		payload = []byte("StatusCode error:" + string(payload))
		// 	}

		// 	//todo 部分章节功能补充4
		// 	//todo 因为预读了数据所以内容重新回写
		// 	c.Set("status_code", resp.StatusCode)
		// 	c.Set("payload", payload)
		// 	resp.Body = ioutil.NopCloser(bytes.NewBuffer(payload))
		// 	resp.ContentLength = int64(len(payload))
		// 	resp.Header.Set("Content-Length", strconv.FormatInt(int64(len(payload)), 10))
		// 	return nil
		// }
		//错误回调 ：关闭real_server时测试，错误回调
		//范围：transport.RoundTrip发生的错误、以及ModifyResponse发生的错误
		// errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
		// 	DB, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc&_journal_mode=WAL")
		// 	if err := DB.Where("time = ? AND server_id = ?", time.Now().Format("2006/01/02"), proxyObj["serverId"]).Update("fail", gorm.Expr("fail + ?", 1)).First(&chartInfo).Error; err != nil {
		// 	}
		// 	DB.Close()
		// }
		// proxy := &httputil.ReverseProxy{Director: director, ErrorHandler: errFunc}
		// proxy.ServeHTTP(c.Writer, c.Request)
		// c.Next()
		// DBNext, _ := gorm.Open("sqlite3", "gateway.sqlite?cache=shared&mode=rwc&_journal_mode=WAL")
		// var chartInfo1 InterfaceEntity.ChartInfo
		// DBNext.Close()
	}
}
