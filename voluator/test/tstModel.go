/*
@Time : 2019/9/30 14:21
@Author : Hermes
@File : tstModel
@Description:
*/
package test

//回应数据结构
type NumberInfo struct {
	ID        int    `json:"id"`             //
	DisplayNo string `json:"display_number"` //
	ChannelID int    `json:"channel_id"`     //
}

type applicationInfo struct {
	AppId                int               `json:"AppID"`       //
	AppName              string            `json:"AppName"`     //
	NotifyBlockInfoArray []notifyBlockInfo `json:"NotifyBlock"` //
	Callback             string            `json:"Callback"`    //
}

type notifyBlockInfo struct {
	ID    int    `json:"id"`    //
	Name  string `json:"name"`  //
	Level int    `json:"level"` //
}

type subject struct {
	NumberInfo      NumberInfo      `json:"number"`      //
	ApplicationInfo applicationInfo `json:"application"` //
}

type rootResponse struct {
	ID        int      `json:"id"`
	Gap       float32  `json:"gap"`
	Data      string   `json:"data"`
	Subject   subject  `json:"subject,omitempty"`
	Timestamp string   `json:"timestamp"`
	StrArr    []string `json:"sa"`
}

type cominfo struct {
	NumberInfo
	Data      string      `json:"data"`
	Timestamp string      `json:"timestamp"`
	Info      WrapperInfo `json:"wrap"`
}

type WrapperInfo struct {
	NumberInfo
	Data string `json:"data"`
}
