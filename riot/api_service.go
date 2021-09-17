package riot

import (
	"riot-cli/pkg"
)

type Summoner struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconId int    `json:"profileIconId"`
	SummonerLevel int    `json:"summonerLevel"`
	RevisionDate  int    `json:"revisionDate"`
}

func GetSummonerByName(client *pkg.Client, sum_name string) (*Summoner, error) {
	path := "summoner/v4/summoners/by-name/" + sum_name
	var sum Summoner
	error := GetService(client, nil, path, &sum)
	return &sum, error
}

type ChampionMastery struct {
	ChampionID                   int64  `json:"championId"`
	ChampionLevel                int64  `json:"championLevel"`
	ChampionPoints               int64  `json:"championPoints"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int64  `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	TokensEarned                 int64  `json:"tokensEarned"`
	SummonerID                   string `json:"summonerId"`
}

func GetChampionMasteriesBySummonerId(client *pkg.Client, sumId string) (*[]ChampionMastery, error) {
	path := "champion-mastery/v4/champion-masteries/by-summoner/" + sumId

	var championMasteryList []ChampionMastery
	error := GetService(client, nil, path, &championMasteryList)
	return &championMasteryList, error
}

type Game struct {
	GameID            int64            `json:"gameId"`
	MapID             int64            `json:"mapId"`
	GameMode          string           `json:"gameMode"`
	GameType          string           `json:"gameType"`
	GameQueueConfigID int64            `json:"gameQueueConfigId"`
	Participants      []Participant    `json:"participants"`
	Observers         Observers        `json:"observers"`
	PlatformID        string           `json:"platformId"`
	BannedChampions   []BannedChampion `json:"bannedChampions"`
	GameStartTime     int64            `json:"gameStartTime"`
	GameLength        int64            `json:"gameLength"`
}

type BannedChampion struct {
	ChampionID int64 `json:"championId"`
	TeamID     int64 `json:"teamId"`
	PickTurn   int64 `json:"pickTurn"`
}

type Observers struct {
	EncryptionKey string `json:"encryptionKey"`
}

type Perks struct {
	PerkIDS      []int64 `json:"perkIds"`
	PerkStyle    int64   `json:"perkStyle"`
	PerkSubStyle int64   `json:"perkSubStyle"`
}

func GetCurrentGameBySummonerId(client *pkg.Client, sumId string) (*Game, error) {
	params := map[string]string{"endIndex": "10"}
	path := "spectator/v4/active-games/by-summoner/" + sumId
	var game Game
	error := GetService(client, params, path, &game)
	return &game, error
}
