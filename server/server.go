package server

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/andreylm/process-manager/models/task"
	RouterFactory "github.com/andreylm/process-manager/router"
	"github.com/gorilla/handlers"
)

// Server - this is the server object
type Server struct {
	Host       string
	Port       int
	Addr       string
	HTTPServer *http.Server
}

//Start - starts the server service
func (s *Server) Start() {
	log.Println("Server started on port", s.Port)
	log.Fatal(s.HTTPServer.ListenAndServe())
}

// NewServer - creates a new server
func NewServer(host string, port int, collection *task.Collection) *Server {
	var server Server
	server.Host = host
	server.Port = port
	server.Addr = host + ":" + strconv.Itoa(port)

	router := RouterFactory.NewRouter(collection)

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PUT", "OPTIONS", "PATCH"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(router.Router))

	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	server.HTTPServer = &http.Server{
		Addr:           server.Addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &server
}
