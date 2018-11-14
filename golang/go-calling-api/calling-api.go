package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "time"
)

type Sensor struct {
	class string `json:"string"`
  teamID int `json:"number"`
  teamName string `json:"string"`
  sensorTemp int `json:"number"`
  thermostatTemp int `json:"number"`
  recommendation int `json:"number"`
}

func main() {

	url := "http://148.100.98.30:3000/api/org.acme.sample.Sensor/1"
  //
	// spaceClient := http.Client{
	// 	Timeout: time.Second * 2, // Maximum of 2 secs
	// }
  //
	// req, err := http.NewRequest(http.MethodGet, url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
  // fmt.Println(req)
	// req.Header.Set("User-Agent", "spacecount-tutorial")
  //
	// res, getErr := spaceClient.Do(req)
	// if getErr != nil {
	// 	log.Fatal(getErr)
	// }
  // fmt.Println(res)
  //
	// body, readErr := ioutil.ReadAll(res.Body)
	// if readErr != nil {
	// 	log.Fatal(readErr)
	// }
  // fmt.Println(body)
	// sensor1 := Sensor{}
	// jsonErr := json.Unmarshal(body, &sensor1)
	// if jsonErr != nil {
	// 	log.Fatal(jsonErr)
	// }
  //
	// fmt.Println(sensor1)


  resp, err := http.Get(url)
  if err != nil {
  	log.Fatal(err)
  }
  fmt.Println(resp.Body)
  body, err := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body));
  var sensor1 map[string]interface{}
  jsonErr := json.Unmarshal(body, &sensor1)
  if jsonErr != nil {
		log.Fatal(jsonErr)
	}
  fmt.Println(sensor1)
  fmt.Println(sensor1["$class"])
}
