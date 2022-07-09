/**
  @author:panliang
  @data:2022/7/8
  @note
**/
package requests

type PrivateMessageRequest struct {
	MsgId       int64       `json:"msg_id" validate:"required"`        // 服务端消息唯一id
	MsgClientId int64       `json:"msg_client_id" validate:"required"` // 客户端消息唯一id
	MsgCode     int         `json:"msg_code" validate:"required"`      // 定义的消息code
	FormID      int64       `json:"form_id" validate:"required"`       // 发消息的人
	ToID        int64       `json:"to_id" validate:"required"`         // 接收消息人的id
	MsgType     int         `json:"msg_type" validate:"required"`      // 消息类型 例如 1.文本 2.语音 3.文件
	ChannelType int         `json:"channel_type" validate:"required"`  // 频道类型 1.私聊 2.频道 3.广播
	Message     string      `json:"message" validate:"required"`       // 消息
	SendTime    string      `json:"send_time"`                         // 消息发送时间
	Data        interface{} `json:"data"`                              // 自定义携带的数据
}
