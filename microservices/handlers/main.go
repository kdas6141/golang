package main

import (
	"fmt"
    "net/http"
    "log"
	"io/ioutil"
//    "github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
//			rw.WriteHeader(http.StatusBadRequest)
//			rw.Write([]byte("Oops"))
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
		log.Printf("Data %s\n", d);
	});

    http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
        log.Println("Goodbye World")
    });

	http.ListenAndServe(":9090", nil)
}


