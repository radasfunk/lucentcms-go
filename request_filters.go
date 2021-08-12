package lucentcmsgo

import "fmt"

func (lr *Request) Where(key string, value interface{}) {
	key = "filter[" + key + "]"
	lr.Filters[key] = value
}

func (lr *Request) makeFilter(key string, value interface{}, op string) {
	lr.Filters[fmt.Sprintf("filter[%s][%s]", op, key)] = value
}

func (lr *Request) OrWhere(key string, value interface{}) {
	lr.makeFilter(key, value, "or")
}

func (lr *Request) In(key string, value string) {
	lr.makeFilter(key, value, "in")
}

func (lr *Request) Regex(key string, value string) {
	lr.makeFilter(key, value, "regex")
}

func (lr *Request) Exists(key string) {
	lr.makeFilter(key, nil, "exists")
}

func (lr *Request) NExists(key string) {
	lr.makeFilter(key, nil, "nexists")
}

func (lr *Request) Eq(key string, value interface{}) {
	lr.makeFilter(key, value, "eq")
}

func (lr *Request) Ne(key string, value interface{}) {
	lr.makeFilter(key, value, "ne")
}

func (lr *Request) Nin(key string, value interface{}) {
	lr.makeFilter(key, value, "nin")
}

func (lr *Request) Lt(key string, value interface{}) {
	lr.makeFilter(key, value, "lt")
}

func (lr *Request) Lte(key string, value interface{}) {
	lr.makeFilter(key, value, "lte")
}

func (lr *Request) Gt(key string, value interface{}) {
	lr.makeFilter(key, value, "gt")
}

func (lr *Request) Gte(key string, value interface{}) {
	lr.makeFilter(key, value, "gte")
}

func (lr *Request) True(key string) {
	lr.makeFilter(key, nil, "true")
}

func (lr *Request) False(key string) {
	lr.makeFilter(key, nil, "false")
}

func (lr *Request) Null(key string) {
	lr.makeFilter(key, nil, "null")
}

func (lr *Request) Nil(key string) {
	lr.makeFilter(key, nil, "null")
}

func (lr *Request) Empty(key string) {
	lr.makeFilter(key, nil, "empty")
}
