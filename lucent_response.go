package lucentcmsgo

import "fmt"

type LucentResponse struct {
	// raw todo add the raw response
	Data           []interface{}
	Errors, Links  []string
	Meta, Included map[string]interface{} // TODO add paginator
}

func makeResponse() {
	fmt.Printf("Get response")
}

func (lrr *LucentResponse) HasErrors() bool {
	return len(lrr.Errors) > 0
}

func (lrr *LucentResponse) GetData() []interface{} {
	return lrr.Data
}

func (lrr *LucentResponse) GetIncluded() map[string]interface{} {
	return lrr.Included
}

// returns all the errors
func (lrr *LucentResponse) GetErrors() []string {
	return lrr.Errors
}

// returns the first error
func (lrr *LucentResponse) Error() string {
	if lrr.HasErrors() {
		return lrr.Errors[0]
	}

	return ""
}
