package gin

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	logger "github.com/cheetah-fun-gs/goplus/multier/multilogger"
	"github.com/cheetah-fun-gs/lego/internal/biz/handler"
	legogin "github.com/cheetah-fun-gs/lego/pkg/svc/gin"
	"github.com/gin-gonic/gin"
)

// middlewareCORS 处理跨域问题
func middlewareCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
	}
}

// middlewareLegoGin 处理 lego gin 的 httpcode 不为 200 的问题
func middlewareLegoGin(c *gin.Context) {
	// Process request
	c.Next()

	if legoErr, ok := c.Get(legogin.LegoHandlerErr); ok {
		err := fmt.Errorf("%v %v", legoErr, c.MustGet(legogin.LegoHandlerMsg))
		commonResp := handler.GetCommonResp(handler.CommonRespCodeServerUnknown, err)
		commonResp.RequestID = c.MustGet(legogin.LegoRequestID).(string)
		c.Set(logRespCode, commonResp.Code)
		c.Set(logRespMsg, commonResp.Msg)
		c.AbortWithStatusJSON(http.StatusOK, commonResp)
	}
}

const (
	logRespCode = "logRespCode"
	logRespMsg  = "logRespMsg"
)

// middlewareLogger 日志
func middlewareLogger(c *gin.Context) {
	// Start timer
	start := time.Now()

	path := c.Request.URL.Path
	method := c.Request.Method
	// Process request
	c.Next()

	if method == "OPTIONS" || method == "HEAD" {
		return
	}

	// Stop timer
	cost := time.Now().Sub(start).Milliseconds()
	clientIP := c.ClientIP()
	httpCode := c.Writer.Status()

	respCode := "-"
	if code, ok := c.Get(logRespCode); ok {
		respCode = strconv.Itoa(code.(int))
	}

	respMsg := "-"
	if msg, ok := c.Get(logRespMsg); ok {
		respMsg = msg.(string)
	}

	requestID := "-"
	if requestID, ok := c.Get(legogin.LegoRequestID); ok {
		requestID = requestID.(string)
	}

	contentSplit := []string{requestID, clientIP, method, path, fmt.Sprintf("%dms", cost),
		strconv.Itoa(httpCode), respCode, respMsg}
	content := strings.Join(contentSplit, ",")

	if httpCode == 200 && (respCode == "0" || respCode == "-") {
		logger.InfoN(nil, "request", content)
	} else {
		logger.WarnN(nil, "request", content)
	}
}
