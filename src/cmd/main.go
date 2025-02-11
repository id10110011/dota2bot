package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/id10110011/dota2bot/src/internal/model"
)

var players = map[string]string{
	"Danek": "412196175",
}

func main() {
	ticker := time.NewTicker(1 * time.Minute) // Интервал: 1 минута
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			showInfoOfMatchByPlayer(players["Danek"], "8168593562")
		}
	}
	//recentMatches("Danek")
}

func recentMatches(name string) {
	resp, err := http.Get("https://api.opendota.com/api/players/" + players[name] + "/matches/?date=1")
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

	var matches []model.PlayerMatch
	err = json.Unmarshal(body, &matches)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range matches {
		fmt.Println(v.CovertToString())
	}

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
	var match model.Match
	err = json.Unmarshal(body, &match)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(match)
}
