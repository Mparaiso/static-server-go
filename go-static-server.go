package gostaticserver

// go-static-server serves files on port 8080
import (
	"log"
	"net/http"
	"os"
)

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Error Getting current directory", err)
	}
	server := http.NewServeMux()
	server.Handle("/", http.FileServer(http.Dir(currentDirectory)))
	log.Fatal("Error serving files.", http.ListenAndServe(":8080", server))
}
