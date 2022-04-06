package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"regexp"
)

func Contains(a []string, x string) int {
    for index, n := range a {
        if x == n {
            return index
        }
    }
    return -1
}

func GetFilesNames() []string {
    files, err := ioutil.ReadDir("/go/files")
    if err != nil {
        log.Fatal(err)
    }

    files_names := make([]string, 0)
    for _, file := range files {
    	files_names = append(files_names, file.Name())
    }

    return files_names
}

func GetFilesNamesWithoutExt(files_names []string) []string {
	files_names_ext := make([]string, 0)
    for _, file_name := range files_names {
    	regex := regexp.MustCompile("[.][a-z]+")
    	file_name = regex.ReplaceAllString(file_name, "")
    	files_names_ext = append(files_names_ext, file_name)
    }
    return files_names_ext
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	files_names := GetFilesNames()
	files_names_ext := GetFilesNamesWithoutExt(files_names)
	index_file := Contains(files_names_ext, r.URL.Path[1:])

	if r.URL.Path[1:] == "" {
		_, _ = fmt.Fprintf(w, "<h1>Hi, there!</h1>")

	    for index, name := range files_names_ext {
	    	_, _ = fmt.Fprintf(w, "<a href='/%s'>%s</a><br>", name, files_names[index])
	    }

	} else if index_file != -1 {
		http.ServeFile(w, r, "/go/files/" + files_names[index_file])
	} else if r.URL.Path[1:] == "error" {
		_, _ = fmt.Fprintf(w, "<h1>404 Error :(</h1>")
	} else {
		http.Redirect(w, r, "/error", http.StatusMovedPermanently)
	}
}

func main() {
	fmt.Printf("Server started")
	http.HandleFunc("/", helloWorldHandler)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}
