# litefeed
Litefeed: Go edition 🐹

A lightweight search tool for Instagram users (up to 50 results)
Featuring: htmx (nofancy css, flex all the way)

For now: Search and Profile pic are only supported
*more comming soon..*

## Why Use Litefeed?
Litefeed is designed for low resource and no bloat access to Instagram user information without unnecessary complexity.

# Setup:

1. *Log In*: Sign in to Instagram on the web
2. *Open Developer Tools*: Press F12 or right-click and select "Inspect"
3. *Navigate to the Network Tab*: Reload the page (Ctrl + R) to capture network activity.
4. *Filter Network Requests*: Click on the "Filter" button and select "Fetch/XHR"
5. *Perform a Search*: Type a keyword in the Instagram search bar.
6. *Find the Request*: Look for the `query` {;} JSON entry in the Network tab.
7. *Check Request Headers*: Click on the request and confirm that the headers include your search term.


*Example Request Headers*: 
- for Search:
  
`{"data":{"context":"blended","include_reel":"true","query":"apple","rank_token":"","search_surface":"web_top_search"},"hasQuery":true}`

- for Profile:
  
`{"id":"apple","render_surface":"PROFILE"}` 


8. Copy as cURL: Use the "Copy as cURL" option and convert it to Go HTTP requests at curlconverter.com
9. Format Data Strings: In your Go code, format the data strings for search and profile:

`data = strings.NewReader(\.....\ + query + \.....\) // for search`

`data = strings.NewReader(\.....\ + id + \.....\) // for profile`

here, interpret \ as ` backtick

# Usage:
Once the setup is complete, run the application with:
`$ go run main.go`


# TODO:
- Simplify Setup
- Implement suggested users below user's profile page.
- Include user bio and additional profile information.

# Contributions
Contributions are welcome! Feel free to open issues or submit pull requests. Your help is appreciated!
