package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/yeboahnanaosei/edith"
	"google.golang.org/grpc"
)

var serverRoot = os.Getenv("EDITHD_SERVER_ROOT")
var workingDir = os.Getenv("EDITHD_WORKING_DIR")
var port = os.Getenv("EDITHD_PORT")
var configFile = filepath.Join(serverRoot, "config.json")

func init() {
	if serverRoot == "" {
		fmt.Fprintf(os.Stderr, "edithd: enviroment variable [EDITHD_SERVER_ROOT] not set\n")
		os.Exit(1)
	}

	if workingDir == "" {
		fmt.Fprintf(os.Stderr, "edithd: enviroment variable [EDITHD_WORKING_DIR] not set\n")
		os.Exit(1)
	}

	// Prepare log file
	logfilePath := filepath.Join(serverRoot, "edithd.log")
	logfile, err := os.OpenFile(logfilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "edithd: failed to load log file: %s\n", err)
		os.Exit(1)
	}
	// defer logfile.Close()
	log.SetOutput(logfile)

	// Make sure config file exists
	_, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		log.Printf("config file `config.json` does not exist in server root [%s]\n", serverRoot)
		fmt.Fprintf(os.Stderr, "edithd: config file [config.json] does not exist in server root [%s]\n", serverRoot)
		os.Exit(1)
	} else if err != nil {
		log.Println("failed to stat config file: ", err)
		fmt.Fprintf(os.Stderr, "edithd: failed to stat config file: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	if port == "" {
		port = "54920"
		log.Println("environment variable EDITHD_PORT not set defaulting to ", port)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Printf("failed to create listener: %s\n", err)
		fmt.Fprintf(os.Stderr, "edithd: failed to create listener: %s\n", err)
		os.Exit(1)
	}

	grpcServer, edithServer := grpc.NewServer(), &service{}
	edith.RegisterEdithServer(grpcServer, edithServer)

	fmt.Println("edithd: listening on port ", port)
	grpcServer.Serve(listener)
}
