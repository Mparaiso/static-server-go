package main

// go-static-server serves files on port 8080
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type HandlerWithLogger struct {
	http.Handler
	*log.Logger
}

// NewHandlerWithLogger returns an handler with a logger
func NewHandlerWithLogger(handler http.Handler, logger *log.Logger) http.Handler {
	return &HandlerWithLogger{handler, logger}
}

// ServeHTTP handles and http request
func (h *HandlerWithLogger) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	h.Logger.Println(request.URL.String())
	h.Handler.ServeHTTP(responseWriter, request)
}

func main() {
	ip := flag.String("ip", "localhost", "ip")
	port := flag.String("port", "8080", "port")
	flag.Parse()
	address := fmt.Sprintf("%s:%s", *ip, *port)
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Error Getting current directory", err)
	}
	server := http.NewServeMux()
	server.Handle("/", NewHandlerWithLogger(http.FileServer(http.Dir(currentDirectory)), log.New(os.Stdout, "log: ", 0)))
	log.Print("Listening on ", address, ".")
	log.Fatal("Error serving files.", http.ListenAndServe(address, server))
}
