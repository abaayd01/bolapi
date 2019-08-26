package main

import (
	"bolapi/internal/pkg/database"
	"log"
)

func init() {
	err := database.InitDB()

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	resetDB()
}

func resetDB() {
	dropAllTables()
	createPriceSnapshotsTable()
	createPriceEvaluationsTable()
}

func createPriceSnapshotsTable() {
	createPriceSnapshotsTableQuery := `
	CREATE TABLE price_snapshots (
	    price_snapshot_id INT GENERATED ALWAYS AS IDENTITY,
	    created_time DATE,
	    price NUMERIC,
	    PRIMARY KEY (price_snapshot_id)
	)
	`

	_, err := database.DB.Exec(createPriceSnapshotsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func createPriceEvaluationsTable() {
	createPriceEvaluationsTableQuery := `
	CREATE TABLE price_evaluations (
	    price_evaluation_id INT GENERATED ALWAYS AS IDENTITY,
	    price_snapshot_id INT REFERENCES public.price_snapshots(price_snapshot_id),
	    created_time DATE,
	    action VARCHAR,
	    evaluation_price NUMERIC,
	    target_exit_price NUMERIC,
	    stop_loss_price NUMERIC,
	    bol_upper NUMERIC,
	    bol_lower NUMERIC,
	    moving_average NUMERIC,
	    PRIMARY KEY (price_evaluation_id, price_snapshot_id)
	)
	`

	_, err := database.DB.Exec(createPriceEvaluationsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func dropAllTables() {
	dropAllTablesQuery := `
	DROP TABLE price_snapshots;
	DROP TABLE price_evaluations;
	`

	_, err := database.DB.Exec(dropAllTablesQuery)
	if err != nil {
		log.Fatal(err)
	}
}
