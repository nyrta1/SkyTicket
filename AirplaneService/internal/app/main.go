package main

import (
	"AirplaneService/internal/controller"
	grpcc "AirplaneService/internal/grpc"
	"AirplaneService/internal/gw"
	"AirplaneService/internal/repository"
	"AirplaneService/pkg/logger"
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "google.golang.org/protobuf/types/known/timestamppb"

	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	grpcAddress = "localhost:50053"
	httpAddress = "localhost:8083"
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

	airplaneRepo := repository.NewAirplaneRepository(dbConn)
	countryRepo := repository.NewCountryRepository(dbConn)
	manufacturerRepo := repository.NewManufacturerRepository(dbConn)

	airportHandler, err := controller.NewBookingHandler(*airplaneRepo)
	if err != nil {
		log.Fatal(err)
	}
	countryHandler, err1 := controller.NewCountryHandler(*countryRepo)
	if err1 != nil {
		log.Fatal(err)
	}
	manufacturerHandler, err2 := controller.NewManufacturerHandler(*manufacturerRepo)
	if err2 != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	grpcc.RegisterAirplaneServiceServer(grpcServer, airportHandler)
	grpcc.RegisterCountryServiceServer(grpcServer, countryHandler)
	grpcc.RegisterManufacturerServiceServer(grpcServer, manufacturerHandler)

	reflection.Register(grpcServer)

	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	log.Printf("gRPC server listening at %v", grpcAddress)

	return grpcServer.Serve(list)
}

func startHttpServer(ctx context.Context) error {
	fmt.Println(ctx)
	log := logger.NewLogger()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	err1 := gw.RegisterAirplaneServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err1 != nil {
		return err1
	}
	err2 := gw.RegisterCountryServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err2 != nil {
		return err2
	}
	err3 := gw.RegisterManufacturerServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err3 != nil {
		return err3
	}

	log.Printf("HTTP server listening at %v", httpAddress)

	return http.ListenAndServe(httpAddress, mux)
}

func ConnectDB() (*sql.DB, error) {
	newLogger := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		newLogger.Printf("Error loading .env file: %v", err)
	}

	connectionString := os.Getenv("PG_URL")
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
