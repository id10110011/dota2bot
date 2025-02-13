package util

import (
	"fmt"
	"slices"
	"strings"

	"github.com/id10110011/dota2bot/src/internal/constants"
	"github.com/id10110011/dota2bot/src/internal/model"
)

func BuildInfoOfPlayer(match *model.Match, playerId int) string {
	sb := strings.Builder{}

	idx := slices.IndexFunc(match.Players, func(player model.Player) bool { return player.AccountID == playerId })

	if idx != -1 {
		player := match.Players[idx]
		sb.WriteString("Данек сейчас сыграл игру в дотку на " + constants.Heroes[player.HeroID] + " под ником: " + player.PersonaName + "\n")
		if player.Win == 1 {
			sb.WriteString("Провезли руину на победу\n")
		} else {
			sb.WriteString("Бездарно проиграл\n")
		}

		sb.WriteString("Стата в катке:\n")
		sb.WriteString(fmt.Sprintf("%s%d%s\n%s%d-%d-%d\n%s%d👾\n%s%d🔶\n",
			"Длительность игры: ", player.Duration/60, " мин.",
			"KDA: ", player.Kills, player.Deaths, player.Assists,
			"Крипочки: ", player.LastHits,
			"GPM: ", player.GoldPerMin))

	}

	return sb.String()
}
