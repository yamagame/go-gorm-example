package csvconv

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func structToJsonMap(v any) (map[string]any, error) {
	byte, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var rawJson map[string]any
	decoder := json.NewDecoder(strings.NewReader(string(byte)))
	if err := decoder.Decode(&rawJson); err != nil {
		log.Fatal(err)
	}
	return rawJson, nil
}

func mapToJsonText(v map[string]any) ([]byte, error) {
	byte, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return byte, nil
}

func getMapValue(values any, key []string, idx int) (interface{}, error) {
	k := key[idx]
	r := regexp.MustCompile(`(.+)\[(\d)\]`)
	m := r.FindStringSubmatch(k)
	n := -1
	if len(m) > 0 {
		k = m[1]
		i, err := strconv.Atoi(m[2])
		if err != nil {
			return nil, err
		}
		n = i
	}
	mvalue := values.(map[string]any)
	if v, ok := mvalue[k]; ok {
		if n >= 0 {
			ar := v.([]any)
			if len(ar) >= n {
				v = ar[n]
			}
		}
		if len(key) == idx+1 {
			return v, nil
		}
		return getMapValue(v, key, idx+1)
	}
	return "", nil
}

// StructToField2 構造体を配列レコードに変換
func StructToField2(v any, keys []string) ([]interface{}, error) {
	ret := []interface{}{}
	jsonMap, err := structToJsonMap(v)
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		v, err := getMapValue(jsonMap, strings.Split(key, "."), 1)
		if err != nil {
			return nil, err
		}
		ret = append(ret, v)
	}
	return ret, nil
}

func setMapValue(values any, key []string, idx int, value interface{}) error {
	mvalue := values.(map[string]any)
	k := key[idx]
	r := regexp.MustCompile(`(.+)\[(\d)\]`)
	m := r.FindStringSubmatch(k)
	n := -1
	if len(m) > 0 {
		k = m[1]
		i, err := strconv.Atoi(m[2])
		if err != nil {
			return err
		}
		n = i
	}
	if n >= 0 {
		if len(key) > idx+1 {
			ar := []map[string]interface{}{}
			if mvalue[k] != nil {
				ar = append(ar, mvalue[k].([]map[string]interface{})...)
			}
			for len(ar) < n+1 {
				ar = append(ar, map[string]interface{}{})
			}
			mvalue[k] = ar
			return setMapValue(ar[n], key, idx+1, value)
		} else {
			if mvalue[k] == nil {
				mvalue[k] = []interface{}{}
			}
			ar := mvalue[k].([]interface{})
			for len(ar) < n+1 {
				ar = append(ar, nil)
			}
			mvalue[k] = ar
			if len(key) == idx+1 {
				ar := mvalue[k].([]interface{})
				ar[n] = value
				return nil
			}
		}
		return setMapValue(mvalue[k], key, idx+1, value)
	}
	if len(key) == idx+1 {
		mvalue[k] = value
		return nil
	}
	if _, ok := mvalue[k]; !ok {
		mvalue[k] = map[string]interface{}{}
	}
	return setMapValue(mvalue[k], key, idx+1, value)
}

// FieldToStruct 配列レコードを構造体に変換
func FieldToStruct2(v any, values map[string]interface{}) error {
	var err error
	jsonMap := map[string]interface{}{}
	for key, value := range values {
		err = setMapValue(jsonMap, strings.Split(key, "."), 1, value)
		if err != nil {
			return err
		}
	}
	jsonText, err := mapToJsonText(jsonMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonText, v)
}
