package repo

import (
	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		)
		VALUES (
			:first_name,
			:last_name,
			:email,
			:password,
			:is_shop_owner
		)
		RETURNING id
	`

	stmt, err := r.dbCon.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(user).Scan(&id)
	if err != nil {
		return nil, err
	}

	user.ID = id
	return &user, nil
}

func (r *userRepo) Find(email, password string) (*domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE email = $1 AND password = $2"
	err := r.dbCon.Get(&user, query, email, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
