package handlers

import (
	"fmt"
	"net/http"
)

type ProjectedHandler struct {
}

func (handler *ProjectedHandler) ServeHTTP(writer http.ResponseWriter, httpRequest *http.Request) {
	fmt.Println(httpRequest)
}
