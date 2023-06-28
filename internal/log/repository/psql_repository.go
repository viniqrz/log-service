package repository

import (
	"database/sql"

	"github.com/viniqrz/gin-crud/internal/log/entity"
)

type LogPSQLRepository struct {
	db *sql.DB
}

func NewLogPSQLRepository(db *sql.DB) *LogPSQLRepository {
	return &LogPSQLRepository{
		db: db,
	}
}

func (r *LogPSQLRepository) FindAll() ([]*entity.Log, error) {
	rows, err := r.db.Query("SELECT id, message, created_at, level FROM logs AS l")
	if (err != nil) {
		return nil, err
	}
	logs := make([]*entity.Log, 0)
	for rows.Next() {
		log := new(entity.Log)
		err := rows.Scan(&log.ID, &log.Message, &log.CreatedAt, &log.Level)
		if (err != nil) {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func (r *LogPSQLRepository) Create(log *entity.Log) (*entity.Log, error) {
	_, err := r.db.Exec("INSERT INTO logs (id, level, message, created_at) VALUES ($1, $2, $3, $4)", log.ID, log.Level, log.Message, log.CreatedAt)
	if err != nil {
		return nil, err
	}
	return log, nil
}