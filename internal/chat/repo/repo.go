package repo

import (
	"context"
	"errors"
	"fmt"
	"service/internal/auth"
	"service/internal/cconstants"
	"service/internal/chat"

	"github.com/jmoiron/sqlx"
)

type postgresRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) chat.Repository {
	return &postgresRepository{db: db}
}

//	GetInfoFromBatch(ctx context.Context, req GetInfoFromBatchRequest) (GetInfoFromBatchResponse, error)
// GetOneStation(ctx context.Context, req GetOneStationRequest) (GetOneStationResponse, error)

func (p *postgresRepository) GetInfoFromBatch(ctx context.Context, req chat.GetInfoFromBatchRequest) (chat.GetInfoFromBatchResponse, error) {
	var (
		query = `SELECT SUM(passenger_count) FROM %[1]s
		 WHERE login=$1
		 LIMIT 1;
		`
		// userDB auth.UserLoginPassword = auth.UserLoginPassword{}

		// vals []any = []any{cconstants.AccessTable}
	)
	requestDB := fmt.Sprintf(query, cconstants.PassengersTable)

	err := p.db.GetContext(ctx, &userDB, requestDB, req.NickName)
	if err != nil {
		return chat.GetInfoFromBatchResponse{}, err
	}

	return chat.GetInfoFromBatchResponse{}, errors.New("wrong password")
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
