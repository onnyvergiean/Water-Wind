package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

func main() {

	for true {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 100
		var data = map[string]interface{}{
			"water" : rand.Intn((max - min + 1) + min),
			"wind" : rand.Intn((max - min + 1) + min),
		}
		requestJson, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res,err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		body,err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(body))

		water := gjson.Get(string(body), "water")
		wind := gjson.Get(string(body), "wind")

		if water.Int() < 5 {
			log.Println("Status Water : Aman")
		}else if water.Int() >= 6 && water.Int() <= 8{
			log.Println("Status Water : Siaga")
		}else{
			log.Println("Status Water : Bahaya")
		}

		if wind.Int() < 6 {
			log.Println("Status Wind  : Aman")
		}else if wind.Int() >= 7 && wind.Int() <= 15{
			log.Println("Status Wind  : Siaga")
		}else {
			log.Println("Status Wind  : Bahaya")
		}
		time.Sleep(15 * time.Second)

	}

}

