package main

import (
	"github.com/DisgoOrg/disgo/core"
	"strings"
)

func onSlashCommand(event *core.SlashCommandEvent) {
	switch event.CommandName {
	case "sounds":
		switch *event.SubCommandName {
		case "add":
			go onSoundAddCommand(event)
		case "remove":
			go onSoundRemoveCommand(event)
		}
	case "soundboard":
		go onSoundBoardCommand(event)
	}
}

func onButtonClick(event *core.ButtonClickEvent) {
	action := strings.Split(event.CustomID, ":")
	switch action[0] {
	case "play":
		onPlayButton(event, action[1])
	case "stop":
		onStopButton(event)
	case "pause":
		onPauseButton(event)
	case "page":
		onPageButton(event, action[1])
	}
}
