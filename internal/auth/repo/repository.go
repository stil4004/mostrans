package repo

import (
	"context"
	"errors"
	"fmt"
	"service/internal/auth"
	"service/internal/cconstants"

	"github.com/jmoiron/sqlx"
)

type postgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) auth.Repository {
	return &postgresRepository{db: db}
}

// CheckLogIn(req CheckLogInRequest) (CheckLogInResponse, error)
// 	GetUserByLogin(req GetUserByLoginRequest) (GetUserByLoginResponse, error)

func (p *postgresRepository) CheckLogIn(ctx context.Context, req auth.CheckLogInRequest) (auth.CheckLogInResponse, error) {
	var (
		query = `SELECT login, password FROM %[1]s
		 WHERE login=$1
		 LIMIT 1;
		`
		userDB auth.UserLoginPassword = auth.UserLoginPassword{}

		// vals []any = []any{cconstants.AccessTable}
	)
	requestDB := fmt.Sprintf(query, cconstants.AccessTable)

	err := p.db.GetContext(ctx, &userDB, requestDB, req.NickName)
	if err != nil {
		return auth.CheckLogInResponse{}, err
	}

	if userDB.Password != "" && userDB.Password == req.Password {
		return auth.CheckLogInResponse{
			Authenticated: true,
		}, nil
	}

	return auth.CheckLogInResponse{
		Authenticated: false,
	}, errors.New("wrong password")
}

func (p *postgresRepository) GetUserByLogin(ctx context.Context, req auth.GetUserByLoginRequest) (auth.GetUserByLoginResponse, error) {
	var (
		query = `
		SELECT name, surname FROM %[1]s
		 WHERE nickname = $1
		 LIMIT 1;
		`
		userDB auth.User = auth.User{}
	)
	requestDB := fmt.Sprintf(query, cconstants.UserTable)

	err := p.db.GetContext(ctx, &userDB, requestDB, req.NickName)
	if err != nil {
		return auth.GetUserByLoginResponse{}, err
	}

	if userDB.Name == "" || userDB.Surname == "" {
		return auth.GetUserByLoginResponse{}, errors.New(" missed some rows ")
	}

	return auth.GetUserByLoginResponse{
		UserResp: userDB,
	}, nil
}

// func (p postgresRepository) CreateAdmin(params *admins.AdminSignUpParams) (int, *cErrors.ResponseErrorModel) {
// 	var adminID int

// 	tx, err := p.db.Begin()
// 	if err != nil {
// 		return adminID, &cErrors.ResponseErrorModel{
// 			InternalCode: cErrors.BeginTransactionError,
// 			StandardCode: http.StatusInternalServerError,
// 		}
// 	}
// 	defer func() {
// 		_ = tx.Rollback()
// 	}()

// 	now := time.Now().UTC()

// 	queryInsertAdmin := `
// 		INSERT INTO admin_db.public.admin
// 		(nickname, password, second_password, registration_date, totp_secret, is_blocked, last_entry, ip)
// 		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
// 		RETURNING admin_id
// 	`

// 	err = tx.QueryRow(queryInsertAdmin, params.Nickname, params.Password, params.Password, now, params.TotpSecret, false, now, params.IP).Scan(&adminID)
// 	if err != nil {
// 		return adminID, &cErrors.ResponseErrorModel{
// 			InternalCode: cErrors.Admins_CreateAdmin_Error,
// 			StandardCode: http.StatusInternalServerError,
// 			Message:      err.Error(),
// 		}
// 	}

// 	queryInsertRole := `
// 		INSERT INTO admin_db.public.admin_roles
// 		(admin_id, admin_role_id)
// 		VALUES($1, $2)
// 	`

// 	for _, role := range params.Roles {
// 		_, err := tx.Exec(queryInsertRole, adminID, role)
// 		if err != nil {
// 			return adminID, &cErrors.ResponseErrorModel{
// 				InternalCode: cErrors.Admins_InsertRole_Error,
// 				StandardCode: http.StatusInternalServerError,
// 			}
// 		}
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return adminID, &cErrors.ResponseErrorModel{
// 			InternalCode: cErrors.CommitTransactionError,
// 			StandardCode: http.StatusInternalServerError,
// 		}
// 	}

// 	return adminID, nil
// }

// func (p postgresRepository) ChangeAdminBlock(params *admins.BlockAdminParams) *cErrors.ResponseErrorModel {
// 	var (
// 		query string = `UPDATE admin_db.public.admin SET is_blocked=$1 WHERE admin_id=$2`
// 	)
// 	res, err := p.db.Exec(query, params.Block, params.AdminID)
// 	if err != nil {
// 		return &cErrors.ResponseErrorModel{
// 			InternalCode: http.StatusInternalServerError,
// 			StandardCode: http.StatusInternalServerError,
// 			Message:      err.Error(),
// 		}
// 	}
// 	insertedRow, err := res.RowsAffected()
// 	if err != nil {
// 		return &cErrors.ResponseErrorModel{
// 			InternalCode: cErrors.Admins_InsertAdminOperation_insertedRow,
// 			StandardCode: http.StatusInternalServerError,
// 		}
// 	}
// 	if insertedRow == 0 {
// 		return &cErrors.ResponseErrorModel{
// 			InternalCode: cErrors.Admins_InsertAdminOperation_insertedRow_Zero,
// 			StandardCode: http.StatusInternalServerError,
// 		}
// 	}
// 	return &cErrors.ResponseErrorModel{}
// }
