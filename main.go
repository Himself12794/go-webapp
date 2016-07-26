package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?id=4466033&APPID=85383b7344bff8b9af7f717eae540519")
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
