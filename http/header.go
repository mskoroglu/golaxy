package http

import (
	"strconv"
	"errors"
	"fmt"
)

type Header map[string][]string

func (h *Header) value(key string) ([]string, error) {
	for k, val := range *h {
		if k == key {
			return val, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Key[%v] not found!", key))
}

func (h *Header) GetStringArray(key string) ([]string, error) {
	value, err := h.value(key)
	return value, err
}

func (h *Header) GetString(key string) (string, error) {
	value, err := h.GetStringArray(key)
	if err != nil {
		return "", err
	}
	return value[0], err
}

func (h *Header) GetInt32Array(key string) ([]int32, error) {
	stringArray, err := h.value(key)
	if err != nil {
		return nil, err
	}
	c := make([]int32, 0)
	for _, v := range stringArray {
		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		value := int32(val)
		c = append(c, value)
	}
	return c, nil
}

func (h *Header) GetInt32(key string) (int32, error) {
	value, err := h.GetInt32Array(key)
	if err != nil {
		return 0, err
	}
	return value[0], err
}

func (h *Header) GetFloat32Array(key string) ([]float32, error) {
	stringArray, err := h.value(key)
	if err != nil {
		return nil, err
	}
	c := make([]float32, 0)
	for _, v := range stringArray {
		val, _ := strconv.Atoi(v)
		value := float32(val)
		c = append(c, value)
	}
	return c, nil
}

func (h *Header) GetFloat32(key string) (float32, error) {
	value, err := h.GetFloat32Array(key)
	if err != nil {
		return 0, err
	}
	return value[0], err
}
