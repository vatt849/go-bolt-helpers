package helpers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func SafeSliceAccess(data interface{}, index int) (val reflect.Value) {
	if data == nil || reflect.TypeOf(data).Kind() != reflect.Slice {
		return
	}
	tmp := reflect.ValueOf(data)
	tmpLen := tmp.Len()
	if tmpLen > 0 && index < 0 {
		val = tmp.Index(tmpLen - 1)
	} else if tmpLen > 0 && index < tmpLen {
		val = tmp.Index(index)
	}
	return
}

func FilterUrl(u string) string {
	parsedUrl, err := ParseUrl(u)

	if err != nil {
		return ""
	}

	return parsedUrl.String()
}

func ParseUrl(u string) (*url.URL, error) {
	return url.ParseRequestURI(u)
}

func StrToIDArr(s string) (IDarr []uint) {
	strArr := strings.Split(s, ",")
	for _, v := range strArr {
		i, _ := strconv.Atoi(v)
		if i > 0 {
			IDarr = append(IDarr, uint(i))
		}
	}

	return
}

func IDArrToStr(IDarr []uint) string {
	var strArr []string

	for _, v := range IDarr {
		strArr = append(strArr, fmt.Sprint(v))
	}

	return strings.Join(strArr, ",")
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func OpenJson(path string, data any) (err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, data)

	return
}

func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return fmt.Sprintf("%x", h.Sum(nil))
}
