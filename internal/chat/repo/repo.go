package repo

import (
	"context"
	"errors"
	"fmt"
	"log"
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

func (p *postgresRepository) GetInfoFromBatch(ctx context.Context, req chat.GetInfoFromBatchRequest) (chat.GetInfoFromBatchResponse, error) {
	var (
		query = `SELECT SUM(passenger_count) ps_count FROM %[1]s
		 WHERE station_id=(
			SELECT station_id FROM %[2]s 
			WHERE station_name=$1
		 ) AND date=$2
		`
		// userDB auth.UserLoginPassword = auth.UserLoginPassword{}
		respDB chat.GetInfoFromBatchRequest = chat.GetInfoFromBatchRequest{}
		time   string
		// vals []any = []any{cconstants.AccessTable}
	)
	requestDB := fmt.Sprintf(query, cconstants.PassengersTable, cconstants.StationsTable)

	if len(req.Periods) == 1 {
		time = req.Periods[0]
	} else if len(req.Periods) > 1 {
		time = req.Periods[len(req.Periods)-1]
	}
	if len(req.Periods) < 1 || len(req.Stations) < 1 {
		return chat.GetInfoFromBatchResponse{}, errors.New("no data given")
	}

	err := p.db.GetContext(ctx, &respDB, requestDB, req.Stations[0], time)
	if err != nil {
		log.Println(err)
		return chat.GetInfoFromBatchResponse{}, err
	}

	return chat.GetInfoFromBatchResponse{}, nil
}

func (p *postgresRepository) GetOneStation(ctx context.Context, req chat.GetOneStationRequest) (chat.GetOneStationResponse, error) {
	var (
		query = `SELECT passenger_count ps_count FROM %[1]s
		 WHERE station_id=(
			SELECT station_id FROM %[2]s 
			WHERE station_name=$1
		 ) AND date=$2
		`
		respDB chat.GetOneStationResponse = chat.GetOneStationResponse{}
	)

	requestDB := fmt.Sprintf(query, cconstants.PassengersTable, cconstants.StationsTable)

	err := p.db.GetContext(ctx, &respDB, requestDB, req.Station, req.Date)
	if err != nil {
		return chat.GetOneStationResponse{}, err
	}
	return chat.GetOneStationResponse{}, nil
}
