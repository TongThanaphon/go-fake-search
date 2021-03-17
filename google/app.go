package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"golang.org/x/talks/content/2016/applicative/google"
)

type response struct {
	Results []google.Result
	Elapsed time.Duration
}

func handleSearch(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)

	query := req.FormValue("q")
	if query == "" {
		http.Error(w, `missing "q" URL parameter`, http.StatusBadRequest)
		return
	}

	start := time.Now()
	results, err := google.Search(query)
	elapsed := time.Since(start)
	resp := response{results, elapsed}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch req.FormValue("output") {
	case "json":
		err = json.NewEncoder(w).Encode(resp)
	case "prettyjson":
		var b []byte
		b, err = json.MarshalIndent(resp, "", " ")
		if err == nil {
			_, err = w.Write(b)
		}
	default:
		err = responseTemplate.Execute(w, resp)
	}
}

func main() {
	http.HandleFunc("/search", handleSearch)
	fmt.Println("serving on http://localhost:8080/search")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

var responseTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results in {{.Elapsed}}</p>
</body>
</html>
`))
