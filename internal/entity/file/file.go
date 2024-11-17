package file

import (
	"encoding/json"
	"github.com/google/uuid"
)

type ID uuid.UUID

type File struct {
	ID          *ID
	Name        string
	Path        string
	ContentType string
	Size        int64
	Metadata    json.RawMessage
	MD5         string
}

func (e *File) GenerateID() {
	newUUID, _ := uuid.NewUUID()
	e.ID = (*ID)(&newUUID)
}
