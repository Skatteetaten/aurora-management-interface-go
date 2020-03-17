package main

import (
	"fmt"
	management "github.com/skatteetaten/aurora-management-interface-go"
	"github.com/skatteetaten/aurora-management-interface-go/env"
	"github.com/skatteetaten/aurora-management-interface-go/example/dummywebserver"
	"github.com/skatteetaten/aurora-management-interface-go/health"
	"net/http"
)

func main() {
	initWebServer()
	http.ListenAndServe(":8080", nil)
}

func initWebServer() {
	// Setting up dummy webserver
	dummywebserver.InitDummyWebServer()

	// Setting up the managementinterface
	managementInterfaceHandler := InitManagementHandler()
	managementInterfaceHandler.StartHTTPListener()
}

// InitManagementHandler initializes the managementinterface with /health and /env endpoints
func InitManagementHandler() *management.RoutingHandler {
	managementHandler := management.CreateRoutingHandler()

	// You probably want to write an application specific health retriever unless you always want it to return with Status: UP
	applicationHealthRetriever := health.GetDefaultHealthRetriever()

	// DefaultApplicationEnvHandler works for basic cases. Override for application specific environment
	applicationEnvRetriever := env.GetDefaultEnvRetriever()
	applicationEnvRetriever.SetKeysToMask([]string{"SomeEnvKeyToMask"})

	// The info endpoint has no default implementation so far. This is a simple stub example, returning an empty JSON.
	applicationInfoHandler := emptyInfoHandler{}

	managementHandler.RouteApplicationHealthRetriever(applicationHealthRetriever)
	managementHandler.RouteApplicationEnvRetriever(applicationEnvRetriever)
	managementHandler.RouteEndPointToHandlerFunc(management.Info, applicationInfoHandler.ServeHTTP)

	return managementHandler
}

type emptyInfoHandler struct{}

func (eih emptyInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	returnJSON := "{}"

	w.Header().Set("Content-Type", "application/vnd.spring-boot.actuator.v3+json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "%s", returnJSON)
}
