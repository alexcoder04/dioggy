package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alexcoder04/friendly/v2"
)

func CloneRepo() error {
	err := friendly.Run([]string{"git", "clone", fmt.Sprintf("https://github.com/%s.git", GITHUB_CLONE)}, "")
	if err != nil {
		return err
	}
	return PrepareRepo()
}

func UpdateRepo() error {
	err := friendly.Run([]string{"git", "pull", "--rebase"}, GetRepoFolder())
	if err != nil {
		return err
	}
	return PrepareRepo()
}

func PrepareRepo() error {
	log.Println("Running prepare command")
	return friendly.Run(PREPARE_COMMAND, GetRepoFolder())
}

func LocalLatestCommit() (string, error) {
	data, err := friendly.GetOutput([]string{"git", "log", "-n", "1", "--pretty=format:%H", "main"}, GetRepoFolder())
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func RemoteLatestCommit() (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/commits/main", GITHUB_CLONE))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	jsonData := RepositoryCommit{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return "", err
	}
	return *jsonData.SHA, nil
}

func WatchUpdates() {
	for {
		local, err := LocalLatestCommit()
		if err != nil {
			log.Println("[updater] Failed to determine latest local commit")
			log.Printf("[updater] Error: %s\n", err.Error())
			log.Println("[updater] WARNING: stopping update thread!")
			// TODO send signal to process
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

			UpdateRunning = true

			err := Stop()
			if err != nil {
				UpdateRunning = false
				log.Println("[updater] Failed to stop the process for update")
				log.Printf("[updater] Error: %s\n", err.Error())
				log.Println("[updater] Retrying to update in 5 minutes...")
				time.Sleep(time.Minute * 5)
				continue
			}

			err = UpdateRepo()
			if err != nil {
				UpdateRunning = false
				CommunicationChannel <- true

				log.Println("[updater] Failed to run git pull")
				log.Printf("[updater] Error: %s\n", err.Error())
				log.Println("[updater] Retrying to update in 5 minutes...")
				time.Sleep(time.Minute * 5)
				continue
			}

			log.Println("[updater] Update successfull, restarting soon")

			UpdateRunning = false
			CommunicationChannel <- true
		}

		log.Println("[updater] Checking for next update in 24 hours...")
		time.Sleep(time.Hour * 24)
	}
}
