package main

import (
	"context"
	"fmt"
	"time"

	g "github.com/AllenDang/giu"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var client_docker *client.Client
var containers = make([]types.Container, 0)

func update_containers_loop() {
	for {
		time.Sleep(time.Duration(time.Duration.Seconds(1)))

		containers_temp, err := client_docker.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}

		containers = containers_temp
		g.Update()
	}
}

func list_label() []g.Widget {
	labels := make([]g.Widget, 10)
	for _, container := range containers {
		labels = append(labels, g.Label(container.ID))
	}
	return labels
}

func run() {
	g.SingleWindow().Layout(
		//g.Label("Docker"),
		list_label()...,
	)
}

func main() {
	fmt.Println("Hello, World!")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	client_docker = cli

	go update_containers_loop()

	containers_temp, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	containers = containers_temp

	window := g.NewMasterWindow("Docker", 640, 360, 0)
	window.Run(run)
}
