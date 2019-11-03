package controller

import (
	"github.com/Thiamath/repo2graph/controller/handlers"
	"github.com/Thiamath/repo2graph/entities"
	"github.com/sirupsen/logrus"
)

func GetDiagram(credentials map[string]string) (diagram *entities.Graph, err error) {
	ghToken := credentials[handlers.TOKEN]

	repositories, err := handlers.GetRepositories(ghToken)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	nodes := craftNodesFromRepositories(repositories)
	return &entities.Graph{
		Nodes: nodes,
		Edges: craftEdges(nodes, customEdgeTransform),
	}, nil
}

func craftNodesFromRepositories(repositories []*entities.Repository) (nodes []entities.Node) {
	nodes = make([]entities.Node, len(repositories))
	for ix, repository := range repositories {
		nodes[ix] = entities.Node{
			Id:    repository.Id,
			Label: repository.Name,
		}
	}
	return nodes
}

func craftEdges(nodes []entities.Node, transform func([]entities.Node) []entities.Edge) (edges []entities.Edge) {
	if transform != nil {
		return transform(nodes)
	}
	return nil
}

func customEdgeTransform(nodes []entities.Node) (edges []entities.Edge) {
	edges = make([]entities.Edge, 0)
	for _, node := range nodes {
		switch node.Id {
		case "nalej/grpc-application-go":
			edges = append(edges, entities.Edge{From: "nalej/grpc-application-go", To: "nalej/system-model"})
		}
	}
	return edges
}
