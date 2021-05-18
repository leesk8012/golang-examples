package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	defer resp.Body.Close()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Elapsed %s\n",elapsed)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Howdy %s\n", body)
}
