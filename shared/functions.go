package shared

import "strconv"

func IntOrNil(param string) *int {
	if len(param) > 0 {
		value, err := strconv.Atoi(param)
		if err != nil {
			return nil
		} else {
			return &value
		}
	}
	return nil
}
