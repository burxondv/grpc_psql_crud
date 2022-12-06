package main

import (
	"context"
	"fmt"
	"net"
	"os/user"

	userService "github.com/burxondv/grpc_psql_crud/genproto/user_crud"
	"github.com/burxondv/grpc_psql_crud/storage"
	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	userService.UnimplementedUserCrudServer
}

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) storage.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	userService.RegisterUserCrudServer(srv, &Server{})
	reflection.Register(srv)

	fmt.Println("server started")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *Server) Create(ctx context.Context, request *userService) (*userService.User, error) {
	query := `
		INSERT INTO users (
			first_name=$1,
			last_name=$2,
            age=$3,
			phone_number=$4
		) VALUES ($1, $2, $3, $4)
		RETURNING first_name, last_name, age, phone_number
	`

	row := ur.db.QueryRow(
		query,
		request.FirstName,
		request.LastName,
        request.Age,
		request.PhoneNumber,
	)
	
	err := row.Scan(
        &request.FirstName,
        &request.LastName,
        &request.Age,
        &request.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}
	
	return &userService.User{FirstName: request.FirstName, LastName: request.LastName, Age: request.Age, PhoneNumber: request.PhoneNumber}, nil
}

func (s *Server) Get(ctx context.Context, request ) (*repo.User, error) {
	var result repo.User

	query := `
		SELECT
			id,
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			username,
			profile_image_url,
			type,
			created_at
		FROM users
		WHERE id=$1
	`

	row := ur.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
		&result.Gender,
		&result.Password,
		&result.Username,
		&result.ProfileImageUrl,
		&result.Type,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

