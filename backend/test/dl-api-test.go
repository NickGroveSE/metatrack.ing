package main

import (
	// "encoding/json"
	// "io"
	// "log"
	// "net/http"
	// "slices"
	// "strconv"
	// "strings"
	// "github.com/NickGroveSE/metatrack.ing/backend/models"
	// "github.com/PuerkitoBio/goquery"
	"context"
	"fmt"

	// "slices"
	// "time"

	deadlock_game_api "github.com/deadlock-api/openapi-clients/go/api"
	deadlock_assets_api "github.com/deadlock-api/openapi-clients/go/assets-api"
)

type DLHero struct {
	ID    int32
	Name  string
	Image string
	Color string
}

type DLHeroStats struct {
	Wins     int64
	Matches  int64
	PickRate float32
	WinRate  float32
}

type DLHeroEntity struct {
	Heroes []DLHero
	Stats  DLHeroStats
}

var DLHeroColors = map[int32]string{
	1:  "#F80700", // Infernus
	2:  "#F4E80B", // Seven
	3:  "#8AA6F6", // Vindicta
	4:  "#1AC78D", // Lady Geist
	6:  "#43A0B3", // Abrams
	7:  "#8D2C7A", // Wraith
	8:  "#295A9D", // McGinnis
	10: "#FD99B6", // Paradox
	11: "#798BA9", // Dynamo
	12: "#C76F58", // Kelvin
	13: "#B36635", // Haze
	14: "#E85835", // Holliday
	15: "#6DF5BE", // Bebop
	16: "#F0D870", // Calico
	17: "#6A8552", // Grey Talon
	18: "#634F44", // Mo & Krill
	19: "#8B2229", // Shiv
	20: "#753B2A", // Ivy
	25: "#4B5467", // Warden
	27: "#8E7BAA", // Yamato
	31: "#145EE3", // Lash
	35: "#34892D", // Viscous
	50: "#6A4245", // Pocket
	52: "#9B7351", // Mirage
	58: "#73B740", // Vyper
	60: "#857FBD", // Sinclair
	63: "#B01A1C", // Mina
	64: "#A81010", // Drifter
	66: "#8C8B82", // Victor
	67: "#939E30", // Paige
	69: "#7C495A", // The Doorman
	72: "#E21EC7", // Billy
}

// func main() {

// 	DLSingleHeroHandler()

// }

func DLSingleHeroHandler() {

	assetsConfig := deadlock_assets_api.NewConfiguration()
	assetsAPIClient := deadlock_assets_api.NewAPIClient(assetsConfig)

	gameConfig := deadlock_game_api.NewConfiguration()
	gameAPIClient := deadlock_game_api.NewAPIClient(gameConfig)

	var hardCodedMinUnixTS int64 = 1763765592
	var hardCodedMaxUnixTS int64 = 1763855999
	var hardCodedMinAvgBadge int32 = 91
	var hardCodedMaxAvgBadge int32 = 116
	var minHeroMatches = 0
	var minHeroMatchesTotal = 0

	idLookup := activeHeroes(assetsAPIClient)

	heroStats(gameAPIClient, idLookup, hardCodedMinUnixTS, hardCodedMaxUnixTS, hardCodedMinAvgBadge, hardCodedMaxAvgBadge, int64(minHeroMatches), int64(minHeroMatchesTotal))

}

func activeHeroes(assetsClient *deadlock_assets_api.APIClient) map[int32]DLHero {

	res, _, err := assetsClient.HeroesAPI.GetHeroesV2HeroesGet(context.Background()).OnlyActive(true).Execute()
	if err != nil {
		fmt.Println("Error")
	}

	idLookup := make(map[int32]DLHero)

	for _, v := range res {
		idLookup[v.Id] = DLHero{v.Id, v.Name, v.Images.GetMinimapImageWebp(), DLHeroColors[v.Id]}
	}

	return idLookup

}

func heroStats(gameClient *deadlock_game_api.APIClient, activeLookup map[int32]DLHero, minUnixTS int64, maxUnixTS int64, minAvgBadge int32, maxAvgBadge int32, minHeroMatches int64, minHeroMatchesTotal int64) {

	var heroEntities []DLHeroEntity
	var runningMatchSum int64 = 0

	res, _, err := gameClient.AnalyticsAPI.HeroStats(context.Background()).MinUnixTimestamp(minUnixTS).MaxUnixTimestamp(maxUnixTS).MinAverageBadge(minAvgBadge).MaxAverageBadge(maxAvgBadge).MinHeroMatches(minHeroMatches).MinHeroMatchesTotal(minHeroMatchesTotal).Execute()
	if err != nil {
		fmt.Println("Error")
	}

	for _, hero := range res {
		heroObj, contains := activeLookup[hero.HeroId]
		if contains {
			runningMatchSum += hero.Matches
			heroEntities = append(heroEntities, DLHeroEntity{Heroes: []DLHero{heroObj}, Stats: DLHeroStats{Wins: hero.Wins, Matches: hero.Matches, PickRate: 0, WinRate: float32(hero.Wins) / float32(hero.Matches)}})
		}
	}

	for i := range heroEntities {
		divisionByNumOfPlayersPerMatch := float32(runningMatchSum) / float32(12)
		heroEntities[i].Stats.PickRate = float32(heroEntities[i].Stats.Matches) / divisionByNumOfPlayersPerMatch
	}

	fmt.Println(heroEntities)

}
