package util

import (
	"encoding/json"
)

/**
 * @Author: bavelee
 * @Date: 2023/3/17 20:24
 * @Desc:
 */

func ToJson(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
