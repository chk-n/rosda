package dash

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type AdminServer struct {
	http httpServer
	mq   messageQueue
	db   persistenceMaster
}

func (a *AdminServer) RegisterRoutes() {
	a.http.Use(a.LoggerMiddleware)
	a.http.Use(a.AuthMiddleware)

	api := chi.NewRouter()

	// manage auth related tasks
	api.Route("/v1/auth", func(r chi.Router) {
		// creates new client token for internal api access
		r.Post("/clients", a.newClientHandler)
	})

	// manage service related tasks
	api.Route("/v1/services", func(r chi.Router) {
		// list all services, search filters in query parameter possible
		r.Get("/", a.getServicesHandler)
		// creates a new service with the dash config file
		r.Post("/", a.createServiceHandler)
		// updates a service with the dash config file
		r.Put("/", a.updateServiceHandler)
		// create new deployment with updated image
		r.Post("/deployments/{service_id}", a.deployServiceHandler)
	})

	api.Route("/v1/nodes", func(r chi.Router) {
		// list all nodes, search filters in query parameter possible
		r.Get("/", a.getNodesHandler)
	})

	api.Route("/v1/cluster", func(r chi.Router) {
		// returns health of cluster (running nodes, total number of node crashes), health of db
		r.Get("/health", a.getClusterHealthHandler)
		// returns used cpu, ram vs total, # of masters and slaves, latency (we might not have this data)
		r.Get("/metrics", a.getClusterMetricsHandler)
	})

	a.http.Mount("/api", api)
}

// ----------- //
// Middlewares //
// ----------- //

func (a *AdminServer) AuthMiddleware(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// Logs requests to persistence layer (ip, api key, user agent, path, timestamp)
func (a *AdminServer) LoggerMiddleware(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// -------- //
// Handlers //
// -------- //

func (a *AdminServer) newClientHandler(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (a *AdminServer) getServicesHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: parse query parameters
	// TODO: add query to fetch all services

}

func (a *AdminServer) deployHandler(w http.ResponseWriter, r *http.Request) {
	imgUrl := r.URL.Query().Get("image_url")

	parts := strings.Split(imgUrl, ":")
	if len(parts) != 3 {
		w.WriteHeader(400)
		w.Write([]byte("invalid url length: did you append image hash or version in format https://registry.com/image:version"))
		return
	} else if parts[0] != "https" {
		w.WriteHeader(400)
		w.Write([]byte("invalid url: container registry needs to be accessible via https"))
		return
	}

	// TODO: proto3 encode image url

	if err := a.mq.Publish(topicNewDeployment, imgUrl); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("check dash logs for more information"))
		return
	}

	w.WriteHeader(200)
}

func (a *AdminServer) imageDownloadHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: get image_id from path
	// TODO: read in db where file is stored
	panic("implement me")
	path := ""
	http.ServeFile(w, r, path)
}
