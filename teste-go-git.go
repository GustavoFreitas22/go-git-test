package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func main() {
	msg, err := getStatus("../pathTest")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)
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

	println(status.String())

	return "Finalized!", nil
}
