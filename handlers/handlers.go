package handlers

import (
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/sirupsen/logrus"
	"kbsbot/handlers/catan"
	"strings"
)

func Greet(ctx *exrouter.Context) {
	ctx.Reply(fmt.Sprintf("Hello %s", ctx.Msg.Author))
}

func GetColonistProfile(ctx *exrouter.Context) {
	var username = ""
	if len(ctx.Args) > 1 {
		username = ctx.Args[1]
	}
	if strings.TrimSpace(username) == "" {
		username = ctx.Msg.Author.Username
	}
	profile, err := catan.GetProfile(username)
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
