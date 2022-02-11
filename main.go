package main

import (
	"github.com/OpenCal-FYDP/PreferenceManagement/internal/service"
	"github.com/OpenCal-FYDP/PreferenceManagement/internal/storage"
	"github.com/OpenCal-FYDP/PreferenceManagement/rpc"
	"log"
	"net/http"
)

func main() {
	svc := service.New(storage.New())
	server := rpc.NewPreferenceManagementServiceServer(svc)
	log.Fatal(http.ListenAndServe(":8080", server))
}
