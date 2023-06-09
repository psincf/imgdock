package main

import (
	g "github.com/AllenDang/giu"
)

func list_containers() g.Layout {
	labels := make([]g.Widget, 10)
	for _, container := range docker_info.containers {
		//labels = append(labels, g.Label(container.Names[0]))
		labels = append(labels, g.Row(
			g.Label(container.Names[0]),
			g.Label(container.Status),
		))
	}
	return labels
}

func list_images() g.Layout {
	labels := make([]g.Widget, 10)
	for _, image := range docker_info.images {
		//labels = append(labels, g.Label(image.ID))

		if len(image.RepoTags) > 0 {
			labels = append(labels, g.Label(image.RepoTags[0]))
		} else {
			labels = append(labels, g.Label("None"))
		}
	}
	return labels
}

func list_volumes() g.Layout {
	labels := make([]g.Widget, 10)
	for _, volume := range docker_info.volumes.Volumes {
		labels = append(labels, g.Label(volume.Name))
	}
	return labels
}

func run_gui() {
	width, height := master_window.GetSize()
	flags_config := g.WindowFlagsNoResize | g.WindowFlagsNoMove | g.WindowFlagsNoCollapse
	g.Window("Containers").Pos(0, 0).Size(float32(width), float32(height/3)).Flags(flags_config).Layout(
		list_containers(),
	)
	g.Window("Images").Pos(0, float32(height/3)).Size(float32(width), float32(height/3)).Flags(flags_config).Layout(
		list_images(),
	)
	g.Window("Volumes").Pos(0, 2*float32(height/3)).Size(float32(width), float32(height/3)).Flags(flags_config).Layout(
		list_volumes(),
	)
}
