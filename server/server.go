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
	pb.UnimplementedTransportServiceServer
}

func (s *server) CreateTransport(ctx context.Context, in *pb.CreateTransportRequest) (*pb.CreateTransportResponse, error) {
	log.Printf("Получен запрос от пользователя на создание транспорта")

	requestID, err := db.CreateTransportRequest(in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ошибка записи в базу данных: %v", err)
	}

	log.Printf("Заявка успешно создана с ID %d", requestID)

	return &pb.CreateTransportResponse{
		Success:     true,
		TransportId: requestID,
	}, nil
}

func (s *server) UpdateTransport(ctx context.Context, in *pb.UpdateTransportRequest) (*pb.UpdateTransportResponse, error) {
	log.Printf("Получен запрос от пользователя на обновление инфы о транспорте")
	if in.TransportId == nil {
		log.Printf("Ошибка: Получен запрос от пользователя с пустым TransportId")
		return nil, status.Errorf(codes.InvalidArgument, "TransportId не может быть пустым")
	}

	success, err := db.UpdateTransportRequest(in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка при обновлении заявки: %v", err)
	}

	log.Printf("Заявка успешно обновлена")

	return &pb.UpdateTransportResponse{
		Success: success,
	}, nil

}

func Start() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTransportServiceServer(s, &server{}) // Регистрация gRPC-сервиса

	reflection.Register(s)
	log.Println("gRPC сервер запущен на порту 50054")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("не удалось запустить gRPC: %v", err)
	}
}
