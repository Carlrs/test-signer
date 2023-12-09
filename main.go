package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"reflect"
	"strconv"
	"test-signer/controller"
	"test-signer/entity"
	"test-signer/util"
)

const DefaultPort = 8080

func main() {
	r := gin.Default()
	sr, err := util.SharedResources{}.Generate()
	if err != nil {
		fmt.Printf("Failed to initialize shared resources. Error: %s\n", err.Error())
		return
	}
	var controllers []controller.Controller
	controllers = append(controllers, controller.SignatureController{})
	for _, c := range controllers {
		c.Init(sr)
		c.Register(r)
	}
	var entities []interface{}
	entities = append(entities, entity.Test{}, entity.Signature{}, entity.QuestionAnswer{})
	for _, e := range entities {
		if err = sr.Db.AutoMigrate(e); err != nil {
			fmt.Printf("Failed to automigratemigrate table for entity %s, error: %s", reflect.TypeOf(e), err.Error())
			return
		}
	}

	port, err := strconv.Atoi(os.Getenv("SIGNER_PORT"))
	if err != nil || port < 1 || port > 65535 {
		port = DefaultPort
	}
	portString := fmt.Sprintf(":%d", port)

	err = r.Run(portString)
	if err != nil {
		fmt.Printf("Error running gin server: %s", err.Error())
	}
}
