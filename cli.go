package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	D "./lib"
)

type Preset struct {
	Name       string `json:"name"`
	Mode       string `json:"mode"`
	Brightness int    `json:"brightness"`
	Red        int    `json:"red"`
	Green      int    `json:"green"`
	Blue       int    `json:"blue"`
}

func main() {
	c := D.Controller{}

	if len(os.Args) < 2 {
		fmt.Println("off | on | set | preset")
		return
	}

	switch os.Args[1] {
	case "set":
		if os.Args[2] == "preset" {
			fmt.Println("set preset")
			presetCmd := flag.NewFlagSet("preset", flag.ExitOnError)
			nameFlag := presetCmd.String("name", "", "the name of the preset")
			err := presetCmd.Parse(os.Args[3:])

			if err != nil {
				log.Fatal(err.Error())
			}

			if len(*nameFlag) != 0 {
				preset := getPresetByName(*nameFlag)

				if len(preset.Name) != 0 {
					if preset.Mode == "static" {
						c.SetStaticColor(preset.Brightness, preset.Red,
							preset.Green, preset.Blue)
					}
				}
			}

		} else if os.Args[2] == "brightness" {
			if len(os.Args) < 4 {
				log.Fatal("You have to provide a brightness value")
			}

			brightness, err := strconv.Atoi(os.Args[3])
			if err != nil {
				log.Fatal(err.Error())
				os.Exit(2)
			}

			c.SetBrightness(&brightness)
		} else {
			setCmd := flag.NewFlagSet("set", flag.ExitOnError)
			brightnessFlag := setCmd.Int("brightness", 100, "the brightness")
			rFlag := setCmd.Int("r", -1, "the red color")
			gFlag := setCmd.Int("g", -1, "the green color")
			bFlag := setCmd.Int("b", -1, "the blue color")
			err := setCmd.Parse(os.Args[2:])

			if err != nil {
				log.Fatal(err.Error())
			}

			if *rFlag > 255 {
				*rFlag = 255
			}
			if *gFlag > 255 {
				*gFlag = 255
			}
			if *bFlag > 255 {
				*bFlag = 255
			}

			c.SetStaticColor(*brightnessFlag, *rFlag, *gFlag, *bFlag)
		}
	case "preset":
		if len(os.Args) <= 2 {
			fmt.Println("list | add | delete")
			return
		}

		presetCmd := os.Args[2]
		handlePreset(presetCmd)
	case "on":
		c.SetOff()
	case "off":
		fmt.Println("off")
	default:
		fmt.Println("off | on | set")
	}
}

func getPresetByName(presetName string) Preset {
	query := fmt.Sprintf("SELECT name, mode, brightness, red, green, blue FROM preset WHERE name='%s';", presetName)
	rows, err := D.Query(query, D.ConnectDB())

	if err != nil {
		log.Fatal(err)
	}

	var name, mode string
	var brightness, red, green, blue int

	rows.Next()
	rows.Scan(&name, &mode, &brightness, &red, &green, &blue)

	return Preset{Name: name, Mode: mode, Brightness: brightness, Red: red, Green: green, Blue: blue}
}

func handlePreset(cmd string) {
	switch cmd {
	case "list":
		rows, err := D.Query("SELECT name FROM preset;", D.ConnectDB())

		if err != nil {
			log.Fatal(err)
			return
		}

		for rows.Next() {
			var name string
			rows.Scan(&name)

			fmt.Println(name)
		}
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		nameFlag := addCmd.String("name", "", "the name of the preset")
		modeFlag := addCmd.String("mode", "", "the mode of the preset")
		brightnessFlag := addCmd.Int("brightness", 100, "the brightness")
		rFlag := addCmd.Int("r", 0, "the red color")
		gFlag := addCmd.Int("g", 0, "the green color")
		bFlag := addCmd.Int("b", 0, "the blue color")
		err := addCmd.Parse(os.Args[3:])

		if err != nil {
			log.Fatal(err)
		}

		if len(*nameFlag) == 0 || len(*modeFlag) == 0 {
			fmt.Println("You have to specify: name, mode, brightness, r, g, b")
			return
		}

		if *modeFlag == "static" {
			query := fmt.Sprintf("INSERT INTO preset(name, mode, brightness, red, green, blue) VALUES('%s', '%s', %d, %d, %d, %d);",
				*nameFlag, *modeFlag, *brightnessFlag, *rFlag, *gFlag, *bFlag)
			D.Query(query, D.ConnectDB())
		}

	case "delete":
		removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
		nameFlag := removeCmd.String("name", "", "the name of the preset to delete")
		err := removeCmd.Parse(os.Args[3:])

		if err != nil {
			log.Fatal(err)
		}

		if len(*nameFlag) == 0 {
			log.Fatal("You have to provide a name of a preset")
		}

		query := fmt.Sprintf("DELETE FROM preset WHERE name='%s';", *nameFlag)
		D.Query(query, D.ConnectDB())

	default:
		fmt.Println("list | add | delete")
	}
}
