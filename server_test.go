package golang_web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello World")
		if err != nil {
			panic(err)
		}
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/dashboard", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Dashboard")
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)

	})

	mux.HandleFunc("/profile", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Profile")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Image")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/images/thumbnails", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "thumbnails")
		if err != nil {
			panic(err)
		}
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Rafel")
}

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello-rafel", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)
}
