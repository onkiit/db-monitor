package user

import (
	"context"

	"github.com/onkiit/db-monitor/lib/db/psql"
)

type pgdb struct{}

func (p pgdb) AddUser(ctx context.Context, user *User) error {
	stmt, err := psql.DB().PrepareContext(ctx, "INSERT INTO USERS.USERS (UUID,NAME,USERNAME,EMAIL,PASSWORD) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, user.UUID, user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil
}

func (p pgdb) GetById(ctx context.Context, uid string) (*User, error) {
	user := new(User)
	err := psql.DB().QueryRowContext(ctx, "SELECT UUID, NAME, USERNAME, EMAIL, DTM_CRT, DTM_UPD FROM USERS.USERS WHERE UUID=$1", uid).Scan(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p pgdb) GetByUsername(ctx context.Context, username string) (*User, error) {
	user := &User{}
	err := psql.DB().QueryRowContext(ctx, "SELECT UUID, NAME, USERNAME, EMAIL, PASSWORD, DTM_CRT, DTM_UPD FROM USERS.USERS WHERE USERNAME=$1", username).Scan(&user.UUID, &user.Name, &user.Username, &user.Email, &user.Password, &user.Created, &user.Modified)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p pgdb) GetByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	err := psql.DB().QueryRowContext(ctx, "SELECT UUID, NAME, USERNAME, EMAIL, PASSWORD, DTM_CRT, DTM_UPD FROM USERS.USERS WHERE EMAIL=$1", email).Scan(&user.UUID, &user.Name, &user.Username, &user.Email, &user.Password, &user.Created, &user.Modified)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func New() Store {
	return pgdb{}
}
