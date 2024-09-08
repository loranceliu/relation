package main

import (
	"flag"
	"fmt"
	"gin/config"
	"gin/model"
	"gin/routes"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var release string

func init() {
	flag.StringVar(&release, "release", "local", "release model, optional local/dev/prod")
}

func main() {
	flag.Parse()

	var conf *config.Config
	var err error

	switch release {
	case "local":
		conf, err = loadConfig("etc/relation-local.yaml")
	case "dev":
		conf, err = loadConfig("etc/relation-dev.yaml")
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		conf, err = loadConfig("etc/relation-prod.yaml")
	}

	if err != nil {
		log.Fatal(err)
	}

	model.InitDatabase(&conf.Mysql)

	e := gin.Default()

	routes.Setup(e)

	s := fmt.Sprintf("%s:%v", conf.Server.Host, conf.Server.Port)
	err = e.Run(s)

	if err != nil {
		return
	}
}

func loadConfig(filePath string) (*config.Config, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var conf config.Config
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Server Name: %s\n", conf.Server.Name)
	fmt.Printf("Server Host: %s\n", conf.Server.Host)
	fmt.Printf("Server Port: %d\n", conf.Server.Port)

	return &conf, nil
}
