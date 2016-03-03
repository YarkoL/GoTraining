package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var url string = "https://blockchain.info/latestblock"

func main() {
	noCheck()
}

func withCheck() {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}

func noCheck() {
	res, _ := http.Get(url)
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
