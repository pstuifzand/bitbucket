package main

import "github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
import "os"
import "log"

import "github.com/davecgh/go-spew/spew"

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	r, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	remotes, err := r.Remotes()

    for _, remote := range remotes {
        spew.Dump(remote)
    }

	return
}
