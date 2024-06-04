package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	_ "github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "google.golang.org/protobuf/types/known/timestamppb"

	"SkyTicket/handlers"
	pb "SkyTicket/pb"
	"SkyTicket/pkg/logger"
	repository "SkyTicket/repo"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	grpcAddress = "localhost:50051"
	httpAddress = "localhost:8080"
)

func main() {
	ctx := context.Background()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := startGrpcServer(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()

		if err := startHttpServer(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}

func startGrpcServer() error {
	log := logger.NewLogger()

	dbConn, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	models := repository.NewModels(dbConn)

	bookingHandler, err := handlers.NewBookingHandler(&models.Booking)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookingManagerServer(grpcServer, bookingHandler)
	reflection.Register(grpcServer)

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	log.Printf("gRPC server listening at %v\n", grpcAddress)

	return grpcServer.Serve(list)
}

func startHttpServer(ctx context.Context) error {
	log := logger.NewLogger()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	err := pb.RegisterBookingManagerHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	log.Printf("HTTP server listening at %v\n", httpAddress)

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
