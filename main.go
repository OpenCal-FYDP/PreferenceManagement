package main

import (
	"github.com/OpenCal-FYDP/PreferenceManagement/internal/preferencemanagementservice"
	"github.com/OpenCal-FYDP/PreferenceManagement/rpc"
	"log"
	"net/http"
)

func main() {
	service := preferencemanagementservice.New()
	server := rpc.NewPreferenceManagementServiceServer(service)
	log.Fatal(http.ListenAndServe(":8080", server))
}
