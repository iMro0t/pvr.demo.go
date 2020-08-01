package pvr

import (
	"unsafe"
)

/*
#include "../kodi/xbmc_pvr_dll.h"
*/
import "C"

var (
	channels []Channel
)

type Channel struct {
	ID        int
	IsRadio   bool
	Number    int
	SubNumber int
	Name      string
	IconPath  string
	IsHidden  bool
	Live      Stream
	Catchup   Stream
}

type Stream struct {
	URL        string
	Properties map[string]string
}

type ChannelGroup struct {
	Name     string
	IsRadio  bool
	Position int
}

func AddChannel(channel Channel) {
	channels = append(channels, channel)
}

func GetChannel(channel *C.cPVR_CHANNEL_t) *Channel {
	for _, c := range channels {
		if c.ID == (int)(channel.iUniqueId) {
			return &c
		}
	}
	return nil
}

//export GetChannelsAmount
func GetChannelsAmount() C.int {
	return (C.int)(len(channels))
}

//export GetChannels
func GetChannels(handle C.ADDON_HANDLE, isRadio C.bool) C.PVR_ERROR {
	for _, channel := range channels {
		PVR.TransferChannelEntry(unsafe.Pointer(handle), channel)
	}
	return C.PVR_ERROR_NO_ERROR
}

//export GetChannelStreamProperties
func GetChannelStreamProperties(channel *C.cPVR_CHANNEL_t, props *C.struct_PVR_NAMED_VALUE, propsCount *C.uint) C.PVR_ERROR {
	if channel == nil || props == nil || propsCount == nil || len(channels) == 0 {
		return C.PVR_ERROR_SERVER_ERROR
	}
	if (int)(*propsCount) < 2 {
		return C.PVR_ERROR_INVALID_PARAMETERS
	}
	if c := GetChannel(channel); c != nil {
		count := PVR.SetProperties(props, c.Live)
		*propsCount = C.uint(count)
		return C.PVR_ERROR_NO_ERROR
	}
	return C.PVR_ERROR_SERVER_ERROR
}

//export GetChannelGroupsAmount
func GetChannelGroupsAmount() C.int {
	return -1
}

//export GetChannelGroups
func GetChannelGroups(handle C.ADDON_HANDLE, isRadio C.bool) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export GetChannelGroupMembers
func GetChannelGroupMembers(handle C.ADDON_HANDLE, group *C.cPVR_CHANNEL_GROUP_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export OpenDialogChannelScan
func OpenDialogChannelScan() C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export DeleteChannel
func DeleteChannel(channel *C.cPVR_CHANNEL_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export RenameChannel
func RenameChannel(channel *C.cPVR_CHANNEL_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export OpenDialogChannelSettings
func OpenDialogChannelSettings(channel *C.cPVR_CHANNEL_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export OpenDialogChannelAdd
func OpenDialogChannelAdd(channel *C.cPVR_CHANNEL_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}
