package handlers

import (
	"github.com/sirupsen/logrus"
	"kbsbot/store"
	"kbsbot/utils"
	"net/http"
)

func Ring(w http.ResponseWriter, r *http.Request) {
	//username:=r.URL.Query().Get("role")
	dg := store.State.GetDiscordSession()

	_, err := dg.ChannelMessageSend("927994006317723688", "<@&940145568582352966>, somebody is at the poker table. Go join them:\n https://gather.town/app/27JfyAUjgcySoLDj/Hostel%20Janta")
	//_,err:=dg.ChannelMessageSend("926761202200215608","<@OPanda>, somebody is at the poker table. Go join them:\n https://gather.town/app/27JfyAUjgcySoLDj/Hostel%20Janta")
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponseHTML(w, "could not send message")
		return
	}

	utils.SuccessResponseHTML(w, "Successfully notified the gamblers")
}
