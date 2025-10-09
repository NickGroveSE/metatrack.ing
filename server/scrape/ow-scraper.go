package scrape

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/NickGroveSE/metatrack.ing/server/models"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	Scrape("PC", "all-maps", "Americas", "Support", "1", "All")

}

func Scrape(input string, owmap string, region string, role string, queue string, rank string) []models.OWHero {

	var heroes []models.OWHero
	var indexes []int

	requestLink := "https://overwatch.blizzard.com/en-us/rates/?input=" + input + "&map=" + owmap + "&region=" + region + "&role=All&rq=" + queue + "&tier=" + rank
	fmt.Println(requestLink)

	res, err := http.Get("https://overwatch.blizzard.com/en-us/rates/?input=" + input + "&map=" + owmap + "&region=" + region + "&role=All&rq=" + queue + "&tier=" + rank)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer res.Body.Close() // Always close the response body

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return heroes
	}

	if role != "All" {
		doc.Find(".hero-role-icon").Each(func(iconIndex int, iconElement *goquery.Selection) {
			srcString, exists := iconElement.Attr("src")

			if exists && strings.Contains(srcString, strings.ToLower(role)) {
				indexes = append(indexes, iconIndex)
			}
		})
	}

	// Extract and print data
	doc.Find(".hero-name").Each(func(index int, element *goquery.Selection) {

		if slices.Contains(indexes, index) {
			heroColor, exists := doc.Find("[alt='" + element.Text() + "']").Attr("style")

			if exists {
				heroColor = heroColor[13 : len(heroColor)-2]
			}

			heroImage, _ := doc.Find("[alt='" + element.Text() + "']").Attr("src")

			replacer := strings.NewReplacer(
				".", "",
				" ", "-",
				"ú", "u",
				":", "",
				"ö", "o",
			)

			preConvertedPickrate, err := strconv.ParseFloat(strings.ReplaceAll(doc.Find("#"+replacer.Replace(strings.ToLower(element.Text()))+"-pickrate-value").Text(), "%", ""), 64)
			if err != nil {
				fmt.Println("Error Parsing Pick Rate "+strings.ReplaceAll(doc.Find("#"+replacer.Replace(strings.ToLower(element.Text()))+"-pickrate-value").Text(), "%", "")+" for "+element.Text(), err)
				return
			}

			preConvertedWinrate, err := strconv.ParseFloat(strings.ReplaceAll(doc.Find("#"+replacer.Replace(strings.ToLower(element.Text()))+"-winrate-value").Text(), "%", ""), 64)
			if err != nil {
				fmt.Println("Error Parsing Win Rate", err)
				return
			}

			hero := models.OWHero{Name: element.Text(), Color: heroColor, Image: heroImage, PickRate: float32(preConvertedPickrate), WinRate: float32(preConvertedWinrate)}
			heroes = append(heroes, hero)

			// fmt.Println("Hero:", heroes[index].Name)
			// fmt.Println("Color:", heroes[index].Color)
			// fmt.Println("Image:", heroes[index].Image)
			// fmt.Println("Pick Rate:", heroes[index].PickRate)
			// fmt.Println("Win Rate:", heroes[index].WinRate)
		}
	})

	return heroes

}
