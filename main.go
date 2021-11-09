package main

import(
	"fmt"
	"os"
	"log"
	"flag"
	D "./lib"
)

func main() {
	if len(os.Args) <  2 {
		fmt.Println("off | on | set | preset")
		return
	}

	switch os.Args[1] {
		case "set":
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

			fmt.Printf("set brightness: %d, red: %d, green: %d, blue: %d\n", *brightnessFlag, *rFlag, *gFlag, *bFlag)
		case "preset":
			db := D.ConnectDB()	

			rows, err := db.Query("SELECT * FROM preset;")
			if err != nil {
				panic(err)
			}

			for rows.Next() {
				var name string
				rows.Scan(&name)

				fmt.Println(name)

			}

			fmt.Println("preset")
		case "on":
			fmt.Println("on")
		case "off":
			fmt.Println("off")
		default: 
			fmt.Println("off | on | set")

	}
}
