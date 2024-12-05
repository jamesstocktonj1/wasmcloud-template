//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world greeting --out gen ./wit
package main

import (
	"fmt"

	"github.com/jamesstocktonj1/wasmcloud-template/components/greeting/gen/jamesstocktonj1/wasmcloud-template/connect"
	"go.wasmcloud.dev/component/wasmcloud"
)

func init() {
	connect.Exports.Greet = greetHandler
}

func greetHandler(req connect.Request) string {
	format := wasmcloud.GetConfigOrDefault("format", "%s, %s!")
	return fmt.Sprintf(format, req.Message, req.Name)
}

func main() {}
