package main

import (
	"log"

	"github.com/robdimsdale/wl"
	"github.com/robdimsdale/wl/logger"
	"github.com/robdimsdale/wl/oauth"
)

type Todos map[uint]wl.Task

func main() {
	client := oauth.NewClient(
		"my_access_token",
		"my_client_id",
		wl.APIURL,
		logger.NewLogger(logger.INFO),
	)

	tasks, err := client.TasksForListID(32000000)
	if err != nil {
		log.Fatal(err)
	}

	todos := map[uint]wl.Task{}
	for _, task := range tasks {
		todos[task.ID] = task
	}

	pos, err := client.TaskPositionsForListID(32000000)
	if err != nil {
		log.Fatal(err)
	}

	slack(pos[0].Values, Todos(todos))
}
