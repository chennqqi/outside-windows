package winuser

import (
	. "fmt"
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	"testing"
	"unsafe"
)

// import "unsafe"

func init() {
	outside.AddApis(WinUserApis)
	// outside.AddApis(WinUserANSIApis)
	outside.AddApis(WinUserUnicodeApis)
}

//TODO(t): Handle output string buffers

func callback(h HWND, l LPARAM) BOOL {
	var t [256]WChar
	GetClassName(h, &t[0], 255)
	// c := CStrToString((uintptr)(unsafe.Pointer(&t[0])))
	// c := UniStrToString((uintptr)(unsafe.Pointer(&t[0])))
	// s := GetWindowTextLength(h)
	// for i := LPARAM(0); i < l; i++ {
	// 	Print("\t")
	// }

	// Print(c)
	// if s != 0 {
	// 	GetWindowText(h, &t[0], 255)
	// 	// Print(" ", CStrToString((uintptr)(unsafe.Pointer(&t[0]))))
	// 	// Print(" ", UniStrToString((uintptr)(unsafe.Pointer(&t[0]))))
	// }
	// Println()
	EnumChildWindows(h, callback, l+1)
	W++
	return 1
}

var W uint32

var tT *testing.T

func wscallback(c *VString, l LPARAM) BOOL {
	if l != 123 {
		return 0
	}
	tT.Log(*c)
	return 1
}

func TestEnumWindows(t *testing.T) {
	tT = t
	if EnumWindows(callback, 0) == 0 {
		t.Fail()
	}
	Println(W, "Windows")
}
func TestEnumDesktops(t *testing.T) {
	tT = t
	h := GetProcessWindowStation()
	if EnumDesktops(h, wscallback, 123) == 0 {
		t.Fail()
	}
}
func TestEnumWindowStations(t *testing.T) {
	tT = t
	if EnumWindowStations(wscallback, 123) == 0 {
		t.Fail()
	}
}

func TestWsprintf(t *testing.T) {
	var o [1000]WChar
	Wsprintf(&o[0], "%d %d %d", 123, 456, 789)
	t.Log(outside.UniStrToString((uintptr)(unsafe.Pointer(&o[0]))))
}
