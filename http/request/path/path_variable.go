package path

import (
	"strconv"
)

type Variable map[string]string

func (v *Variable) value(key string) string {
	for k, val := range *v {
		if k == key {
			return val
		}
	}
	return ""
}

func (v *Variable) GetString(key string) string {
	return v.value(key)
}

func (v *Variable) GetInt32(key string) int32 {
	val, _ := strconv.Atoi(v.value(key))
	return int32(val)
}

func (v *Variable) GetFloat32(key string) float32 {
	val, _ := strconv.Atoi(v.value(key))
	return float32(val)
}
