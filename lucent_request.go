package lucentcmsgo 

import "fmt"

type LucentRequest struct {
	Method, EndPoint string
	Headers          map[string]string
	Data             interface{}
}

func (lr *LucentRequest) AddHeaders(headers ...map[string]string) {
	fmt.Printf("add headers to request")
}

func (lr *LucentRequest) Send() {
	fmt.Printf("send the request")
}