#include "workaround.h"
#include "../kodi/xbmc_pvr_dll.h"

CGoHelper_libXBMC_addon CGoHelper_libXBMC_addonInit()
{
    ADDON::CHelper_libXBMC_addon *XBMC = new ADDON::CHelper_libXBMC_addon;
    return (void *)XBMC;
}

bool CGoHelper_libXBMC_addonRegisterMe(CGoHelper_libXBMC_addon XBMC, void *handle)
{
    ADDON::CHelper_libXBMC_addon *xbmc = (ADDON::CHelper_libXBMC_addon *)XBMC;
    return xbmc->RegisterMe(handle);
}

void CGoHelper_libXBMC_addonLog(CGoHelper_libXBMC_addon XBMC, int loglevel, char *format)
{
    ADDON::CHelper_libXBMC_addon *xbmc = (ADDON::CHelper_libXBMC_addon *)XBMC;
    ADDON::addon_log_t lvl = (ADDON::addon_log_t)loglevel;
    xbmc->Log(lvl, format);
}

void CGoHelper_libXBMC_addonFree(CGoHelper_libXBMC_addon XBMC)
{
    ADDON::CHelper_libXBMC_addon *xbmc = (ADDON::CHelper_libXBMC_addon *)XBMC;
    delete xbmc;
}

CGoHelper_libXBMC_pvr CGoHelper_libXBMC_pvrInit()
{
    CHelper_libXBMC_pvr *PVR = new CHelper_libXBMC_pvr;
    return (void *)PVR;
}

bool CGoHelper_libXBMC_pvrRegisterMe(CGoHelper_libXBMC_pvr PVR, void *handle)
{
    CHelper_libXBMC_pvr *pvr = (CHelper_libXBMC_pvr *)PVR;
    return pvr->RegisterMe(handle);
}

void CGoHelper_libXBMC_pvrTransferChannelEntry(CGoHelper_libXBMC_pvr PVR, void *handle, cPVR_CHANNEL_t *entry)
{
    CHelper_libXBMC_pvr *pvr = (CHelper_libXBMC_pvr *)PVR;
    cADDON_HANDLE_t hdl = (cADDON_HANDLE_t)handle;
    pvr->TransferChannelEntry(hdl, entry);
}

void CGoHelper_libXBMC_pvrSetProperty(PVR_NAMED_VALUE *properties, int index, char *key, char *value)
{
    strncpy(properties[index].strName, key, sizeof(properties[index].strName) - 1);
    strncpy(properties[index].strValue, value, sizeof(properties[index].strValue) - 1);
}

void CGoHelper_libXBMC_pvrFree(CGoHelper_libXBMC_pvr PVR)
{
    CHelper_libXBMC_pvr *pvr = (CHelper_libXBMC_pvr *)PVR;
    delete pvr;
}

// void get_addon(void *ptr)
// {
//     AddonInstance_PVR *pClient = (AddonInstance_PVR *)(ptr);
//     pClient->toAddon.addonInstance = nullptr; // used in future

//     pClient->toAddon.GetAddonCapabilities = GetAddonCapabilities;
//     pClient->toAddon.GetStreamProperties = GetStreamProperties;
//     pClient->toAddon.GetConnectionString = GetConnectionString;
//     pClient->toAddon.GetBackendName = GetBackendName;
//     pClient->toAddon.GetBackendVersion = GetBackendVersion;
//     pClient->toAddon.GetDriveSpace = GetDriveSpace;
//     pClient->toAddon.OpenDialogChannelScan = OpenDialogChannelScan;
//     pClient->toAddon.MenuHook = CallMenuHook;

//     pClient->toAddon.GetEPGForChannel = GetEPGForChannel;
//     pClient->toAddon.IsEPGTagRecordable = IsEPGTagRecordable;
//     pClient->toAddon.IsEPGTagPlayable = IsEPGTagPlayable;
//     pClient->toAddon.GetEPGTagEdl = GetEPGTagEdl;
//     pClient->toAddon.GetEPGTagStreamProperties = GetEPGTagStreamProperties;

//     pClient->toAddon.GetChannelGroupsAmount = GetChannelGroupsAmount;
//     pClient->toAddon.GetChannelGroups = GetChannelGroups;
//     pClient->toAddon.GetChannelGroupMembers = GetChannelGroupMembers;

//     pClient->toAddon.GetChannelsAmount = GetChannelsAmount;
//     pClient->toAddon.GetChannels = GetChannels;
//     pClient->toAddon.DeleteChannel = DeleteChannel;
//     pClient->toAddon.RenameChannel = RenameChannel;
//     pClient->toAddon.OpenDialogChannelSettings = OpenDialogChannelSettings;
//     pClient->toAddon.OpenDialogChannelAdd = OpenDialogChannelAdd;

//     pClient->toAddon.GetRecordingsAmount = GetRecordingsAmount;
//     pClient->toAddon.GetRecordings = GetRecordings;
//     pClient->toAddon.DeleteRecording = DeleteRecording;
//     pClient->toAddon.UndeleteRecording = UndeleteRecording;
//     pClient->toAddon.DeleteAllRecordingsFromTrash = DeleteAllRecordingsFromTrash;
//     pClient->toAddon.RenameRecording = RenameRecording;
//     pClient->toAddon.SetRecordingLifetime = SetRecordingLifetime;
//     pClient->toAddon.SetRecordingPlayCount = SetRecordingPlayCount;
//     pClient->toAddon.SetRecordingLastPlayedPosition = SetRecordingLastPlayedPosition;
//     pClient->toAddon.GetRecordingLastPlayedPosition = GetRecordingLastPlayedPosition;
//     pClient->toAddon.GetRecordingEdl = GetRecordingEdl;

//     pClient->toAddon.GetTimerTypes = GetTimerTypes;
//     pClient->toAddon.GetTimersAmount = GetTimersAmount;
//     pClient->toAddon.GetTimers = GetTimers;
//     pClient->toAddon.AddTimer = AddTimer;
//     pClient->toAddon.DeleteTimer = DeleteTimer;
//     pClient->toAddon.UpdateTimer = UpdateTimer;

//     pClient->toAddon.OpenLiveStream = OpenLiveStream;
//     pClient->toAddon.CloseLiveStream = CloseLiveStream;
//     pClient->toAddon.ReadLiveStream = ReadLiveStream;
//     pClient->toAddon.SeekLiveStream = SeekLiveStream;
//     pClient->toAddon.LengthLiveStream = LengthLiveStream;
//     pClient->toAddon.SignalStatus = SignalStatus;
//     pClient->toAddon.GetDescrambleInfo = GetDescrambleInfo;
//     pClient->toAddon.GetChannelStreamProperties = GetChannelStreamProperties;
//     pClient->toAddon.GetRecordingStreamProperties = GetRecordingStreamProperties;
//     pClient->toAddon.CanPauseStream = CanPauseStream;
//     pClient->toAddon.PauseStream = PauseStream;
//     pClient->toAddon.CanSeekStream = CanSeekStream;
//     pClient->toAddon.SeekTime = SeekTime;
//     pClient->toAddon.SetSpeed = SetSpeed;

//     pClient->toAddon.OpenRecordedStream = OpenRecordedStream;
//     pClient->toAddon.CloseRecordedStream = CloseRecordedStream;
//     pClient->toAddon.ReadRecordedStream = ReadRecordedStream;
//     pClient->toAddon.SeekRecordedStream = SeekRecordedStream;
//     pClient->toAddon.LengthRecordedStream = LengthRecordedStream;

//     pClient->toAddon.DemuxReset = DemuxReset;
//     pClient->toAddon.DemuxAbort = DemuxAbort;
//     pClient->toAddon.DemuxFlush = DemuxFlush;
//     pClient->toAddon.DemuxRead = DemuxRead;

//     pClient->toAddon.GetBackendHostname = GetBackendHostname;

//     pClient->toAddon.IsTimeshifting = IsTimeshifting;
//     pClient->toAddon.IsRealTimeStream = IsRealTimeStream;

//     pClient->toAddon.SetEPGTimeFrame = SetEPGTimeFrame;

//     pClient->toAddon.OnSystemSleep = OnSystemSleep;
//     pClient->toAddon.OnSystemWake = OnSystemWake;
//     pClient->toAddon.OnPowerSavingActivated = OnPowerSavingActivated;
//     pClient->toAddon.OnPowerSavingDeactivated = OnPowerSavingDeactivated;
//     pClient->toAddon.GetStreamTimes = GetStreamTimes;

//     pClient->toAddon.GetStreamReadChunkSize = GetStreamReadChunkSize;
// }