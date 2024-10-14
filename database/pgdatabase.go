package database

import (
	"context"
	"log"
	"users/server/users/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgDatabase struct {
	db *pgxpool.Pool
}

func InitDatabase(connString string) *PgDatabase {
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v \n", err)
	}
	log.Println("Connected to database")
	return &PgDatabase{db}
}

func (pg PgDatabase) GetUserById(id int) (models.User, error) {
	var user models.User
	err := pg.db.QueryRow(context.Background(),
		"SELECT id, name, age FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Age)
	return user, err
}

func (pg PgDatabase) PostUser(user models.User) (int, error) {
	var userID int
	err := pg.db.QueryRow(context.Background(),
		"INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id",
		user.Name, user.Age).Scan(&userID)
	return userID, err
}

func (pg PgDatabase) PutUser(user models.User) error {
	_, err := pg.db.Exec(context.Background(),
		"UPDATE users Set name = $1, age = $2 WHERE id=$3",
		user.Name, user.Age, user.ID)
	return err
}

func (pg PgDatabase) DeleteUser(id int) error {
	_, err := pg.db.Exec(context.Background(),
		"DELETE FROM users WHERE id=$1", id)
	return err
}
