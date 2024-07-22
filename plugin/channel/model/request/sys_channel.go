package request

type ChannelRequest struct {
	AppId string `json:"appId" form:"appId"`
}
type UpdateChannelRequest struct {
	AppId string `json:"appId" form:"appId"`
	Field string `json:"field" form:"field"`
}
