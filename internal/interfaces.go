package internal

import (
	"awesomeProject/internal/models"
	"context"
)

type RecordHandler interface {
	FetchRecords(ctx context.Context, input *models.RecordInp) (*[]models.RecordResp, error)
}
