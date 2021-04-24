package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

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
	flag.StringVar(&keyPath, "key-path", filepath.Join(os.Getenv("HOME"), ".ssh", "deploy"), "Key path, default is ~/.ssh/deploy/")
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

	generateKey(keyPath, repoAlias, force)

	sshConfig := fmt.Sprintf(
		"Host %s\n    Hostname %s\n    IdentityFile=%s\n\n",
		repoAlias, repo.Host, getKeyFileName(keyPath, repoAlias)+".id_rsa",
	)
	err = updateSSHConfig(sshConfig)
	if err != nil {
		fmt.Printf("Key generated, buy error occur while update ssh config: %s", err.Error())
		return
	}
	fmt.Printf(
		"Deploy key generated, the public key is stored in %s\n\nYour new repo url is: %s@%s:%s/%s.git\n",
		getKeyFileName(keyPath, repoAlias)+".id_rsa.pub", repo.SshUser, repoAlias, repo.Owner, repo.Name,
	)
}

func file_exist(filename string) (bool, error) {
	if _, err := os.Stat(filename); err == nil {
		// file exist
		return true, nil
	} else if os.IsNotExist(err) {
		// file not exist
		return false, nil
	} else {
		return false, err
	}
}

func getKeyFileName(_keyPath, _repoAlias string) string {
	return filepath.Join(_keyPath, _repoAlias)
}

func generateKey(_keyPath, _repoAlias string, _force bool) {
	keypath_ex, err := file_exist(_keyPath)
	if err != nil {
		exitWithMessage(err.Error())
	}
	if !keypath_ex {
		os.Mkdir(_keyPath, 0700)
	}

	privateKeyFile := getKeyFileName(_keyPath, _repoAlias) + ".id_rsa"
	publicKeyFile := privateKeyFile + ".pub"
	prikey_ex, err := file_exist(privateKeyFile)
	if err != nil {
		exitWithMessage(err.Error())
	}
	if prikey_ex && !_force {
		exitWithMessage("Private key exist. You can use -f to overwrite it")
	}
	pubkey_ex, err := file_exist(publicKeyFile)
	if err != nil {
		exitWithMessage(err.Error())
	}
	if pubkey_ex && !_force {
		exitWithMessage("Public key exist. You can use -f to overwrite it")
	}

	kp, err := utils.NewKeyPair(bitSize)
	if err != nil {
		exitWithMessage(err.Error())
	}
	err = kp.WriteToFile(privateKeyFile)
	if err != nil {
		exitWithMessage("Error occur while saving key file: " + err.Error())
	}
}

func updateSSHConfig(config string) error {
	sshConfigFile := filepath.Join(os.Getenv("HOME"), ".ssh", "config")
	f, err := os.OpenFile(sshConfigFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(config); err != nil {
		return err
	}

	return nil
}
