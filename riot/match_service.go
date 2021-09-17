package riot

import (
	"github.com/yucealiosman/riot-cli/pkg"
)

type MatchDetail struct {
	GameID                int64                 `json:"gameId"`
	PlatformID            PlatformID            `json:"platformId"`
	GameCreation          int64                 `json:"gameCreation"`
	GameDuration          int64                 `json:"gameDuration"`
	QueueID               int64                 `json:"queueId"`
	MapID                 int64                 `json:"mapId"`
	SeasonID              int64                 `json:"seasonId"`
	GameVersion           string                `json:"gameVersion"`
	GameMode              string                `json:"gameMode"`
	GameType              string                `json:"gameType"`
	Teams                 []Team                `json:"teams"`
	Participants          []Participant         `json:"participants"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
}

type ParticipantIdentity struct {
	ParticipantID int64  `json:"participantId"`
	Player        Player `json:"player"`
}

type Player struct {
	PlatformID        PlatformID `json:"platformId"`
	AccountID         string     `json:"accountId"`
	SummonerName      string     `json:"summonerName"`
	SummonerID        string     `json:"summonerId"`
	CurrentPlatformID PlatformID `json:"currentPlatformId"`
	CurrentAccountID  string     `json:"currentAccountId"`
	MatchHistoryURI   string     `json:"matchHistoryUri"`
	ProfileIcon       int64      `json:"profileIcon"`
}

type Participant struct {
	TeamID                   int64         `json:"teamId"`
	Spell1ID                 int64         `json:"spell1Id"`
	Spell2ID                 int64         `json:"spell2Id"`
	ChampionID               int64         `json:"championId"`
	ProfileIconID            int64         `json:"profileIconId"`
	SummonerName             string        `json:"summonerName"`
	Bot                      bool          `json:"bot"`
	SummonerID               string        `json:"summonerId"`
	GameCustomizationObjects []interface{} `json:"gameCustomizationObjects"`
	Perks                    Perks         `json:"perks"`
	Stats                    Stats         `json:"stats"`
	Timeline                 Timeline      `json:"timeline"`
}

type Stats struct {
	ParticipantID                   int64 `json:"participantId"`
	Win                             bool  `json:"win"`
	Item0                           int64 `json:"item0"`
	Item1                           int64 `json:"item1"`
	Item2                           int64 `json:"item2"`
	Item3                           int64 `json:"item3"`
	Item4                           int64 `json:"item4"`
	Item5                           int64 `json:"item5"`
	Item6                           int64 `json:"item6"`
	Kills                           int64 `json:"kills"`
	Deaths                          int64 `json:"deaths"`
	Assists                         int64 `json:"assists"`
	LargestKillingSpree             int64 `json:"largestKillingSpree"`
	LargestMultiKill                int64 `json:"largestMultiKill"`
	KillingSprees                   int64 `json:"killingSprees"`
	LongestTimeSpentLiving          int64 `json:"longestTimeSpentLiving"`
	DoubleKills                     int64 `json:"doubleKills"`
	TripleKills                     int64 `json:"tripleKills"`
	QuadraKills                     int64 `json:"quadraKills"`
	PentaKills                      int64 `json:"pentaKills"`
	UnrealKills                     int64 `json:"unrealKills"`
	TotalDamageDealt                int64 `json:"totalDamageDealt"`
	MagicDamageDealt                int64 `json:"magicDamageDealt"`
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`
	TrueDamageDealt                 int64 `json:"trueDamageDealt"`
	LargestCriticalStrike           int64 `json:"largestCriticalStrike"`
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions"`
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`
	TotalHeal                       int64 `json:"totalHeal"`
	TotalUnitsHealed                int64 `json:"totalUnitsHealed"`
	DamageSelfMitigated             int64 `json:"damageSelfMitigated"`
	DamageDealtToObjectives         int64 `json:"damageDealtToObjectives"`
	DamageDealtToTurrets            int64 `json:"damageDealtToTurrets"`
	VisionScore                     int64 `json:"visionScore"`
	TimeCCingOthers                 int64 `json:"timeCCingOthers"`
	TotalDamageTaken                int64 `json:"totalDamageTaken"`
	MagicalDamageTaken              int64 `json:"magicalDamageTaken"`
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`
	TrueDamageTaken                 int64 `json:"trueDamageTaken"`
	GoldEarned                      int64 `json:"goldEarned"`
	GoldSpent                       int64 `json:"goldSpent"`
	TurretKills                     int64 `json:"turretKills"`
	InhibitorKills                  int64 `json:"inhibitorKills"`
	TotalMinionsKilled              int64 `json:"totalMinionsKilled"`
	NeutralMinionsKilled            int64 `json:"neutralMinionsKilled"`
	NeutralMinionsKilledTeamJungle  int64 `json:"neutralMinionsKilledTeamJungle"`
	NeutralMinionsKilledEnemyJungle int64 `json:"neutralMinionsKilledEnemyJungle"`
	TotalTimeCrowdControlDealt      int64 `json:"totalTimeCrowdControlDealt"`
	ChampLevel                      int64 `json:"champLevel"`
	VisionWardsBoughtInGame         int64 `json:"visionWardsBoughtInGame"`
	SightWardsBoughtInGame          int64 `json:"sightWardsBoughtInGame"`
	WardsPlaced                     int64 `json:"wardsPlaced"`
	WardsKilled                     int64 `json:"wardsKilled"`
	FirstBloodKill                  bool  `json:"firstBloodKill"`
	FirstBloodAssist                bool  `json:"firstBloodAssist"`
	FirstTowerKill                  bool  `json:"firstTowerKill"`
	FirstTowerAssist                bool  `json:"firstTowerAssist"`
	CombatPlayerScore               int64 `json:"combatPlayerScore"`
	ObjectivePlayerScore            int64 `json:"objectivePlayerScore"`
	TotalPlayerScore                int64 `json:"totalPlayerScore"`
	TotalScoreRank                  int64 `json:"totalScoreRank"`
	PlayerScore0                    int64 `json:"playerScore0"`
	PlayerScore1                    int64 `json:"playerScore1"`
	PlayerScore2                    int64 `json:"playerScore2"`
	PlayerScore3                    int64 `json:"playerScore3"`
	PlayerScore4                    int64 `json:"playerScore4"`
	PlayerScore5                    int64 `json:"playerScore5"`
	PlayerScore6                    int64 `json:"playerScore6"`
	PlayerScore7                    int64 `json:"playerScore7"`
	PlayerScore8                    int64 `json:"playerScore8"`
	PlayerScore9                    int64 `json:"playerScore9"`
	Perk0                           int64 `json:"perk0"`
	Perk0Var1                       int64 `json:"perk0Var1"`
	Perk0Var2                       int64 `json:"perk0Var2"`
	Perk0Var3                       int64 `json:"perk0Var3"`
	Perk1                           int64 `json:"perk1"`
	Perk1Var1                       int64 `json:"perk1Var1"`
	Perk1Var2                       int64 `json:"perk1Var2"`
	Perk1Var3                       int64 `json:"perk1Var3"`
	Perk2                           int64 `json:"perk2"`
	Perk2Var1                       int64 `json:"perk2Var1"`
	Perk2Var2                       int64 `json:"perk2Var2"`
	Perk2Var3                       int64 `json:"perk2Var3"`
	Perk3                           int64 `json:"perk3"`
	Perk3Var1                       int64 `json:"perk3Var1"`
	Perk3Var2                       int64 `json:"perk3Var2"`
	Perk3Var3                       int64 `json:"perk3Var3"`
	Perk4                           int64 `json:"perk4"`
	Perk4Var1                       int64 `json:"perk4Var1"`
	Perk4Var2                       int64 `json:"perk4Var2"`
	Perk4Var3                       int64 `json:"perk4Var3"`
	Perk5                           int64 `json:"perk5"`
	Perk5Var1                       int64 `json:"perk5Var1"`
	Perk5Var2                       int64 `json:"perk5Var2"`
	Perk5Var3                       int64 `json:"perk5Var3"`
	PerkPrimaryStyle                int64 `json:"perkPrimaryStyle"`
	PerkSubStyle                    int64 `json:"perkSubStyle"`
	StatPerk0                       int64 `json:"statPerk0"`
	StatPerk1                       int64 `json:"statPerk1"`
	StatPerk2                       int64 `json:"statPerk2"`
}

type Timeline struct {
	ParticipantID               int64        `json:"participantId"`
	CreepsPerMinDeltas          PerMinDeltas `json:"creepsPerMinDeltas"`
	XPPerMinDeltas              PerMinDeltas `json:"xpPerMinDeltas"`
	GoldPerMinDeltas            PerMinDeltas `json:"goldPerMinDeltas"`
	CSDiffPerMinDeltas          PerMinDeltas `json:"csDiffPerMinDeltas"`
	XPDiffPerMinDeltas          PerMinDeltas `json:"xpDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     PerMinDeltas `json:"damageTakenPerMinDeltas"`
	DamageTakenDiffPerMinDeltas PerMinDeltas `json:"damageTakenDiffPerMinDeltas"`
	Role                        Role         `json:"role"`
	Lane                        Lane         `json:"lane"`
}

type PerMinDeltas struct {
	The010 float64 `json:"0-10"`
}

type Team struct {
	TeamID               int64  `json:"teamId"`
	Win                  string `json:"win"`
	FirstBlood           bool   `json:"firstBlood"`
	FirstTower           bool   `json:"firstTower"`
	FirstInhibitor       bool   `json:"firstInhibitor"`
	FirstBaron           bool   `json:"firstBaron"`
	FirstDragon          bool   `json:"firstDragon"`
	FirstRiftHerald      bool   `json:"firstRiftHerald"`
	TowerKills           int64  `json:"towerKills"`
	InhibitorKills       int64  `json:"inhibitorKills"`
	BaronKills           int64  `json:"baronKills"`
	DragonKills          int64  `json:"dragonKills"`
	VilemawKills         int64  `json:"vilemawKills"`
	RiftHeraldKills      int64  `json:"riftHeraldKills"`
	DominionVictoryScore int64  `json:"dominionVictoryScore"`
	Bans                 []Ban  `json:"bans"`
}

type Ban struct {
	ChampionID int64 `json:"championId"`
	PickTurn   int64 `json:"pickTurn"`
}

type PlatformID string

const (
	Euw1          PlatformID = "EUW1"
	PlatformIDTr1 PlatformID = "tr1"
	Tr1           PlatformID = "TR1"
)

type Lane string

const (
	None Lane = "NONE"
)

type Role string

const (
	Duo        Role = "DUO"
	DuoSupport Role = "DUO_SUPPORT"
)

func GetMatch(client pkg.Client, matchId string) (*MatchDetail, error) {
	path := "match/v4/matches/" + matchId

	var matchDetail MatchDetail
	error := GetService(client, nil, path, &matchDetail)
	return &matchDetail, error
}

type Match struct {
	GameId     int    `json:"gameId"`
	Role       string `json:"role"`
	Season     int    `json:"season"`
	PlatformId string `json:"platformId"`
	Champion   int    `json:"champion"`
	Queue      int    `json:"queue"`
	Lane       string `json:"lane"`
	TimeStamp  int    `json:"timeStamp"`
}

type MatchList struct {
	MatchList []Match `json:"matches"`
}

func GetMatchList(client pkg.Client, accountId string) (*MatchList, error) {
	params := map[string]string{"endIndex": "10"}
	path := "match/v4/matchlists/by-account/" + accountId
	var matchListResponse MatchList
	error := GetService(client, params, path, &matchListResponse)
	return &matchListResponse, error
}
