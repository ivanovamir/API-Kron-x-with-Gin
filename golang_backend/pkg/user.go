package pkg

import (
	"fmt"
	"strconv"
)

func UserIdChecker(userId string) (int, error) {

	userUUID, _ := strconv.Atoi(userId)

	switch userUUID {
	case 0:
		return 0, fmt.Errorf("invalid user id")
	default:
		return userUUID, nil
	}
}
