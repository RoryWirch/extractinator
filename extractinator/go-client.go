package main

import (
	"log"
	"github.com/jhump/protoreflect/grpcreflect"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"github.com/fullstorydev/grpcurl"

)

func main() {
	var conn *grpc.ClientConn
	var descSource grpcurl.DescriptorSource
	var myRefClient *grpcreflect.Client
	var fileSource grpcurl.DescriptorSource 
	
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	ctx := context.Background()

	myRefClient = grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(conn))
	myRefSource := grpcurl.DescriptorSourceFromServer(ctx, myRefClient)


}
