package main

import (
	"github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var log = logrus.New()

func init() {
	appId := os.Getenv("GAE_DEPLOYMENT_ID")
	if appId != "" {
		// On app engine
		log.Formatter = stackdriver.NewFormatter(
			stackdriver.WithService(os.Getenv("GAE_SERVICE")),
			stackdriver.WithVersion(os.Getenv("GAE_VERSION")),
		)
	} else {
		// Local
		log.Formatter = &logrus.TextFormatter{ForceColors: true}
	}
	log.Info("Logging initialized")
}

func httpLogger(logger *logrus.Logger, r *http.Request) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"httpRequest": map[string]interface{}{
			"method":    r.Method,
			"url":       r.URL.String(),
			"userAgent": r.Header.Get("User-Agent"),
			"referrer":  r.Header.Get("Referer"),
		},
	})
}
