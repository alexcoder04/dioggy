package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"

	"github.com/alexcoder04/friendly/v2/ffiles"
)

func ProcessPid() (int, error) {
	pidBytes, err := ioutil.ReadFile(GetPidFile())
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(pidBytes))
}

func GetPidFile() string {
	rd, err := ffiles.GetRuntimeDir()
	if err != nil {
		panic("Failed to get runtime dir: " + err.Error())
	}
	return path.Join(rd, fmt.Sprintf("dioggy-process-%d.pid", os.Getpid()))
}

func ProcessRunning() bool {
	if !ffiles.IsFile(GetPidFile()) {
		return false
	}

	pid, err := ProcessPid()
	if err != nil {
		return false
	}

	_, err = os.FindProcess(pid)
	return err == nil
}

func CreatePidFile(pid int) error {
	return ffiles.WriteNewFile(GetPidFile(), fmt.Sprintf("%d", pid))
}

func DeletePidFile() error {
	return os.Remove(GetPidFile())
}

func Stop() error {
	log.Println("Stopping process")

	pid, err := ProcessPid()
	if err != nil {
		return err
	}

	err = syscall.Kill(pid, syscall.SIGTERM)
	if err != nil {
		return err
	}
	err = DeletePidFile()
	if err != nil {
		return err
	}
	return nil
}

func Run() error {
	cmd := exec.Command(EXEC_COMMAND[0], EXEC_COMMAND[1:]...)
	cmd.Dir = GetRepoFolder()
	cmd.Env = append(os.Environ(), fmt.Sprintf("DIOGGY_PID=%d", os.Getpid()))

	err := cmd.Start()
	if err != nil {
		return err
	}

	err = CreatePidFile(cmd.Process.Pid)
	if err != nil {
		return err
	}

	log.Println("Process started successfully")

	err = cmd.Wait()
	if err.Error() == "signal: terminated" {
		return nil
	}

	if ffiles.IsFile(GetPidFile()) {
		err := DeletePidFile()
		if err != nil {
			log.Println("Warning: failed to delete pid file")
			log.Printf("Error: %s\n", err.Error())
			return err
		}
	}
	return err
}
