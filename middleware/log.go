package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang-CICD/handler"
	"golang-CICD/lib"
	"io/ioutil"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var blw bodyLogWriter
		var reqBodyStr interface{}
		blw = bodyLogWriter{bodyBuf: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		httpMethod := c.Request.Method
		httpProto := c.Request.Proto
		httpPath := c.Request.URL.Path
		httpHeader := c.Request.Header
		uuid := c.Request.Header.Get("X-Authorize-Uuid")
		realIp := c.Request.Header.Get("X-Real-Ip")
		httpClientIP := c.Request.RemoteAddr

		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		c.Next()
		body := blw.bodyBuf.Bytes()
		resp := handler.Response{}
		_ = json.Unmarshal(body, &resp)
		_ = json.Unmarshal(reqBody, &reqBodyStr)
		logData := map[string]interface{}{
			"RequestMethod":   httpMethod,
			"RequestHeader":   httpHeader,
			"RequestPath":     httpPath,
			"RequestProto":    httpProto,
			"RequestBody":     reqBodyStr,
			"RequestClientIP": httpClientIP,
			"Resp":            resp,
		}
		ugly, _ := json.Marshal(logData)
		fields := logrus.Fields{
			"RequestPath":      httpPath,
			"httpMethod":       httpMethod,
			"RequestHeader":    httpHeader,
			"RequestClientIP":  httpClientIP,
			"X-Authorize-Uuid": uuid,
			"X-Real-Ip":        realIp,
			"Req":              "",
			"Resp":             "logRes",
			"httpStatusCode":   c.Writer.Status(),
			"RequestBody":      fmt.Sprintf("%s", reqBody),
			"resp.code":        resp.Code,
			"resp.msg":         resp.Msg,
		}
		lib.Logger.WithFields(fields).Infof("%s", ugly)
	}
}
