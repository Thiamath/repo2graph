package internal

import (
	"context"
	r2g "github.com/Thiamath/repo2graph/github"
	"github.com/google/go-github/github"
	"log"
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
	ghToken := r.URL.Query()["GITHUB_TOKEN"]
	ghClient := github.NewClient(oauthClient)
	getReposCtx, getReposCancel := context.WithTimeout(ctx, time.Minute)
	repositories := r2g.GetRepositories(ghClient, ctx)
	getReposCancel()
	nodes := r2g.CraftNodes(repositories)
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
