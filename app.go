package main

import (
        "log"
        "net/http"
        "net"
        "fmt"
        "os"
        "time"
        "github.com/gorilla/mux"
)

//REST API Section
func apiServer(){
        fmt.Printf("[API Server Section]\n")
        router := mux.NewRouter().StrictSlash(true)
        router.HandleFunc("/", apiServerStatus)
        router.HandleFunc("/lookup/{query}", ipLookUp).Methods("GET")
        log.Fatal(http.ListenAndServe(":28080", router))
        fmt.Println("["+time.Now().Format(time.RFC3339)+"]API Server Started")
}
func apiServerStatus(w http.ResponseWriter, r *http.Request){
        message :="["+time.Now().Format(time.RFC3339)+"][API Server Is Online]\n"
        fmt.Fprintf(w,message)
}

func ipLookUp(w http.ResponseWriter, r *http.Request)
        vars := mux.Vars(r)
        lookUpQuery := vars["query"]
        ips, err := net.LookupIP(lookUpQuery)
        if err != nil {
                fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
                fmt.Fprintf(w,"No IP Found\n")
        }
        for _, ip := range ips {
                fmt.Println("%s\n", ip.String())
                fmt.Fprintf(w,ip.String()+"\n")
        }
}

func main() {
    fmt.Printf("[Application Started]\n")
    apiServer()
    fmt.Printf("[Application Exit]\n")
}
{
