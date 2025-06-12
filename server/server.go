package server

import (
	"context"
	// "fmt"
	"log"
	"net"

	"TransportManagementService/db"
	pb "TransportManagementService/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *server) CreateRequest(ctx context.Context, in *pb.CreateTransportRequest) (*pb.CreateTransportResponse, error) {
	log.Printf("Получен запрос от пользователя на создание транспорта")

	// if err := in.Validate(); err != nil {
	// 	return nil, status.Errorf(codes.InvalidArgument, "ошибка валидации: %v", err)
	// }

	// Вставка в БД
	requestID, err := db.CreateTransportRequest(in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ошибка записи в базу данных: %v", err)
	}

	log.Printf("Заявка успешно создана с ID %d", requestID)

	return &pb.CreateTransportResponse{
		TransportId: requestID,
	}, nil
}

func Start() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterNotificationServiceServer(s, &server{}) // Регистрация gRPC-сервиса

	reflection.Register(s)
	log.Println("gRPC сервер запущен на порту 50054")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("не удалось запустить gRPC: %v", err)
	}
}
