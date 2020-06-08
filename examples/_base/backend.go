package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Backend)
	http.ListenAndServe(":8080", nil)
}

func Backend(w http.ResponseWriter, r *http.Request) {
	var standard = fmt.Sprintf("http://%s%s", r.Host, r.URL.Path)
	var forward = "//nowhere"
	var quintessence = standard
	if r.Header.Get("X-Forwarded-Host") != "" {
		forward = fmt.Sprintf("%s://%s%s", r.Header.Get("X-Forwarded-Proto"), r.Header.Get("X-Forwarded-Host"), r.URL.Path)
		quintessence = forward
	}
	w.Header().Set("Location", quintessence)

	fmt.Fprintln(w, "<!DOCTYPE html>")
	fmt.Fprintln(w, "<head></head>")
	fmt.Fprintln(w, "<body>")

	fmt.Fprintln(w, "According to standard headers, I am at: ")
	fmt.Fprintf(w, "<a href=\"%s\">%s</a><br />", standard, standard)

	fmt.Fprintln(w, "According to forward headers, I am at: ")
	fmt.Fprintf(w, "<a href=\"%s\">%s</a><br />", forward, forward)

	fmt.Fprintln(w, "I thus set the Location header to: ")
	fmt.Fprintf(w, "<pre><code>Location: %s</code></pre>", quintessence)
	fmt.Fprintln(w, "Without setting a redirection HTTP status code, luckily this won't have any effect. But you can view the actual Location header your browser received by using its console.")

	fmt.Fprintln(w, "</body></html>")
}
