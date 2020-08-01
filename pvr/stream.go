package pvr

/*
#include "../kodi/xbmc_pvr_dll.h"
*/
import "C"

//export IsRealTimeStream
func IsRealTimeStream() C.bool {
	return false
}

//export DemuxReset
func DemuxReset() {}

//export DemuxFlush
func DemuxFlush() {}

//export OpenLiveStream
func OpenLiveStream(_ *C.cPVR_CHANNEL_t) C.bool {
	return false
}

//export CloseLiveStream
func CloseLiveStream() {}

//export ReadLiveStream
func ReadLiveStream(pBuffer *C.uchar, iBufferSize C.uint) C.int {
	return 0
}

//export SeekLiveStream
func SeekLiveStream(iPosition C.longlong, iWhence C.int) C.longlong {
	return -1
}

//export LengthLiveStream
func LengthLiveStream() C.longlong {
	return -1
}

//export DemuxAbort
func DemuxAbort() {}

//export DemuxRead
func DemuxRead() *C.DemuxPacket {
	return (*C.DemuxPacket)(C.NULL)
}

//export PauseStream
func PauseStream(bPaused C.bool) {}

//export CanPauseStream
func CanPauseStream() C.bool { return false }

//export CanSeekStream
func CanSeekStream() C.bool { return false }

//export SeekTime
func SeekTime(time C.double, backwards C.bool, startpts *C.double) C.bool { return false }

//export SetSpeed
func SetSpeed(C.int) {}

//export IsTimeshifting
func IsTimeshifting() C.bool { return false }

//export GetStreamProperties
func GetStreamProperties(*C.PVR_STREAM_PROPERTIES) C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }

//export GetStreamTimes
func GetStreamTimes(*C.PVR_STREAM_TIMES) C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }

//export GetStreamReadChunkSize
func GetStreamReadChunkSize(chunksize *C.int) C.PVR_ERROR { return C.PVR_ERROR_NOT_IMPLEMENTED }
