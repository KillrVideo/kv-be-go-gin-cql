package models

import (
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

type Video struct {
	videoid              gocql.UUID
	userid               gocql.UUID
	name                 string
	description          string
	location             string
	previewImageLocation string
	vector               gocql.VectorType
	addedDate            time.Time
}
