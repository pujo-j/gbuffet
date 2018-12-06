package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"time"
)

type Config struct {
	AuthConfig AuthConfig
	DbConfig   DbConfig
}

type DbConfig struct {
	ConnectionName string
	User           string
	Password       string
}

var db *sql.DB

func main() {

	router := httprouter.New()

	router.ServeFiles("/openapi/*filepath", http.Dir("./openapi"))
	var configFile *os.File
	var err error
	if os.Getenv("GAE_DEPLOYMENT_ID") == "" {
		// Dev env
		configFile, err = os.Open("config.dev.json")
		if err != nil {
			panic(err.Error())
		}
	} else {
		configFile, err = os.Open("config.json")
		if err != nil {
			panic(err.Error())
		}
	}

	config := Config{}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	_ = configFile.Close()
	auth := NewAuthMiddleware(router, config.AuthConfig, []string{
		"email", "profile", "openid",
		"https://www.googleapis.com/auth/admin.directory.group.readonly",
		"https://www.googleapis.com/auth/iam.test",
	}, MakeUserDataFunction(config.AuthConfig.Project))
	auth.AddPaths(router)
	var DSN string
	if os.Getenv("GAE_DEPLOYMENT_ID") != "" {
		DSN = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/control-tower?parseTime=true", config.DbConfig.User, config.DbConfig.Password, config.DbConfig.ConnectionName)
	} else {
		DSN = fmt.Sprintf("%s:%s@/control-tower?parseTime=true", config.DbConfig.User, config.DbConfig.Password)
	}
	db, err = sql.Open("mysql", DSN)
	if err != nil {
		log.WithField("DSN", DSN).WithError(err).Fatal("Connecting to Mysql")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.PingContext(ctx)
	if err != nil {
		log.WithField("DSN", DSN).WithError(err).Fatal("Connecting to Mysql")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	RouteProjectRequests(router, auth)
	RouteUtils(router, auth)

	log.WithField("port", port).Info("Listening")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), auth))
}
