package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/id10110011/dota2bot/src/internal/model"
)

func RefreshPlayerHistory(id int) (int, error) {
	url := fmt.Sprintf("%s%d%s", "https://api.opendota.com/api/players/", id, "/refresh")
	resp, err := http.Post(url, "application/json", nil)
	defer resp.Body.Close()

	return resp.StatusCode, err
}

func RecentMatches(id int) (int, error) {
	url := fmt.Sprintf("%s%d%s", "https://api.opendota.com/api/players/", id, "/matches/?date=1")
	resp, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	var matches []model.PlayerMatch
	err = json.Unmarshal(body, &matches)
	if err != nil {
		return -1, err
	}

	return matches[0].MatchId, nil
}

func ShowInfoOfMatchByPlayer(matchId int) (*model.Match, error) {
	url := fmt.Sprint("https://api.opendota.com/api/matches/", matchId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var match model.Match
	err = json.Unmarshal(body, &match)
	if err != nil {
		return nil, err
	}
	return &match, nil
}
