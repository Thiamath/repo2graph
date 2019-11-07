package controller

import (
	"github.com/Thiamath/repo2graph/controller/handlers"
	"github.com/Thiamath/repo2graph/entities"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

func GetDiagram(credentials map[string]string) (diagram *entities.Graph, err *entities.Error) {
	ghToken := credentials[handlers.TOKEN]

	repositories, err := handlers.GetRepositories(ghToken)
	if err != nil {
		log.Error(err)
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
	// FIRST PASS
	grpcNodes := make(map[string]entities.Node, 0)
	serviceNodes := make(map[string]entities.Node, 0)
	for _, node := range nodes {
		matched, _ := regexp.MatchString("grpc-.+-go", node.Id)
		if matched {
			grpcNodes[node.Id] = node
		} else {
			serviceNodes[node.Id] = node
		}
	}

	// SECOND PASS
	for _, node := range grpcNodes {
		submatch := strings.Split(node.Id, "-")
		serviceName := strings.Join(submatch[1:len(submatch)-1], "-")
		switch {
		case submatch[0] == "grpc" && submatch[len(submatch)-1] == "go": // grpc proto
			service, exists := serviceNodes[serviceName]
			if exists {
				edges = append(edges, entities.Edge{
					From:   node.Id,
					To:     service.Id,
					Arrows: "to;middle",
				})
			} else {
				edges = append(edges, entities.Edge{
					From:   node.Id,
					To:     "system-model",
					Arrows: "to;middle",
				})
			}
		}
	}
	return edges
}
