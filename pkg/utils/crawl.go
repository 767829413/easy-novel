package utils

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// NormalizeURL 将相对路径转换为绝对路径，并确保 URL 格式正确
func NormalizeURL(s, host string) string {

	// 检查 s 是否已经是一个完整的 URL
	if strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") {
		return cleanURL(s)
	}

	// 如果 s 不是完整 URL，则将其与 host 组合
	baseURL, err := url.Parse(host)
	if err != nil {
		return ""
	}

	relativeURL, err := url.Parse(s)
	if err != nil {
		return ""
	}

	return cleanURL(baseURL.ResolveReference(relativeURL).String())
}

// cleanURL 清理 URL，移除多余的斜杠和处理编码
func cleanURL(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return s
	}

	// 确保路径不以 '//' 开始
	u.Path = strings.TrimPrefix(u.Path, "//")

	// 移除路径中的连续斜杠
	for strings.Contains(u.Path, "//") {
		u.Path = strings.ReplaceAll(u.Path, "//", "/")
	}

	// 确保查询参数正确编码
	q := u.Query()
	u.RawQuery = q.Encode()

	return u.String()
}

// BuildParams builds a map of parameters from a JSON string and a keyword
func BuildParams(body string, keyword string) (map[string]string, error) {
	if body == "" {
		return nil, errors.New("empty body provided")
	}
	if keyword == "" {
		return nil, errors.New("empty keyword provided")
	}

	params := make(map[string]string)

	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &jsonMap)
	if err != nil {
		return nil, err
	}

	for key, value := range jsonMap {
		if key == "kw" {
			params[value.(string)] = keyword
		} else {
			params[key] = value.(string)
		}
	}

	return params, nil
}

// BuildCookies builds a map of cookies from a JSON string
func BuildCookies(cookies string) (map[string]string, error) {
	if cookies == "" {
		return nil, errors.New("empty cookies string provided")
	}

	cookieMap := make(map[string]string)

	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(cookies), &jsonMap)
	if err != nil {
		return nil, err
	}

	for key, value := range jsonMap {
		cookieMap[key] = value.(string)
	}

	return cookieMap, nil
}

// BuildMethod returns the corresponding HTTP method
func BuildMethod(method string) string {
	if method == "" {
		return http.MethodPost // 默认返回 POST 方法
	}

	switch strings.ToLower(method) {
	case "get":
		return http.MethodGet
	case "post":
		return http.MethodPost
	case "put":
		return http.MethodPut
	case "delete":
		return http.MethodDelete
	case "patch":
		return http.MethodPatch
	case "head":
		return http.MethodHead
	case "options":
		return http.MethodOptions
	case "trace":
		return http.MethodTrace
	default:
		return http.MethodPost
	}
}

// RandomSleep sleeps for a random duration between min and max milliseconds
func RandomSleep(min, max int64) error {
	if min < 0 || max < 0 {
		return errors.New("min and max must be non-negative")
	}
	if min > max {
		return errors.New("min must be less than or equal to max")
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	duration := random.Int63n(max-min+1) + min
	time.Sleep(time.Duration(duration) * time.Millisecond)
	return nil
}
