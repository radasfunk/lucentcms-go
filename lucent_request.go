package lucentcmsgo

import "fmt"

type LucentRequest struct {
	Method, EndPoint string
	Headers          map[string]string
	Data             interface{}
}

func (lr *LucentRequest) AddHeaders(headers map[string]string) {

	for key, value := range headers {
		// will not allow users to change protected headers
		if _, ok := protectedHeaders[key]; ok {
			continue
		}

		lr.Headers[key] = value
	}
}

func (lr *LucentRequest) Send() {
	fmt.Printf("send the request")
}
