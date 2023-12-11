package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"strconv"
	"test-signer/controller"
	"test-signer/entity"
	"test-signer/util"
)

const DefaultPort = 8080

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load dotenv")
		return
	}
	r := gin.Default()
	sr := &util.SharedResources{}
	err := sr.Generate()
	if err != nil {
		fmt.Printf("Failed to initialize shared resources. Error: %s\n", err.Error())
		return
	}
	var entities []interface{}
	entities = append(entities, entity.Test{}, entity.Signature{}, entity.QuestionAnswer{})
	for _, e := range entities {
		if err = sr.Db.AutoMigrate(e); err != nil {
			fmt.Printf("Failed to automigrate table for entity %s, error: %s", reflect.TypeOf(e), err.Error())
			return
		}
	}

	var controllers []controller.Controller
	controllers = append(controllers, controller.SignatureController{})
	for _, c := range controllers {
		c.Init(sr)
		c.Register(r)
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
