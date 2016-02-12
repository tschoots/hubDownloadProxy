package main 

import (
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"bytes"

)

func main() {
	var buffer bytes.Buffer
	files, _ := ioutil.ReadDir("./data")
	for _, f := range files {
		url := fmt.Sprintf("/%s", f.Name())
		file_path := fmt.Sprintf("data/%s", f.Name())
		fmt.Printf("url : %s\nfile_path : %s\n\n", url, file_path)
		http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, file_path)
			})
		buffer.WriteString(f.Name())
		buffer.WriteString("\n")
	}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    		w.Write(buffer.Bytes())
    	})

    // now see if a port environment variable is set to overide default port 8080
    var port string
    if port = fmt.Sprintf(":%s", os.Getenv("LISTEN_PORT")); strings.Compare(port, ":") == 0 {
    	port = ":8080"
    }
    fmt.Printf("port %s\n", port)

	if err := http.ListenAndServe(port, nil);  err != nil {
		log.Fatal("ListenAndServer: " , err)
	}

}

