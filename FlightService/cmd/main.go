package main

import (
	"SkyTicket/FlightService/handlers"
	repository "SkyTicket/FlightService/repo"
	"SkyTicket/pkg/logger"
	"SkyTicket/proto/pb"
	"context"
	"database/sql"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

const (
	grpcAddress = "localhost:50053"
	httpAddress = "localhost:8082"
)

func main() {
	ctx := context.Background()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := startGrpcServer(); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		if err := startHttpServer(ctx); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	wg.Wait()
}

func startGrpcServer() error {
	// Connect to the database
	dbConn, err := ConnectDB()
	if err != nil {
		return fmt.Errorf("database connection error: %v", err)
	}
	defer dbConn.Close()

	// Create repository models
	models := repository.NewModels(dbConn)

	// Create a new flight handler with the flight repository
	flightHandler, err := handlers.NewFlightHandler(models.Flight)
	if err != nil {
		return fmt.Errorf("failed to create flight handler: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register flight handler with the gRPC server
	pb.RegisterFlightManagerServer(grpcServer, flightHandler)

	// Register reflection service on gRPC server
	reflection.Register(grpcServer)

	// Listen for incoming connections on the specified address
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Serve gRPC server
	log.Printf("gRPC server listening at %s\n", grpcAddress)
	return grpcServer.Serve(listener)
}

func startHttpServer(ctx context.Context) error {
	// Create a new logger
	log := logger.NewLogger()

	// Create a new ServeMux
	mux := runtime.NewServeMux()

	// Create gRPC dial options
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register flight manager handler from endpoint
	err := pb.RegisterFlightManagerHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("failed to register HTTP handler: %v", err)
	}

	// Listen and serve HTTP requests
	log.Printf("HTTP server listening at %s\n", httpAddress)
	return http.ListenAndServe(httpAddress, mux)
}
func ConnectDB() (*sql.DB, error) {
	newLogger := logger.NewLogger()
	err := godotenv.Load()
	if err != nil {
		newLogger.Printf("Error loading .env file: %v", err)
	}

	connectionString := os.Getenv("DB_CONNECTION_STRING")
	if connectionString == "" {
		return nil, fmt.Errorf("DB_CONNECTION_STRING environment variable not set")
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	return db, nil
}
