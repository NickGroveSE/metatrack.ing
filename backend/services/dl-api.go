package services

import (
	// "encoding/json"
	// "io"
	// "log"
	// "net/http"

	// "strconv"
	// "strings"
	// "github.com/NickGroveSE/metatrack.ing/backend/models"
	// "github.com/PuerkitoBio/goquery"
	"context"
	"fmt"
	"math"
	"slices"
	"sort"

	// "time"

	"github.com/NickGroveSE/metatrack.ing/backend/models"
	deadlock_game_api "github.com/deadlock-api/openapi-clients/go/api"
	deadlock_assets_api "github.com/deadlock-api/openapi-clients/go/assets-api"
)

var DLHeroColors = map[int32]string{
	1:         "#F80700", // Infernus
	2:         "#F4E80B", // Seven
	3:         "#8AA6F6", // Vindicta
	4:         "#1AC78D", // Lady Geist
	6:         "#43A0B3", // Abrams
	7:         "#8D2C7A", // Wraith
	8:         "#295A9D", // McGinnis
	10:        "#FD99B6", // Paradox
	11:        "#798BA9", // Dynamo
	12:        "#C76F58", // Kelvin
	13:        "#B36635", // Haze
	14:        "#E85835", // Holliday
	15:        "#6DF5BE", // Bebop
	16:        "#F0D870", // Calico
	17:        "#6A8552", // Grey Talon
	18:        "#634F44", // Mo & Krill
	19:        "#8B2229", // Shiv
	20:        "#753B2A", // Ivy
	25:        "#4B5467", // Warden
	27:        "#8E7BAA", // Yamato
	31:        "#145EE3", // Lash
	35:        "#34892D", // Viscous
	50:        "#6A4245", // Pocket
	52:        "#9B7351", // Mirage
	58:        "#73B740", // Vyper
	60:        "#857FBD", // Sinclair
	63:        "#B01A1C", // Mina
	64:        "#A81010", // Drifter
	66:        "#8C8B82", // Victor
	67:        "#939E30", // Paige
	69:        "#7C495A", // The Doorman
	72:        "#E21EC7", // Billy
	999999999: "#FBBF24", // Default (reserved for new heroes that dont have a color assigned yet)
}

// func main() {

// 	DLSingleHeroHandler()

// }

func DLSCHandler(minUnixTS int64, maxUnixTS int64, minAvgBadge int32, maxAvgBadge int32, minHeroMatches int64, minHeroMatchesTotal int64) []models.DLHeroEntity {

	var heroEntities []models.DLHeroEntity

	assetsConfig := deadlock_assets_api.NewConfiguration()
	assetsAPIClient := deadlock_assets_api.NewAPIClient(assetsConfig)

	gameConfig := deadlock_game_api.NewConfiguration()
	gameAPIClient := deadlock_game_api.NewAPIClient(gameConfig)

	// var hardCodedMinUnixTS int64 = 1763765592
	// var hardCodedMaxUnixTS int64 = 1763855999
	// var hardCodedMinAvgBadge int32 = 91
	// var hardCodedMaxAvgBadge int32 = 116
	// var minHeroMatches = 0
	// var minHeroMatchesTotal = 0

	idLookup := activeHeroes(assetsAPIClient)

	heroEntities = heroStats(gameAPIClient, idLookup, minUnixTS, maxUnixTS, minAvgBadge, maxAvgBadge, minHeroMatches, minHeroMatchesTotal)

	return heroEntities

}

func DLComboAdditionHandler(minUnixTS int64, maxUnixTS int64, minAvgBadge int32, maxAvgBadge int32, minMatches int32, combo []int32) models.DLHeroEntity {

	assetsConfig := deadlock_assets_api.NewConfiguration()
	assetsAPIClient := deadlock_assets_api.NewAPIClient(assetsConfig)

	gameConfig := deadlock_game_api.NewConfiguration()
	gameAPIClient := deadlock_game_api.NewAPIClient(gameConfig)

	// var hardCodedMinUnixTS int64 = 1763765592
	// var hardCodedMaxUnixTS int64 = 1763855999
	// var hardCodedMinAvgBadge int32 = 91
	// var hardCodedMaxAvgBadge int32 = 116
	// var minHeroMatches = 0
	// var minHeroMatchesTotal = 0

	idLookup := activeHeroes(assetsAPIClient)

	return heroCombStats(gameAPIClient, idLookup, minUnixTS, maxUnixTS, minAvgBadge, maxAvgBadge, minMatches, combo)

}

func activeHeroes(assetsClient *deadlock_assets_api.APIClient) map[int32]models.DLHero {

	res, _, err := assetsClient.HeroesAPI.GetHeroesV2HeroesGet(context.Background()).OnlyActive(true).Execute()
	if err != nil {
		fmt.Println("Error")
	}

	idLookup := make(map[int32]models.DLHero)

	for _, v := range res {
		idLookup[v.Id] = models.DLHero{ID: v.Id, Name: v.Name, Image: v.Images.GetMinimapImageWebp(), Color: DLHeroColors[v.Id]}
	}

	return idLookup

}

func heroStats(gameClient *deadlock_game_api.APIClient, activeLookup map[int32]models.DLHero, minUnixTS int64, maxUnixTS int64, minAvgBadge int32, maxAvgBadge int32, minHeroMatches int64, minHeroMatchesTotal int64) []models.DLHeroEntity {

	var heroEntities []models.DLHeroEntity
	var runningMatchSum int64 = 0

	res, _, err := gameClient.AnalyticsAPI.HeroStats(context.Background()).MinUnixTimestamp(minUnixTS).MaxUnixTimestamp(maxUnixTS).MinAverageBadge(minAvgBadge).MaxAverageBadge(maxAvgBadge).MinHeroMatches(minHeroMatches).MinHeroMatchesTotal(minHeroMatchesTotal).Execute()
	if err != nil {
		fmt.Println("Error")
	}

	for _, hero := range res {
		heroObj, contains := activeLookup[hero.HeroId]
		if contains {
			runningMatchSum += hero.Matches
			heroEntities = append(heroEntities, models.DLHeroEntity{Heroes: []models.DLHero{heroObj}, Stats: models.DLHeroStats{Wins: hero.Wins, Matches: hero.Matches, PickRate: 0, WinRate: roundFloat32((float32(hero.Wins)/float32(hero.Matches))*100, 2)}})
		}
	}

	for i := range heroEntities {
		divisionByNumOfPlayersPerMatch := float32(runningMatchSum) / float32(12)
		heroEntities[i].Stats.PickRate = roundFloat32((float32(heroEntities[i].Stats.Matches)/divisionByNumOfPlayersPerMatch)*100, 2)
	}

	sort.Slice(heroEntities, func(i, j int) bool {
		return heroEntities[i].Stats.PickRate > heroEntities[j].Stats.PickRate
	})

	return heroEntities

}

func heroCombStats(gameClient *deadlock_game_api.APIClient, activeLookup map[int32]models.DLHero, minUnixTS int64, maxUnixTS int64, minAvgBadge int32, maxAvgBadge int32, minMatches int32, combo []int32) models.DLHeroEntity {

	var heroComboEntity models.DLHeroEntity
	var biDirectionalAvgs []float32

	res, _, err := gameClient.AnalyticsAPI.HeroCombStats(context.Background()).MinUnixTimestamp(minUnixTS).MaxUnixTimestamp(maxUnixTS).MinAverageBadge(minAvgBadge).MaxAverageBadge(maxAvgBadge).MinMatches(int32(minMatches)).CombSize(int32(len(combo))).Execute()
	if err != nil {
		fmt.Println("Error")
	}

	for _, heroCombo := range res {
		if slices.Equal(combo, heroCombo.HeroIds) {
			heroComboEntity = models.DLHeroEntity{Heroes: comboFilter(activeLookup, combo), Stats: models.DLHeroStats{Wins: heroCombo.Wins, Matches: heroCombo.Matches, PickRate: 0, WinRate: roundFloat32((float32(heroCombo.Wins)/float32(heroCombo.Matches))*100, 2)}}
		}
		// fmt.Println(heroCombo.HeroIds)
		// fmt.Println(heroCombo.Matches)
	}

	for _, heroEntity := range comboFilterforHeroStats(heroStats(gameClient, activeLookup, minUnixTS, maxUnixTS, minAvgBadge, maxAvgBadge, int64(minMatches), int64(minMatches)), combo) {

		biDirectionalAvgs = append(biDirectionalAvgs, float32(heroComboEntity.Stats.Matches)/float32(heroEntity.Stats.Matches))
		fmt.Println(float32(heroComboEntity.Stats.Matches))
		fmt.Println(float32(heroEntity.Stats.Matches))

	}

	heroComboEntity.Stats.PickRate = roundFloat32(biDirectionalAvgCalc(biDirectionalAvgs)*100, 2)
	fmt.Println(roundFloat32(biDirectionalAvgCalc(biDirectionalAvgs)*100, 2))

	return heroComboEntity

}

func roundFloat32(val float32, precision uint) float32 {
	ratio := float32(math.Pow(10, float64(precision)))
	return float32(math.Round(float64(val*ratio))) / ratio
}

func comboFilter(activeLookup map[int32]models.DLHero, combo []int32) []models.DLHero {

	var filtered []models.DLHero
	for _, hero := range activeLookup {
		if slices.Contains(combo, hero.ID) {
			filtered = append(filtered, hero)
		}
	}

	return filtered

}

func comboFilterforHeroStats(heroStatsArray []models.DLHeroEntity, combo []int32) []models.DLHeroEntity {

	var filtered []models.DLHeroEntity
	for _, heroEntity := range heroStatsArray {
		if slices.Contains(combo, heroEntity.Heroes[0].ID) {
			filtered = append(filtered, heroEntity)
		}
	}

	return filtered

}

func biDirectionalAvgCalc(biDA []float32) float32 {
	var sum float32
	for _, val := range biDA {
		sum += val
	}
	return sum / float32(len(biDA))
}
