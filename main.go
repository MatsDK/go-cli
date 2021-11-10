package main

import(
	"fmt"
	"os"
	"log"
	"flag"
	D "./lib"
)

type Preset struct {
	name string `json:"name"`
	mode string `json:"mode"`
	brightness int `json:"brightness"`
	red int `json:"red"`
	green int `json:"green"`
	blue int `json:"blue"`

}

func main() {
	if len(os.Args) <  2 {
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

					if len(preset.name) != 0 {
						if preset.mode == "static" {
							setStatic(preset.brightness, preset.red, preset.green, preset.blue)
						}
					}
				}


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

				setStatic(*brightnessFlag, *rFlag, *gFlag, *bFlag)
			}
		case "preset":
			if len(os.Args) <= 2 {
				fmt.Println("list | add | set | delete")
				return
			}

			presetCmd := os.Args[2]
			handlePreset(presetCmd)
		case "on":
			fmt.Println("on")
		case "off":
			fmt.Println("off")
		default:
			fmt.Println("off | on | set")

	}
}

func setStatic(brightness int, red int, green int, blue int) {
	fmt.Printf("set brightness: %d, red: %d, green: %d, blue: %d\n", brightness, red, green, blue)
}

func getPresetByName(presetName string) Preset {
	query := fmt.Sprintf("SELECT name, mode, brightness, red, green, blue FROM preset WHERE name='%s';", presetName)
	rows := D.Query(query, D.ConnectDB())

	rows.Next() 
	var name, mode string
	var brightness, red, green, blue int
	rows.Scan(&name, &mode, &brightness, &red, &green, &blue)

	return Preset{name: name, mode: mode, brightness: brightness, red: red, green: green, blue:blue}
}

func handlePreset(cmd string) {
	db := D.ConnectDB()

	switch cmd {
		case "list":
			rows, err := db.Query("SELECT * FROM preset;")
			if err != nil {
				panic(err)
			}

			for rows.Next() {
				var name string
				rows.Scan(&name)

				fmt.Println(name)
			}
		default:
			fmt.Println("list | add | set| delete")
	}
}
