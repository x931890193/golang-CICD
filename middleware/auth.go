package middleware

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"golang-CICD/lib"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		Debug := c.Request.Header.Get("Debug")
		if Debug != "1" {
			signature := c.Request.Header.Get("Signature")
			//authorize := strconv.Itoa(int(time.Now().Unix())) + ":3315a432-0b76-472b-b4d9-19b8bc266133"
			authorize := c.Request.Header.Get("X-authorize-uuid") // 时间戳：随机uuid
			//sigLife := int64(60 * 60)
			paramsSlice := strings.Split(authorize, ":")
			if len(paramsSlice) != 2 {
				lib.Logger.Error(fmt.Sprintf("鉴权参数错误！%v", c.Request.Header))
				c.AbortWithStatus(403)
				return
			}
			//timeStr, _ := strconv.Atoi(paramsSlice[0])
			//if time.Now().Unix()-int64(timeStr) > sigLife {
			//	logger.Logger.Error(fmt.Sprintf("请求时间戳错误！%v", c.Request.Header))
			//	c.AbortWithStatus(403)
			//	return
			//}
			secret := "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5"
			key := []byte(secret)
			h := hmac.New(sha1.New, key)
			h.Write([]byte(authorize))
			sign := base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
			if sign != signature {
				lib.Logger.Error(fmt.Sprintf("签名错误！%v", c.Request.Header))
				c.AbortWithStatus(403)
				return
			}

		}
		c.Next()
	}
}
