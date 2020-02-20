package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

type githubRepo struct {
	PushedAt string `json:"pushed_at"`
}

func main() {
	resp, err := http.Get("https://api.github.com/repos/steamdatabase/steamtracking")

	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var r githubRepo
	err = json.Unmarshal(bytes, &r)

	if err != nil {
		panic(err)
	}

	fmt.Println(r.PushedAt)

	lastTime, _ := ioutil.ReadFile("lasttime.txt")

	if string(lastTime) != r.PushedAt {
		err = ioutil.WriteFile("lasttime.txt", []byte(r.PushedAt), 0644)

		if err != nil {
			panic(err)
		}

		exec.Command("./trigger.sh")
	}
}
