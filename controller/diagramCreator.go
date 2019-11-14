package controller

import (
	"github.com/Thiamath/repo2graph/controller/handlers"
	"github.com/Thiamath/repo2graph/entities"
	"github.com/pelletier/go-toml"
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

	// Split between grpc nodes and service nodes
	grpcNodes := make(map[string]entities.Node, 0)
	grpcTransient := make(map[string]string, 0)
	serviceNodes := make(map[string]entities.Node, 0)
	for _, node := range nodes {
		matched, _ := regexp.MatchString("grpc-.+-go", node.Id)
		if matched {
			grpcNodes[node.Id] = node
		} else {
			serviceNodes[node.Id] = node
		}
	}

	// Link gRPC nodes
	for _, node := range grpcNodes {
		submatch := strings.Split(node.Id, "-")
		serviceName := strings.Join(submatch[1:len(submatch)-1], "-")
		switch {
		case submatch[0] == "grpc" && submatch[len(submatch)-1] == "go": // grpc proto
			service, exists := serviceNodes[serviceName]
			if exists {
				//edges = append(edges, entities.Edge{
				//	From:   node.Id,
				//	To:     service.Id,
				//	Arrows: "to;middle",
				//})
				grpcTransient[node.Id] = service.Id
			} else {
				//edges = append(edges, entities.Edge{
				//	From:   node.Id,
				//	To:     "system-model",
				//	Arrows: "to;middle",
				//})
				grpcTransient[node.Id] = "system-model"
			}
		}
	}

	i := 1
	for _, node := range serviceNodes {
		log.Info("Revised ", i, " of ", len(serviceNodes))
		i++
		content, _ := handlers.GetFileFromRepository(ghToken, node.LinkedRepository, "/Gopkg.lock")
		tomlContent, _ := toml.Load(content)
		get := tomlContent.Get("projects")
		if get != nil {
			constraints := get.([]*toml.Tree)
			systemModel := false
			for _, constraint := range constraints {
				name := constraint.Get("name").(string)
				split := strings.Split(name, "/")
				if len(split) > 2 {
					//host := split[0]
					owner := split[1]
					repo := split[2]
					//version := constraint.Get("version")
					//log.Debug("name: ", name, "\thost: ", host, "\tversion: ", version)
					if owner == "nalej" {
						sink, transitive := grpcTransient[repo]
						if transitive {
							if sink == "system-model" {
								if !systemModel {
									edges = append(edges, entities.Edge{
										From:   node.Id,
										To:     sink,
										Arrows: "to;middle",
									})
									systemModel = sink == "system-model"
								}
							} else {
								edges = append(edges, entities.Edge{
									From:   node.Id,
									To:     sink,
									Arrows: "to;middle",
								})
							}
						} else {
							edges = append(edges, entities.Edge{
								From:   node.Id,
								To:     repo,
								Arrows: "to;middle",
							})
						}
					}
				}
			}
		}
	}

	return edges
}
