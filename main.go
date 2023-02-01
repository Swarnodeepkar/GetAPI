package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", dataApi)
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}

func dataApi(w http.ResponseWriter, r *http.Request) {
	url := "https://random-data-api.com/api/v2/users"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responsestr := string(responseData)
	fmt.Println(w, responsestr)
}
