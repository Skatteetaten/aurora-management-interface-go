package main

import (
	"github.com/sirupsen/logrus"
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
	managementHandler, err := management.CreateRoutingHandler()
	if err != nil {
		logrus.Fatalf("Error while creating management interface router: %s", err)
	}
	// Override health.ApplicationHealthHandler.Statusfunc with application specific health check
	appHealthHandlerStatusFunc := health.ApplicationHealthHandler{Statusfunc: health.DefaultHealthHandlerStatusFunc}

	// DefaultApplicationEnvHandler works for basic cases. Override for application specific environment
	appEnvHandler := env.DefaultApplicationEnvHandler()
	appEnvHandler.SetKeysToMask([]string{"SomeEnvKeyToMask"})

	managementHandler.RouteEndPointHandlerFunc(management.Health, appHealthHandlerStatusFunc.ServeHTTP)
	managementHandler.RouteEndPointHandlerFunc(management.Env, appEnvHandler.ServeHTTP)

	return managementHandler
}
