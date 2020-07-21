package test

import "sync"

type SdSvcInfo struct {
	NodeID string
	IP     string
}

type ICallTaskCache interface {
	GetTaskName() string
	GetTaskId() string
}

type SessionCache struct {
	startTs     int64  //时间戳，nano
	Name        string `json:"name"`
	CallCap     int    `json:"callCap"`
	SpeakCap    int    `json:"speakCap"`
	IsMeeting   bool   `json:"isMeeting"`
	LockSession bool   `json:"lockSession"`

	ValueLocker sync.RWMutex //针对操作值的锁，主要是结构体和string
	appId       string
	sessionId   string
	userId      string
	BllType     int `json:"bllType"`
	SwitchInfo  *SdSvcInfo

	WillDelete       bool `json:"willDelete"`
	LastSyncHangup   int  `json:"lastSyncHangup"` //最后一通电话同步挂标志，0不同步挂 1同步挂
	hangupAllDoCount int32
	hangupAllDoTime  int64

	taskLocker sync.RWMutex
	//taskCaches []ICallTaskCache

	callStateLocker sync.RWMutex
	//callStates      []*bllCommModel.CallDetailState
	busyCallIds      []string
}
