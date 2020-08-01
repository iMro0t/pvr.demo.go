package pvr

/*
#include "workaround.h"
#include "../kodi/libXBMC_addon.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type GoHelper_libXBMC_addon struct {
	DEBUG  C.int
	INFO   C.int
	NOTICE C.int
	ERROR  C.int
	addon  C.CGoHelper_libXBMC_addon
}

func NewlibXBMC_addon() GoHelper_libXBMC_addon {
	return GoHelper_libXBMC_addon{
		DEBUG:  C.LOG_DEBUG,
		INFO:   C.LOG_INFO,
		NOTICE: C.LOG_NOTICE,
		ERROR:  C.LOG_ERROR,
		addon:  C.CGoHelper_libXBMC_addonInit(),
	}
}

func (xbmc GoHelper_libXBMC_addon) RegisterMe(hdl unsafe.Pointer) C.bool {
	return C.CGoHelper_libXBMC_addonRegisterMe(xbmc.addon, hdl)
}

func (xbmc GoHelper_libXBMC_addon) Log(loglevel C.int, format ...interface{}) {
	C.CGoHelper_libXBMC_addonLog(xbmc.addon, loglevel, C.CString(fmt.Sprint(format...)))
}

func (xbmc GoHelper_libXBMC_addon) Free() {
	C.CGoHelper_libXBMC_addonFree(xbmc.addon)
}
