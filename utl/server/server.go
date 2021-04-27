package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stolarskis/goPlant/pkg/transport"
	"goji.io"
	"goji.io/pat"
)

type ServerConfig struct {
	Host string
	Port string
}

var ServerStg ServerConfig

func NewHTTP(m *goji.Mux) {

	h := transport.Handler{}

	m.HandleFunc(pat.Post("/data/moisture"), h.MoistureData)
	// m.HandleFunc(pat.Post("/data/temperature"))
	// m.HandleFunc(pat.Post("/data/light"))
}

func Start(m *goji.Mux) {

	if (ServerStg == ServerConfig{}) {
		log.Fatal("Server settings have not been initialized.")
	}

	fmt.Println(ServerStg.Host, ServerStg.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", ServerStg.Host, ServerStg.Port), m)
	if err != nil {
		log.Fatal(err)
	}
}
