/* Simple Web Server to get the files and and report if file is not present.
 The tree structure if as follows,
.
├── api
│ └── v1
│     └── getfile
│         ├── file1.txt
│         ├── file2.txt
│         ├── file3.txt
│         └── Senorita.avi
├── main.go
├── README.md
└── test.txt

Basically we have handler to check the dir location of api/v1/getfile.
Many ToDo yet to be implemented.

To run please visit URL:
go run main.go and then open any one URL in browser window.

1. http://localhost:8080/api/v1/getfile/<fileName>
2. http://localhost:8080/test 
3. http://localhost:8080/api/v1/getfile/Senorita.avi
*/

package main

import(

	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
)

// Page is a simple struct to store page data
type Page struct {

	Title string
	Body []byte
}

// save saves/writes the file
// Currently not in use
func (p *Page) save() error {

	filename	:= p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage loads the file. Returns error if file is not present.
func loadPage(title string) (p *Page, err error) {

	p = &Page{}
	filename	:= title //+ ".txt"

	body, err := ioutil.ReadFile(filename)
	if err != nil {	return nil, err	}

	p.Title = title
	p.Body	= body
	//return &Page{Title: title, Body: body}, nil

	return
}

// apiHandler handles the required path requests
func apiHandler(w http.ResponseWriter, r *http.Request) {

	title	:= r.URL.Path[len("/"):]

	switch{

	case strings.HasSuffix(title, ".txt"):

		p, err := loadPage(title)
		if err != nil {
			strErr := "\nThe Status Code:" + fmt.Sprint(http.StatusNotFound) + ". File Not Found."
			http.Error(w, err.Error() + strErr , http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", strings.TrimPrefix(p.Title, "api/v1/getfile/"), p.Body)

	case strings.HasSuffix(title, ".avi"):

		fileName	:= strings.TrimPrefix(title, "api/v1/getfile/")

		// For Downloading the file
		w.Header().Add("Content-Disposition", "Attachment;" + fileName)
		// w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		// w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
		
		http.ServeFile(w, r, title)
	}
}

func main() {

	http.HandleFunc("/", apiHandler)
	http.ListenAndServe(":8080", nil)
}
