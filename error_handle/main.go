package main

import (
	"net/http"
	"zshanjun/go-exercise/error_handle/listing"
	"os"
	"log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("fatal error: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			//处理用户错误
			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errorWrapper(listing.ListingHandler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
