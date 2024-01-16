package main

import (
	"log"
	"os"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	if len(os.Args) != 2 {
		log.Print("Usage: luapaint <file.png>")
		return
	}
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(os.Args[1])
	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(text)
}
