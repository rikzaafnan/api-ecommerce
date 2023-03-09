package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const visitor = "visitor"

var Reset = "\033[0m"
var Green = "\033[32m"

func LoggingMiddlewareCustom() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request
		t1 := time.Now()
		c.Next()
		t2 := time.Now()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)

		ss := Green + c.Request.Method + Reset
		// s := fmt.Sprintf(" %q (%s) %v %s", r.URL.String(), email, t2.Sub(t1), truncateString(r.Header.Get("User-Agent"), 15))
		s := fmt.Sprintf(" %q (%s) %s %s, status: %d", c.Request.URL.String(), "email", t2.Sub(t1), c.Request.Header.Get("User-Agent"), status)

		if status >= 299 {
			log.Error(ss + s)
		} else {
			log.Info(ss + s)
		}

	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request
		t1 := time.Now().UTC().Local()
		c.Next()
		t2 := time.Now().UTC().Local()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)

		ss := Green + c.Request.Method + Reset
		// s := fmt.Sprintf(" %q (%s) %v %s", r.URL.String(), email, t2.Sub(t1), truncateString(r.Header.Get("User-Agent"), 15))
		s := fmt.Sprintf(" %q (%s) %s %s, status: %d, ip address : %s", c.Request.URL.String(), "email", t2.Sub(t1), c.Request.Header.Get("User-Agent"), status, c.Request.RemoteAddr)

		// log.Info(fmt.Sprintf(" %q (%s) %s %s, status: %d", c.Request.URL.String(), "email", t2.Sub(t1), c.Request.Header.Get("User-Agent"), status))
		// log.Info(ss)

		if status >= 299 {
			log.Error(ss + s)
			log.Info(c.Errors.Errors())
			log.Info(c.Errors.String())
		} else {
			log.Info(ss + s)
		}

	}
}

func LoggingMiddlewareCustomTesting() gin.HandlerFunc {
	return func(c *gin.Context) {

		makeLogEntry(c).Info("incoming request")
		c.Next()

	}
}

func makeLogEntry(c *gin.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request.Method,
		"uri":    c.Request.URL,
		"ip":     c.Request.RemoteAddr,
	})
}

func errorHandler(err error, c *gin.Context) {
	// report, ok := err.(*gin.Error)
	// if ok {
	// 	report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	// } else {
	// 	report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	// makeLogEntry(c).Error(report.Message)
	// c.HTML(report.Code, report.Message.(string))
}
