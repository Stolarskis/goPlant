package server

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/inconshreveable/log15"
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
		log.Crit("Server settings have not been initialized.")
		os.Exit(1)
	}

	log.Debug("Server Host: " + ServerStg.Host + " Server Port: " + ServerStg.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", ServerStg.Host, ServerStg.Port), m)
	if err != nil {
		log.Crit("Failed to start server: " + err.Error())
	}
	log.Debug("Server Started Successfully")
}
