package repository

import (
	"context"
	"os"
	"url-shortener/model"

	"github.com/jackc/pgx/v5"
)

type PostgresRepository struct {
	db *pgx.Conn
}

func InitDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	schema := `CREATE TABLE IF NOT EXISTS urls (
    	hash VARCHAR(6) PRIMARY KEY,
    	link TEXT,
    	created_at TIMESTAMP,
    	expires_at TIMESTAMP,
      click_counts int
    );`
	_, err = conn.Exec(context.Background(), schema)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewPostgresRepository() (PostgresRepository, error) {
	dbConn, err := InitDB()
	if err != nil {
		return PostgresRepository{}, err
	}
	return PostgresRepository{
		db: dbConn,
	}, nil
}

func (r *PostgresRepository) LoadURL(shortcode string) (model.Url, error) {
	query := `SELECT hash, link, expires_at FROM urls WHERE hash = $1;`
	row := r.db.QueryRow(context.Background(), query, shortcode)
	var url model.Url
	err := row.Scan(&url.Hash, &url.Link, &url.ExpiresAt)
	return url, err
}

func (r *PostgresRepository) SaveURL(url model.Url) error {
	query := `INSERT INTO urls(hash, link, created_at, expires_at, click_counts) 
            values($1, $2, $3, $4, $5);`
	_, err := r.db.Exec(context.Background(),
		query, url.Hash, url.Link, url.CreatedAt, url.ExpiresAt, url.ClickCounts)
	return err
}

func (r *PostgresRepository) IncrementClicks(hash string) error {
	query := `UPDATE urls SET click_counts = click_counts + 1 WHERE hash = $1;`
	_, err := r.db.Exec(context.Background(), query, hash)
	return err
}

func (r *PostgresRepository) LoadAll(showMostClicked bool) (map[string]model.Url, error) {
	urls := make(map[string]model.Url)
	query := `SELECT hash, link, created_at, expires_at, click_counts FROM urls`
	if showMostClicked {
		query += ` ORDER BY click_counts DESC LIMIT 5`
	}
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var url model.Url
		err := rows.Scan(&url.Hash, &url.Link, &url.CreatedAt, &url.ExpiresAt, &url.ClickCounts)
		if err != nil {
			return nil, err
		}
		urls[url.Hash] = url
	}
	return urls, err
}
