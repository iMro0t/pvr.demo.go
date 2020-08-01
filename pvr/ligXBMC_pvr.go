package pvr

import (
	"unsafe"
)

/*
#include "workaround.h"
*/
import "C"

type GoHelper_libXBMC_pvr struct {
	addon C.CGoHelper_libXBMC_pvr
}

func NewlibXBMC_pvr() GoHelper_libXBMC_pvr {
	return GoHelper_libXBMC_pvr{addon: C.CGoHelper_libXBMC_pvrInit()}
}

func (pvr GoHelper_libXBMC_pvr) RegisterMe(hdl unsafe.Pointer) C.bool {
	return C.CGoHelper_libXBMC_pvrRegisterMe(pvr.addon, hdl)
}

func (pvr GoHelper_libXBMC_pvr) TransferChannelEntry(handle unsafe.Pointer, channel Channel) {
	entry := C.struct_PVR_CHANNEL{
		iUniqueId: (C.uint)(channel.ID),
		bIsRadio:  (C.bool)(channel.IsRadio),
		// iChannelNumber:    (C.uint)(channel.Number),
		// iSubChannelNumber: (C.uint)(channel.SubNumber),
		strChannelName: StrToCCharArr(channel.Name),
		strIconPath:    StrToCCharArr(channel.IconPath),
		bIsHidden:      (C.bool)(channel.IsHidden),
	}
	C.CGoHelper_libXBMC_pvrTransferChannelEntry(pvr.addon, handle, &entry)
}

func (pvr GoHelper_libXBMC_pvr) SetProperties(properties *C.struct_PVR_NAMED_VALUE, stream Stream) int {
	C.CGoHelper_libXBMC_pvrSetProperty(properties, C.int(0), C.CString(C.PVR_STREAM_PROPERTY_STREAMURL), C.CString(stream.URL))
	index := 1
	for k, v := range stream.Properties {
		ck := C.CString(k)
		cv := C.CString(v)
		C.CGoHelper_libXBMC_pvrSetProperty(properties, C.int(index), ck, cv)
		C.free(unsafe.Pointer(ck))
		C.free(unsafe.Pointer(cv))
		index++
	}
	return index
}

func (pvr GoHelper_libXBMC_pvr) Free() {
	C.CGoHelper_libXBMC_pvrFree(pvr.addon)
}
