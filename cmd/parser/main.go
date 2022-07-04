package main

import (
	"fmt"
	"os"
	"std/sunduck/parser/internal/urls"
)

func main() {
	// для ввода через панель https://ya.ru
	//for _, url := range os.Args[1:] {

	// слайс линков из заданного списка
	links := urls.Catalog()     // ..
	for _, url := range links { //  https://ya.ru

		// поиск линков на странице
		var answerlinks []string
		answerlinks = urlfinder(url)
		for _, link := range answerlinks {
			fmt.Println("links on page = ", url)
			fmt.Println(link)
		}
	}
}

func urlfinder(url string) []string {
	links, err := urls.findLinks(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
	}
	return links
}
