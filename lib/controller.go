package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const BASE_URL = "http://192.168.0.164:5000"

type Controller struct {
}

func (c *Controller) SetBrightness(brightness *int) {
	fmt.Printf("set brightness: %d\n", *brightness)
}

func (c *Controller) SetStaticColor(brightness int, red int, green int, blue int) {
	const postURL = BASE_URL + "/setStaticColor"
	fmt.Printf("set brightness: %d, red: %d, green: %d, blue: %d\n", brightness, red, green, blue)

	postBody, _ := json.Marshal(map[string]int{
		"r": red,
		"g": green,
		"b": blue,
	})

	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post(postURL, "application/json", responseBody)

	if err != nil {
		log.Fatal(err)
	}

}

func (c *Controller) SetOff() {
	fmt.Println("Turn off")
	const postURL = BASE_URL + "/off"

	postBody, _ := json.Marshal(map[string]int{})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post(postURL, "application/json", responseBody)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Controller) SetOn() {
	fmt.Println("Turn on")
	const postURL = BASE_URL + "/on"

	postBody, _ := json.Marshal(map[string]int{})
	responseBody := bytes.NewBuffer(postBody)
	_, err := http.Post(postURL, "application/json", responseBody)
	if err != nil {
		log.Fatal(err)
	}
}
