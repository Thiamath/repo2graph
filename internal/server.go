package internal

import (
	"encoding/json"
	"github.com/Thiamath/repo2graph/controller"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func getDiagramData(w http.ResponseWriter, r *http.Request) {
	credentials := make(map[string]string)
	credentials["GITHUB_TOKEN"] = r.URL.Query().Get("GITHUB_TOKEN")

	diagram, _err := controller.GetDiagram(credentials)
	if _err != nil {
		log.Error(_err)
		render, _ := json.Marshal(_err)
		w.WriteHeader(_err.ErrorCode)
		_, _ = w.Write(render)
		return
	}

	render, err := json.Marshal(diagram)
	if err != nil {
		log.Error(err)
	}
	_, _ = w.Write(render)
}

func StartServer() {
	http.HandleFunc("/getDiagramData", getDiagramData)

	web := http.FileServer(http.Dir("internal/web/"))
	http.Handle("/", http.StripPrefix("/", web))

	staticFileServer := http.FileServer(http.Dir("internal/web/static/"))
	http.Handle("/internal/web/static/", http.StripPrefix("/internal/web/static/", staticFileServer))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
