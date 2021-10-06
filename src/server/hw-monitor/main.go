package main

import (
	"fmt"
	"github.com/detecc/detecctor/cache"
	"github.com/detecc/detecctor/database"
	"github.com/detecc/detecctor/server/plugin"
	"github.com/detecc/detecctor/shared"
	"log"
)

func init() {
	hwPlugin := &HardwareMonitorPlugin{}
	plugin.Register(hwPlugin.GetCmdName(), hwPlugin)
}

type HardwareMonitorPlugin struct {
	plugin.Handler
}

func (e HardwareMonitorPlugin) GetCmdName() string {
	return "/get-hw-status"
}

func (e HardwareMonitorPlugin) Response(payload shared.Payload) shared.Reply {
	log.Println(payload)
	var content = "An unexpected error occurred."

	chatId, isFound := cache.Memory().Get(payload.Id)
	if !isFound {
		log.Println("not found")
		return shared.Reply{}
	}

	if payload.Success {
		hwInfo := payload.Data.(map[string]interface{})

		content = fmt.Sprintf("ID: %s", payload.ServiceNodeKey)
		content = fmt.Sprintf("CPU usage (in percent): %.2f\n", hwInfo["cpu"])
		content = content + fmt.Sprintf("Memory available: %.2f MB\n", hwInfo["mem-available"])
		content = content + fmt.Sprintf("Memory used: %.2f MB\n", hwInfo["mem-used"])
		content = content + fmt.Sprintf("Total memory: %.2f MB\n", hwInfo["mem-total"])
	}

	return shared.Reply{
		ChatId:    chatId.(int64),
		ReplyType: shared.TypeMessage,
		Content:   content,
	}
}

func (e HardwareMonitorPlugin) Execute(args ...string) ([]shared.Payload, error) {
	var payloads []shared.Payload
	for _, arg := range args {
		// check if the desired client exists
		key, err := database.GetClientWithServiceNodeKey(arg)
		if err != nil {
			continue
		}

		payloads = append(payloads, shared.Payload{
			Id:             "",
			ServiceNodeKey: key.ServiceNodeKey,
			Data:           nil,
			Command:        e.GetCmdName(),
			Success:        true,
			Error:          "",
		})
	}
	log.Println(args)
	return payloads, nil
}

func (e HardwareMonitorPlugin) GetMetadata() plugin.Metadata {
	return plugin.Metadata{
		Type:       plugin.PluginTypeServerClient,
		Middleware: []string{},
	}
}
