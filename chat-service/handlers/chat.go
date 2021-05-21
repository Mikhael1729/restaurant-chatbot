package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Chat struct {
	logger *log.Logger
}

func NewChat(logger *log.Logger) *Chat {
	return &Chat{logger}
}

func (chat *Chat) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	chat.logger.Println("All messages have been retrieved")
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Lo sentimos, hubo un inconveniente con nuestros servidores.", http.StatusBadRequest)
	}

	// Return the same provided data.
	fmt.Fprintf(rw, "Data %s\n", data)
}
