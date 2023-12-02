package todorepository

import (
	"context"
	"database/sql"
	"log"

	todoservicev1 "github.com/Kimoto-Norihiro/connect-todo/server/api/todoservice/v1"
)

type TODORepository struct {
	db *sql.DB
}

func NewTODORepository(db *sql.DB) *TODORepository {
	return &TODORepository{
		db: db,
	}
}

func (r TODORepository) CreateTODO(ctx context.Context, tx *sql.Tx, title string) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO todos (title, done) VALUES (?, ?)", title, false)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r TODORepository) ListTODOs(ctx context.Context, db *sql.DB) ([]*todoservicev1.TODO, error) {
	todos := []*todoservicev1.TODO{}
	rows, err := db.QueryContext(ctx, "SELECT * FROM todos")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		todo := &todoservicev1.TODO{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r TODORepository) UpdateTODO(ctx context.Context, tx *sql.Tx, todo *todoservicev1.TODO) error {
	_, err := tx.ExecContext(ctx, "UPDATE todos SET title = ?, done = ? WHERE id = ?", todo.Title, todo.Done, todo.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r TODORepository) DeleteTODO(ctx context.Context, db *sql.DB, id int32) error {
	_, err := db.ExecContext(ctx, "DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
