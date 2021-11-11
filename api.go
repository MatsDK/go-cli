package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	D "./lib"
	"github.com/gorilla/mux"
)

type Preset struct {
	Name       string `json:"name"`
	Mode       string `json:"mode"`
	Brightness int    `json:"brightness"`
	Red        int    `json:"red"`
	Green      int    `json:"green"`
	Blue       int    `json:"blue"`
}

func getPresets(w http.ResponseWriter, r *http.Request) {
	var presets []Preset

	query := fmt.Sprintf("SELECT name, mode, brightness, red, green, blue FROM preset;")
	rows := D.Query(query, D.ConnectDB())

	for rows.Next() {
		var name, mode string
		var brightness, red, green, blue int
		rows.Scan(&name, &mode, &brightness, &red, &green, &blue)

		presets = append(presets, Preset{Name: name, Mode: mode, Brightness: brightness, Red: red, Green: green, Blue: blue})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(presets)
}

func getPreset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	presetName := params["name"]

	query := fmt.Sprintf("SELECT name, mode, brightness, red, green, blue FROM preset WHERE name='%s';", presetName)
	rows := D.Query(query, D.ConnectDB())

	var name, mode string
	var brightness, red, green, blue, counter int
	for rows.Next() {
		rows.Scan(&name, &mode, &brightness, &red, &green, &blue)
		counter++
	}

	if counter == 0 {
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(Preset{Name: name, Mode: mode, Brightness: brightness, Red: red, Green: green, Blue: blue})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/presets", getPresets).Methods("GET")
	r.HandleFunc("/api/presets/{name}", getPreset).Methods("GET")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
