package main

import (
	"fmt"
	"io"
	"net/http"
)

var players = map[string]string{
	"Danek": "412196175",
}

func main() {
	recentMatches("Danek")
	showInfoOfMatchByPlayer(players["Danek"], "8093496511")
}

func recentMatches(name string) {
	resp, err := http.Get("https://api.opendota.com/api/players/" + players[name] + "/matches/?date=18")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println(body)
	// var matches []PlayerMatch
	// err = json.Unmarshal(body, &matches)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// for _, v := range matches {
	// 	fmt.Println(v.toString())
	// }

}

func showInfoOfMatchByPlayer(playerId string, matchId string) {
	fmt.Println(playerId)
	resp, err := http.Get("https://api.opendota.com/api/matches/" + matchId)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println(body)
	// var match Match
	// err = json.Unmarshal(body, &match)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(match)
}
