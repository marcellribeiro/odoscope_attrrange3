package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/marcellribeiro/odoscope_attrrange3/calculator/delivery"
	"github.com/marcellribeiro/odoscope_attrrange3/entities"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.New()
	var config entities.Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return
	}
	delivery.NewCalculatorHandler(r, config)

	r.Run()
}
