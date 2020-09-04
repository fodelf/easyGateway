package proxy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := "http://172.23.0.187:9990/"
		u, err := url.Parse(target)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(u.Host)
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "172.23.0.187:9990"
			req.URL.Path = "/cloud-platform/system/searchAllSystem"
		}
		//更改内容
		modifyFunc := func(resp *http.Response) error {
			//todo 部分章节功能补充2
			//todo 兼容websocket
			if strings.Contains(resp.Header.Get("Connection"), "Upgrade") {
				return nil
			}
			var payload []byte
			var readErr error

			//todo 部分章节功能补充3
			//todo 兼容gzip压缩
			if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
				gr, err := gzip.NewReader(resp.Body)
				if err != nil {
					return err
				}
				payload, readErr = ioutil.ReadAll(gr)
				resp.Header.Del("Content-Encoding")
			} else {
				payload, readErr = ioutil.ReadAll(resp.Body)
			}
			if readErr != nil {
				return readErr
			}

			//异常请求时设置StatusCode
			if resp.StatusCode != 200 {
				payload = []byte("StatusCode error:" + string(payload))
			}

			//todo 部分章节功能补充4
			//todo 因为预读了数据所以内容重新回写
			c.Set("status_code", resp.StatusCode)
			c.Set("payload", payload)
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(payload))
			resp.ContentLength = int64(len(payload))
			resp.Header.Set("Content-Length", strconv.FormatInt(int64(len(payload)), 10))
			return nil
		}
		//错误回调 ：关闭real_server时测试，错误回调
		//范围：transport.RoundTrip发生的错误、以及ModifyResponse发生的错误
		errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
			//todo record error log
			fmt.Println(err)
		}
		proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc, ErrorHandler: errFunc}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
