package server

import (
	"flag"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"kbsbot/handlers"
	"kbsbot/models"
	"kbsbot/store"
	"net/http"
	"os"
)

func StartServer() {
	writeToFile := flag.Bool("f", false, "write logs to file")
	flag.Parse()

	// parse config
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	viper.SetEnvPrefix("KBS")

	//set log level
	level, err := logrus.ParseLevel(viper.GetString("Logging.Level"))
	if err != nil {
		level = logrus.ErrorLevel
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)

	// create log file if enabled
	if *writeToFile {
		f, err := os.OpenFile("/var/log/kbsbot.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		logrus.SetOutput(f)
	}

	// create database connections and other dependencies
	store.State = store.NewRealStore(models.Config{})

	dg := store.State.GetDiscordSession()
	handlers.AddRoutes(dg)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	http.HandleFunc("/api/ring", handlers.Ring)
	http.ListenAndServe("0.0.0.0:3333", http.DefaultServeMux)

}
