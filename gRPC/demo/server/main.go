package main

// Here I wrote 2 functions - GetUser and DelUser. All other functions were in example.
// This code doesn't include 'grpc/status' because it will be covered in future lessons.

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sort"
	"sync"

	// импортируем пакет со сгенерированными protobuf-файлами
	pb "demo/proto"
)

//task:
// Реализуйте методы GetUser и DelUser.
// Следите за тем, чтобы типы входящих параметров и возвращаемых значений были
// определены правильно, в соответствии со спецификацией, описанной в demo.proto.

// UsersServer поддерживает все необходимые методы сервера.
type UsersServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedUsersServer

	// используем sync.Map для хранения пользователей
	users sync.Map
}

// AddUser реализует интерфейс добавления пользователя.
func (s *UsersServer) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	var response pb.AddUserResponse

	if _, ok := s.users.Load(in.User.Email); ok {
		response.Error = fmt.Sprintf("Пользователь с email %s уже существует", in.User.Email)
	} else {
		s.users.Store(in.User.Email, in.User)
	}
	return &response, nil
}

// ListUsers реализует интерфейс получения списка пользователей.
func (s *UsersServer) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	var list []string

	s.users.Range(func(key, _ interface{}) bool {
		list = append(list, key.(string))
		return true
	})
	// сортируем слайс из email
	sort.Strings(list)

	offset := int(in.Offset)
	end := int(in.Offset + in.Limit)
	if end > len(list) {
		end = len(list)
	}
	if offset >= end {
		offset = 0
		end = 0
	}
	response := pb.ListUsersResponse{
		Count:  int32(len(list)),
		Emails: list[offset:end],
	}
	return &response, nil
}

func (s *UsersServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, ok := s.users.Load(in.Email)
	if !ok {
		response := &pb.GetUserResponse{
			Error: "404: user not found",
		}
		return response, nil
	}

	toReturn, ok := user.(*pb.User)
	if !ok {
		response := &pb.GetUserResponse{
			Error: "500: internal server error",
		}
		return response, nil
	}
	return &pb.GetUserResponse{
		User:  toReturn,
		Error: "",
	}, nil
}

func (s *UsersServer) DelUser(ctx context.Context, in *pb.DelUserRequest) (*pb.DelUserResponse, error) {
	_, ok := s.users.LoadAndDelete(in.Email)
	if !ok {
		response := &pb.DelUserResponse{
			Error: "404: user not found",
		}
		return response, nil
	}
	response := &pb.DelUserResponse{}
	return response, nil
}

func main() {
	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()
	// регистрируем сервис
	pb.RegisterUsersServer(s, &UsersServer{})

	fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
