package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Repo properties definition
type Repo struct {
	RawUrl  string
	Host    string
	Port    int
	Scheme  string
	SshUser string
	Owner   string
	Name    string
}

func (r *Repo) ParseFromUrl(url string) error {
	isMatch, parsedUrl := r.parseSSHUrl(url)
	if !isMatch {
		return fmt.Errorf("parse url failed: %s is not a ssh url", url)
	}

	var err error

	r.Host = parsedUrl["host"]
	r.Port, err = strconv.Atoi(parsedUrl["port"])
	if err != nil {
		r.Port = 22
	}
	r.Scheme = "ssh"
	r.SshUser = parsedUrl["user"]
	r.Owner, r.Name = parsePath(parsedUrl["path"])
	if r.Owner == "" || r.Name == "" {
		return fmt.Errorf("invalid url: missing owner or name part(s)")
	}
	r.Name = strings.Split(r.Name, ".git")[0]
	r.RawUrl = url

	return nil
}

func (r *Repo) parseSSHUrl(url string) (isMatch bool, urlMap map[string]string) {
	reg := regexp.MustCompile(`^(?P<user>.*?)@(?P<host>.*?):(?:(?P<port>.*?)/)?(?P<path>.*?/.*?)$`)

	match := reg.FindStringSubmatch(url)

	if len(match) == 0 {
		isMatch = false
		return
	}

	isMatch = true

	urlMap = make(map[string]string)
	for i, name := range reg.SubexpNames() {
		if i > 0 && i <= len(match) {
			urlMap[name] = match[i]
		}
	}
	return
}

// GetHostAlias return a alias of this repo
func (r *Repo) GetAlias(withoutOwner bool) string {
	if withoutOwner {
		return r.Host + "-" + r.Name
	}

	return r.Host + "-" + r.Owner + "-" + r.Name
}

func parsePath(path string) (owner, repo string) {
	parsed := strings.Split(path, "/")
	if len(parsed) != 2 {
		return
	}

	owner = parsed[0]
	repo = parsed[1]
	return
}
