package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/id10110011/dota2bot/src/internal/model"
	"github.com/id10110011/dota2bot/src/internal/request"
	"github.com/id10110011/dota2bot/src/internal/util"
)

var players = map[string]int{
	"Danek": 412196175,
}

var lastMatch = 0

func main() {
	config, err := util.LoadConfig("./src")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Println(err.Error())
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	for {
		select {
		case <-ticker.C:
			match := requestToAPI()
			if match != nil {
				info := util.BuildInfoOfPlayer(match, players["Danek"])
				msg := tgbotapi.NewMessage(config.ChatID, info)
				bot.Send(msg)
			}
		}
	}

}

func requestToAPI() *model.Match {
	code, err := request.RefreshPlayerHistory(players["Danek"])
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Refresh player history with status code: %d", code)

	matchId, err := request.RecentMatches(players["Danek"])
	if err != nil {
		log.Println(err.Error())
	}

	if matchId != -1 && matchId != lastMatch {
		lastMatch = matchId
		match, err := request.ShowInfoOfMatchByPlayer(matchId)
		if err != nil {
			log.Println(err.Error())
		}

		return match
	}

	return nil
}
