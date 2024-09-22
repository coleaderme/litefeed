package routes

import (
	"fmt"
	"html/template"
	"igpp/workers"
	"io"
	"net/http"
)

// Json Templates
// json for Profile `{"id":"%s","render_surface":"PROFILE"}`
// json for Search  `{"data":{"context":"blended","include_reel":"true","query":"%s","rank_token":"","search_surface":"web_top_search"},"hasQuery":true}`

// global client
var client = &http.Client{
	Timeout: 0, // to prevent immediate crash. reason?
}

type PageData struct {
	Message string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 1. data for the template
	data := PageData{
		Message: "LiteFeed 🐹",
	}

	// 2. Load the template file
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// 3. Send html
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	query := r.FormValue("search")
	fmt.Printf("CLIENT: %s\n", query)
	users, err := workers.Search(query, client)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Load the template file
	tmpl, err := template.ParseFiles("templates/grid.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// Send combined html
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// if id == "" {
	// 	http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
	// 	return
	// }
	fmt.Printf("RECEIVED: %s\n", id)
	user, err := workers.Profile(id, client)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// load tmpl
	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	// send
	if err := tmpl.Execute(w, user); err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	// some corner were cut to improve speed. //
	targetURL := r.URL.Query().Get("url")
	// SKIP CHECKING IF EMPTY
	// if targetURL == "" {
	// 	http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
	// 	return
	// }
	// fmt.Printf("proxy: %s\n", targetURL)
	// Make a request to the target URL
	resp, err := client.Get(targetURL)
	if err != nil {
		http.Error(w, "Failed to fetch URL", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Set the appropriate headers OR SKIP
	// for key, values := range resp.Header {
	// 	for _, value := range values {
	// 		w.Header().Add(key, value)
	// 	}
	// }
	// NO NEED TO EXPLICITLY CALL
	// Set the status code from the response
	// w.WriteHeader(resp.StatusCode)

	// Copy the response body to the ResponseWriter
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
