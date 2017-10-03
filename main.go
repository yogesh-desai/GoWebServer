package main

import(

	"fmt"
	"io/ioutil"
	"net/http"
	
)

type Page struct {

	Title string
	Body []byte
}

func (p *Page) save() error {

	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {

	filename:= title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {

		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/"):]
	p, err := loadPage(title)
	if err != nil {
		strErr := "\nThe Status Code:" + fmt.Sprint(http.StatusNotFound) + ".  File Not Found."
		http.Error(w, err.Error() + strErr , http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {

	//	http.HandleFunc("/api/v1/getfile/", apiHandler)
	http.HandleFunc("/", apiHandler)
	http.ListenAndServe(":8080", nil)

}
