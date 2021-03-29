package main

import (
	"flag"
)

var bitSize int
var force bool
var keyPath string

func init() {
	flag.IntVar(&bitSize, "bits", 4096, "RSA key bits")
	flag.BoolVar(&force, "f", false, "Generate key anyway")
	flag.StringVar(&keyPath, "key-path", "", "Key path, default is ~/.ssh/deploy/")
}

func main() {
	flag.Parse()
}
