/* Simple Web Server to get the files and and report if file is not present.
 The tree structure if as follows,
.
├── api
│ └── v1
│     └── getfile
│         ├── file1.txt
│         ├── file2.txt
│         └── file3.txt
├── main.go
├── README.md
└── test.txt

Basically we have handler to check the dir location of api/v1/getfile.
Many ToDo yet to be implemented.


To run please visit URL:
go run main.go and then open any one URL in browser window.

1. http://localhost:8080/api/v1/getfile/<fileName>
2. http://localhost:8080/test 
*/

package main

import(

	"fmt"
	"io/ioutil"
	"net/http"
	
)

// Page is a simple struct to store page data
type Page struct {

	Title string
	Body []byte
}

// save saves/writes the file
func (p *Page) save() error {

	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage loads the file. Returns error if file is not present.
func loadPage(title string) (*Page, error) {

	filename:= title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {

		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// apiHandler handles the required path requests
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
