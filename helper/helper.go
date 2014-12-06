package helper

//// #include <stdio.h>
//// #include <stdlib.h>
import (
    
	"encoding/base64"
	"crypto/md5"
	"fmt"
	"os"
	"time"
	//"unsafe"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var Base64_ = base64.NewEncoding(base64Table)
/*
func Alloc(size uintptr) *byte {
	return (*byte)(C.malloc(C.size_t(size)))
}

func Free(ptr *byte) {
	//C.free(unsafe.Pointer(ptr))
}
*/
func GMT(t time.Time) string {
	gmt := t.Format("Mon, 02 Jan 2006 15:04:05 GMT")
	return gmt
}

func FileExists(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func DirExists(dir string) bool {
	d, e := os.Stat(dir)
	switch {
	case e != nil:
		return false
	case !d.IsDir():
		return false
	}

	return true
}

func Md5(in string) string {
	hash := md5.New()
	hash.Write([]byte(in))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Base64Encode(in string) string {
	return Base64_.EncodeToString([]byte(in))
}

func Base64Decode(in string) string {
	str,_ := Base64_.DecodeString(in)
	return 	string(str)
}
