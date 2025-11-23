### Structure

Frontend will include a "Single Hero & Combos" tab & a "Matchups" tab

Rank filters with our combinations in the hero panel may be tricky

DLSHOCHandler() {
    call active heroes
    call hero stats
    assign colors
}

activeHeroes() {
    fetch active heroes
}

heroStats() {
    fetch hero stats
    use active heroes to filter heroes
    calculate win rates
    calculate pick rates
}


### API Calls

https://api.deadlock-api.com/v1/analytics/hero-stats?min_unix_timestamp=1761350091&max_unix_timestamp=1763769599&min_average_badge=91&max_average_badge=116&min_hero_matches=0&min_hero_matches_total=0
- divide wins by matches for win rate
- sum heroes' match numbers divide by 12, then divide each hero's number of matches by the original calculation for pick rate

https://assets.deadlock-api.com/v2/heroes?only_active=true
- receive ids of active heroes

http://localhost:8080/deadlock?min-unix-ts=1763765592&max-unix-ts=1763855999&min-badge=91&max-badge=116&min-hero-matches=0&min-hero-matches-total=0



### Color Assignment

heroColors := map[string]int{
    "1" : "#F80700", // Infernus
    "2" : "#F4E80B", // Seven
    "3" : "#8AA6F6", // Vindicta
    "4" : "#1AC78D", // Lady Geist
    "6" : "#43A0B3", // Abrams
    "7" : "#8D2C7A", // Wraith
    "8" : "#295A9D", // McGinnis
    "10" : "#FD99B6", // Paradox
    "11" : "#798BA9", // Dynamo
    "12" : "#C76F58", // Kelvin
    "13" : "#B36635", // Haze
    "14" : "#E85835", // Holliday
    "15" : "#6DF5BE", // Bebop
    "16" : "#F0D870", // Calico
    "17" : "#6A8552", // Grey Talon
    "18" : "#634F44", // Mo & Krill
    "19" : "#8B2229", // Shiv
    "20" : "#753B2A", // Ivy
    "25" : "#4B5467", // Warden
    "27" : "#8E7BAA", // Yamato
    "31" : "#145EE3", // Lash
    "35" : "#34892D", // Viscous
    "50" : "#6A4245", // Pocket
    "52" : "#9B7351", // Mirage
    "58" : "#73B740", // Vyper
    "60" : "#857FBD", // Sinclair
    "63" : "#B01A1C", // Mina
    "64" : "#A81010", // Drifter
    "66" : "#8C8B82", // Victor
    "67" : "#939E30", // Paige
    "69" : "#7C495A", // The Doorman
    "72" : "#E21EC7" // Billy
}

"" : "", //