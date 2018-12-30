package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"os"
	"log"
	"encoding/json"
)

//Struct
type IssLocation struct {
	Message     string `json:"message"`
	Timestamp   int64    `json:"timestamp"`
	IssPosition struct {
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	} `json:"iss_position"`
}

// sample
// 20181230152434
// http://api.open-notify.org/iss-now.json

// {
//   "message": "success",
//   "timestamp": 1546151073,
//   "iss_position": {
//     "longitude": "-167.9689",
//     "latitude": "-51.5217"
//   }
// }

func main() {

	for range time.Tick(3 * time.Second) {
		fmt.Println("Tick!!")

		response, err := http.Get("http://api.open-notify.org/iss-now.json")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		json_byte := ([]byte)(data)
		iss_location := new(IssLocation)

		json.Unmarshal(json_byte, &iss_location)
		fmt.Println(time.Unix(iss_location.Timestamp, 0))
		fmt.Println(iss_location.IssPosition)
	}
}