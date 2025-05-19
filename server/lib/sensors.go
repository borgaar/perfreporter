package sensors

import (
	"encoding/json"
	"os"
)

func ReadSensorData() (map[string]any, error) {
	bytes, err := os.ReadFile("/server/data.json")

	if err != nil {
		return nil, err
	}

	var output map[string]any

	parseErr := json.Unmarshal(bytes, &output)

	if parseErr != nil {
		return nil, parseErr
	}

	return output, nil
}
