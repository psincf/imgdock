package main

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type DockerInfo struct {
	client     *client.Client
	containers []types.Container
	images     []types.ImageSummary
	volumes    volume.ListResponse
}

func new_docker_info() DockerInfo {
	info := DockerInfo{}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	info.client = cli

	return info
}

func (info *DockerInfo) update(ctx context.Context) error {
	containers, err := info.client.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return err
	}

	info.containers = containers

	images, err := info.client.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return err
	}

	info.images = images

	volumes, err := info.client.VolumeList(ctx, volume.ListOptions{})
	if err != nil {
		return err
	}

	info.volumes = volumes

	return nil
}
