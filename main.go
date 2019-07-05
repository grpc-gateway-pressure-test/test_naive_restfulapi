package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var g_content string

func ReadContent(level string) string {
	path := ""
	if level == "1" {
		path = "./1B.txt"
	} else if level == "2" {
		path = "./1KB.txt"
	} else if level == "3" {
		path = "./1MB.txt"
	}
	fmt.Println(path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error : %s", err)
		return ""
	}
	return string(bytes)
}

type StringMessage struct {
	Value string
}

func Echo(w http.ResponseWriter, r *http.Request) {
	result := &StringMessage{}
	result.Value = g_content

	byte, _ := json.Marshal(result)
	fmt.Fprintf(w, string(byte))
}

func main() {
	//level := os.Args[1]
	level := "3"
	g_content = ReadContent(level)
	http.HandleFunc("/v1/echo", Echo)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("ListenAndServe err", err)
	}
}
