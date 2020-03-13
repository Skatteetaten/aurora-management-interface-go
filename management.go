package management

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

// DefaultPort is the default port for the management interface
const DefaultPort = "8081"

// RoutingHandler is a router for the management interface
type RoutingHandler struct {
	host           string
	port           string
	managementMux  *http.ServeMux
	managementSpec *ManagementinterfaceSpec
}

// CreateRoutingHandler creates a router to handle the management interface requests on default port (8081)
func CreateRoutingHandler() (*RoutingHandler, error) {
	return CreateRoutingHandlerForPort(DefaultPort)
}

// CreateRoutingHandlerForPort creates a router to handle the management interface requests on a specific port
func CreateRoutingHandlerForPort(port string) (*RoutingHandler, error) {
	hostIP, err := getOwnHostIP()
	if err != nil {
		return nil, err
	}

	managementSpec := NewManagementinterfaceSpec()
	managementMux := http.NewServeMux()

	mrh := RoutingHandler{
		host:           hostIP.String(),
		port:           port,
		managementMux:  managementMux,
		managementSpec: managementSpec,
	}
	mrh.RouteEndPointHandlerFunc(Management, mrh.managementHandler)

	return &mrh, nil
}

func (mrh RoutingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mrh.managementMux.ServeHTTP(w, r)
}

// StartHTTPListener starts a http.ListenAndServe process for the router on the configured port
func (mrh RoutingHandler) StartHTTPListener() {
	managementPort := fmt.Sprintf(":%s", mrh.port)
	go http.ListenAndServe(managementPort, mrh)
}

// RouteEndPointHandlerFunc routes endpoint to a handlerfunc
func (mrh RoutingHandler) RouteEndPointHandlerFunc(endPointType EndPointType, handlerfunc func(http.ResponseWriter, *http.Request)) error {
	usedefaultpathstring := ""
	return mrh.RouteEndPointHandlerFuncToPath(endPointType, usedefaultpathstring, handlerfunc)
}

// RouteEndPointHandlerFuncToPath routes endpoint to a handlerfunc on a specified, non-default path
func (mrh RoutingHandler) RouteEndPointHandlerFuncToPath(endPointType EndPointType, path string, handlerfunc func(http.ResponseWriter, *http.Request)) error {
	endpoint, err := newEndPoint(endPointType, mrh.host, mrh.port, handlerfunc)
	if err != nil {
		return err
	}

	if path != "" {
		endpoint.setPath(path)
	}
	mrh.managementSpec.mapEndpoint(*endpoint)
	mrh.managementMux.HandleFunc(endpoint.path, handlerfunc)

	return nil
}

func (mrh RoutingHandler) managementHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	managementJSON, err := mrh.managementSpec.createManagementJSON()

	if err != nil {
		message := "Error while creating JSON for management interface"
		logrus.Errorf("%s: %s", message, err)
		w.WriteHeader(http.StatusInternalServerError)
		responseJSON, _ := json.Marshal(map[string]string{
			"error": message,
			"cause": fmt.Sprintf("%v", err),
		})
		_, _ = fmt.Fprintf(w, "%s", responseJSON)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "%s", managementJSON)
}

func getOwnHostIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logrus.Debugf("Error while retrieving own IP")
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}
