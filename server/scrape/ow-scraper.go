package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/NickGroveSE/metatrack.ing/server/models"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	var heroes []models.OWHero

	res, err := http.Get("https://overwatch.blizzard.com/en-us/rates/?input=PC&map=all-maps&region=Americas&role=All&rq=1&tier=All")
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

	// Extract and print data
	doc.Find(".hero-name").Each(func(index int, element *goquery.Selection) {

		replacer := strings.NewReplacer(
			".", "",
			" ", "-",
			"ú", "u",
			":", "",
			"ö", "o",
		)

		preConvertedPickrate, err := strconv.ParseFloat(strings.ReplaceAll(doc.Find("#"+replacer.Replace(strings.ToLower(element.Text()))+"-pickrate-value").Text(), "%", ""), 64)
		if err != nil {
			fmt.Println("Error Parsing Pick Rate", err)
			return
		}

		preConvertedWinrate, err := strconv.ParseFloat(strings.ReplaceAll(doc.Find("#"+replacer.Replace(strings.ToLower(element.Text()))+"-winrate-value").Text(), "%", ""), 64)
		if err != nil {
			fmt.Println("Error Parsing Win Rate", err)
			return
		}

		hero := models.OWHero{Name: element.Text(), PickRate: float32(preConvertedPickrate), WinRate: float32(preConvertedWinrate)}
		heroes = append(heroes, hero)

		fmt.Println("Hero:", heroes[index].Name)
		fmt.Println("Pick Rate:", heroes[index].PickRate)
		fmt.Println("Win Rate:", heroes[index].WinRate)
	})
}
