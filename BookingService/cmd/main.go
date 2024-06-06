package main

import (
	"SkyTicket/pkg/logger"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"SkyTicket/BookingService/handlers"
	repository "SkyTicket/BookingService/repo"
	"SkyTicket/proto/pb"
	"database/sql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	dbConn, err := ConnectDB()
	if err != nil {
		return fmt.Errorf("database connection error: %v", err)
	}
	defer dbConn.Close()
	models := repository.NewModels(dbConn)

	// Dial user gRPC server
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dial user gRPC server: %v", err)
	}
	defer conn.Close()
	userClient := pb.NewUserManagerClient(conn)

	// Dial flight gRPC server
	conn2, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dial flight gRPC server: %v", err)
	}
	defer conn2.Close()
	flightClient := pb.NewFlightManagerClient(conn2)

	bookingHandler, err := handlers.NewBookingHandler(&models.Booking, flightClient, userClient)
	if err != nil {
		return fmt.Errorf("failed to create booking handler: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookingManagerServer(grpcServer, bookingHandler)
	reflection.Register(grpcServer)

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
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
		return fmt.Errorf("failed to register HTTP handler: %v", err)
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
