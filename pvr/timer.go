package pvr

/*
#cgo CFLAGS: -Wno-error=implicit-function-declaration
#include "../kodi/xbmc_pvr_dll.h"
*/
import "C"

//export GetTimerTypes
func GetTimerTypes(types *C.PVR_TIMER_TYPE, size *C.int) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export GetTimersAmount
func GetTimersAmount() C.int {
	return -1
}

//export GetTimers
func GetTimers(handle C.ADDON_HANDLE) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export AddTimer
func AddTimer(timer *C.cPVR_TIMER_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export DeleteTimer
func DeleteTimer(timer *C.cPVR_TIMER_t, bForceDelete C.bool) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export UpdateTimer
func UpdateTimer(timer *C.cPVR_TIMER_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}
