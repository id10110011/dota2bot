package model

import "fmt"

type PlayerMatch struct {
	MatchId      int  `json:"match_id"`
	PlayerSlot   int  `json:"player_slot"`
	RadiantWin   bool `json:"radiant_win"`
	HeroId       int  `json:"hero_id"`
	StartTime    int  `json:"start_time"`
	Duration     int  `json:"duration"`
	GameMode     int  `json:"game_mode"`
	LobbyType    int  `json:"email"`
	Version      int  `json:"version"`
	Kills        int  `json:"kills"`
	Deaths       int  `json:"deaths"`
	Assists      int  `json:"assists"`
	AverageRank  int  `json:"average_rank"`
	Xpm          int  `json:"xp_per_min"`
	Gpm          int  `json:"gold_per_min"`
	HeroDamage   int  `json:"hero_damage"`
	TowerDamage  int  `json:"tower_damage"`
	HeroHealing  int  `json:"hero_healing"`
	LastHits     int  `json:"last_hits"`
	Lane         int  `json:"lane"`
	LaneRole     int  `json:"lane_role"`
	IsRoaming    bool `json:"is_roaming"`
	Cluster      bool `json:"cluster"`
	LeaverStatus int  `json:"leaver_status"`
	PartySize    int  `json:"party_size"`
	HeroVariant  int  `json:"hero_variant"`
}

func (match PlayerMatch) toString() string {
	return fmt.Sprintf(
		"MatchId: %d\t"+
			"PlayerSlot: %d\t"+
			"RadiantWin: %t\t"+
			"HeroId: %d\t"+
			"StartTime: %d\t"+
			"Duration: %d\t"+
			"GameMode: %d\t"+
			"LobbyType: %d\t"+
			"Version: %d\t"+
			"Kills: %d\t"+
			"Deaths: %d\t"+
			"Assists: %d\t"+
			"AverageRank: %d\t"+
			"Xpm: %d\t"+
			"Gpm: %d\t"+
			"HeroDamage: %d\t"+
			"TowerDamage: %d\t"+
			"HeroHealing: %d\t"+
			"LastHits: %d\t"+
			"Lane: %d\t"+
			"LaneRole: %d\t"+
			"IsRoaming: %t\t"+
			"Cluster: %t\t"+
			"LeaverStatus: %d\t"+
			"PartySize: %d\t"+
			"HeroVariant: %d",
		match.MatchId, match.PlayerSlot, match.RadiantWin, match.HeroId, match.StartTime, match.Duration, match.GameMode,
		match.LobbyType, match.Version, match.Kills, match.Deaths, match.Assists, match.AverageRank, match.Xpm, match.Gpm,
		match.HeroDamage, match.TowerDamage, match.HeroHealing, match.LastHits, match.Lane, match.LaneRole, match.IsRoaming,
		match.Cluster, match.LeaverStatus, match.PartySize, match.HeroVariant,
	)
}
