package main


import (
  "net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)


type JsonStruct struct {
  Sted string
}

func main() {
  resp, err := http.Get("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
  fmt.Println("contents", contents)
	data := []JsonStruct{}
	err = json.Unmarshal(contents, &data)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println( "unmarshalled json", data)
	fmt.Println("\n contents: ", len(string(contents)), "\n \n")
	for i, v := range data {
		fmt.Println(i, "\n \n \n \n")
		fmt.Println(v.Sted)
	}
}
