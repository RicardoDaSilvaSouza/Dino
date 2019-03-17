package dynowebportal

import (
	"fmt"
	"log"
	"net/http"
)

// RunWebPortal starts running the dino web portal on address parameter
func RunWebPortal(address string) error {
	http.HandleFunc("/", rootHandler)
	log.Printf("Starting Dino Web Portal at %s ...", address)
	return http.ListenAndServe(address, nil)
}

func rootHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "This is the Dino Web Portal %s", req.RemoteAddr)
}
