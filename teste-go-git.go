package main

import (
	"strings"

	"github.com/go-git/go-git/v5"
)

func main() {
	getStatus("./")
}

func getStatus(path string) (string, error) {

	repo, err := git.PlainOpen(path)
	if err == git.ErrRepositoryNotExists {
		return "Repository not found", nil
	}

	if err != nil {
		return err.Error(), err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err.Error(), err
	}

	status, err := worktree.Status()
	if err != nil {
		return err.Error(), err
	}

	arrayStatus := strings.Split(status.String(), "\n")

	println(arrayStatus)

	return "Finalized!", nil
}
