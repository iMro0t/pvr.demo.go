package pvr

/*
#include "../kodi/xbmc_pvr_dll.h"
*/
import "C"

//export GetRecordingsAmount
func GetRecordingsAmount(deleted C.bool) C.int {
	return -1
}

//export GetRecordings
func GetRecordings(handle C.ADDON_HANDLE, deleted C.bool) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export GetRecordingStreamProperties
func GetRecordingStreamProperties(recording *C.cPVR_RECORDING_t, props *C.PVR_NAMED_VALUE, propsCount *C.uint) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export OpenRecordedStream
func OpenRecordedStream(recording *C.cPVR_RECORDING_t) C.bool {
	return false
}

//export CloseRecordedStream
func CloseRecordedStream() {}

//export ReadRecordedStream
func ReadRecordedStream(pBuffer *C.uchar, iBufferSize C.uint) C.int {
	return 0
}

//export SeekRecordedStream
func SeekRecordedStream(iPosition C.longlong, iWhence C.int) C.longlong {
	return 0
}

//export LengthRecordedStream
func LengthRecordedStream() C.longlong {
	return 0
}

//export DeleteRecording
func DeleteRecording(recording *C.cPVR_RECORDING_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export RenameRecording
func RenameRecording(recording *C.cPVR_RECORDING_t) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export SetRecordingPlayCount
func SetRecordingPlayCount(recording *C.cPVR_RECORDING_t, count C.int) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export SetRecordingLastPlayedPosition
func SetRecordingLastPlayedPosition(recording *C.cPVR_RECORDING_t, lastplayedposition C.int) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export GetRecordingLastPlayedPosition
func GetRecordingLastPlayedPosition(recording *C.cPVR_RECORDING_t) C.int {
	return -1
}

//export GetRecordingEdl
func GetRecordingEdl(_ *C.cPVR_RECORDING_t, _ *C.PVR_EDL_ENTRY, _ *C.int) C.PVR_ERROR {
	return C.PVR_ERROR_NOT_IMPLEMENTED
}

//export UndeleteRecording
func UndeleteRecording(recording *C.cPVR_RECORDING_t) C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }

//export DeleteAllRecordingsFromTrash
func DeleteAllRecordingsFromTrash() C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }

//export GetDescrambleInfo
func GetDescrambleInfo(*C.PVR_DESCRAMBLE_INFO) C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }

//export SetRecordingLifetime
func SetRecordingLifetime(*C.cPVR_RECORDING_t) C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }
