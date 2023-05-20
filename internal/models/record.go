package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type RecordResp struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalMarks uint      `json:"totalMarks"`
}

type Record struct {
	ID        uint      `json:"id" gorm:"primarykey;autoIncrement"`
	Name      string    `json:"name"`
	Marks     string    `json:"marks"`
	CreatedAt time.Time `json:"createdAt"`
}

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type RecordInp struct {
	StartDate CustomTime `json:"startDate"`
	EndDate   CustomTime `json:"endDate"`
	MinCount  uint       `json:"minCount"`
	MaxCount  uint       `json:"maxCount"`
}

func parseTime(t CustomTime) time.Time {
	pTime, err := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	if err != nil {
		fmt.Println(err)
	}
	return pTime
}

func (inp RecordInp) FetchRecordsByCreatedAt(DB *gorm.DB) (error, []Record) {
	var records []Record
	date1 := parseTime(inp.StartDate)
	date2 := parseTime(inp.EndDate)
	var query string
	var values []interface{}
	if date1.Equal(date2) {
		query = "created_at = ? "
		values = append(values, date1)
	} else if date1.After(date2) {
		return nil, nil
	} else {
		query = "created_at BETWEEN ? AND ?"
		values = append(values, date1)
		values = append(values, date2)
	}
	res := DB.Where(query, values...).Find(&records)
	if res.Error != nil {
		return res.Error, records
	}
	return nil, records
}
