package base

import "encoding/json"

func unmarshalJSONToMap(b []byte) (map[string]any, error) {
	if len(b) == 0 || string(b) == "{}" || string(b) == "null" {
		return nil, nil
	}
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
