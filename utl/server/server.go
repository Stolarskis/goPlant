package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/stolarskis/goPlant/pkg/transport"
	"goji.io"
	"goji.io/pat"
)

func NewHTTP(m *goji.Mux) {

	h := transport.Handler{}

	m.HandleFunc(pat.Post("/data/moisture"), h.MoistureData)
	// m.HandleFunc(pat.Post("/data/temperature"))
	// m.HandleFunc(pat.Post("/data/light"))
}

func Start(m *goji.Mux) error {
	h := os.Getenv("GOPLANT_HOST")
	p := os.Getenv("GOPLANT_PORT")
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", h, p), m)
	if err != nil {
		return err
	}

	return nil
}
