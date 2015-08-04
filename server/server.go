package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pgrunde/back/parse"
)

type Settings struct {
	Domain    string `json:"domain"`
	Port      int64  `json:"port"`
	ProxyPort int64  `json:"proxy-port"`
	Cookie    Cookie `json:"cookie"`
}

type Cookie struct {
	Age      int64  `json:"age"`
	HttpOnly bool   `json:"http-only"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}

type Server struct {
	Settings
}

func New(s Settings) *Server {
	server := Server{Settings: s}
	http.HandleFunc("/shootingtracker", info)
	return &server
}

func (server Server) GetPort() string {
	return ":" + strconv.Itoa(int(server.Port))
}

func (server Server) ListenAndServe() error {
	log.Printf("Server starting on address %s\n", server.GetPort())
	return http.ListenAndServe(server.GetPort(), nil)
}

func info(w http.ResponseWriter, r *http.Request) {
	year := getYear(r)
	resp, err := http.Get("http://shootingtracker.com/w/index.php?title=Mass_Shootings_in_" + year + "&action=edit")
	if err != nil {
		fmt.Fprintf(w, "Dangole Internet: %s", err)
		return
	}
	defer resp.Body.Close()
	shootings := parse.BuildShootings(resp.Body)
	data, err := json.Marshal(shootings)
	if err != nil {
		fmt.Fprintf(w, "Dangole Internet: %s", err)
		return
	}
	w.Write(data)
}

func getYear(r *http.Request) string {
	values := r.URL.Query()
	if len(values["year"]) == 1 {
		if values["year"][0] == "2013" {
			return "2013"
		}
		if values["year"][0] == "2014" {
			return "2014"
		}
	}
	return "2015"
}
