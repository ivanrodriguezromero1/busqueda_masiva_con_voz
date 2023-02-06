package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
)

var staticDir string
var serverPort string

func main() {
	staticDir := "dist"
	router := chi.NewRouter()
	router.Get("/", func(resp http.ResponseWriter, req *http.Request) {
		http.ServeFile(resp, req, staticDir+"/index.html")
	})
	FileServer(router, "/", http.Dir(staticDir))
	// serverPort = os.Getenv("PORT")
	// if len(serverPort) == 0 {
	// 	serverPort = "3000"
	// }
	serverPort = os.Args[2]
	fmt.Println("Comenzando a escuchar al servidor en puerto ", serverPort)
	fmt.Println("Sirviendo contenido estático desde ", staticDir)
	fmt.Println("Enlace: http://localhost:" + serverPort + " ")
	http.ListenAndServe(":"+serverPort, router)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer no permite parámetros de URL")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fs.ServeHTTP(resp, req)
	}))
}
