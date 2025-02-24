package api

import (
	"fmt"
	"net/http"
)

func HelloFromServer(w http.ResponseWriter, h *http.Request) {
	fmt.Fprintf(w, "HELLO!")
}
