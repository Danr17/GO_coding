package main

import (
	"net/http"
	"sync"
	"text/template"
)

func htmlPage(pinger *WebPing) http.HandlerFunc {
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
		})

		retrieve := pinger.sites

		data := ListItemPage{
			PageTitle: "Web TCP Ping",
			Items:     retrieve,
		}

		tmpl.Execute(w, data)
	}

}
