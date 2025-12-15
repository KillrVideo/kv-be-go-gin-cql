package repository

import (
	"fmt"
	"killrvideo/go-backend-astra-cql/models"

	apachegocql "github.com/apache/cassandra-gocql-driver/v2"
)

type VideoDAL struct {
	DB *apachegocql.Session
}

func NewVideoDAL(session *apachegocql.Session) *VideoDAL {
	return &VideoDAL{
		DB: session,
	}
}

func (r *VideoDAL) GetVideo(id apachegocql.UUID) (*models.Video, error) {
	video := &models.Video{Videoid: id}
	//var vector []float32

	err1 := r.DB.Query(
		"SELECT userid, name, description, location, preview_image_location, added_date, content_features FROM videos WHERE videoid = ?",
		id,
	).Scan(&video.Userid, &video.Name, &video.Description, &video.Location, &video.PreviewImageLocation, &video.AddedDate, &video.ContentFeatures)

	if err1 != nil {
		return nil, fmt.Errorf("query has failed: %w", err1)
	}

	return video, nil
}
