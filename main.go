package main

import (
    "html/template"
    "net/http"
    //"log"
    "sync"
)

// PageData is a struct to hold dynamic data for the HTML templates
type PageData struct {
    Description string
}

var (
    templates *template.Template
    once sync.Once
)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/reserve", reserveHandler)
    http.HandleFunc("/club-information", clubInfoHandler)
    http.HandleFunc("/payments", paymentsHandler)
    http.HandleFunc("/documents", documentsHandler)
    http.HandleFunc("/custom.css", cssHandler)
    http.Handle("/pics/", http.StripPrefix("/pics", http.FileServer(http.Dir("./pics"))))
    http.ListenAndServe(":8080", nil)
}

func initTemplates() {
    templates = template.Must(template.ParseFiles("header.html", "home.html", "reserve.html", "club-information.html", "payments.html", "documents.html", "footer.html"))
}


func renderTemplate(w http.ResponseWriter, tmplName string, data PageData) {
    once.Do(initTemplates)

    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    err := templates.ExecuteTemplate(w, "header.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = templates.ExecuteTemplate(w, tmplName, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = templates.ExecuteTemplate(w, "footer.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func cssHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/css; charset=utf-8")
    http.ServeFile(w, r, "./custom.css")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    renderTemplate(w, "home.html", PageData{Description: "Welcome to home"})
}

func reserveHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "reserve.html", PageData{Description: "This is the Reserve page"})
}

func clubInfoHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "club-information.html", PageData{Description: "This is the Club Information page"})
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "payments.html", PageData{Description: "This is the Payments page"})
}

func documentsHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "documents.html", PageData{Description: "This is the Documents page"})
}



