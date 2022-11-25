package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
