package repository

import (
	"database/sql"
	"log"
	"math"

	"enigmacamp.com/blog-apps/model"
	"enigmacamp.com/blog-apps/shared/shared_model"
)

type BlogRepository interface {
	Create(payload model.Blog) (model.Blog, error)
	List(page, size int) ([]model.Blog, shared_model.Paging, error)
}

type blogRepository struct {
	db *sql.DB
}

func (b *blogRepository) Create(payload model.Blog) (model.Blog, error) {
	sql := `INSERT INTO blogs (user_id, title, body) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := b.db.QueryRow(sql, payload.UserId, payload.Title, payload.Body).Scan(&payload.ID, &payload.CreatedAt, &payload.UpdatedAt)
	if err != nil {
		log.Println("blogRepository.Create.QueryRow:", err.Error())
		return model.Blog{}, err
	}

	return payload, nil
}

func (b *blogRepository) List(page int, size int) ([]model.Blog, shared_model.Paging, error) {
	var tasks []model.Blog
	offset := (page - 1) * size
	sql := `SELECT id, user_id, title, body, created_at, updated_at FROM blogs ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := b.db.Query(sql, size, offset)
	if err != nil {
		log.Println("taskRepository.Query:", err.Error())
		return nil, shared_model.Paging{}, err
	}
	for rows.Next() {
		var task model.Blog
		err := rows.Scan(
			&task.ID,
			&task.UserId,
			&task.Title,
			&task.Body,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			log.Println("blogRepository.List.Rows.Next():", err.Error())
			return nil, shared_model.Paging{}, err
		}

		tasks = append(tasks, task)
	}

	totalRows := 0
	if err := b.db.QueryRow("SELECT COUNT(*) FROM blogs").Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}

	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return tasks, paging, nil

}

func NewBlogRepository(db *sql.DB) BlogRepository {
	return &blogRepository{db: db}
}
