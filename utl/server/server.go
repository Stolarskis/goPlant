package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/stolarskis/goPlant/utl/config"

	log "github.com/inconshreveable/log15"
	"github.com/stolarskis/goPlant/pkg/transport"
	"goji.io"
	"goji.io/pat"
)

func NewHTTP(m *goji.Mux) {

	h := transport.Handler{}

	m.HandleFunc(pat.Post("/data/moisture"), h.MoistureData)
	m.HandleFunc(pat.Post("/data/soilTemp"), h.SoilTempData)
	m.HandleFunc(pat.Post("/data/airTemp"), h.AirTempData)
	m.HandleFunc(pat.Post("/data/humidity"), h.HumidityData)
	m.HandleFunc(pat.Post("/data/light"), h.LightData)
}

func Start(configPath string, m *goji.Mux) {

	_, svcConf, err := config.GetDbSettings(configPath)
	if err != nil {
		log.Crit(err.Error())
	}

	if (svcConf == config.SvcConfig{}) {
		log.Crit("Server settings have not been initialized.")
		os.Exit(1)
	}

	log.Debug("Server Host: " + svcConf.Host + " Server Port: " + svcConf.Port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", svcConf.Host, svcConf.Port), m)
	if err != nil {
		log.Crit("Failed to start server: " + err.Error())
	}
	log.Debug("Server Started Successfully")
}
