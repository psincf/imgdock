package main

import (
	"context"
	"time"

	g "github.com/AllenDang/giu"
)

var ctx context.Context
var docker_info DockerInfo
var master_window *g.MasterWindow

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
	ctx = context.Background()
	docker_info = new_docker_info()

	go update_loop(ctx)

	master_window = g.NewMasterWindow("Docker", 1280, 720, 0)
	master_window.Run(run)
}
