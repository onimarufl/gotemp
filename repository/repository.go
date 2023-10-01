package repository

import "github.com/jmoiron/sqlx"

type Repository interface {
	InsertData(req InsertDataRequest) error
	InquiryDataByID(id int) (User, error)
	InquiryAllData() ([]User, error)
	UpdateDataByID(req UpdateDataRequest) (UpdateDataRequest, error)
	DeleteDataByID(id int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository {
	return repository{db: db}
}

func (r repository) InsertData(req InsertDataRequest) error {

	const query = `INSERT INTO user_test (firstname,lastname, age)VALUES ($1, $2, $3)`

	tx := r.db.MustBegin()
	tx.MustExec(query, req.Firstname, req.Lastname, req.Age)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r repository) InquiryDataByID(id int) (User, error) {

	userResp := User{}

	const query = `SELECT * FROM user_test u WHERE u.id = $1 `

	err := r.db.Get(&userResp, query, id)
	if err != nil {
		return userResp, err
	}

	return userResp, nil
}

func (r repository) InquiryAllData() ([]User, error) {

	userResp := []User{}
	rows, err := r.db.Query("SELECT * FROM user_test")
	if err != nil {
		return userResp, err
	}

	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.Firstname,
			&user.Lastname,
			&user.Age,
		)
		if err != nil {
			panic(err)
		}

		userResp = append(userResp, user)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return userResp, nil
}

func (r repository) UpdateDataByID(req UpdateDataRequest) (UpdateDataRequest, error) {

	const query = `UPDATE user_test SET firstname=$1,lastname=$2, age=$3 WHERE id=$4`

	tx := r.db.MustBegin()
	tx.MustExec(query, req.Firstname, req.Lastname, req.Age, req.ID)
	if err := tx.Commit(); err != nil {
		return UpdateDataRequest{}, err
	}

	return req, nil
}

func (r repository) DeleteDataByID(id int) error {
	_, err := r.db.Exec("DELETE FROM user_test u WHERE u.id = $1 ", id)
	if err != nil {
		return err
	}

	return nil
}
