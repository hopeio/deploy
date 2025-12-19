package main

import (
	"net/http"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/converter"
	httpx "github.com/hopeio/gox/net/http"
	"github.com/hopeio/gox/net/http/binding"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req converter.Request
		binding.Bind(r, &req)
		// TODO:
		httpx.RespondSuccess(r.Context(), w, &drone.Config{
			Data: "",
		})
	}))
}
