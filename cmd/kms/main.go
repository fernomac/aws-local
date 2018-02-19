package main

import (
	"net/http"

	"github.com/fernomac/aws-local/pkg/kms"
)

func main() {
	http.ListenAndServe("localhost:8080", kms.NewHandler(kms.New()))
}
