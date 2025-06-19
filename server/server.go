package server

import (
	"context"

	// "fmt"
	"log"
	"net"

	"TransportManagementService/db"
	pb "TransportManagementService/proto"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/yar1k3x/JWTValidation/jwt"
	"github.com/yar1k3x/JWTValidation/middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *server) GetTransportInfo(ctx context.Context, in *pb.GetTransportInfoRequest) (*pb.GetTransportInfoResponse, error) {
	log.Printf("Получен запрос от пользователя на получение инфы о транспорте")

	transport, err := db.GetTransportRequest(in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка при получении заявки: %v", err)
	}

	log.Printf("Заявка успешно получена")

	return &pb.GetTransportInfoResponse{
		Transports: transport,
	}, nil
}

func (s *server) CreateTransportLog(ctx context.Context, in *pb.CreateTransportLogRequest) (*pb.CreateTransportLogResponse, error) {
	log.Printf("Получен запрос от пользователя на создание транспортного ТО")

	requestID, err := db.CreateTransportLogRequest(in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ошибка записи в базу данных: %v", err)
	}

	log.Printf("Заявка успешно создана с ID %d", requestID)

	return &pb.CreateTransportLogResponse{
		Success: true,
	}, nil
}

func (s *server) GetTransportLogsInfo(ctx context.Context, in *pb.GetTransportLogsInfoRequest) (*pb.GetTransportLogsInfoResponse, error) {
	if in.TransportId != nil {
		log.Printf("Получен запрос от пользователя на получение инфы о транспортном ТО (#%d):", in.TransportId.Value)
	} else {
		log.Printf("Получен запрос от пользователя на получение всех транспортных ТО")
	}

	transportLogs, err := db.GetTransportLogsRequest(in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка при получении заявки: %v", err)
	}

	log.Printf("Заявка успешно получена")

	return &pb.GetTransportLogsInfoResponse{
		TransportLogs: transportLogs,
	}, nil
}

func (s *server) GetTransportType(ctx context.Context, in *emptypb.Empty) (*pb.GetTransportTypeResponse, error) {
	log.Printf("Получен запрос от пользователя на получения запросов")

	types, err := db.GetTransportTypes()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ошибка получения данных из базы данных: %v", err)
	}

	log.Printf("Заявки успешно получены")

	return &pb.GetTransportTypeResponse{
		Types: types,
	}, nil
}

func Start() {
	//jwt.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	jwt.JWTSecretKey = "ZuxooEpNl7MgUUbnxGntsBvSxEnizlgsDfTvOBGamck"
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(middleware.AuthMiddleware),
		),
	)

	pb.RegisterTransportServiceServer(s, &server{}) // Регистрация gRPC-сервиса

	reflection.Register(s)
	log.Println("gRPC сервер запущен на порту 50054")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("не удалось запустить gRPC: %v", err)
	}
}
