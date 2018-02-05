package listing

import (
	"net/http"
	"os"
	"io/ioutil"
)

func ListingHandler(writer http.ResponseWriter, request *http.Request) error {
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
