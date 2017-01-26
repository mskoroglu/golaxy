package request

import (
	"strconv"
	"errors"
	"fmt"
)

// HTTP istekleri ile gelen parametreleri döndürür.
// Tip dönüşümleriyle uğraşmak zorunda kalınmadan, uygun metod kullanılarak veri elde edilebilir.
// Aynı ada sahip parametreleri elde edebilmek için ilgili tipin array döndüren metodu kullanılabilir.
type Param map[string][]string

func (p *Param) value(key string) ([]string, error) {
	for k, val := range *p {
		if k == key {
			return val, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Key[%v] not found!", key))
}

func (p *Param) GetStringArray(key string) ([]string, error) {
	value, err := p.value(key)
	return value, err
}

func (p *Param) GetString(key string) (string, error) {
	value, err := p.GetStringArray(key)
	if err != nil {
		return "", err
	}
	return value[0], err
}

func (p *Param) GetInt32Array(key string) ([]int32, error) {
	stringArray, err := p.value(key)
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

func (p *Param) GetInt32(key string) (int32, error) {
	value, err := p.GetInt32Array(key)
	if err != nil {
		return 0, err
	}
	return value[0], err
}

func (p *Param) GetFloat32Array(key string) ([]float32, error) {
	stringArray, err := p.value(key)
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

func (p *Param) GetFloat32(key string) (float32, error) {
	value, err := p.GetFloat32Array(key)
	if err != nil {
		return 0, err
	}
	return value[0], err
}
