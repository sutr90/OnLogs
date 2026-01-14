package docker

import (
	"context"

	"github.com/moby/moby/client"
)

type DockerService struct {
    Client *client.Client
}

type ContainerNamesResult struct {
	Name string
	ID string
}

func (s *DockerService) GetContainerNames(ctx context.Context) ([]ContainerNamesResult, error) {
    containers, err := s.Client.ContainerList(ctx, client.ContainerListOptions{All: true})
    if err != nil {
        return nil, err
    }
    
    var res []ContainerNamesResult
    for _, c := range containers.Items {
		res = append(res, ContainerNamesResult{Name: c.Names[0], ID: c.ID})
    }
    return res, nil
}


func (s *DockerService) GetContainerImageNameByContainerID(ctx context.Context, containerID string) (string, error) {
    res, err := s.Client.ContainerInspect(ctx, containerID, client.ContainerInspectOptions{})
    if err != nil {
        return "", err
    }
    
    return res.Container.Image, nil
}
