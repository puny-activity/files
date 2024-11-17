package history

import (
	"github.com/google/uuid"
	"github.com/puny-activity/files/internal/entity/file"
	"time"
)

type ID uuid.UUID

type Row struct {
	ID          *ID
	FileID      file.ID
	PerformedAt time.Time
}

func (e *Row) GenerateID() {
	newUUID, _ := uuid.NewUUID()
	e.ID = (*ID)(&newUUID)
}
