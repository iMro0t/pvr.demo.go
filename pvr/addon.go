package pvr

/*
#cgo CFLAGS: -Wno-error=implicit-function-declaration
#include "../kodi/xbmc_pvr_dll.h"
*/
import "C"
import (
	"unsafe"
)

var (
	XBMC        GoHelper_libXBMC_addon
	PVR         GoHelper_libXBMC_pvr
	addonStatus C.ADDON_STATUS = C.ADDON_STATUS_UNKNOWN
	clientPath  string
	Call        func()
)

//export ADDON_GetTypeVersion
func ADDON_GetTypeVersion(addonType C.int) *C.char {
	cs := C.CString("0.0.0")
	switch addonType {
	case C.ADDON_GLOBAL_MAIN:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_MAIN)
	case C.ADDON_GLOBAL_GENERAL:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_GENERAL)
	case C.ADDON_GLOBAL_GUI:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_GUI)
	case C.ADDON_GLOBAL_AUDIOENGINE:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_AUDIOENGINE)
	case C.ADDON_GLOBAL_FILESYSTEM:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_FILESYSTEM)
	case C.ADDON_GLOBAL_NETWORK:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_NETWORK)

	// addon type instances
	case C.ADDON_INSTANCE_AUDIODECODER:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_AUDIODECODER)
	case C.ADDON_INSTANCE_GAME:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_GAME)
	case C.ADDON_INSTANCE_IMAGEDECODER:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_IMAGEDECODER)
	case C.ADDON_INSTANCE_INPUTSTREAM:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_INPUTSTREAM)
	case C.ADDON_INSTANCE_PERIPHERAL:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_PERIPHERAL)
	case C.ADDON_INSTANCE_PVR:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_PVR)
	case C.ADDON_INSTANCE_VFS:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_VFS)
	case C.ADDON_INSTANCE_VIDEOCODEC:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_VIDEOCODEC)
	}
	return cs
}

//export ADDON_GetTypeMinVersion
func ADDON_GetTypeMinVersion(addonType C.int) *C.char {
	cs := C.CString("0.0.0")
	switch addonType {
	case C.ADDON_GLOBAL_MAIN:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_MAIN_MIN)
	case C.ADDON_GLOBAL_GENERAL:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_GENERAL_MIN)
	case C.ADDON_GLOBAL_GUI:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_GUI_MIN)
	case C.ADDON_GLOBAL_AUDIOENGINE:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_AUDIOENGINE_MIN)
	case C.ADDON_GLOBAL_FILESYSTEM:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_FILESYSTEM_MIN)
	case C.ADDON_GLOBAL_NETWORK:
		cs = C.CString(C.ADDON_GLOBAL_VERSION_NETWORK_MIN)

	// addon type instances
	case C.ADDON_INSTANCE_AUDIODECODER:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_AUDIODECODER_MIN)
	case C.ADDON_INSTANCE_GAME:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_GAME_MIN)
	case C.ADDON_INSTANCE_IMAGEDECODER:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_IMAGEDECODER_MIN)
	case C.ADDON_INSTANCE_INPUTSTREAM:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_INPUTSTREAM_MIN)
	case C.ADDON_INSTANCE_PERIPHERAL:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_PERIPHERAL_MIN)
	case C.ADDON_INSTANCE_PVR:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_PVR_MIN)
	case C.ADDON_INSTANCE_VFS:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_VFS_MIN)
	case C.ADDON_INSTANCE_VIDEOCODEC:
		cs = C.CString(C.ADDON_INSTANCE_VERSION_VIDEOCODEC_MIN)
	}
	return cs
}

//export ADDON_ReadSettings
func ADDON_ReadSettings() {}

//export ADDON_Create
func ADDON_Create(handle unsafe.Pointer, properties unsafe.Pointer) C.ADDON_STATUS {
	if handle == nil || properties == nil {
		return C.ADDON_STATUS_UNKNOWN
	}
	XBMC = NewlibXBMC_addon()
	if !XBMC.RegisterMe(handle) {
		XBMC.Free()
		return C.ADDON_STATUS_PERMANENT_FAILURE
	}
	PVR = NewlibXBMC_pvr()
	if !PVR.RegisterMe(handle) {
		PVR.Free()
		XBMC.Free()
		return C.ADDON_STATUS_PERMANENT_FAILURE
	}
	XBMC.Log(XBMC.DEBUG, "Creating the PVR demo add-on")
	Call()
	addonStatus = C.ADDON_STATUS_OK
	return C.ADDON_STATUS_OK
}

//export ADDON_GetStatus
func ADDON_GetStatus() C.ADDON_STATUS {
	return addonStatus
}

//export ADDON_Destroy
func ADDON_Destroy() {
	addonStatus = C.ADDON_STATUS_UNKNOWN
}

//export ADDON_SetSetting
func ADDON_SetSetting(settingName *C.char, settingValue unsafe.Pointer) C.ADDON_STATUS {
	return C.ADDON_STATUS_OK
}

//export get_addon
func get_addon(ptr unsafe.Pointer) {
	var pClient *C.AddonInstance_PVR = (*C.AddonInstance_PVR)(ptr)
	pClient.toAddon.addonInstance = (C.KODI_HANDLE)(C.nullPTR)

	pClient.toAddon.GetAddonCapabilities = (*[0]byte)(C.GetAddonCapabilities)
	pClient.toAddon.GetStreamProperties = (*[0]byte)(C.GetStreamProperties)
	pClient.toAddon.GetConnectionString = (*[0]byte)(C.GetConnectionString)
	pClient.toAddon.GetBackendName = (*[0]byte)(C.GetBackendName)
	pClient.toAddon.GetBackendVersion = (*[0]byte)(C.GetBackendVersion)
	pClient.toAddon.GetDriveSpace = (*[0]byte)(C.GetDriveSpace)
	pClient.toAddon.OpenDialogChannelScan = (*[0]byte)(C.OpenDialogChannelScan)
	pClient.toAddon.MenuHook = (*[0]byte)(C.CallMenuHook)

	pClient.toAddon.GetEPGForChannel = (*[0]byte)(C.GetEPGForChannel)
	pClient.toAddon.IsEPGTagRecordable = (*[0]byte)(C.IsEPGTagRecordable)
	pClient.toAddon.IsEPGTagPlayable = (*[0]byte)(C.IsEPGTagPlayable)
	pClient.toAddon.GetEPGTagEdl = (*[0]byte)(C.GetEPGTagEdl)
	pClient.toAddon.GetEPGTagStreamProperties = (*[0]byte)(C.GetEPGTagStreamProperties)

	pClient.toAddon.GetChannelGroupsAmount = (*[0]byte)(C.GetChannelGroupsAmount)
	pClient.toAddon.GetChannelGroups = (*[0]byte)(C.GetChannelGroups)
	pClient.toAddon.GetChannelGroupMembers = (*[0]byte)(C.GetChannelGroupMembers)

	pClient.toAddon.GetChannelsAmount = (*[0]byte)(C.GetChannelsAmount)
	pClient.toAddon.GetChannels = (*[0]byte)(C.GetChannels)
	pClient.toAddon.DeleteChannel = (*[0]byte)(C.DeleteChannel)
	pClient.toAddon.RenameChannel = (*[0]byte)(C.RenameChannel)
	pClient.toAddon.OpenDialogChannelSettings = (*[0]byte)(C.OpenDialogChannelSettings)
	pClient.toAddon.OpenDialogChannelAdd = (*[0]byte)(C.OpenDialogChannelAdd)

	pClient.toAddon.GetRecordingsAmount = (*[0]byte)(C.GetRecordingsAmount)
	pClient.toAddon.GetRecordings = (*[0]byte)(C.GetRecordings)
	pClient.toAddon.DeleteRecording = (*[0]byte)(C.DeleteRecording)
	pClient.toAddon.UndeleteRecording = (*[0]byte)(C.UndeleteRecording)
	pClient.toAddon.DeleteAllRecordingsFromTrash = (*[0]byte)(C.DeleteAllRecordingsFromTrash)
	pClient.toAddon.RenameRecording = (*[0]byte)(C.RenameRecording)
	pClient.toAddon.SetRecordingLifetime = (*[0]byte)(C.SetRecordingLifetime)
	pClient.toAddon.SetRecordingPlayCount = (*[0]byte)(C.SetRecordingPlayCount)
	pClient.toAddon.SetRecordingLastPlayedPosition = (*[0]byte)(C.SetRecordingLastPlayedPosition)
	pClient.toAddon.GetRecordingLastPlayedPosition = (*[0]byte)(C.GetRecordingLastPlayedPosition)
	pClient.toAddon.GetRecordingEdl = (*[0]byte)(C.GetRecordingEdl)

	pClient.toAddon.GetTimerTypes = (*[0]byte)(C.GetTimerTypes)
	pClient.toAddon.GetTimersAmount = (*[0]byte)(C.GetTimersAmount)
	pClient.toAddon.GetTimers = (*[0]byte)(C.GetTimers)
	pClient.toAddon.AddTimer = (*[0]byte)(C.AddTimer)
	pClient.toAddon.DeleteTimer = (*[0]byte)(C.DeleteTimer)
	pClient.toAddon.UpdateTimer = (*[0]byte)(C.UpdateTimer)

	pClient.toAddon.OpenLiveStream = (*[0]byte)(C.OpenLiveStream)
	pClient.toAddon.CloseLiveStream = (*[0]byte)(C.CloseLiveStream)
	pClient.toAddon.ReadLiveStream = (*[0]byte)(C.ReadLiveStream)
	pClient.toAddon.SeekLiveStream = (*[0]byte)(C.SeekLiveStream)
	pClient.toAddon.LengthLiveStream = (*[0]byte)(C.LengthLiveStream)
	pClient.toAddon.SignalStatus = (*[0]byte)(C.SignalStatus)
	pClient.toAddon.GetDescrambleInfo = (*[0]byte)(C.GetDescrambleInfo)
	pClient.toAddon.GetChannelStreamProperties = (*[0]byte)(C.GetChannelStreamProperties)
	pClient.toAddon.GetRecordingStreamProperties = (*[0]byte)(C.GetRecordingStreamProperties)
	pClient.toAddon.CanPauseStream = (*[0]byte)(C.CanPauseStream)
	pClient.toAddon.PauseStream = (*[0]byte)(C.PauseStream)
	pClient.toAddon.CanSeekStream = (*[0]byte)(C.CanSeekStream)
	pClient.toAddon.SeekTime = (*[0]byte)(C.SeekTime)
	pClient.toAddon.SetSpeed = (*[0]byte)(C.SetSpeed)

	pClient.toAddon.OpenRecordedStream = (*[0]byte)(C.OpenRecordedStream)
	pClient.toAddon.CloseRecordedStream = (*[0]byte)(C.CloseRecordedStream)
	pClient.toAddon.ReadRecordedStream = (*[0]byte)(C.ReadRecordedStream)
	pClient.toAddon.SeekRecordedStream = (*[0]byte)(C.SeekRecordedStream)
	pClient.toAddon.LengthRecordedStream = (*[0]byte)(C.LengthRecordedStream)

	pClient.toAddon.DemuxReset = (*[0]byte)(C.DemuxReset)
	pClient.toAddon.DemuxAbort = (*[0]byte)(C.DemuxAbort)
	pClient.toAddon.DemuxFlush = (*[0]byte)(C.DemuxFlush)
	pClient.toAddon.DemuxRead = (*[0]byte)(C.DemuxRead)

	pClient.toAddon.GetBackendHostname = (*[0]byte)(C.GetBackendHostname)

	pClient.toAddon.IsTimeshifting = (*[0]byte)(C.IsTimeshifting)
	pClient.toAddon.IsRealTimeStream = (*[0]byte)(C.IsRealTimeStream)

	pClient.toAddon.SetEPGTimeFrame = (*[0]byte)(C.SetEPGTimeFrame)

	pClient.toAddon.OnSystemSleep = (*[0]byte)(C.OnSystemSleep)
	pClient.toAddon.OnSystemWake = (*[0]byte)(C.OnSystemWake)
	pClient.toAddon.OnPowerSavingActivated = (*[0]byte)(C.OnPowerSavingActivated)
	pClient.toAddon.OnPowerSavingDeactivated = (*[0]byte)(C.OnPowerSavingDeactivated)
	pClient.toAddon.GetStreamTimes = (*[0]byte)(C.GetStreamTimes)

	pClient.toAddon.GetStreamReadChunkSize = (*[0]byte)(C.GetStreamReadChunkSize)
}
