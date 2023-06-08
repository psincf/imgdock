package main

import (
	"context"
	"time"

	g "github.com/AllenDang/giu"
)

var docker_info DockerInfo

func update_loop(ctx context.Context) {
	for {
		time.Sleep(time.Duration(time.Duration.Seconds(1)))

		docker_info.update(ctx)
		g.Update()
	}
}

func run() {
	run_gui()
}

func main() {
	ctx := context.Background()
	docker_info = new_docker_info()

	go update_loop(ctx)

	window := g.NewMasterWindow("Docker", 640, 360, 0)
	window.Run(run)
}
