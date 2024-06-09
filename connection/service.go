package connection

import (
	"fmt"
	"log"
)

/*** 链路保持 ***/

// LinkKeepHandler 处理链路保持相关的命令
type LinkKeepHandler struct{}

func (h *LinkKeepHandler) Execute(cmd Command) {
	switch cmd.Type {
	case "LK":
		// 链路保持命令不需要参数，直接回复相同的消息
		fmt.Printf("[DW*%s*%s*%s]\n", cmd.DeviceID, cmd.Params, cmd.Type)
	case "PING":
		// PING命令用于检测终端是否绑定
		// 假设我们检查绑定状态并设置bindStatus变量
		bindStatus := "0" // 假设终端未绑定，设置为"0"
		// 实际应用中，您需要根据实际情况来设置bindStatus的值
		fmt.Printf("[DW*%s*%s*%s,%s]\n", cmd.DeviceID, cmd.Params, cmd.Type, bindStatus)
	case "KA":
		// KA命令包含日期、步数、翻滚次数和电量百分数
		// 此处只是简单地回复相同的消息类型
		fmt.Printf("[DW*%s*%s*%s]\n", cmd.DeviceID, cmd.Params, cmd.Type)
	default:
		fmt.Println("Unknown command type for LinkKeepHandler")
	}
}

/*** 位置数据上报 ***/

// LocationDataHandler 处理位置数据上报相关的命令
type LocationDataHandler struct{}

func (h *LocationDataHandler) Execute(cmd Command) {
	switch cmd.Type {
	case "UD", "UD_CDMA", "UP":
		// 这里我们简单地打印出位置数据，实际应用中应该将数据存储或进一步处理
		fmt.Printf("Received location data: [DeviceID: %s, Type: %s, Params: %s]\n", cmd.DeviceID, cmd.Type, cmd.Params)
	default:
		log.Printf("Unknown command type for LocationDataHandler: %s", cmd.Type)
	}
	// 注意：根据命令说明，平台不需要回复位置数据上报命令
}
