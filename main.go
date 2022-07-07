package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// package main
var (
	cfg config
)

// config
type config struct {
	Url   string   `yaml:"url"`
	Files []string `yaml:"files"`
}

// init
func init() {
	viper.SetConfigFile("./etc/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("read config error")
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Println("unmarshal error")
	}
}

// main
func main() {
	e := gin.Default()

	e.GET("picture", func(c *gin.Context) {
		c.JSON(http.StatusOK, cfg)
	})

	e.Run(":8071")
}
