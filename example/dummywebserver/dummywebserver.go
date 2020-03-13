package dummywebserver

import (
	"fmt"
	"net/http"
)

// initDummyWebServer initializes dummy web server for example
func InitDummyWebServer() {
	routeHandler := createRouter()
	http.Handle("/", routeHandler)
}

func createRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", roothandler)
	return router
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Dummmyserver says hi at %s!", r.Host)
}
