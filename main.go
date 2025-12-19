package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// hello := "Hello There"
	// helloPointer := &hello
	// fmt.Println("Point to hello")
	// fmt.Println(helloPointer)
	// fmt.Println(*helloPointer)

	peopleURL := getSwapiRespone().People
	swapiPeople := getSwapiPeople(peopleURL)
	for _, people := range swapiPeople {
		fmt.Println(people.Name)
	}
	// fmt.Println(swapiPeople[0].Name)
}

func getSwapiRespone() SwapiResponse {
	resp, err := http.Get("https://swapi.info/api/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))

	var swapiRespone SwapiResponse
	if err := json.Unmarshal([]byte(body), &swapiRespone); err != nil {
		slog.Error("Couldn't handle the JSON", "err", err)
		os.Exit(1)
	}

	// fmt.Println(swapiRespone)

	return swapiRespone
}

func getSwapiPeople(url string) []SwapiPeople {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var swapiPeople []SwapiPeople
	if err := json.Unmarshal([]byte(body), &swapiPeople); err != nil {
		// log.Fatalln(err)
		slog.Error("Couldn't handle the JSON", "err", err)
		os.Exit(1)
	}

	return swapiPeople
}
