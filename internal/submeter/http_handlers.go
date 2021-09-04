package submeter

import (
	"fmt"
	"log"
	"net/http"
)

func Submit(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "mocked submission response")
	if err != nil {
		log.Fatal("error while trying to write response for submission")
	}
}
