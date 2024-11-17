package history

import "fmt"

type ActionType string

const (
	Unknown ActionType = "UNKNOWN"
	Create  ActionType = "CREATE"
	Update  ActionType = "UPDATE"
	Delete  ActionType = "DELETE"
)

var actionTypes = map[string]ActionType{
	"CREATE": Create,
	"UPDATE": Update,
	"DELETE": Delete,
}

func New(typeName string) (ActionType, error) {
	actionType, ok := actionTypes[typeName]
	if !ok {
		return Unknown, fmt.Errorf("unknown action type: %s", typeName)
	}
	return actionType, nil
}

func (e ActionType) String() string {
	return string(e)
}
