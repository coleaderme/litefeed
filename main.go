// HUGE Thanks to this handy tool https://transform.tools/json-to-go

package main

import (
	"fmt"
	"igpp/routes"
	"net/http"
)

func main() {
	// static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// routes
	http.HandleFunc("/", routes.IndexHandler)
	http.HandleFunc("/search", routes.SearchHandler)
	http.HandleFunc("/proxy", routes.ProxyHandler)
	http.HandleFunc("/profile", routes.ProfileHandler)

	port := ":3000" // visible to all devices on same network
	fmt.Printf("Listening: http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
	
	
}
