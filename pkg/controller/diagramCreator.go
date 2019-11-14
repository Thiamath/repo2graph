package controller

import (
	handlers2 "github.com/Thiamath/repo2graph/pkg/controller/handlers"
	"github.com/Thiamath/repo2graph/pkg/entities"
	"github.com/pelletier/go-toml"
	log "github.com/sirupsen/logrus"
	"strings"
)

func GetDiagram(credentials map[string]string) (diagram *entities.Graph, err *entities.Error) {
	ghToken := credentials[handlers2.TOKEN]

	repositories, err := handlers2.GetRepositories(ghToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	nodes := craftNodesFromRepositories(repositories)
	return &entities.Graph{
		Nodes: nodes,
		Edges: craftEdges(ghToken, nodes, customEdgeTransform),
	}, nil
}

func craftNodesFromRepositories(repositories []*entities.Repository) (nodes []entities.Node) {
	nodes = make([]entities.Node, len(repositories))
	for ix, repository := range repositories {
		nodes[ix] = entities.Node{
			Id:               repository.Id,
			Label:            repository.Name,
			LinkedRepository: repository,
		}
	}
	return nodes
}

func craftEdges(ghToken string, nodes []entities.Node, transform func(string, []entities.Node) []entities.Edge) (edges []entities.Edge) {
	if transform != nil {
		return transform(ghToken, nodes)
	}
	return nil
}

func customEdgeTransform(ghToken string, nodes []entities.Node) (edges []entities.Edge) {
	edges = make([]entities.Edge, 0)

	i := 1
	for _, node := range nodes {
		log.Info("Revised ", i, " of ", len(nodes))
		i++
		content, _ := handlers2.GetFileFromRepository(ghToken, node.LinkedRepository, "/Gopkg.lock")
		tomlContent, _ := toml.Load(content)
		get := tomlContent.Get("projects")
		if get != nil {
			dependencies := get.([]*toml.Tree)
			for _, dependency := range dependencies {
				name := dependency.Get("name").(string)
				split := strings.Split(name, "/")
				if len(split) > 2 {
					//host := split[0]
					//owner := split[1]
					repo := split[2]
					//version := dependency.Get("version")
					//log.Debug("name: ", name, "\thost: ", host, "\tversion: ", version)
					edges = append(edges, entities.Edge{
						From:   node.Id,
						To:     repo,
						Arrows: "to;middle",
					})
				}
			}
		}
	}

	return edges
}
