package ext

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	skipURLPrefixes = []string{
		"/api",
		"/static",
		"/favicon",
	}

	logger *logrus.Logger
	lock   sync.RWMutex
)

var shutDownFuncs []func()

func Shutdown() {
	for _, fn := range shutDownFuncs {
		fn()
	}
}

// 日志provider
func InitLogger() {
	fi, err := os.OpenFile("http.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	var closer io.Closer
	closer = fi
	shutDownFuncs = append(shutDownFuncs, func() {
		closer.Close()
	})

	logger = logrus.New()
	logger.Out = fi
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	go func() {
		tk := time.NewTicker(8 * 86400 * time.Second) // 8天生成一个日志文件.
		for range tk.C {
			lock.Lock()
			fi.Close()
			os.Rename("http.log", "http.log.before")
			fi, _ = os.OpenFile("http.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
			closer = fi
			logger.SetOutput(fi)
			lock.Unlock()
		}
	}()
}

func Logger() gin.HandlerFunc {
	// 生产环境下，打印req body 但是不打印 res body
	return func(c *gin.Context) {
		lock.RLock()
		defer lock.RUnlock()

		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()

		for _, v := range skipURLPrefixes {
			if strings.HasPrefix(c.Request.RequestURI, v) {
				return
			}
		}

		// 结束时间
		end := time.Now()

		// 执行时间
		rt := end.Sub(start) / time.Millisecond
		// 状态码
		code := c.Writer.Status()

		// 请求IP
		ip := ClientIP(c)

		refer := c.Request.Header.Get("referer")

		logger.WithFields(logrus.Fields{
			"status":   code,
			"method":   c.Request.Method,
			"url":      c.Request.URL.String(),
			"ip":       ip,
			"ua":       c.Request.UserAgent(),
			"duration": rt,
			"refer":    refer,
		}).Info()
	}
}
