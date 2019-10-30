package internal

import (
	"context"
	"encoding/json"
	r2g "github.com/Thiamath/repo2graph/github"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	ctx context.Context
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("This is an example server.\n"))
}

func getNodes(w http.ResponseWriter, r *http.Request) {
	ghToken := r.URL.Query().Get("GITHUB_TOKEN")

	ghClient := r2g.GetNewClientFromToken(ghToken)

	getReposCtx, getReposCancel := context.WithTimeout(ctx, time.Minute)
	repositories, err := r2g.GetRepositories(ghClient, getReposCtx)
	getReposCancel()
	if err != nil {
		log.Error(err)
	}

	nodes := r2g.CraftNodes(repositories)
	render, err := json.Marshal(nodes)
	if err != nil {
		log.Error(err)
	}
	_, _ = w.Write(render)
}

func StartServer() {
	ctx = context.Background()
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/getNodes", getNodes)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
