/*
 * Copyright (c) 2011-2013 by naehrwert
 * This file is released under the GPLv2.
 */

#ifndef _FRONTEND_H_
#define _FRONTEND_H_

#if __cplusplus
#define export extern "C"
#else
#define export
#endif

// if windows, we have to use __desclspec(dllexport)
#ifdef _WIN32
#define export extern "C" __declspec(dllexport)
#endif

#include "types.h"

export void frontend_print_infos(s8 const *file);
export void set_disc_encrypt_options();
export void set_npdrm_encrypt_options();
export void frontend_decrypt(s8 *file_in, s8 *file_out);
export void frontend_encrypt(s8 *file_in, s8 *file_out);

export void set_npdrm_content_id(s8 *content_id);
export u8 *get_content_id(s8 const *file);

#endif
