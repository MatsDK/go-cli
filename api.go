package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	var presets []Preset

	query := fmt.Sprintf("SELECT name, mode, brightness, red, green, blue FROM preset;")
	rows, err := D.Query(query, D.ConnectDB())

	if err != nil {
		json.NewEncoder(w).Encode(nil)
		return
	}

	for rows.Next() {
		var name, mode string
		var brightness, red, green, blue int
		rows.Scan(&name, &mode, &brightness, &red, &green, &blue)

		presets = append(presets, Preset{Name: name, Mode: mode, Brightness: brightness, Red: red, Green: green, Blue: blue})
	}

	json.NewEncoder(w).Encode(presets)
}

func getPreset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	presetName := params["name"]

	query := fmt.Sprintf("SELECT name, mode, brightness, red, green, blue FROM preset WHERE name='%s';", presetName)
	rows, err := D.Query(query, D.ConnectDB())

	if err != nil {
		json.NewEncoder(w).Encode(nil)
		return
	}

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

func createPreset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p Preset

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf("INSERT INTO preset(name, mode, brightness, red, green, blue) VALUES('%s', '%s', %d, %d, %d, %d) RETURNING name, mode, brightness, red, green, blue; ",
		p.Name, p.Mode, p.Brightness, p.Red, p.Green, p.Blue,
	)
	rows, err := D.Query(query, D.ConnectDB())

	if err != nil {
		json.NewEncoder(w).Encode(nil)
		return
	}

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

func deletePreset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	presetName := params["name"]

	query := fmt.Sprintf("DELETE FROM preset WHERE name='%s'",
		presetName,
	)
	if _, err := D.Query(query, D.ConnectDB()); err != nil {
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewEncoder(w).Encode(true)
}

func setPreset(c *D.Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		presetName := params["name"]

		query := fmt.Sprintf("SELECT name, mode, brightness, red, green, blue FROM preset WHERE name='%s'",
			presetName,
		)
		rows, err := D.Query(query, D.ConnectDB())

		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
			return
		}

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

		if mode == "static" {
			c.SetStaticColor(brightness, red, green, blue)
		}

		json.NewEncoder(w).Encode(true)
	}
}

func setBrigthness(c *D.Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		brightness, err := strconv.Atoi(params["value"])
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		c.SetBrightness(&brightness)

		json.NewEncoder(w).Encode(true)
	}
}

func main() {
	r := mux.NewRouter()
	c := D.Controller{}

	r.HandleFunc("/api/presets", getPresets).Methods("GET")
	r.HandleFunc("/api/presets/{name}", getPreset).Methods("GET")
	r.HandleFunc("/api/presets", createPreset).Methods("POST")
	r.HandleFunc("/api/presets/{name}", deletePreset).Methods("DELETE")

	r.HandleFunc("/api/set/preset/{name}", setPreset(&c)).Methods("POST")
	r.HandleFunc("/api/set/brightness/{value}", setBrigthness(&c)).Methods("POST")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
