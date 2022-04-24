package httpServer

import (
	"net/http"
)

func RunServer() {
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
