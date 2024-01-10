package main

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/gin-contrib/gzip"
)

type requestLog struct {
	ClientIP     string        `json:"client_ip"`
	TimeStamp    time.Time     `json:"time_stamp"`
	Method       string        `json:"method"`
	Path         string        `json:"path"`
	RequestProto string        `json:"request_proto"`
	StatusCode   int           `json:"status_code"`
	Latency      time.Duration `json:"latency"`
	UserAgent    string        `json:"user_agent"`
	ErrorMessage string        `json:"error_message"`
	RequestId    string        `json:"request_id"`
	Body         string        `json:"body"`
}

func setupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	f, _ := os.OpenFile("gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// log request body
		var body []byte
		if param.Request.Body != nil {
			body, _ = io.ReadAll(param.Request.Body)
			param.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		reqLog := requestLog{
			ClientIP:     param.ClientIP,
			TimeStamp:    param.TimeStamp,
			Method:       param.Method,
			Path:         param.Path,
			RequestProto: param.Request.Proto,
			StatusCode:   param.StatusCode,
			Latency:      param.Latency / 1e6,
			UserAgent:    param.Request.UserAgent(),
			ErrorMessage: param.ErrorMessage,
			RequestId:    param.Request.Header.Get("X-Request-Id"),
			Body:         string(body),
		}
		// to json
		marshal, _ := json.Marshal(reqLog)
		return string(marshal) + "\n"
	}))

	router.Use(gin.Recovery())

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	store := persistence.NewInMemoryStore(time.Second)
	router.Use(
		requestid.New(
			requestid.WithCustomHeaderStrKey("X-Request-Id"),
		),
	)

	router.Any("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Cached Page
	router.Any("/cache_ping", cache.CachePage(store, time.Second*5, func(c *gin.Context) {
		data := map[string]interface{}{
			"message": "pong 咚咚咚",
			"tag":     "<br>",
			"date":    time.Now().Format(time.DateTime),
		}

		c.PureJSON(http.StatusOK, data)
	}))
	return router
}

func main() {
	router := setupRouter()

	srv := &http.Server{
		Addr:           "0.0.0.0:9090",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 3 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
