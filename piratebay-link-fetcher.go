package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var limit = 5

func printResults(links []string) {
	if len(links) == 0 {
		fmt.Println("No results! Sorry :(")
	} else {
		printLinks(links)
	}
}

func printLinks(links []string) {
	fmt.Println("Here are links you are looking for. Enjoy!\n")

	for _, link := range links {
		fmt.Println(link, "\n")
	}
}

func main() {

	commandParams := os.Args
	if len(commandParams) < 2 {
		fmt.Println("ERROR!")
		fmt.Println("You need tell me what you are looking for!")
		os.Exit(1)
	}

	if len(commandParams) >= 4 {
		limit, _ = strconv.Atoi(commandParams[3])
	}

	var search SearchParams
	search.Create(commandParams[1:])
	url := search.ConstructUrl()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	links := GetMagnetLinks(resp.Body, limit)
	printResults(links)

}
