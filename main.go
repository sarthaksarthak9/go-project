package main

import (
	"fmt"
	"encoding/json"
    "log"
	"net/http"
    "time"
	details "go-microservices/details"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
    response := map[string]string{
        "status" : "up",
        
        "timestamp" : time.Now().String(),
    }
    json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")	
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching the details")
	hostname, err := details.GetHostname()
	if err != nil{
		panic(err)
	}
	IP, _ := details.GetIp()
	fmt.Println(hostname, IP)
	response := map[string]string{
        "hostname" : hostname,
        
        "ip" : IP.String(),
    }
    json.NewEncoder(w).Encode(response)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	
	log.Println("Server has started")
	log.Fatal(http.ListenAndServe(":80", r))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request)
// func main() {
// 	http.HandleFunc("/", rootHandler)

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println("web server has started")
// 	http.ListenAndServe(":80", nil)
// }

// package main

// import (
//     "fmt"
//     "rsc.io/quote"
// )

// func rectProps(length, width float64)(perimeter, area float64){    //to return two values
//     perimeter = 2 * (length + width)
//     area = length * width

//     return // perimeter, area
// }

// func main() {
//     var x int = 10
//     var y, z = 2, 3
//     t := 10
//     name := "devOps"
//     check := true

//     fmt.Println("Hello, World!")
//     fmt.Println(quote.Go())
//     fmt.Println(x, y, z, t, name, check)

//     p, a := rectProps(4, 5)

//     fmt.Printf("The perimeter is %f and area is %f", p, a)
// }
