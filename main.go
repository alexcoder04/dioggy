package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alexcoder04/friendly/v2/ffiles"
)

var (
	GITHUB_CLONE    string   = GetGithubClone()
	PREPARE_COMMAND []string = GetPrepareCommand()
	EXEC_COMMAND    []string = GetExecCommand()

	EnableDiscordNotifications *bool = flag.Bool("enable-discord-notifications", false, "send discord mesages")

	UpdateRunning        bool      = false
	CommunicationChannel chan bool = make(chan bool)
)

func init() {
	flag.Parse()
}

func main() {
	log.Println("Started the dioggy daemon")
	log.Printf("Watching process: %s - %s\n", GITHUB_CLONE, strings.Join(EXEC_COMMAND, " "))

	if ProcessRunning() {
		log.Println("Process already running, attempting to stop...")

		err := Stop()
		if err != nil {
			log.Println("Failed to stop the process for restart")
			log.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	}

	if !ffiles.IsDir(GetRepoFolder()) {
		log.Println("Repo not exists yet, attempting to clone...")
		SendToDiscord("Cloning repo")

		err := CloneRepo()
		if err != nil {
			log.Println("Failed to clone the repo")
			log.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	}

	log.Println("Starting update thread...")
	go WatchUpdates()

	log.Println("Starting signal thread...")
	go WatchSignal()

	for {
		if UpdateRunning {
			log.Println("Update running, waiting to finish...")
			<-CommunicationChannel
		}

		log.Println("Starting process")
		err := Run()
		if err != nil {
			log.Println("Process failed to run")
			log.Printf("Error: %s\n", err.Error())
			log.Println("Trying to restart in 10 seconds...")
			time.Sleep(time.Second * 10)
			continue
		}
	}
}
