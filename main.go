package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/andreylm/process-manager/models/task"
	ServerFactory "github.com/andreylm/process-manager/server"
	"github.com/joho/godotenv"
)

var host string
var port int

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	e := godotenv.Load(exPath + "/configs/main.conf") //Load .env file
	// e := godotenv.Load("./configs/main.conf") //Load .env file
	host = os.Getenv("host")

	rawPort := os.Getenv("port")

	if len(rawPort) > 0 {
		var err error
		port, err = strconv.Atoi(rawPort)
		if err != nil {
			panic(err)
		}
	} else {
		port = 8000
	}

	if e != nil {
		fmt.Print(e)
	}

}

func main() {
	taskCollection := task.CreateCollection()

	server := ServerFactory.NewServer(host, port, taskCollection)
	server.Start()
}
