package student

import (
	"awesomeProject/internal/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type record struct {
	DB *gorm.DB
}

func New(db *gorm.DB) record {
	return record{DB: db}
}

func marksSum(marks []uint) uint {
	var sum uint
	for _, val := range marks {
		sum += val
	}
	return sum
}

func (r record) FetchRecords(ctx context.Context, input *models.RecordInp) (*[]models.RecordResp, error) {
	fmt.Println(input)
	err, records := input.FetchRecordsByCreatedAt(r.DB)
	if err != nil {
		return nil, errors.New("content not found")
	}
	var result []models.RecordResp
	for _, val := range records {

		var resp models.RecordResp
		type mark struct {
			Marks []uint `json:"marks"`
		}
		var m mark

		_ = json.Unmarshal([]byte(val.Marks), &m)
		sum := marksSum(m.Marks)
		if input.MinCount <= sum && input.MaxCount >= sum {
			resp.ID = val.ID
			resp.CreatedAt = val.CreatedAt
			resp.TotalMarks = sum
			result = append(result, resp)
		}

	}
	return &result, nil
}
