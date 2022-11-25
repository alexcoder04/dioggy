package main

import (
	"os"
	"path"
	"strings"
)

func GetGithubClone() string {
	if clone := os.Getenv("GITHUB_CLONE"); clone != "" {
		return clone
	}
	return "alexcoder04/if-schleife-bot"
}

func GetExecCommand() []string {
	if cmd := os.Getenv("EXEC_COMMAND"); cmd != "" {
		return strings.Split(cmd, " ")
	}
	return []string{"npm", "start"}
}

func GetPrepareCommand() []string {
	if cmd := os.Getenv("PREPARE_COMMAND"); cmd != "" {
		return strings.Split(cmd, " ")
	}
	return []string{"npm", "install"}
}

func GetRepoFolder() string {
	return path.Base(GITHUB_CLONE)
}
