package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// "slices"
	// "strconv"
	// "strings"
	"github.com/NickGroveSE/metatrack.ing/backend/models"
	// "github.com/PuerkitoBio/goquery"
)

const activeHeroRequest = "https://assets.deadlock-api.com/v2/heroes?only_active=true"

func main() {

	DLSingleHeroHandler()

}

func DLSingleHeroHandler() {

	resp, err := http.Get(activeHeroRequest)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", body)

	var activeHeroes []models.DLHero

	err = json.Unmarshal(body, &activeHeroes)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Printf("IDL %d, Name: %s, Image: %s", activeHeroes[0].ID, activeHeroes[0].Name, activeHeroes[0].Image)

}
