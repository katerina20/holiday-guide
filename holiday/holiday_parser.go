package holiday

import (
	"encoding/json"
	"fmt"
)

func parseHolidayFromJson(content []byte) []Holiday {
	var holidays []Holiday
	var err = json.Unmarshal(content, &holidays)
	if err != nil {
		fmt.Println(err)
	}
	return holidays
}
