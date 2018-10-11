package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type job struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Salary      int64  `json:"salary"`
	State       string `json:"state"`
	City        string `json:"city"`
}

func getJobs(db *sql.DB, start, count int) ([]job, error) {
	statement := fmt.Sprintf("SELECT * FROM job LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	jobs := []job{}

	for rows.Next() {
		var u job
		if err := rows.Scan(&u.ID, &u.Title, &u.Description, &u.Salary, &u.State, &u.City); err != nil {
			return nil, err
		}
		jobs = append(jobs, u)
	}

	return jobs, nil
}

func (u *job) getJobID(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM job WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.ID, &u.Title, &u.Description, &u.Salary, &u.State, &u.City)
}

func (u *job) createJob(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO job(id, title, description, salary, state, city) VALUES(%d, '%s', '%s', %d, '%s', '%s')", u.ID, u.Title, u.Description, u.Salary, u.State, u.City)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *job) deleteJob(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM job WHERE id=%d", u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *job) updateJob(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE job SET title='%s', description='%s', salary=%d, state='%s', city='%s' WHERE id=%d", u.Title, u.Description, u.Salary, u.State, u.City, u.ID)
	_, err := db.Exec(statement)
	return err
}
