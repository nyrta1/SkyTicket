package main

import (
	"SkyTicket/handlers"
	"SkyTicket/pb"
	"SkyTicket/pkg/logger"
	repository "SkyTicket/repo"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	log := logger.NewLogger()
	dbConn, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	models := repository.NewModels(dbConn)

	bookingHandler, err := handlers.NewBookingHandler(&models.Booking, &models.Flight)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookingManagerServer(grpcServer, bookingHandler)
	reflection.Register(grpcServer)

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		log.Infof("gRPC server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	r := gin.Default()
	log.Infof("HTTP server listening at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run HTTP server: %v", err)
	}
}
func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
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
