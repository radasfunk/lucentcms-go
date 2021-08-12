package lucentcmsgo

import "fmt"

func (lr *LucentRequest) Where(key string, value interface{}) {
	key = "filter[" + key + "]"
	lr.Filters[key] = value
}

func (lr *LucentRequest) makeFilter(key string, value interface{}, op string) {
	lr.Filters[fmt.Sprintf("filter[%s][%s]", op, key)] = value
}

func (lr *LucentRequest) OrWhere(key string, value interface{}) {
	lr.makeFilter(key, value, "or")
}

func (lr *LucentRequest) In(key string, value string) {
	lr.makeFilter(key, value, "in")
}

func (lr *LucentRequest) Regex(key string, value string) {
	lr.makeFilter(key, value, "regex")
}

func (lr *LucentRequest) Exists(key string) {
	lr.makeFilter(key, nil, "exists")
}

func (lr *LucentRequest) NExists(key string) {
	lr.makeFilter(key, nil, "nexists")
}

func (lr *LucentRequest) Eq(key string, value interface{}) {
	lr.makeFilter(key, value, "eq")
}

func (lr *LucentRequest) Ne(key string, value interface{}) {
	lr.makeFilter(key, value, "ne")
}

func (lr *LucentRequest) Nin(key string, value interface{}) {
	lr.makeFilter(key, value, "nin")
}

func (lr *LucentRequest) Lt(key string, value interface{}) {
	lr.makeFilter(key, value, "lt")
}

func (lr *LucentRequest) Lte(key string, value interface{}) {
	lr.makeFilter(key, value, "lte")
}

func (lr *LucentRequest) Gt(key string, value interface{}) {
	lr.makeFilter(key, value, "gt")
}

func (lr *LucentRequest) Gte(key string, value interface{}) {
	lr.makeFilter(key, value, "gte")
}

func (lr *LucentRequest) True(key string) {
	lr.makeFilter(key, nil, "true")
}

func (lr *LucentRequest) False(key string) {
	lr.makeFilter(key, nil, "false")
}

func (lr *LucentRequest) Null(key string) {
	lr.makeFilter(key, nil, "null")
}

func (lr *LucentRequest) Nil(key string) {
	lr.makeFilter(key, nil, "null")
}

func (lr *LucentRequest) Empty(key string) {
	lr.makeFilter(key, nil, "empty")
}
