package utils

import (
	"log"
	"os"
	"path"
)

func GetAbsDir(a ...string) string {
	p, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// this sucks need better way to get abs path to base package
	if p[len(p)-len(PARENT_PACKAGE):] != PARENT_PACKAGE {
		p = path.Dir(p)
	}
	for _, v := range a {
		p = path.Join(p, v)
	}
	return p
}
