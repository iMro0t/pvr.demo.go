/*
 *  Copyright (C) 2005-2018 Team Kodi
 *  This file is part of Kodi - https://kodi.tv
 *
 *  SPDX-License-Identifier: GPL-2.0-or-later
 *  See LICENSES/README.md for more information.
 */

#pragma once

#include <stdarg.h> /* va_list, va_start, va_arg, va_end */
#ifdef __cplusplus
#include <cstdlib>
#include <cstring>
#include <stdexcept>
#include <string>
#include <vector>
#else
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#endif

#ifndef TARGET_WINDOWS
#ifndef __cdecl
#define __cdecl
#endif
#ifndef __declspec
#define __declspec(X)
#endif
#endif

#undef ATTRIBUTE_PACKED
#undef PRAGMA_PACK_BEGIN
#undef PRAGMA_PACK_END

#if defined(__GNUC__)
#define ATTRIBUTE_PACKED __attribute__((packed))
#define PRAGMA_PACK 0
#define ATTRIBUTE_HIDDEN __attribute__((visibility("hidden")))
#endif

#if !defined(ATTRIBUTE_PACKED)
#define ATTRIBUTE_PACKED
#define PRAGMA_PACK 1
#endif

#if !defined(ATTRIBUTE_HIDDEN)
#define ATTRIBUTE_HIDDEN
#endif

#ifdef _MSC_VER
#define ATTRIBUTE_FORCEINLINE __forceinline
#elif defined(__GNUC__)
#define ATTRIBUTE_FORCEINLINE inline __attribute__((__always_inline__))
#elif defined(__CLANG__)
#if __has_attribute(__always_inline__)
#define ATTRIBUTE_FORCEINLINE inline __attribute__((__always_inline__))
#else
#define ATTRIBUTE_FORCEINLINE inline
#endif
#else
#define ATTRIBUTE_FORCEINLINE inline
#endif

#include "versions.h"

#ifdef __cplusplus
namespace kodi
{
  namespace addon
  {
    class CAddonBase;
  }
} // namespace kodi
namespace kodi
{
  namespace addon
  {
    class IAddonInstance;
  }
} // namespace kodi
#endif

#ifdef __cplusplus
extern "C"
{
#endif

  //==============================================================================
  /// Standard undefined pointer handle
  typedef void *KODI_HANDLE;
  //------------------------------------------------------------------------------

  //==============================================================================
  ///
  typedef enum ADDON_STATUS
  {
    ///
    ADDON_STATUS_OK,

    ///
    ADDON_STATUS_LOST_CONNECTION,

    ///
    ADDON_STATUS_NEED_RESTART,

    ///
    ADDON_STATUS_NEED_SETTINGS,

    ///
    ADDON_STATUS_UNKNOWN,

    /// permanent failure, like failing to resolve methods
    ADDON_STATUS_PERMANENT_FAILURE,

    /* internal used return error if function becomes not used from child on
   * addon */
    ADDON_STATUS_NOT_IMPLEMENTED
  } ADDON_STATUS;
  //------------------------------------------------------------------------------

  //==============================================================================
  /// @todo remove start with ADDON_* after old way on libXBMC_addon.h is removed
  ///
  typedef enum AddonLog
  {
    ///
    ADDON_LOG_DEBUG = 0,

    ///
    ADDON_LOG_INFO = 1,

    ///
    ADDON_LOG_NOTICE = 2,

    ///
    ADDON_LOG_WARNING = 3,

    ///
    ADDON_LOG_ERROR = 4,

    ///
    ADDON_LOG_SEVERE = 5,

    ///
    ADDON_LOG_FATAL = 6
  } AddonLog;
  //------------------------------------------------------------------------------

  /*!
 * @brief Handle used to return data from the PVR add-on to CPVRClient
 */
  struct ADDON_HANDLE_STRUCT
  {
    void *callerAddress; /*!< address of the caller */
    void *dataAddress;   /*!< address to store data in */
    int dataIdentifier;  /*!< parameter to pass back when calling the callback */
  };
  typedef struct ADDON_HANDLE_STRUCT *ADDON_HANDLE;

/*
 * To have a on add-on and kodi itself handled string always on known size!
 */
#define ADDON_STANDARD_STRING_LENGTH 1024
#define ADDON_STANDARD_STRING_LENGTH_SMALL 256

  /*
 * Callback function tables from addon to Kodi
 * Set complete from Kodi!
 */
  struct AddonToKodiFuncTable_kodi;
  struct AddonToKodiFuncTable_kodi_audioengine;
  struct AddonToKodiFuncTable_kodi_filesystem;
  struct AddonToKodiFuncTable_kodi_network;
  struct AddonToKodiFuncTable_kodi_gui;
  typedef struct AddonToKodiFuncTable_Addon
  {
    // Pointer inside Kodi, used on callback functions to give related handle
    // class, for this ADDON::CAddonDll inside Kodi.
    KODI_HANDLE kodiBase;

    // Function addresses used for callbacks from addon to Kodi
    void (*free_string)(void *kodiBase, char *str);
    void (*free_string_array)(void *kodiBase, char **arr, int numElements);
    char *(*get_addon_path)(void *kodiBase);
    char *(*get_base_user_path)(void *kodiBase);
    void (*addon_log_msg)(void *kodiBase, const int loglevel, const char *msg);

    bool (*get_setting_bool)(void *kodiBase, const char *id, bool *value);
    bool (*get_setting_int)(void *kodiBase, const char *id, int *value);
    bool (*get_setting_float)(void *kodiBase, const char *id, float *value);
    bool (*get_setting_string)(void *kodiBase, const char *id, char **value);

    bool (*set_setting_bool)(void *kodiBase, const char *id, bool value);
    bool (*set_setting_int)(void *kodiBase, const char *id, int value);
    bool (*set_setting_float)(void *kodiBase, const char *id, float value);
    bool (*set_setting_string)(void *kodiBase, const char *id, const char *value);

    struct AddonToKodiFuncTable_kodi *kodi;
    struct AddonToKodiFuncTable_kodi_audioengine *kodi_audioengine;
    struct AddonToKodiFuncTable_kodi_filesystem *kodi_filesystem;
    struct AddonToKodiFuncTable_kodi_gui *kodi_gui;
    struct AddonToKodiFuncTable_kodi_network *kodi_network;

    void *(*get_interface)(void *kodiBase, const char *name, const char *version);
  } AddonToKodiFuncTable_Addon;

  /*
 * Function tables from Kodi to addon
 */
  typedef struct KodiToAddonFuncTable_Addon
  {
    void (*destroy)();
    ADDON_STATUS(*get_status)
    ();
    ADDON_STATUS(*create_instance)
    (int instanceType, const char *instanceID, KODI_HANDLE instance, KODI_HANDLE *addonInstance, KODI_HANDLE parent);
    void (*destroy_instance)(int instanceType, KODI_HANDLE instance);
    ADDON_STATUS(*set_setting)
    (const char *settingName, const void *settingValue);
    ADDON_STATUS(*create_instance_ex)
    (int instanceType, const char *instanceID, KODI_HANDLE instance, KODI_HANDLE *addonInstance, KODI_HANDLE parent, const char *version);
  } KodiToAddonFuncTable_Addon;

  // #ifdef __cplusplus
  //   /*
  //  * Main structure passed from kodi to addon with basic information needed to
  //  * create add-on.
  //  */
  //   typedef struct AddonGlobalInterface
  //   {
  //     // String with full path where add-on is installed (without his name on end)
  //     // Set from Kodi!
  //     const char *libBasePath;

  //     // Pointer of first created instance, used in case this add-on goes with single way
  //     // Set from Kodi!
  //     KODI_HANDLE firstKodiInstance;

  //     // Pointer to master base class inside add-on
  //     // Set from addon header!
  //     kodi::addon::CAddonBase *addonBase;

  //     // Pointer to a instance used on single way (together with this class)
  //     // Set from addon header!
  //     kodi::addon::IAddonInstance *globalSingleInstance;

  //     // Callback function tables from addon to Kodi
  //     // Set from Kodi!
  //     AddonToKodiFuncTable_Addon *toKodi;

  //     // Function tables from Kodi to addon
  //     // Set from addon header!
  //     KodiToAddonFuncTable_Addon *toAddon;
  //   } AddonGlobalInterface;
  // #endif

#ifdef __cplusplus
} /* extern "C" */
#endif