package server

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"kbsbot/handlers"
	"os"
	"os/signal"
	"syscall"
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
	//store.State = store.NewRealStore(config)

	dg, err := discordgo.New("Bot " + viper.GetString("Bot.APIToken"))
	if err != nil {
		logrus.Fatal(err)
	}

	handlers.AddRoutes(dg)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
