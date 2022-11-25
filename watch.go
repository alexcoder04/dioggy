package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WatchSignal() {
	for {
		channel := make(chan os.Signal, 1)
		signal.Notify(channel, syscall.Signal(syscall.SIGUSR1))

		log.Println("[signal] Watching for SIGUSR1 to update")
		<-channel

		if UpdateRunning {
			continue
		}

		log.Println("[signal] Got signal to update manually")
		err := RunUpdate("signal")
		if err != nil {
			log.Println("[signal] Update failed")
			SendToDiscord("[signal] Update failed")
			continue
		}
	}
}

func WatchUpdates() {
	for {
		local, err := LocalLatestCommit()
		if err != nil {
			log.Println("[updater] Failed to determine latest local commit")
			log.Printf("[updater] Error: %s\n", err.Error())
			log.Println("[updater] WARNING: stopping update thread!")
			SendToDiscord("[updater] WARNING: cannot determine latest local commit, stopping update thread!")
			return
		}

		remote, err := RemoteLatestCommit()
		if err != nil {
			log.Println("[updater] Failed to determine latest remote commit")
			log.Printf("[updater] Error: %s\n", err.Error())
			log.Println("[updater] Retrying to check remote in 5 minutes...")
			time.Sleep(time.Minute * 5)
			continue
		}

		if local != remote {
			log.Println("[updater] Latest local and remote commits doesn't match, updating...")

			err := RunUpdate("updater")
			if err != nil {
				log.Println("[updater] Retrying to update in 5 minutes...")
				SendToDiscord("[updater] Update failed")
				time.Sleep(time.Minute * 5)
				continue
			}
		}

		log.Println("[updater] Checking for next update in 24 hours...")
		time.Sleep(time.Hour * 24)
	}
}

func RunUpdate(context string) error {
	UpdateRunning = true

	err := Stop()
	if err != nil {
		UpdateRunning = false
		log.Printf("[%s] Failed to stop the process for update", context)
		log.Printf("[%s] Error: %s\n", context, err.Error())
		return err
	}

	err = UpdateRepo()
	if err != nil {
		UpdateRunning = false
		CommunicationChannel <- true

		log.Printf("[%s] Failed to run git pull", context)
		log.Printf("[%s] Error: %s\n", context, err.Error())
		return err
	}

	log.Printf("[%s] Update successfull", context)
	SendToDiscord(fmt.Sprintf("[%s] Update successful", context))

	UpdateRunning = false
	CommunicationChannel <- true

	return nil

}
