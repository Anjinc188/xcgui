package wke

import (
	"math"
	"syscall"
	"unsafe"
)

import (
	"github.com/codyguo/xcgui/xc"
)

var (
	// Functions
	xWeb_Init               *syscall.Proc
	xWeb_UnInit             *syscall.Proc
	xWeb_Create             *syscall.Proc
	xWeb_SetProxy           *syscall.Proc
	xWeb_LoadUrl            *syscall.Proc
	xWeb_PostUrl            *syscall.Proc
	xWeb_LoadHtmlFromText   *syscall.Proc
	xWeb_RunJs              *syscall.Proc
	xWeb_GlobalExec         *syscall.Proc
	xWeb_Zoom               *syscall.Proc
	xWeb_GetZoom            *syscall.Proc
	xWeb_ZoomReset          *syscall.Proc
	xWeb_SetEditable        *syscall.Proc
	xWeb_GetStringA         *syscall.Proc
	xWeb_GetStringW         *syscall.Proc
	xWeb_SetStringA         *syscall.Proc
	xWeb_SetStringW         *syscall.Proc
	xWeb_GetCookie          *syscall.Proc
	xWeb_SetCookieEnabled   *syscall.Proc
	xWeb_IsCookieEnabled    *syscall.Proc
	xWeb_GoBack             *syscall.Proc
	xWeb_GoForward          *syscall.Proc
	xWeb_GetVersion         *syscall.Proc
	xWeb_GetVersionString   *syscall.Proc
	xWeb_SetMediaVolume     *syscall.Proc
	xWeb_GetMediaVolume     *syscall.Proc
	xWeb_SetUserAgentA      *syscall.Proc
	xWeb_SetUserAgentW      *syscall.Proc
	xWeb_Sleep              *syscall.Proc
	xWeb_Wake               *syscall.Proc
	xWeb_IsAwake            *syscall.Proc
	xWeb_EditorSelectAll    *syscall.Proc
	xWeb_EditorCopy         *syscall.Proc
	xWeb_EditorCut          *syscall.Proc
	xWeb_EditorPaste        *syscall.Proc
	xWeb_EditorDelete       *syscall.Proc
	xWeb_CanGoForward       *syscall.Proc
	xWeb_CanGoBack          *syscall.Proc
	xWeb_OnTitleChanged     *syscall.Proc
	xWeb_OnURLChanged       *syscall.Proc
	xWeb_OnNavigation       *syscall.Proc
	xWeb_OnLoadingFinish    *syscall.Proc
	xWeb_IsLoadingSucceeded *syscall.Proc
	xWeb_IsLoadingFailed    *syscall.Proc
	xWeb_IsLoadingCompleted *syscall.Proc
	xWeb_IsDocumentReady    *syscall.Proc
	xWeb_StopLoading        *syscall.Proc
	xWeb_Reload             *syscall.Proc
	xWeb_JsBindFuction      *syscall.Proc
	xWeb_JsBindGetter       *syscall.Proc
	xWeb_JsBindSetter       *syscall.Proc
	xWeb_JsToInt            *syscall.Proc
	xWeb_jsToFloat          *syscall.Proc
	xWeb_jsToDouble         *syscall.Proc
	xWeb_jsToBoolean        *syscall.Proc
	xWeb_JsToTempStringW    *syscall.Proc
	xWeb_JsToTempStringA    *syscall.Proc
	xWeb_JsArg              *syscall.Proc
	xWeb_JsInt              *syscall.Proc
	xWeb_JsDouble           *syscall.Proc
	xWeb_JsBoolean          *syscall.Proc
	xWeb_JsUndefined        *syscall.Proc
	xWeb_JsNull             *syscall.Proc
	xWeb_JsTrue             *syscall.Proc
	xWeb_JsFalse            *syscall.Proc
	xWeb_JsStringA          *syscall.Proc
	xWeb_jsStringW          *syscall.Proc
	xWeb_JsGlobalObject     *syscall.Proc
	xWeb_JsGet              *syscall.Proc
	xWeb_JsSet              *syscall.Proc
	xWeb_JsGetAt            *syscall.Proc
	xWeb_JsSetAt            *syscall.Proc
)

func init() {
	// Functions
	xWeb_Init = wkeWebDLL.MustFindProc("XWeb_Init")
	xWeb_UnInit = wkeWebDLL.MustFindProc("XWeb_UnInit")
	xWeb_Create = wkeWebDLL.MustFindProc("XWeb_Create")
	xWeb_SetProxy = wkeWebDLL.MustFindProc("XWeb_SetProxy")
	xWeb_LoadUrl = wkeWebDLL.MustFindProc("XWeb_LoadUrl")
	xWeb_PostUrl = wkeWebDLL.MustFindProc("XWeb_PostUrl")
	xWeb_LoadHtmlFromText = wkeWebDLL.MustFindProc("XWeb_LoadHtmlFromText")
	xWeb_RunJs = wkeWebDLL.MustFindProc("XWeb_RunJs")
	xWeb_GlobalExec = wkeWebDLL.MustFindProc("XWeb_GlobalExec")
	xWeb_Zoom = wkeWebDLL.MustFindProc("XWeb_Zoom")
	xWeb_GetZoom = wkeWebDLL.MustFindProc("XWeb_GetZoom")
	xWeb_ZoomReset = wkeWebDLL.MustFindProc("XWeb_ZoomReset")
	xWeb_SetEditable = wkeWebDLL.MustFindProc("XWeb_SetEditable")
	xWeb_GetStringA = wkeWebDLL.MustFindProc("XWeb_GetStringA")
	xWeb_GetStringW = wkeWebDLL.MustFindProc("XWeb_GetStringW")
	xWeb_SetStringA = wkeWebDLL.MustFindProc("XWeb_SetStringA")
	xWeb_SetStringW = wkeWebDLL.MustFindProc("XWeb_SetStringW")
	xWeb_GetCookie = wkeWebDLL.MustFindProc("XWeb_GetCookie")
	xWeb_SetCookieEnabled = wkeWebDLL.MustFindProc("XWeb_SetCookieEnabled")
	xWeb_IsCookieEnabled = wkeWebDLL.MustFindProc("XWeb_IsCookieEnabled")
	xWeb_GoBack = wkeWebDLL.MustFindProc("XWeb_GoBack")
	xWeb_GoForward = wkeWebDLL.MustFindProc("XWeb_GoForward")
	xWeb_GetVersion = wkeWebDLL.MustFindProc("XWeb_GetVersion")
	xWeb_GetVersionString = wkeWebDLL.MustFindProc("XWeb_GetVersionString")
	xWeb_SetMediaVolume = wkeWebDLL.MustFindProc("XWeb_SetMediaVolume")
	xWeb_GetMediaVolume = wkeWebDLL.MustFindProc("XWeb_GetMediaVolume")
	xWeb_SetUserAgentA = wkeWebDLL.MustFindProc("XWeb_SetUserAgentA")
	xWeb_SetUserAgentW = wkeWebDLL.MustFindProc("XWeb_SetUserAgentW")
	xWeb_Sleep = wkeWebDLL.MustFindProc("XWeb_Sleep")
	xWeb_Wake = wkeWebDLL.MustFindProc("XWeb_Wake")
	xWeb_IsAwake = wkeWebDLL.MustFindProc("XWeb_IsAwake")
	xWeb_EditorSelectAll = wkeWebDLL.MustFindProc("XWeb_EditorSelectAll")
	xWeb_EditorCopy = wkeWebDLL.MustFindProc("XWeb_EditorCopy")
	xWeb_EditorCut = wkeWebDLL.MustFindProc("XWeb_EditorCut")
	xWeb_EditorPaste = wkeWebDLL.MustFindProc("XWeb_EditorPaste")
	xWeb_EditorDelete = wkeWebDLL.MustFindProc("XWeb_EditorDelete")
	xWeb_CanGoForward = wkeWebDLL.MustFindProc("XWeb_CanGoForward")
	xWeb_CanGoBack = wkeWebDLL.MustFindProc("XWeb_CanGoBack")
	xWeb_OnTitleChanged = wkeWebDLL.MustFindProc("XWeb_OnTitleChanged")
	xWeb_OnURLChanged = wkeWebDLL.MustFindProc("XWeb_OnURLChanged")
	xWeb_OnNavigation = wkeWebDLL.MustFindProc("XWeb_OnNavigation")
	xWeb_IsLoadingSucceeded = wkeWebDLL.MustFindProc("XWeb_IsLoadingSucceeded")
	xWeb_OnLoadingFinish = wkeWebDLL.MustFindProc("XWeb_OnLoadingFinish")
	xWeb_IsLoadingFailed = wkeWebDLL.MustFindProc("XWeb_IsLoadingFailed")
	xWeb_IsLoadingCompleted = wkeWebDLL.MustFindProc("XWeb_IsLoadingCompleted")
	xWeb_IsDocumentReady = wkeWebDLL.MustFindProc("XWeb_IsDocumentReady")
	xWeb_StopLoading = wkeWebDLL.MustFindProc("XWeb_StopLoading")
	xWeb_Reload = wkeWebDLL.MustFindProc("XWeb_Reload")
	xWeb_JsBindFuction = wkeWebDLL.MustFindProc("XWeb_JsBindFuction")
	xWeb_JsBindGetter = wkeWebDLL.MustFindProc("XWeb_JsBindGetter")
	xWeb_JsBindSetter = wkeWebDLL.MustFindProc("XWeb_JsBindSetter")
	xWeb_JsToInt = wkeWebDLL.MustFindProc("XWeb_JsToInt")
	xWeb_jsToFloat = wkeWebDLL.MustFindProc("XWeb_jsToFloat")
	xWeb_jsToDouble = wkeWebDLL.MustFindProc("XWeb_jsToDouble")
	xWeb_jsToBoolean = wkeWebDLL.MustFindProc("XWeb_jsToBoolean")
	xWeb_JsToTempStringW = wkeWebDLL.MustFindProc("XWeb_JsToTempStringW")
	xWeb_JsToTempStringA = wkeWebDLL.MustFindProc("XWeb_JsToTempStringA")
	xWeb_JsArg = wkeWebDLL.MustFindProc("XWeb_JsArg")
	xWeb_JsInt = wkeWebDLL.MustFindProc("XWeb_JsInt")
	xWeb_JsDouble = wkeWebDLL.MustFindProc("XWeb_JsDouble")
	xWeb_JsBoolean = wkeWebDLL.MustFindProc("XWeb_JsBoolean")
	xWeb_JsUndefined = wkeWebDLL.MustFindProc("XWeb_JsUndefined")
	xWeb_JsNull = wkeWebDLL.MustFindProc("XWeb_JsNull")
	xWeb_JsTrue = wkeWebDLL.MustFindProc("XWeb_JsTrue")
	xWeb_JsFalse = wkeWebDLL.MustFindProc("XWeb_JsFalse")
	xWeb_JsStringA = wkeWebDLL.MustFindProc("XWeb_JsStringA")
	xWeb_jsStringW = wkeWebDLL.MustFindProc("XWeb_jsStringW")
	xWeb_JsGlobalObject = wkeWebDLL.MustFindProc("XWeb_JsGlobalObject")
	xWeb_JsGet = wkeWebDLL.MustFindProc("XWeb_JsGet")
	xWeb_JsSet = wkeWebDLL.MustFindProc("XWeb_JsSet")
	xWeb_JsGetAt = wkeWebDLL.MustFindProc("XWeb_JsGetAt")
	xWeb_JsSetAt = wkeWebDLL.MustFindProc("XWeb_JsSetAt")

	// 默认初始化
	XWeb_Init()
}

func XWeb_Init() {
	xWeb_Init.Call()
}

func XWeb_UnInit() {
	xWeb_UnInit.Call()
}

func XWeb_Create(x, y, cx, cy int, hParent xc.HXCGUI) uintptr {
	ret, _, _ := xWeb_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return uintptr(ret)
}

func XWeb_SetProxy(hWeb uintptr, proxyType WkeProxyType, hostName string, port uint32, username, pwd string) {
	xWeb_SetProxy.Call(
		hWeb,
		uintptr(proxyType),
		xc.StringToUintPtr(hostName),
		uintptr(port),
		xc.StringToUintPtr(username),
		xc.StringToUintPtr(pwd))
}

func XWeb_LoadUrl(hWeb uintptr, url string) {
	xWeb_LoadUrl.Call(
		hWeb,
		xc.StringToUintPtr(url))
}

func XWeb_PostUrl(hWeb uintptr, url, postData string, postLen int) {
	xWeb_PostUrl.Call(
		hWeb,
		xc.StringToUintPtr(url),
		uintptr(postLen))
}

func XWeb_LoadHtmlFromText(hWeb uintptr, htmlData string) {
	xWeb_LoadHtmlFromText.Call(
		hWeb,
		xc.StringToUintPtr(htmlData))
}

func XWeb_RunJs(hWeb uintptr, jsText string) {
	xWeb_RunJs.Call(
		hWeb,
		xc.StringToUintPtr(jsText))
}

func XWeb_GlobalExec(hWeb uintptr) uintptr {
	ret, _, _ := xWeb_GlobalExec.Call(hWeb)

	return ret
}

func XWeb_Zoom(hWeb uintptr, f float32) {
	zoom := math.Float32bits(f)
	xWeb_Zoom.Call(
		hWeb,
		uintptr(zoom))
}

func XWeb_GetZoom(hWeb uintptr) float32 {
	ret, _, _ := xWeb_GetZoom.Call(hWeb)

	return float32(ret)
}

func XWeb_ZoomReset(hWeb uintptr) {
	xWeb_ZoomReset.Call(hWeb)
}

func XWeb_SetEditable(hWeb uintptr, editable bool) {
	xWeb_SetEditable.Call(
		hWeb,
		uintptr(xc.BoolToBOOL(editable)))
}

func XWeb_GetStringA(stringA uintptr) string {
	ret, _, _ := xWeb_GetStringA.Call(stringA)

	return xc.UTF8PtrToSting(ret)
}

func XWeb_GetStringW(stringW uintptr) string {
	ret, _, _ := xWeb_GetStringW.Call(stringW)

	return xc.UINTptrToString(ret)
}

/*
http://baike.baidu.com/link?url=gFiwJdUF0mAgY-pLBvflyvwAdr5mXESqBBj7S_pTg2UDaB_w6Ees6O8fQJ3pUQwktt4ADu80z5eW9G4t5eHV4_
*/
func XWeb_SetStringA(stringA uintptr, str string, lenA int) {
	strA := xc.StringToBytePtr(str)
	xWeb_SetStringA.Call(
		stringA,
		uintptr(unsafe.Pointer(strA)),
		uintptr(lenA))
}

func XWeb_SetStringW(stringW uintptr, str string, lenW int) {
	xWeb_SetStringW.Call(
		stringW,
		xc.StringToUintPtr(str),
		uintptr(lenW))
}

func XWeb_GetCookie(hWeb uintptr) string {
	ret, _, _ := xWeb_GetCookie.Call(hWeb)

	return xc.UINTptrToString(ret)
}

func XWeb_SetCookieEnabled(hWeb uintptr, enable bool) {
	xWeb_SetCookieEnabled.Call(
		hWeb,
		uintptr(xc.BoolToBOOL(enable)))
}

func XWeb_IsCookieEnabled(hWeb uintptr) bool {
	ret, _, _ := xWeb_IsCookieEnabled.Call(hWeb)

	return ret == xc.TRUE
}

func XWeb_GoBack(hWeb uintptr) {
	xWeb_GoBack.Call(hWeb)
}

func XWeb_GoForward(hWeb uintptr) {
	xWeb_GoForward.Call(hWeb)
}

func XWeb_GetVersion() uint32 {
	ret, _, _ := xWeb_GetVersion.Call()

	return uint32(ret)
}

func XWeb_GetVersionString() string {
	ret, _, _ := xWeb_GetVersionString.Call()

	return xc.UTF8PtrToSting(ret)
}

func XWeb_SetMediaVolume(hWeb uintptr, volume float32) {
	xWeb_SetMediaVolume.Call(
		hWeb,
		uintptr(volume))
}

func XWeb_GetMediaVolume(hWeb uintptr) float32 {
	ret, _, _ := xWeb_GetMediaVolume.Call(hWeb)

	return float32(ret)
}

func XWeb_SetUserAgentA(hWeb uintptr, userAgentStr string) {
	userAgent := syscall.StringBytePtr(userAgentStr)
	xWeb_SetUserAgentA.Call(
		hWeb,
		uintptr(unsafe.Pointer(userAgent)))
}

func XWeb_SetUserAgentW(hWeb uintptr, userAgent string) {
	xWeb_SetUserAgentW.Call(
		hWeb,
		xc.StringToUintPtr(userAgent))
}

func XWeb_Sleep(hWeb uintptr) {
	xWeb_Sleep.Call(hWeb)
}

func XWeb_Wake(hWeb uintptr) {
	xWeb_Wake.Call(hWeb)
}

func XWeb_IsAwake(hWeb uintptr) bool {
	ret, _, _ := xWeb_IsAwake.Call(hWeb)

	return ret == xc.TRUE
}

func XWeb_EditorSelectAll(hWeb uintptr) {
	xWeb_EditorSelectAll.Call(hWeb)
}

func XWeb_EditorCopy(hWeb uintptr) {
	xWeb_EditorCopy.Call(hWeb)
}

func XWeb_EditorCut(hWeb uintptr) {
	xWeb_EditorCut.Call(hWeb)
}

func XWeb_EditorPaste(hWeb uintptr) {
	xWeb_EditorPaste.Call(hWeb)
}

func XWeb_EditorDelete(hWeb uintptr) {
	xWeb_EditorDelete.Call(hWeb)
}

func XWeb_CanGoForward(hWeb uintptr) bool {
	ret, _, _ := xWeb_CanGoForward.Call(hWeb)

	return ret == xc.TRUE
}

func XWeb_CanGoBack(hWeb uintptr) bool {
	ret, _, _ := xWeb_CanGoBack.Call(hWeb)

	return ret == xc.TRUE
}

func XWeb_OnTitleChanged(hWeb, callback, param uintptr) {
	xWeb_OnTitleChanged.Call(
		hWeb,
		callback,
		param)
}

func XWeb_OnURLChanged(hWeb, callback, param uintptr) {
	xWeb_OnURLChanged.Call(
		hWeb,
		callback,
		param)
}

func XWeb_OnNavigation(hWeb, callback, param uintptr) {
	xWeb_OnNavigation.Call(
		hWeb,
		callback,
		param)
}

func XWeb_OnLoadingFinish(hWeb, callback, param uintptr) {
	xWeb_OnLoadingFinish.Call(
		hWeb,
		callback,
		param)
}

func XWeb_IsLoadingSucceeded(hWeb uintptr) bool {
	ret, _, _ := xWeb_IsLoadingSucceeded.Call(hWeb)

	return ret == xc.TRUE
}

func XWeb_IsLoadingFailed(hWeb uintptr) bool {
	ret, _, _ := xWeb_IsLoadingFailed.Call(hWeb)

	return ret == xc.TRUE
}

func XWeb_IsLoadingCompleted(hWeb uintptr) bool {
	ret, _, _ := xWeb_IsLoadingCompleted.Call(hWeb)

	return ret == xc.TRUE
}

func XWeb_IsDocumentReady(hWeb uintptr) bool {
	ret, _, _ := xWeb_IsDocumentReady.Call(hWeb)

	return ret == xc.TRUE
}

func XWeb_StopLoading(hWeb uintptr) {
	xWeb_StopLoading.Call(hWeb)
}

func XWeb_Reload(hWeb uintptr) {
	xWeb_Reload.Call(hWeb)
}

func XWeb_JsBindFuction(funcName string, pFunc uintptr, argCount uint32) {
	xWeb_JsBindFuction.Call(
		xc.StringToUintPtr(funcName),
		pFunc,
		uintptr(argCount))
}

func XWeb_JsBindGetter(funcName string, pFunc uintptr) {
	xWeb_JsBindGetter.Call(
		xc.StringToUintPtr(funcName),
		pFunc)
}

func XWeb_JsBindSetter(funcName string, pFunc uintptr) {
	xWeb_JsBindSetter.Call(
		xc.StringToUintPtr(funcName),
		pFunc)
}

func XWeb_JsToInt(es uintptr, value int64) int {
	ret, _, _ := xWeb_JsToInt.Call(
		es,
		uintptr(value))

	return int(ret)
}

func XWeb_jsToFloat(es uintptr, value int64) float32 {
	ret, _, _ := xWeb_jsToFloat.Call(
		es,
		uintptr(value))

	return float32(ret)
}

func XWeb_jsToDouble(es uintptr, value int64) float64 {
	ret, _, _ := xWeb_jsToDouble.Call(
		es,
		uintptr(value))

	return float64(ret)
}

func XWeb_jsToBoolean(es uintptr, value int64) bool {
	ret, _, _ := xWeb_jsToBoolean.Call(
		es,
		uintptr(value))

	return ret == xc.TRUE
}

func XWeb_JsToTempStringW(es uintptr, value int64) string {
	ret, _, _ := xWeb_JsToTempStringW.Call(
		es,
		uintptr(value))

	return xc.UINTptrToString(ret)
}

func XWeb_JsToTempStringA(es uintptr, value int64) string {
	ret, _, _ := xWeb_JsToTempStringA.Call(
		es,
		uintptr(value))

	return xc.UTF8PtrToSting(ret)
}

func XWeb_JsArg(es uintptr, argIdx int) int64 {
	ret, _, _ := xWeb_JsArg.Call(
		es,
		uintptr(argIdx))

	return int64(ret)
}

func XWeb_JsInt(n int) int64 {
	ret, _, _ := xWeb_JsInt.Call(uintptr(n))

	return int64(ret)
}

func XWeb_JsDouble(d float32) int64 {
	ret, _, _ := xWeb_JsDouble.Call(uintptr(d))

	return int64(ret)
}

func XWeb_JsBoolean(b bool) int64 {
	ret, _, _ := xWeb_JsBoolean.Call(uintptr(xc.BoolToBOOL(b)))

	return int64(ret)
}

func XWeb_JsUndefined() int64 {
	ret, _, _ := xWeb_JsUndefined.Call()

	return int64(ret)
}

func XWeb_JsNull() int64 {
	ret, _, _ := xWeb_JsNull.Call()

	return int64(ret)
}

func XWeb_JsTrue() int64 {
	ret, _, _ := xWeb_JsTrue.Call()

	return int64(ret)
}

func XWeb_JsFalse() int64 {
	ret, _, _ := xWeb_JsFalse.Call()

	return int64(ret)
}

func XWeb_JsStringA(es uintptr, str string) int64 {
	strA := xc.StringToBytePtr(str)
	ret, _, _ := xWeb_JsStringA.Call(
		es,
		uintptr(unsafe.Pointer(strA)))

	return int64(ret)
}

func XWeb_jsStringW(es uintptr, str string) int64 {
	ret, _, _ := xWeb_jsStringW.Call(
		es,
		xc.StringToUintPtr(str))

	return int64(ret)
}

func XWeb_JsGlobalObject(es uintptr) int64 {
	ret, _, _ := xWeb_JsGlobalObject.Call(es)

	return int64(ret)
}

func XWeb_JsGet(es uintptr, object int64, prop string) int64 {
	ret, _, _ := xWeb_JsGet.Call(
		es,
		uintptr(object),
		xc.StringToUintPtr(prop))

	return int64(ret)
}

func XWeb_JsSet(es uintptr, object int64, prop string, value int64) {
	xWeb_JsSet.Call(
		es,
		uintptr(object),
		xc.StringToUintPtr(prop),
		uintptr(value))
}

func XWeb_JsGetAt(es uintptr, object int64, index int) int64 {
	ret, _, _ := xWeb_JsGetAt.Call(
		es,
		uintptr(object),
		uintptr(index))

	return int64(ret)
}

func XWeb_JsSetAt(es uintptr, object int64, index int, value int64) {
	xWeb_JsSetAt.Call(
		es,
		uintptr(object),
		uintptr(index),
		uintptr(value))
}

//
