package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
		// panic("ups")
	})

	server := http.Server{
		Addr:    "localhost:3030", // kalua 3000 error 5xx
		Handler: router,
	}

	server.ListenAndServe()

}
