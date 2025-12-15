package controllers

import (
	repo "killrvideo/go-backend-astra-cql/repository"
	"net/http"

	apachegocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/gin-gonic/gin"
)

type VideoController struct {
	videoDAL repo.VideoDAL
}

func NewVideoController(session *apachegocql.Session) *VideoController {
	return &VideoController{
		videoDAL: *repo.NewVideoDAL(session),
	}
}

func (vc *VideoController) GetVideo(c *gin.Context) {
	id, err1 := apachegocql.ParseUUID(c.Param("id"))
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
	}

	video, err2 := vc.videoDAL.GetVideo(id)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
	}

	c.JSON(http.StatusOK, video)
}
