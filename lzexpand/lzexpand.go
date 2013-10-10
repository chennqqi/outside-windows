// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package lzexpand provides API definitions for accessing
//the lz functions in kernel32.dll.
package lzexpand

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/kernel32"
)

var (
	LZStart func() T.INT

	LZDone func() T.VOID

	CopyLZFile func(T.INT, T.INT) T.LONG

	LZCopy func(T.INT, T.INT) T.LONG

	LZInit func(T.INT) T.INT

	GetExpandedName func(VString, OVString) T.INT

	LZOpenFile func(OVString, *T.OFSTRUCT, T.WORD) T.INT

	LZSeek func(T.INT, T.LONG, T.INT) T.LONG

	LZRead func(T.INT, T.AString, T.INT) T.INT

	LZClose func(T.INT) T.VOID
)

var LZExpandANSIApis = Apis{
	{"GetExpandedNameA", &GetExpandedName},
	{"LZOpenFileA", &LZOpenFile},
}

var LZExpandUnicodeApis = Apis{
	{"GetExpandedNameW", &GetExpandedName},
	{"LZOpenFileW", &LZOpenFile},
}

var LZExpandApis = Apis{
	{"CopyLZFile", &CopyLZFile},
	{"LZClose", &LZClose},
	{"LZCopy", &LZCopy},
	{"LZDone", &LZDone},
	{"LZInit", &LZInit},
	{"LZRead", &LZRead},
	{"LZSeek", &LZSeek},
	{"LZStart", &LZStart},
}
