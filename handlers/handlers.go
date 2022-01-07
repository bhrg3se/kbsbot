package handlers

import (
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
)

func Greet(ctx *exrouter.Context) {
	ctx.Reply(fmt.Sprintf("Hello %s", ctx.Msg.Author))
}
