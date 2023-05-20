package server

import (
	"awesomeProject/internal/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) FetchRecordsView(c *gin.Context) {
	ctx := c.Request.Context()
	var input *models.RecordInp
	bytes, err := ioutil.ReadAll(c.Request.Body)
	err = json.Unmarshal(bytes, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  0,
			"msg":   "Missing Params or Date format is Not Correct",
			"error": err,
		})

		return
	}
	resp, err := s.RecordHandler.FetchRecords(ctx, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "records": resp, "msg": "success"})
	return
}
