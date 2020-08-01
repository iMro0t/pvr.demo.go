#include "../kodi/libXBMC_addon.h"
#include "../kodi/xbmc_pvr_dll.h"
#include "../kodi/libXBMC_pvr.h"

#ifdef __cplusplus
extern "C"
{
#endif
    typedef void *CGoHelper_libXBMC_addon;
    CGoHelper_libXBMC_addon CGoHelper_libXBMC_addonInit(void);
    bool CGoHelper_libXBMC_addonRegisterMe(CGoHelper_libXBMC_addon, void *handle);
    void CGoHelper_libXBMC_addonLog(CGoHelper_libXBMC_addon, int loglevel, char *format);
    void CGoHelper_libXBMC_addonFree(CGoHelper_libXBMC_addon);

    typedef void *CGoHelper_libXBMC_pvr;
    CGoHelper_libXBMC_pvr CGoHelper_libXBMC_pvrInit(void);
    bool CGoHelper_libXBMC_pvrRegisterMe(CGoHelper_libXBMC_pvr, void *handle);
    void CGoHelper_libXBMC_pvrTransferChannelEntry(CGoHelper_libXBMC_pvr, void *handle, cPVR_CHANNEL_t *entry);
    void CGoHelper_libXBMC_pvrSetProperty(PVR_NAMED_VALUE *properties, int index, char *key, char *value);
    void CGoHelper_libXBMC_pvrFree(CGoHelper_libXBMC_pvr);

    // void get_addon(void *ptr);

#ifdef __cplusplus
}
#endif