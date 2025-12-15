package models

import (
	"time"

	apachegocql "github.com/apache/cassandra-gocql-driver/v2"
)

type Video struct {
	Videoid              apachegocql.UUID
	Userid               apachegocql.UUID
	Name                 string
	Description          string
	Location             string
	PreviewImageLocation string
	ContentFeatures      [384]float32
	AddedDate            time.Time
}
