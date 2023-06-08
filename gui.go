package main

import g "github.com/AllenDang/giu"

func list_containers() []g.Widget {
	labels := make([]g.Widget, 10)
	for _, container := range docker_info.containers {
		labels = append(labels, g.Label(container.ID))
	}
	return labels
}

func list_images() []g.Widget {
	labels := make([]g.Widget, 10)
	for _, image := range docker_info.images {
		labels = append(labels, g.Label(image.ID))
	}
	return labels
}

func list_volumes() []g.Widget {
	labels := make([]g.Widget, 10)
	for _, volume := range docker_info.volumes.Volumes {
		labels = append(labels, g.Label(volume.Name))
	}
	return labels
}

func run_gui() {
	g.Window("Containers").Pos(0, 0).Layout(
		g.Column(
			list_containers()...,
		),
	)
	g.Window("Images").Pos(100, 0).Layout(
		g.Column(
			list_images()...,
		),
	)
	g.Window("Volumes").Pos(100, 0).Layout(
		g.Column(
			list_volumes()...,
		),
	)
}
