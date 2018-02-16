package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"errors"
	"os"
	"fmt"
)

type testingUserError string

func (u testingUserError) Error() string {
	return u.Message()
}

func (u testingUserError) Message() string {
	return string(u)
}

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h appHandler
	code int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// 测试ErrWrapper函数（单元测试）
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errorWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.zshanjun.com",
			nil,
		)
		f(response, request)

		verifyRespnse(response.Result(), tt.code, tt.message, t)
	}
}

// 测试整个服务
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errorWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		verifyRespnse(response, tt.code, tt.message, t)
	}
}

func verifyRespnse(response *http.Response, expectedCode int, expectedMessage string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectedCode || body != expectedMessage {
		t.Errorf("expect (%d, %s); got (%d, %s)", expectedCode, expectedMessage, response.StatusCode, body)
	}
}