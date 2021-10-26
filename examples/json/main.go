package main

import (
	"fmt"
	"os"
	"reflect"
	"encoding/json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file(path * string) []byte {
	f, err := os.Open(*path)
	defer f.Close()
	check(err)

	stat, _ := f.Stat()
	// fmt.Println(reflect.ValueOf(stat).Type())
	buf_size := stat.Size()

	buffer := make([] byte, buf_size)
	_ , err = f.Read(buffer)
	check(err)

	return buffer
}

func main() {
	fmt.Println("test")

	path := "./complex.json"
	buffer := read_file(&path)
	// fmt.Println(string(buffer))

	var json_map map[string] interface {}
	err := json.Unmarshal(buffer, &json_map)
	check(err)
	fmt.Println(reflect.ValueOf(json_map).Type())

	// fmt.Println(json_map)
	// fmt.Println(json_map["items"])
	fmt.Println(json_map["items"].(map[string]interface{})["item"].(string))
}