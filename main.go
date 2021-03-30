package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/YianAndCode/github-deploy-keys-manager/utils"
)

var bitSize int
var force bool
var keyPath string
var repoUrl string
var repoAlias string

func init() {
	flag.IntVar(&bitSize, "bits", 4096, "RSA key bits")
	flag.BoolVar(&force, "f", false, "Generate key anyway")
	flag.StringVar(&keyPath, "key-path", "", "Key path, default is ~/.ssh/deploy/")
	flag.StringVar(&repoUrl, "repo", "", "Repo url(ssh)")
	flag.StringVar(&repoAlias, "alias", "", "Repo alias")
}

func exitWithMessage(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func exitWithTips(errMsgs ...string) {
	flag.PrintDefaults()
	fmt.Println()

	if len(errMsgs) > 0 {
		for _, msg := range errMsgs {
			fmt.Println("\033[31m" + msg + "\033[0m")
		}
	}

	os.Exit(1)
}

func main() {
	flag.Parse()

	if repoUrl == "" {
		exitWithTips("Repo url is required")
	}

	var repo utils.Repo
	err := repo.ParseFromUrl(repoUrl)
	if err != nil {
		exitWithMessage(err.Error())
	}

	if repoAlias == "" {
		repoAlias = repo.GetAlias(true)
	}
}
