package catan

import (
	"encoding/json"
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strings"
)

func GetColonistProfile(ctx *exrouter.Context) {
	var username = ""
	if len(ctx.Args) > 1 {
		username = ctx.Args[1]
	}
	if strings.TrimSpace(username) == "" {
		username = ctx.Msg.Author.Username
	}
	profile, err := getProfile(username)
	if err != nil {
		logrus.Error(err)
		ctx.Reply("Sorry! Something is wrong. \n Maybe the username is incorrect?")
		return
	}
	ctx.Reply(fmt.Sprintf("Username: %s\nKarma: %s\nTotal Games: %d \nWins Percent: %s",
		profile.Username,
		profile.Karma,
		profile.TotalGames,
		profile.WinPercent,
	))
}

func getProfile(username string) (*respStruct, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://colonist.io/api/profile/"+username, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(req.URL)

	var r respStruct
	err = json.NewDecoder(resp.Body).Decode(&r)

	return &r, err
}

type respStruct struct {
	Username     string        `json:"username"`
	Karma        string        `json:"karma"`
	PresalePacks []interface{} `json:"presalePacks"`
	Achievements []interface{} `json:"achievements"`
	HideData     bool          `json:"hideData"`
	GameDatas    []struct {
		Players []struct {
			UserID          string `json:"userId"`
			Username        string `json:"username"`
			AccessLevel     int    `json:"accessLevel"`
			Rank            int    `json:"rank"`
			Points          int    `json:"points"`
			Finished        bool   `json:"finished"`
			QuitWithPenalty bool   `json:"quitWithPenalty"`
			IsHuman         bool   `json:"isHuman"`
			PlayerColor     int    `json:"playerColor"`
		} `json:"players"`
		Setting struct {
			ModeSetting        int    `json:"modeSetting"`
			ExtensionSetting   int    `json:"extensionSetting"`
			ScenarioSetting    int    `json:"scenarioSetting"`
			MapSetting         int    `json:"mapSetting"`
			DiceSetting        int    `json:"diceSetting"`
			VictoryPointsToWin int    `json:"victoryPointsToWin"`
			KarmaActive        bool   `json:"karmaActive"`
			CardDiscardLimit   int    `json:"cardDiscardLimit"`
			MaxPlayers         int    `json:"maxPlayers"`
			GameSpeed          string `json:"gameSpeed"`
			BotSpeed           string `json:"botSpeed"`
			HideBankCards      bool   `json:"hideBankCards"`
			FriendlyRobber     bool   `json:"friendlyRobber"`
			Version            int    `json:"version"`
			ID                 string `json:"id"`
			GameType           int    `json:"gameType"`
			PrivateGame        bool   `json:"privateGame"`
			NetworkActive      bool   `json:"networkActive"`
		} `json:"setting"`
		ID        string `json:"id"`
		Finished  bool   `json:"finished"`
		TurnCount int    `json:"turnCount"`
		StartTime string `json:"startTime"`
		Duration  string `json:"duration"`
	} `json:"gameDatas"`
	WinsInLast100Games                 int    `json:"winsInLast100Games"`
	TotalFinishedGamesInLast100Games   int    `json:"totalFinishedGamesInLast100Games"`
	TotalUnfinishedGamesInLast100Games int    `json:"totalUnfinishedGamesInLast100Games"`
	PointsInLast100Games               int    `json:"pointsInLast100Games"`
	TotalGames                         int    `json:"totalGames"`
	WinPercent                         string `json:"winPercent"`
	PointsPerGame                      string `json:"pointsPerGame"`
	Items                              []struct {
		ActiveInStore bool   `json:"activeInStore"`
		Category      int    `json:"category"`
		Type          int    `json:"type"`
		Name          string `json:"name"`
		Description   string `json:"description,omitempty"`
		Price         int    `json:"price"`
		ImagePath     int    `json:"imagePath"`
	} `json:"items"`
}
