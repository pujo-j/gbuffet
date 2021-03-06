package main

import (
	"database/sql"
	"encoding/json"
	"github.com/pujo-j/gogae"
	"net/http"
	"os"
)

var db *sql.DB

type Config struct {
	BillingAccount string
	gogae.GogaeConfig
}

var config = Config{}

func main() {
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
	jsonParser := json.NewDecoder(configFile)

	err = jsonParser.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	_ = configFile.Close()
	g, err := gogae.InitGogae(config.GogaeConfig, MakeUserDataFunction(config.AuthConfig.Project))

	RouteUtils(g)
	RouteProjectRequests(g)

	g.Router.ServeFiles("/openapi/*filepath", http.Dir("./openapi"))
	db = g.Db
	g.Start()
}
