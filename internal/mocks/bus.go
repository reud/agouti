package mocks

import "encoding/json"

type Bus struct {
	SendCall struct {
		Endpoint string
		Method   string
		BodyJSON []byte
		Result   string
		Err      error
	}
}

func (b *Bus) Send(method, endpoint string, body, result interface{}) error {
	b.SendCall.Method = method
	b.SendCall.Endpoint = endpoint
	b.SendCall.BodyJSON, _ = json.Marshal(body)
	if result != nil {
		json.Unmarshal([]byte(b.SendCall.Result), result)
	}
	return b.SendCall.Err
}

func (b *Bus) SendAndGetJsonStr(method, endpoint string, body interface{}, resultStr *string) error {
	b.SendCall.Method = method
	b.SendCall.Endpoint = endpoint
	b.SendCall.BodyJSON, _ = json.Marshal(body)
	if resultStr != nil {
		*resultStr = b.SendCall.Result
	}
	return b.SendCall.Err
}
