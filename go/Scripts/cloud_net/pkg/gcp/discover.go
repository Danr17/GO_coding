package gcp

import (
	"log"
	"fmt"
	"context"
	"google.golang.org/api/option"
	raw "google.golang.org/api/compute/v1"
)


//Discover retrieve a subset of information from the GCP project
//that can be used to list/watch various network resources
//
//The discovered resources are saved in a local file for persistence, 
//to lower the number of requests.
//The file is rewritten every time the command is run
func Discover(projectID string) {
	ctx := context.Background()
	opts := []option.ClientOption{}

	client, err := NewClient(ctx, opts...)
	if err != nil {
		panic("could not create the client")
	}

	//resp, err := client.svc.Projects.Get(projectID).Do()
	req := client.svc.Networks.List(projectID)
	// if err != nil {
	// 	log.Printf("coulf not retrive the resource: %v", err)
	// }
	if err := req.Pages(ctx, func(page *raw.NetworkList) error {
                for _, subnet := range page.Items {
                        // TODO: Change code below to process each `firewall` resource:
                        fmt.Printf("%#v\n", subnet.RoutingConfig)
                }
                return nil
        }); err != nil {
                log.Fatal(err)
        }
}
