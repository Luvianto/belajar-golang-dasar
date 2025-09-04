package commonutils

import "github.com/google/uuid"

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

// func ParseUUID(stringUUID string) (uuid.UUID, bool) {
// 	if stringUUID == "" || stringUUID == "null" {
// 		return uuid.Nil, true
// 	}
// 	parsedUUID, err := uuid.Parse(stringUUID)
// 	if err != nil {
// 		return uuid.Nil, true
// 	}
// 	return parsedUUID, false
// }
