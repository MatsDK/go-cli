package lib

import "fmt"

type Controller struct {
}

func (c *Controller) SetBrightness(brightness *int) {
	fmt.Printf("set brightness: %d\n", *brightness)
}

func (c *Controller) SetStaticColor(brightness int, red int, green int, blue int) {
	fmt.Printf("set brightness: %d, red: %d, green: %d, blue: %d\n", brightness, red, green, blue)
}

func (c *Controller) SetOff() {
	fmt.Println("Turn off")
}
