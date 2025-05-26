//go:generate go tool wit-bindgen-go generate --world endpoint --out gen ./wit
package main

import (
	"encoding/json"
	"net/http"

	"github.com/jamesstocktonj1/wasmcloud-template/components/endpoint/gen/jamesstocktonj1/wasmcloud-template/connect"
	"go.wasmcloud.dev/component/net/wasihttp"
)

func init() {
	wasihttp.HandleFunc(handleRequest)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	res := connect.Greet(connect.Request{
		Message: "Hello",
		Name:    "World",
	})

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(map[string]string{
		"message": res,
	})
	w.WriteHeader(http.StatusOK)
}

func main() {}
