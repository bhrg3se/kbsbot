package handlers

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
	"kbsbot/handlers/catan"
	"strings"
)

const PREFIX = "--"

func AddRoutes(s *discordgo.Session) {
	router := exrouter.New()

	router.On("greet", Greet)
	router.On("check", catan.GetColonistProfile)

	// Match the regular expression user(name)?
	//router.OnMatch("username", dgrouter.NewRegexMatcher("/^(hello)/i"), func(ctx *exrouter.Context) {
	//	ctx.Reply("Your username is " + ctx.Msg.Author.Username)
	//})

	router.On("sub", nil).
		On("sub2", func(ctx *exrouter.Context) {
			ctx.Reply("sub2 called with arguments:\n", strings.Join(ctx.Args, ";"))
		}).
		On("sub3", func(ctx *exrouter.Context) {
			ctx.Reply("sub3 called with arguments:\n", strings.Join(ctx.Args, ";"))
		})

	router.Default = router.On("help", func(ctx *exrouter.Context) {
		var f func(depth int, r *exrouter.Route) string
		f = func(depth int, r *exrouter.Route) string {
			text := ""
			for _, v := range r.Routes {
				text += strings.Repeat("  ", depth) + v.Name + " : " + v.Description + "\n"
				text += f(depth+1, &exrouter.Route{Route: v})
			}
			return text
		}
		ctx.Reply("```" + f(0, router) + "```")
	}).Desc("prints this help menu")

	// Add message handler
	s.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		router.FindAndExecute(s, PREFIX, s.State.User.ID, m.Message)
	})

}
