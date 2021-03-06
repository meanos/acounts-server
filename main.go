package main

import (
	"encoding/json"
	"meanos.io/AccountsServer/auth"
	"meanos.io/AccountsServer/db"
	"meanos.io/AccountsServer/server"
	"fmt"
	"github.com/meanOs/AMS"
	beatrix "github.com/meanOs/Beatrix"
	tm "github.com/meanOs/TokenMaster"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Config struct {
	BeatrixToken   string `json:"beatrix-token"`
	BeatrixChannel string `json:"beatrix-channelID"`
	MongoURI       string `json:"mongo-uri"`
	SQLURI         string `json:"sql-uri"`
	WebsiteURL     string `json:"website_url"`
}

func LoadConfig() Config {
	var config Config
	configFile, err := os.Open("ams-credentials.json")
	if err != nil {
		log.Panic(err)
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	fmt.Println(config)
	err = configFile.Close()
	if err != nil {
		log.Panic(err)
	}
	return config
}

func Init(cx Config, r *gin.Engine) {
	beatrix.Init("AccountsServer", cx.BeatrixToken, cx.BeatrixChannel)
	AMS.Init(cx.MongoURI, cx.SQLURI)
	tm.Init(cx.MongoURI)
	db.Init(cx.MongoURI)
	auth.Init()
	server.Init(r, cx.WebsiteURL)
}

func main() {
	r := gin.Default()
	cx := LoadConfig()

	r.LoadHTMLGlob("static/*.html")
	r.Static("/assets", "./static/assets")

	Init(cx, r)

	r.Run(":80")
}
