package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	// "sync"

	v1 "sina/pb"

	"github.com/jackc/pgx/pgtype"
	// "github.com/jackc/pgx/v5/pgtype"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port         = "5432"
	dbConnection = "host=localhost port=5432 user=postgres password=password dbname=users sslmode=disable"
)

type server struct {
	v1.UnimplementedGetUsersServer
	db *sql.DB
}

func (s *server) GetUsers(ctx context.Context, req *v1.UserRequest) (*v1.UserResponse, error) {

	if req.UserId == 0 {
		query := "SELECT * FROM users LIMIT 100"
		rows, err := s.db.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var users []*v1.User
		for rows.Next() {
			var user v1.User
			var ts pgtype.Timestamp
			err := rows.Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex, &ts)
			if err != nil {
				return nil, err
			}
			user.CreatedAt = ts.Time.String()
			users = append(users, &user)
		}

		response := &v1.UserResponse{
			Users:     users,
			MessageId: req.MessageId,
		}

		return response, nil
	}

	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", req.UserId)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*v1.User
	for rows.Next() {
		var user v1.User
		var ts pgtype.Timestamp
		err := rows.Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex, &ts)
		if err != nil {
			return nil, err
		}
		user.CreatedAt = ts.Time.String()
		users = append(users, &user)
	}

	// Create the UserResponse message with the queried users and message_id
	response := &v1.UserResponse{
		Users:     users,
		MessageId: req.MessageId,
	}

	return response, nil
}

func (s *server) GetUsersWithSqlInject(ctx context.Context, req *v1.UserRequestWithSqlInject) (*v1.UserResponse, error) {
	if req.UserId == "" {
		query := "SELECT * FROM users LIMIT 100"
		rows, err := s.db.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var users []*v1.User
		for rows.Next() {
			var user v1.User
			var ts pgtype.Timestamp
			err := rows.Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex, &ts)
			if err != nil {
				return nil, err
			}
			user.CreatedAt = ts.Time.String()
			users = append(users, &user)
		}

		response := &v1.UserResponse{
			Users:     users,
			MessageId: req.MessageId,
		}

		return response, nil

	}
	query := "SELECT * FROM users WHERE id = " + req.UserId
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*v1.User
	for rows.Next() {
		var user v1.User
		var ts pgtype.Timestamp
		err := rows.Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex, &ts)
		if err != nil {
			return nil, err
		}
		user.CreatedAt = ts.Time.String()
		users = append(users, &user)
	}

	response := &v1.UserResponse{
		Users:     users,
		MessageId: req.MessageId,
	}

	return response, nil
}

func main() {
	db, err := sql.Open("postgres", dbConnection)
	if err != nil {
		log.Fatalf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa %v", err)
	}

	defer db.Close()
	lis, err := net.Listen("tcp", ":5062")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	v1.RegisterGetUsersServer(s, &server{db: db})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// wg := sync.WaitGroup{}
	// wg.Add(2)

	// go func() {
	// 	lis, err := net.Listen("tcp", ":5062")
	// 	if err != nil {
	// 		log.Fatalf("failed to listen: %v", err)
	// 	}

	// 	s := grpc.NewServer()
	// 	v1.RegisterGetUsersServer(s, &server{db: db})
	// 	reflection.Register(s)
	// 	if err := s.Serve(lis); err != nil {
	// 		log.Fatalf("failed to serve: %v", err)
	// 	}
	// 	wg.Done()
	// }()
	// go func() {
	// 	lis, err := net.Listen("tcp", ":5061")
	// 	if err != nil {
	// 		log.Fatalf("failed to listen: %v", err)
	// 	}

	// 	s2 := grpc.NewServer()
	// 	v1.RegisterGetUsersWithSqlInjectServer(s2, &injectserver{db: db})
	// 	reflection.Register(s2)
	// 	if err := s2.Serve(lis); err != nil {
	// 		log.Fatalf("failed to serve: %v", err)
	// 	}

	// 	wg.Done()
	// }()

	// wg.Wait()
}
