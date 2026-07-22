package httpclient

import (
	"net/url"
	"reflect"
	"strconv"
)

// SortDirection specifies the ordering of list results supported by CI APIs.
type SortDirection string

const (
	SortAscending  SortDirection = "asc"
	SortDescending SortDirection = "desc"
)

// AppendQueryParams builds a query string from the struct fields of opts
// that are tagged with "query". Only non-zero-value fields are included.
// Use the "query-default" tag to supply a value when the field is zero-valued.
//
// Example:
//
//	type ListOptions struct {
//		Page    int    `query:"page"`
//		PerPage int    `query:"per_page" query-default:"20"`
//		Status  string `query:"status"`
//	}
func AppendQueryParams(path string, opts any) string {
	if opts == nil {
		return path
	}

	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return path
	}

	q := url.Values{}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		ft := t.Field(i)

		tag := ft.Tag.Get("query")
		if tag == "" {
			continue
		}

		defaultVal := ft.Tag.Get("query-default")

		switch field.Kind() {
		case reflect.String:
			s := field.String()
			if s == "" && defaultVal != "" {
				s = defaultVal
			}
			if s != "" {
				q.Set(tag, s)
			}
		case reflect.Int, reflect.Int64:
			n := field.Int()
			if n == 0 && defaultVal != "" {
				n, _ = strconv.ParseInt(defaultVal, 10, 64)
			}
			if n > 0 {
				q.Set(tag, strconv.FormatInt(n, 10))
			}
		case reflect.Bool:
			if field.Bool() {
				q.Set(tag, "true")
			}
		}
	}

	if len(q) == 0 {
		return path
	}

	return path + "?" + q.Encode()
}
