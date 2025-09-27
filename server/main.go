package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://www.mtggoldfish.com/metagame/standard#paper")
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer res.Body.Close() // Always close the response body

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalf("Error reading response body: %v", err)
	// }

	fmt.Println(res)

	// Extract and print data
	doc.Find("tr").Each(func(index int, element *goquery.Selection) {
		title := element.Text()
		fmt.Println("Title:", title)
	})
}
