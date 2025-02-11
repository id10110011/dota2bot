package model

import (
	"time"
)

// Main structure representing the match data
type Match struct {
	MatchID               uint64         `json:"match_id"`
	BarracksStatusDire    int            `json:"barracks_status_dire"`
	BarracksStatusRadiant int            `json:"barracks_status_radiant"`
	Chat                  []Chat         `json:"chat"`
	Cluster               int            `json:"cluster"`
	Cosmetics             Cosmetics      `json:"cosmetics"`
	DireScore             int            `json:"dire_score"`
	DraftTimings          []DraftTiming  `json:"draft_timings"`
	Duration              int            `json:"duration"`
	Engine                int            `json:"engine"`
	FirstBloodTime        int            `json:"first_blood_time"`
	GameMode              int            `json:"game_mode"`
	HumanPlayers          int            `json:"human_players"`
	LeagueID              int            `json:"leagueid"`
	LobbyType             int            `json:"lobby_type"`
	MatchSeqNum           int            `json:"match_seq_num"`
	NegativeVotes         int            `json:"negative_votes"`
	Objectives            []Objective    `json:"objectives"`
	PicksBans             []PickBan      `json:"picks_bans"`
	PositiveVotes         int            `json:"positive_votes"`
	RadiantGoldAdv        []int          `json:"radiant_gold_adv"`
	RadiantScore          int            `json:"radiant_score"`
	RadiantWin            bool           `json:"radiant_win"`
	RadiantXPAdv          []int          `json:"radiant_xp_adv"`
	StartTime             int            `json:"start_time"`
	TeamFights            []TeamFight    `json:"teamfights"`
	TowerStatusDire       int            `json:"tower_status_dire"`
	TowerStatusRadiant    int            `json:"tower_status_radiant"`
	Version               int            `json:"version"`
	ReplaySalt            int            `json:"replay_salt"`
	SeriesID              int            `json:"series_id"`
	SeriesType            int            `json:"series_type"`
	RadiantTeam           Team           `json:"radiant_team"`
	DireTeam              Team           `json:"dire_team"`
	League                League         `json:"league"`
	Skill                 int            `json:"skill"`
	Players               []Player       `json:"players"`
	Patch                 int            `json:"patch"`
	Region                int            `json:"region"`
	AllWordCounts         map[string]int `json:"all_word_counts"`
	MyWordCounts          map[string]int `json:"my_word_counts"`
	Throw                 int            `json:"throw"`
	Comeback              int            `json:"comeback"`
	Loss                  int            `json:"loss"`
	Win                   int            `json:"win"`
	ReplayURL             string         `json:"replay_url"`
}

// Structs for nested data
type Chat struct {
	Time       int    `json:"time"`
	Unit       string `json:"unit"`
	Key        string `json:"key"`
	Slot       int    `json:"slot"`
	PlayerSlot int    `json:"player_slot"`
}

type Cosmetics struct {
	Property1 int `json:"property1"`
	Property2 int `json:"property2"`
}

type DraftTiming struct {
	Order          int  `json:"order"`
	Pick           bool `json:"pick"`
	ActiveTeam     int  `json:"active_team"`
	HeroID         int  `json:"hero_id"`
	PlayerSlot     int  `json:"player_slot"`
	ExtraTime      int  `json:"extra_time"`
	TotalTimeTaken int  `json:"total_time_taken"`
}

type Objective struct {
	// Can be expanded depending on further details from the API
}

type PickBan struct {
	IsPick bool `json:"is_pick"`
	HeroID int  `json:"hero_id"`
	Team   int  `json:"team"`
	Order  int  `json:"order"`
}

type Team struct {
	// Can be expanded depending on further details from the API
}

type League struct {
	// Can be expanded depending on further details from the API
}

type TeamFight struct {
	// Can be expanded depending on further details from the API
}

type Player struct {
	MatchID                 uint64            `json:"match_id"`
	PlayerSlot              int               `json:"player_slot"`
	AbilityUpgradesArr      []int             `json:"ability_upgrades_arr"`
	AbilityUses             map[string]int    `json:"ability_uses"`
	AbilityTargets          map[string]int    `json:"ability_targets"`
	DamageTargets           map[string]int    `json:"damage_targets"`
	AccountID               int               `json:"account_id"`
	Actions                 map[string]int    `json:"actions"`
	AdditionalUnits         []Unit            `json:"additional_units"`
	Assists                 int               `json:"assists"`
	Backpack0               int               `json:"backpack_0"`
	Backpack1               int               `json:"backpack_1"`
	Backpack2               int               `json:"backpack_2"`
	BuybackLog              []BuybackLog      `json:"buyback_log"`
	CampsStacked            int               `json:"camps_stacked"`
	ConnectionLog           []ConnectionLog   `json:"connection_log"`
	CreepsStacked           int               `json:"creeps_stacked"`
	Damage                  map[string]int    `json:"damage"`
	DamageInflictor         map[string]int    `json:"damage_inflictor"`
	DamageInflictorReceived map[string]int    `json:"damage_inflictor_received"`
	DamageTaken             map[string]int    `json:"damage_taken"`
	Deaths                  int               `json:"deaths"`
	Denies                  int               `json:"denies"`
	DnT                     []int             `json:"dn_t"`
	Gold                    int               `json:"gold"`
	GoldPerMin              int               `json:"gold_per_min"`
	GoldReasons             map[string]int    `json:"gold_reasons"`
	GoldSpent               int               `json:"gold_spent"`
	GoldT                   []int             `json:"gold_t"`
	HeroDamage              int               `json:"hero_damage"`
	HeroHealing             int               `json:"hero_healing"`
	HeroHits                map[string]int    `json:"hero_hits"`
	HeroID                  int               `json:"hero_id"`
	Item0                   int               `json:"item_0"`
	Item1                   int               `json:"item_1"`
	Item2                   int               `json:"item_2"`
	Item3                   int               `json:"item_3"`
	Item4                   int               `json:"item_4"`
	Item5                   int               `json:"item_5"`
	ItemUses                map[string]int    `json:"item_uses"`
	KillStreaks             map[string]int    `json:"kill_streaks"`
	Killed                  map[string]int    `json:"killed"`
	KilledBy                map[string]int    `json:"killed_by"`
	Kills                   int               `json:"kills"`
	KillsLog                []KillsLog        `json:"kills_log"`
	LanePos                 map[string]int    `json:"lane_pos"`
	LastHits                int               `json:"last_hits"`
	LeaverStatus            int               `json:"leaver_status"`
	Level                   int               `json:"level"`
	LHT                     []int             `json:"lh_t"`
	LifeState               map[string]int    `json:"life_state"`
	MaxHeroHit              map[string]int    `json:"max_hero_hit"`
	MultiKills              map[string]int    `json:"multi_kills"`
	Obs                     map[string]int    `json:"obs"`
	ObsLeftLog              []ObsLeftLog      `json:"obs_left_log"`
	ObsLog                  []ObsLog          `json:"obs_log"`
	ObsPlaced               int               `json:"obs_placed"`
	PartyID                 int               `json:"party_id"`
	PermanentBuffs          []PermanentBuff   `json:"permanent_buffs"`
	HeroVariant             int               `json:"hero_variant"`
	Pings                   int               `json:"pings"`
	Purchase                map[string]int    `json:"purchase"`
	PurchaseLog             []PurchaseLog     `json:"purchase_log"`
	RunePickups             int               `json:"rune_pickups"`
	Runes                   map[string]int    `json:"runes"`
	RunesLog                []RunesLog        `json:"runes_log"`
	Sen                     map[string]int    `json:"sen"`
	SenLeftLog              []SenLeftLog      `json:"sen_left_log"`
	SenLog                  []SenLog          `json:"sen_log"`
	SenPlaced               int               `json:"sen_placed"`
	Stuns                   int               `json:"stuns"`
	Times                   []int             `json:"times"`
	TowerDamage             int               `json:"tower_damage"`
	XpPerMin                int               `json:"xp_per_min"`
	XpReasons               map[string]int    `json:"xp_reasons"`
	XpT                     []int             `json:"xp_t"`
	PersonaName             string            `json:"personaname"`
	Name                    string            `json:"name"`
	LastLogin               time.Time         `json:"last_login"`
	RadiantWin              bool              `json:"radiant_win"`
	StartTime               int               `json:"start_time"`
	Duration                int               `json:"duration"`
	Cluster                 int               `json:"cluster"`
	LobbyType               int               `json:"lobby_type"`
	GameMode                int               `json:"game_mode"`
	Patch                   int               `json:"patch"`
	Region                  int               `json:"region"`
	IsRadiant               bool              `json:"isRadiant"`
	Win                     int               `json:"win"`
	Lose                    int               `json:"lose"`
	TotalGold               int               `json:"total_gold"`
	TotalXp                 int               `json:"total_xp"`
	KillsPerMin             int               `json:"kills_per_min"`
	KDA                     float64           `json:"kda"`
	Abandons                int               `json:"abandons"`
	NeutralKills            int               `json:"neutral_kills"`
	TowerKills              int               `json:"tower_kills"`
	CourierKills            int               `json:"courier_kills"`
	LaneKills               int               `json:"lane_kills"`
	HeroKills               int               `json:"hero_kills"`
	ObserverKills           int               `json:"observer_kills"`
	SentryKills             int               `json:"sentry_kills"`
	RoshanKills             int               `json:"roshan_kills"`
	NecronomiconKills       int               `json:"necronomicon_kills"`
	AncientKills            int               `json:"ancient_kills"`
	BuybackCount            int               `json:"buyback_count"`
	ObserverUses            int               `json:"observer_uses"`
	SentryUses              int               `json:"sentry_uses"`
	LaneEfficiency          float64           `json:"lane_efficiency"`
	LaneEfficiencyPct       float64           `json:"lane_efficiency_pct"`
	Lane                    int               `json:"lane"`
	LaneRole                int               `json:"lane_role"`
	IsRoaming               bool              `json:"is_roaming"`
	PurchaseTime            map[string]int    `json:"purchase_time"`
	FirstPurchaseTime       map[string]int    `json:"first_purchase_time"`
	ItemWin                 map[string]int    `json:"item_win"`
	ItemUsage               map[string]int    `json:"item_usage"`
	PurchaseTPScroll        int               `json:"purchase_tpscroll"`
	ActionsPerMin           int               `json:"actions_per_min"`
	LifeStateDead           int               `json:"life_state_dead"`
	RankTier                int               `json:"rank_tier"`
	Cosmetics               []CosmeticItem    `json:"cosmetics"`
	Benchmarks              map[string]int    `json:"benchmarks"`
	NeutralTokensLog        []NeutralTokenLog `json:"neutral_tokens_log"`
}

type Unit struct {
	// Can be expanded based on data specifics
}

type BuybackLog struct {
	Time       int `json:"time"`
	Slot       int `json:"slot"`
	PlayerSlot int `json:"player_slot"`
}

type ConnectionLog struct {
	Time       int    `json:"time"`
	Event      string `json:"event"`
	PlayerSlot int    `json:"player_slot"`
}

type KillsLog struct {
	Time int    `json:"time"`
	Key  string `json:"key"`
}

type ObsLeftLog struct {
	// Can be expanded based on data specifics
}

type ObsLog struct {
	// Can be expanded based on data specifics
}

type PermanentBuff struct {
	// Can be expanded based on data specifics
}

type PurchaseLog struct {
	Time    int    `json:"time"`
	Key     string `json:"key"`
	Charges int    `json:"charges"`
}

type RunesLog struct {
	Time int `json:"time"`
	Key  int `json:"key"`
}

type SenLeftLog struct {
	// Can be expanded based on data specifics
}

type SenLog struct {
	// Can be expanded based on data specifics
}

type CosmeticItem struct {
	ItemID          int       `json:"item_id"`
	Name            string    `json:"name"`
	Prefab          string    `json:"prefab"`
	CreationDate    time.Time `json:"creation_date"`
	ImageInventory  string    `json:"image_inventory"`
	ImagePath       string    `json:"image_path"`
	ItemDescription string    `json:"item_description"`
	ItemName        string    `json:"item_name"`
	ItemRarity      string    `json:"item_rarity"`
	ItemTypeName    string    `json:"item_type_name"`
	UsedByHeroes    string    `json:"used_by_heroes"`
}

type NeutralTokenLog struct {
	Time int    `json:"time"`
	Key  string `json:"key"`
}
