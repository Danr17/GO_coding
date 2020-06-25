package webping

import (
	"log"
	"net/http"
	"sync"
	"text/template"
)

//HTMLPage is the handler for the main page
func HTMLPage(pinger *WebPing) http.HandlerFunc {
	var (
		init sync.Once
		tmpl *template.Template
		err  error
	)

	type ListItemPage struct {
		PageTitle string
		Items     []*website
	}

	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tmpl, err = template.ParseFiles("index.html")
			if err != nil {
				log.Fatalln("Couldn't parse the index.html file")
			}
		})

		retrieve := pinger.sites

		data := ListItemPage{
			PageTitle: "Web TCP Ping",
			Items:     retrieve,
		}

		tmpl.Execute(w, data)
	}

}
