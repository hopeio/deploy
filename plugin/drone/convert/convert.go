package main

import (
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/converter"
	httpi "github.com/hopeio/gox/net/http"
	"github.com/hopeio/gox/net/http/binding"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req converter.Request
		binding.Bind(r, &req)
		// TODO:
		httpi.RespSuccessData(w, &drone.Config{
			Data: "",
		})
	}))
}
