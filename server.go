package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/*p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))*/
	http.HandleFunc("/view/", viewHandler)       //route /view handled by viewHandler function
	log.Fatal(http.ListenAndServe(":8080", nil)) // listen on port 8080, if err log.Fatal handles
}

type Page struct {
	Title string
	Body  []byte //io libraries expect byte type
}

func (p *Page) save() error { //returns an error because ioutil.WriteFile return type is error
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) //0600 means create with read-write permissions for current user only
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) //remember functions can return two values
	if err != nil {                        //handle if there's and error
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil // if no error
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]                         // take from the end of /view/ in the request url to end
	p, _ := loadPage(title)                                     //load a page with that title from the url
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body) //write the html with page strings to w (responseWriter)
}
