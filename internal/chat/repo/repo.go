package repo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"service/internal/cconstants"
	"service/internal/chat"
	"time"

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
		query = `select passenger_count  from %[1]s pf 
		where station_id = (
			select station_id from %[2]s s 
			where s.station_name = $1
		) and date = $2;
		`
		respDB     chat.GetInfoFromBatchResponse = chat.GetInfoFromBatchResponse{}
		timePeriod string
	)
	requestDB := fmt.Sprintf(query, cconstants.PassengersTable, cconstants.StationsTable)

	if len(req.Periods) == 1 {
		timePeriod = req.Periods[0]
		t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", timePeriod)
		if err != nil {
			log.Println("time parse 1 err", err)
			return chat.GetInfoFromBatchResponse{}, nil
		}
		timePeriod = t.Format("2006-01-02")
	} else if len(req.Periods) > 1 {
		timePeriod = req.Periods[len(req.Periods)-1]
		t, err := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", timePeriod)
		if err != nil {
			log.Println("time parse 1 err", err)
			return chat.GetInfoFromBatchResponse{}, nil
		}
		timePeriod = t.Format("2006-01-02")
	}
	if len(req.Periods) < 1 || len(req.Stations) < 1 {
		return chat.GetInfoFromBatchResponse{}, errors.New("no data given")
	}
	log.Println("starting exec", "'"+req.Stations[0]+"'", "'"+timePeriod+"'")

	rows, err := p.db.QueryContext(ctx, requestDB, req.Stations[0], timePeriod)
	if err != nil {
		log.Println(err)
		return chat.GetInfoFromBatchResponse{}, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&respDB.PeopleFlow)
	}

	log.Println("AGA ", respDB.PeopleFlow)
	return chat.GetInfoFromBatchResponse{
		PeopleFlow: respDB.PeopleFlow,
		Periods:    req.Periods,
		Stations:   req.Stations[0],
	}, nil
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
