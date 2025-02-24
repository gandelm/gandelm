package main

import (
	"github.com/gandelm/gandelm/internal/container/command"
)

func main() {
	com := command.Command{}
	com.HelmUnInstall("sample")
}
