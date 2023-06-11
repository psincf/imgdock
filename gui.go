package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/docker/docker/api/types"
)

type SelectedContainer struct {
	container types.Container
	logs_str  string
}

var selected SelectedContainer

func list_containers() g.Layout {
	labels := make([]g.Widget, 10)
	rows := make([]*g.TableRowWidget, 0)

	for _, container := range docker_info.containers {
		c := container

		rows = append(rows, g.TableRow(
			g.Selectable(container.Names[0]).Flags(g.SelectableFlagsSpanAllColumns).OnClick(func() {
				selected.container = c
				selected.logs_str = ""

				fmt.Println(c.ID)
			}),
			g.Label(container.Status),
		))
	}

	labels = append(labels, g.Table().Columns(
		g.TableColumn("name"),
		g.TableColumn("status"),
	).Rows(rows...))

	return labels
}

func list_images() g.Layout {
	labels := make([]g.Widget, 10)
	rows := make([]*g.TableRowWidget, 0)

	for _, image := range docker_info.images {
		name := "None"
		if len(image.RepoTags) > 0 {
			name = image.RepoTags[0]
		}
		rows = append(rows, g.TableRow(
			g.Selectable(name).Flags(g.SelectableFlagsSpanAllColumns).OnClick(func() {

			}),
		))

	}

	labels = append(labels, g.Table().Columns(
		g.TableColumn("name"),
	).Rows(rows...))
	return labels
}

func list_volumes() g.Layout {
	labels := make([]g.Widget, 10)
	rows := make([]*g.TableRowWidget, 0)

	for _, volume := range docker_info.volumes.Volumes {
		rows = append(rows, g.TableRow(
			g.Selectable(volume.Name).Flags(g.SelectableFlagsSpanAllColumns).OnClick(func() {

			}),
		))
	}

	labels = append(labels, g.Table().Columns(
		g.TableColumn("name"),
	).Rows(rows...))

	return labels
}

func run_gui() {
	width, height := master_window.GetSize()
	flags_config := g.WindowFlagsNoResize | g.WindowFlagsNoMove | g.WindowFlagsNoCollapse
	g.Window("Containers").Pos(0, 0).Size(float32(width/4), float32(height/3)).Flags(flags_config).Layout(
		list_containers(),
	)
	g.Window("Images").Pos(0, float32(height/3)).Size(float32(width/4), float32(height/3)).Flags(flags_config).Layout(
		list_images(),
	)
	g.Window("Volumes").Pos(0, float32(2*height/3)).Size(float32(width/4), float32(height/3)).Flags(flags_config).Layout(
		list_volumes(),
	)

	if selected.container.ID != "" && selected.logs_str == "" {
		logs, err := docker_info.client.ContainerLogs(ctx, selected.container.ID, types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Timestamps: true,
			Follow:     true,
			Details:    true,
		})
		if err != nil {
			panic(err)
		}
		log_str := make([]byte, 1_000)
		n, err := logs.Read(log_str)
		if err != nil {
			panic(err)
		}
		if n != 0 {
			selected.logs_str = string(log_str[0:n])
		}
	}

	g.Window("Main").Pos(float32(width/4), 0).Size(float32(3*width/4), float32(height)).Flags(flags_config).Layout(
		g.CodeEditor().ShowWhitespaces(false).Text(selected.logs_str),
	)
}
