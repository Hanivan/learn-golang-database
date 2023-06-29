package repository

import (
	"context"
	"fmt"
	learn_golang_database "learn-golang-database"
	"learn-golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(learn_golang_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository01@test.com",
		Comment: "I inserted using repository pattern",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(learn_golang_database.GetConnection())
	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 99)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(learn_golang_database.GetConnection())
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
