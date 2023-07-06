package main

import "net/http"

type indexHandler struct{}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index!"))
}

type aboutHandler struct{}

func (ah *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func main() {
	ih := indexHandler{}
	ah := aboutHandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil, // DefaultServeMux
		//Handler: http.FileServer(http.Dir("wwwroot")), // http.FileServer
	}
	// 处理多个Handler
	http.Handle("/index", &ih)
	http.Handle("/about", &ah)
	http.Handle("/welcome", http.HandlerFunc(welcome))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	server.ListenAndServe()
}
