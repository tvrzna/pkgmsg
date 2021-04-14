package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"github.com/tvrzna/pkgtray/checker"
)

const version = "0.1.0"

var conf *config

func main() {
	conf = loadConfig()

	for true {
		go doRefresh(conf)
		time.Sleep(time.Duration(conf.intervalSeconds) * time.Second)
	}
}

func doRefresh(conf *config) {
	log.Print("refresh")

	updates := checker.CheckPackages(conf.pkgmanager.pkgType, conf.pkgmanagerPath)

	var msg string
	if updates > 0 {
		msg = strconv.Itoa(updates) + " updates available."
	} else if updates == 0 {
		msg = "System is up to date."
	} else {
		msg = "Checking for updates failed."
	}

	log.Println(msg)

	if updates > 0 {
		showMessage(msg)
	}
}

func showMessage(msg string) {
	path, err := exec.LookPath("notify-send")
	if err != nil {
		log.Print(err)
		return
	}

	_, err = exec.Command(path, "pkgmsg", msg).Output()

	if err != nil {
		log.Print(err)
	}
}

func getVersion() string {
	return fmt.Sprintf("pkgmsg %s\nhttps://github.com/tvrzna/pkgmsg\n\nReleased under the MIT License.", version)
}
