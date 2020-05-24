package holiday

import (
	"encoding/json"
)

func parseHolidayFromJson(content []byte) []Holiday {
	var holidays []Holiday
	json.Unmarshal(content, &holidays)
	return holidays
}
