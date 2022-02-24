package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// Flags for CLI options
var (
	addr = flag.String("addr", "localhost:50051", "The address the client will connect to. Default is localhost:50051")
	out = flag.String("out", "csv-extract", "Name for the output file")
)

func main() {
	flag.Parse()
	// Set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	var descSource grpcurl.DescriptorSource

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reflClient := grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(conn))
	reflSource := grpcurl.DescriptorSourceFromServer(ctx, reflClient)
	
	descSource = reflSource
	
	// Set up for using InvokeRPC()
	verbosityLevel := 0
	symbol := "api.Registry/ListBundles"
	in := strings.NewReader("")
	var headers []string
	options := grpcurl.FormatOptions{
		EmitJSONDefaultFields: true,
		IncludeTextSeparator:  true,
		AllowUnknownFields:    true,
	}

	rf, formatter, err := grpcurl.RequestParserAndFormatter(grpcurl.Format("json"), descSource, in, options)
	if err != nil {
		log.Fatalf("Could not construct request parser and formatter")
	}
	handler := &grpcurl.DefaultEventHandler{
		Out:            os.Stdout,
		Formatter:      formatter,
		VerbosityLevel: verbosityLevel,
	}

	err = grpcurl.InvokeRPC(ctx, descSource, conn, symbol, headers, handler, rf.Next)
	if err != nil{
		log.Fatalf("Failed to InvokeRPC: %v", err)
	}
	grpcurl.PrintStatus(os.Stderr, handler.Status, formatter)
}