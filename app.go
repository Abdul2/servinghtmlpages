package main

/*

based on Alex Edwards blog
http://www.alexedwards.net/blog/serving-static-sites-with-go

abdul 

2/9/2015

*/

import (
  "html/template"
  "log"
  "net/http"
  "path"
  "os"
)

func main() {

  //FileServer creates a handle to enable
  // static directory filesystem to 
  // to respond to http requests 	
  fs := http.FileServer(http.Dir("static"))
  
  //register the FileServer to respond to requests with 
  // with prefix /static/
  //since the static directory is at root of
  //project we strip /static/ from http request
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  //all requests other than /static/ will be picked up
  //by serveTemplate
  http.HandleFunc("/", serveTemplate)

  log.Println("Listening...")
  
  //nil means use the default handler
  //listen to port 3000
  http.ListenAndServe(":3000", nil)
}


//serveTemplate builds paths to layout template and
//template correspponding to request (example.html)
//bundles the templates into a set and executes them
func serveTemplate(w http.ResponseWriter, r *http.Request) {
  
  //define a clean file path 
  lp := path.Join("templates", "layout.html")
  fp := path.Join("templates", r.URL.Path)
  
  
  // Return a 404 if the template doesn't exist
  // Stat returns Fileinfo or error
  // Fileinfo is used to check if fb is refrencing a directory as
  //supposed to a file 
  
  info, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }
	// Return a 404 if the request is for a directory
  if info.IsDir() {
    http.NotFound(w, r)
    return
  }

  
  //Parse file creaates a new template consiting of the parsed content
  // of the two input templates. if creation of the new template fails, error is returned.	
  tmpl, err := template.ParseFiles(lp, fp)
  
  
  if err != nil {
    log.Println(err.Error())
    // Respond with error message and code
    http.Error(w, http.StatusText(500), 500)
    return
  }
  
  
  //render "layout" template and write out into http.ResponseWriter
  if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }


}