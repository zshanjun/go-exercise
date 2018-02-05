package listing

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

type userError string

func (u userError) Error() string {
	return u.Message()
}

func (u userError) Message() string {
	return string(u)
}

func ListingHandler(writer http.ResponseWriter, request *http.Request) error {
	if index := strings.Index(request.URL.Path, "/error_handle/"); index != 0 {
		return userError("path must start with error_handle")
	}
	filename := request.URL.Path[len("/error_handle/"):]
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
