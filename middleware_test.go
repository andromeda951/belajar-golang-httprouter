package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

type Middleware struct {
	http.Handler // nama field sama dengan nama tipe datanya yaitu Handler
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive request")
	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleware := &Middleware{
		Handler: router,
	}
	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	bytes, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(bytes))
}
