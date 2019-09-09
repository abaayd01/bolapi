package repositories

import (
	"bolapi/internal/pkg/database"
	"bolapi/internal/pkg/models"
	"log"
)

type PriceSnapshotRepository interface {
	Insert(priceSnapshot *models.PriceSnapshot) error
	GetRecords() error
}

type priceSnapshotRepository struct {
	DB database.DBInterface
}

func NewPriceSnapshotRepository(DB database.DBInterface) *priceSnapshotRepository {
	return &priceSnapshotRepository{DB: DB}
}

func (p *priceSnapshotRepository) Insert(priceSnapshot *models.PriceSnapshot) (*models.PriceSnapshot, error) {
	query := `
		INSERT INTO price_snapshots (created_time, price) VALUES ($1, $2) RETURNING price_snapshot_id
	`

	var id int
	err := p.DB.QueryRow(query, priceSnapshot.CreatedTime, priceSnapshot.Price).Scan(&id)

	if err != nil {
		log.Printf("Error inserting PriceSnapshot")
		return nil, err
	}

	priceSnapshot.Id = id
	return priceSnapshot, nil
}

// GetLatest returns up to the limit number of price snapshots.
// By default snapshots are returned in reverse chronological order, that is,
// the most recent snapshot is at index 0, and the oldest snapshot is
// in the last position in the returned slice.
func (p *priceSnapshotRepository) GetLatest(limit int) (*models.PriceSnapshots, error) {
	query := `
		SELECT * FROM price_snapshots ORDER BY created_time DESC LIMIT $1
	`

	var priceSnapshots models.PriceSnapshots

	rows, err := p.DB.Queryx(query, limit)

	if err != nil {
		log.Printf("Error retrieving PriceSnapshots")
		return nil, err
	}

	if rows == nil {
		log.Printf("Error nil rows returned from query")
		return nil, err
	}

	for rows.Next() {
		var priceSnapshot models.PriceSnapshot
		err = rows.StructScan(&priceSnapshot)
		priceSnapshots = append(priceSnapshots, priceSnapshot)
	}

	if err != nil {
		log.Printf("Error retrieving PriceSnapshots")
		return nil, err
	}

	return &priceSnapshots, nil
}
