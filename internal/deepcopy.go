package internal

import "encoding/json"

func DeepCopy[T any](v T) T {
	data, _ := json.Marshal(v)

	var out T
	json.Unmarshal(data, &out)
	return out
}