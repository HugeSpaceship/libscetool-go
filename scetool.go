package libscetool_go

// This shit is not thread safe, actually safe, or safe for human habitation, please deal with it as I'm not rewriting SCETool

import "C"
import (
	"unsafe"

	_ "github.com/HugeSpaceship/libscetool-go/scetool"
)

// #cgo LDFLAGS: -lz
// #include <stdlib.h>
// #include <stdio.h>
// #include "scetool/main.h"
// #include "scetool/frontend.h"
// #include "scetool/keys.h"
import "C"

func scetoolInit() {
	ret, err := C.libscetool_init()
	if err != nil {
		panic(err)
	}
	if ret != 0 {
		panic("Failed to init SCETool")
	}
}

func PrintInfo(path string) {
	cpath := C.CString(path + "\000") // This is what we like to call in the business safety :sunglasses:
	defer C.free(unsafe.Pointer(cpath))
	C.frontend_print_infos(cpath)

}

func Decrypt(inPatch, outPath string) {
	cInPath := C.CString(inPatch + "\000") // This is what we like to call in the business safety :sunglasses:
	defer C.free(unsafe.Pointer(cInPath))
	cOutPath := C.CString(outPath + "\000") // This is what we like to call in the business safety :sunglasses:
	defer C.free(unsafe.Pointer(cOutPath))

	C.frontend_decrypt(cInPath, cOutPath)
}

func Encrypt(inPatch, outPath string) {
	cInPath := C.CString(inPatch + "\000") // This is what we like to call in the business safety :sunglasses:
	defer C.free(unsafe.Pointer(cInPath))
	cOutPath := C.CString(outPath + "\000") // This is what we like to call in the business safety :sunglasses:
	defer C.free(unsafe.Pointer(cOutPath))

	C.frontend_encrypt(cInPath, cOutPath)
}

func SetRapDirectory(path string) {
	cPath := C.CString(path + "\000") // This is what we like to call in the business safety :sunglasses:
	defer C.free(unsafe.Pointer(cPath))

	C.rap_set_directory(cPath)
}

// SetIDPSKey sets the IDPS inside SCETool for use with signing digital games.
func SetIDPSKey(idps []byte) {
	cIdps := C.CBytes(idps)
	defer C.free(unsafe.Pointer(cIdps))

	C.set_idps_key((*C.uchar)(unsafe.Pointer(cIdps)))
}

func SetActDatFilePath(path string) {
	cPath := C.CString(path + "\000") // This is what we like to call in the business safety :sunglasses:
	defer C.free(unsafe.Pointer(cPath))

	C.set_act_dat_file_path(cPath)
}

func SetRifFilePath(path string) {
	cPath := C.CString(path + "\000") // This is what we like to call in the business safety :sunglasses:
	defer C.free(unsafe.Pointer(cPath))

	C.set_rif_file_path(cPath)
}

func SetDiscSettings() {
	C.set_disc_encrypt_options()
}

func SetNPDRMSettings() {
	C.set_npdrm_encrypt_options()
}

func SetNPDRMContentID(contentID string) {
	cContentID := C.CString(contentID + "\000") // This is what we like to call in the business safety :sunglasses:
	defer C.free(unsafe.Pointer(cContentID))

	C.set_npdrm_content_id(cContentID)
}

func GetContentID(path string) string {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	cContentID := C.get_content_id(cPath)
	defer C.free(unsafe.Pointer(cContentID))
	contentID := C.GoString((*C.char)(unsafe.Pointer(cContentID)))

	return contentID
}
