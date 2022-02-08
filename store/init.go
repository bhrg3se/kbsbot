package store

import (
	"crypto/rsa"
	"github.com/bwmarrin/discordgo"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"kbsbot/models"
)

type Store interface {
	GetConfig() models.Config
	GetJWTPrivateKey() *rsa.PrivateKey
	GetJWTPublicKey() *rsa.PublicKey
	GetDiscordSession() *discordgo.Session
}

var State Store

// NewRealStore creates a new store with all dependencies like database
func NewRealStore(config models.Config) RealStore {

	dg, err := discordgo.New("Bot " + viper.GetString("BOT.APITOKEN"))
	if err != nil {
		logrus.Fatal(err)
	}
	err = dg.Open()
	if err != nil {
		logrus.Fatal(err)
	}

	return RealStore{
		dg:     dg,
		config: config,
	}
}
