package util

/**
 * @Author: bavelee
 * @Date: 2023/3/18 0:12
 * @Desc:
 */

func ToIntPtr(i int) *int {
	return &i
}

func ToStringPtr(s string) *string {
	return &s
}
