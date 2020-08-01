package pvr

/*
#include "../kodi/xbmc_pvr_dll.h"
*/
import "C"

//export GetEPGForChannel
func GetEPGForChannel(handle C.ADDON_HANDLE, channel *C.cPVR_CHANNEL_t, start C.time_t, end C.time_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export IsEPGTagPlayable
func IsEPGTagPlayable(_ *C.EPG_TAG, isPlayable *C.bool) C.PVR_ERROR {
	*isPlayable = false
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export GetEPGTagStreamProperties
func GetEPGTagStreamProperties(tag *C.EPG_TAG, props *C.PVR_NAMED_VALUE, propsCount *C.uint) C.PVR_ERROR {
	if tag == nil || props == nil || propsCount == nil {
		return C.PVR_ERROR_SERVER_ERROR
	}

	// if propsCount < 1 {

	// }
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export IsEPGTagRecordable
func IsEPGTagRecordable(*C.EPG_TAG, *C.bool) C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }

//export GetEPGTagEdl
func GetEPGTagEdl(epgTag *C.cEPG_TAG_t, edl *C.PVR_EDL_ENTRY, size *C.int) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export SetEPGTimeFrame
func SetEPGTimeFrame(C.int) C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }
