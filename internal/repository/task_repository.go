package repository

import (
	"context"

	"github.com/alikurb12/todo-app-go/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) GetALL(ctx context.Context) ([]model.Task, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT * FROM tasks ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}

func (r *TaskRepository) GetById(ctx context.Context, id int64) (*model.Task, error) {
	var task model.Task
	err := r.db.QueryRow(ctx, "SELECT * FROM tasks WHERE id = $1", id).
		Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `INSERT INTO tasks (title, description, completed)
			  VALUES ($1, $2, $3)
			  RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(ctx, query, task.Title, task.Description, task.Completed).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	return err
}

func (r *TaskRepository) Update(ctx context.Context, task *model.Task) error {
	query := `UPDATE tasks 
			  SET title=$1, description=$2, completed=$3, updated_at=NOW()
			  WHERE id=$4
			  RETURNING updated_at`
	err := r.db.QueryRow(ctx, query, task.Title, task.Description, task.Completed, task.ID).
		Scan(&task.UpdatedAt)
	return err
}

func (r *TaskRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM tasks WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
