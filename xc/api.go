package xc

import (
	// "fmt"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"
)

const (
	TRUE  = 1
	FALSE = 0
	NULL  = 0
)

const (
	xcDll = "XCGUI.dll"
)

var (
	XCDLL *syscall.DLL
)

var (
	// Functions
	xC_UnicodeToAnsi         *syscall.Proc
	xC_AnsiToUnicode         *syscall.Proc
	xC_DebugToFileInfo       *syscall.Proc
	xC_IsHELE                *syscall.Proc
	xC_IsHWINDOW             *syscall.Proc
	xC_IsShape               *syscall.Proc
	xC_IsHXCGUI              *syscall.Proc
	xC_hWindowFromHWnd       *syscall.Proc
	xC_IsSViewExtend         *syscall.Proc
	xC_GetObjectType         *syscall.Proc
	xC_GetObjectByID         *syscall.Proc
	xC_RectInRect            *syscall.Proc
	xC_CombineRect           *syscall.Proc
	xC_ShowLayoutFrame       *syscall.Proc
	xC_SetLayoutFrameColor   *syscall.Proc
	xC_EnableErrorMessageBox *syscall.Proc
	xInitXCGUI               *syscall.Proc
	xRunXCGUI                *syscall.Proc
	xExitXCGUI               *syscall.Proc
)

type POINT struct {
	X, Y int
}

type RECT struct {
	Left, Top, Right, Bottom int32
}

type SIZE struct {
	CX, CY int32
}

type BorderSize_ struct {
	Left, Top, Right, Bottom int32
}

type (
	BOOL    int32
	HRESULT int32
)

// 1.初始化UI库
func init() {
	runtime.LockOSThread()

	if FileExist(xcDll) {
		XCDLL = syscall.MustLoadDLL(xcDll)
	} else if FileExist("lib/" + xcDll) {
		XCDLL = syscall.MustLoadDLL("lib/" + xcDll)
	} else if FileExist("../lib/" + xcDll) {
		XCDLL = syscall.MustLoadDLL("../lib/" + xcDll)
	} else {
		panic("xcgui library not found")
	}

	// Functions
	xC_UnicodeToAnsi = XCDLL.MustFindProc("XC_UnicodeToAnsi")
	xC_AnsiToUnicode = XCDLL.MustFindProc("XC_AnsiToUnicode")
	xC_DebugToFileInfo = XCDLL.MustFindProc("XC_DebugToFileInfo")
	xC_IsHELE = XCDLL.MustFindProc("XC_IsHELE")
	xC_IsHWINDOW = XCDLL.MustFindProc("XC_IsHWINDOW")
	xC_IsShape = XCDLL.MustFindProc("XC_IsShape")
	xC_IsHXCGUI = XCDLL.MustFindProc("XC_IsHXCGUI")
	xC_hWindowFromHWnd = XCDLL.MustFindProc("XC_hWindowFromHWnd")
	xC_IsSViewExtend = XCDLL.MustFindProc("XC_IsSViewExtend")
	xC_GetObjectType = XCDLL.MustFindProc("XC_GetObjectType")
	xC_GetObjectByID = XCDLL.MustFindProc("XC_GetObjectByID")
	xC_RectInRect = XCDLL.MustFindProc("XC_RectInRect")
	xC_CombineRect = XCDLL.MustFindProc("XC_CombineRect")
	xC_ShowLayoutFrame = XCDLL.MustFindProc("XC_ShowLayoutFrame")
	xC_SetLayoutFrameColor = XCDLL.MustFindProc("XC_SetLayoutFrameColor")
	xC_EnableErrorMessageBox = XCDLL.MustFindProc("XC_EnableErrorMessageBox")
	xInitXCGUI = XCDLL.MustFindProc("XInitXCGUI")
	xRunXCGUI = XCDLL.MustFindProc("XRunXCGUI")
	xExitXCGUI = XCDLL.MustFindProc("XExitXCGUI")

	// *******************************************
	// @Author: cody.guo
	// @Date: 2015-11-7 09:40:36
	// @Function: XInitXCGUI
	// @Description: 初始化界面库.
	// @Calls: XInitXCGUI
	// @Input: pText 保留参数.
	// @Return: 成功返回TRUE否则返回FALSE.
	// *******************************************
	ret, _, _ := xInitXCGUI.Call(StringToUintPtr("XCGUI Library For Go"))
	// XCGUI的返回值: true 为 1 ，false 为 0
	if ret != TRUE {
		panic("XInitXCGUI call failed.")
	}
}

//由于在init已经调用过了，这里只留个函数名字
func XInitXCGUI() bool {
	//	pText := ""

	//	if len(arg) == 1 {
	//		pText = arg[0]
	//	}

	//	ret, _, _ := xInitXCGUI.Call(
	//		StringToUintPtr(pText),
	//		0,
	//		0)
	//	return ret == TRUE
	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 09:49:55
// @Function: XCUnicodeToAnsi
// @Description: Unicode转换Ansi编码,.
// @Calls: XC_UnicodeToAnsi
// @Input: [in] pIn 指向待转换的Unicode字符串指针. [in] inLen pIn字符数量.
//         [out] pOut 指向接收转换后的Ansi字符串缓冲区指针. [in] outLen pOut缓冲区大小,字节单位.
// @Return: 如果成功,返回写入接收缓冲区字节数量.
// *******************************************
func XCUnicodeToAnsi(pIn string, inLen int, pOut string, outLen int) int {
	ret, _, _ := xC_UnicodeToAnsi.Call(
		StringToUintPtr(pIn),
		uintptr(inLen),
		StringToUintPtr(pOut),
		uintptr(outLen))

	return int(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 09:56:23
// @Function: XCAnsiToUnicode
// @Description: Ansi转换Unicode编码,.
// @Calls: XC_AnsiToUnicode
// @Input: [in] pIn 指向待转换的Ansi字符串指针. [in] inLen pIn字符数量.
//         [out] pOut 指向接收转换后的Unicode字符串缓冲区指针. [in] pOut缓冲区大小,字符wchar_t单位.
// @Return: 如果成功,返回写入接收缓冲区字符wchar_t数量.
// *******************************************
func XCAnsiToUnicode(pIn string, inLen int, pOut string, outLen int) int {
	ret, _, _ := xC_AnsiToUnicode.Call(
		StringToUintPtr(pIn),
		uintptr(inLen),
		StringToUintPtr(pOut),
		uintptr(outLen))

	return int(ret)
}

// 打印调试信息到文件xcgui_debug.txt.
func XCDebugToFileInfo(pInfo string) {
	// fmt.Println("正在打印debug", pInfo)
	p := (*struct {
		pInfo uintptr
		len   int
	})(unsafe.Pointer(&pInfo))
	// fmt.Println("p.info:", p.pInfo)

	xC_DebugToFileInfo.Call(p.pInfo)

	// str := syscall.StringToUTF16(pInfo)
	// s := syscall.UTF16ToString(str)
	// fmt.Println("str长度:", len(str), str, s)
	// m := str[0 : len(str)-1]
	// xC_DebugToFileInfo.Call(uintptr(unsafe.Pointer(&m)))
	// fmt.Println(pInfo)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 17:33:08
// @Function: XCIsHELE
// @Description: 判断是否为元素句柄.
// @Calls: XC_IsHELE
// @Input: hEle 元素句柄.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XCIsHELE(hEle HXCGUI) bool {
	ret, _, _ := xC_IsHELE.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 17:35:34
// @Function: XCIsHWINDOW
// @Description: 判断是否为窗口句柄.
// @Calls: XC_IsHWINDOW
// @Input: hWindow 窗口句柄.
// @Return: 成功返回TRUE,否则相反.
// *******************************************
func XCIsHWINDOW(hWindow HXCGUI) bool {
	ret, _, _ := xC_IsHWINDOW.Call(uintptr(hWindow))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 17:37:47
// @Function: XCIsShape
// @Description: 判断是否为形状对象.
// @Calls: XC_IsShape
// @Input: hShape 形状对象句柄.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XCIsShape(hShape HXCGUI) bool {
	ret, _, _ := xC_IsShape.Call(uintptr(hShape))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 17:43:41
// @Function: XCIsHXCGUI
// @Description: 判断句柄是否拥有该类型.
// @Calls: XC_IsHXCGUI
// @Input: hXCGUI 炫彩句柄. nType 句柄类型. 参考objectType.go中的 XC_OBJECT_TYPE.
// @Return: 成功返回TRUE否则返回FALSE.
// *******************************************
func XCIsHXCGUI(hXCGUI HXCGUI, nType XC_OBJECT_TYPE) bool {
	ret, _, _ := xC_IsHXCGUI.Call(
		uintptr(hXCGUI),
		uintptr(nType))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 18:22:44
// @Function: XChWindowFromHWnd
// @Description: 通过窗口HWND句柄获取HWINDOW句柄.
// @Calls: XC_hWindowFromHWnd
// @Input: hWnd 窗口HWND句柄.
// @Return: 返回HWINDOW句柄.
// *******************************************
func XChWindowFromHWnd(hWnd HWND) HWINDOW {
	ret, _, _ := xC_hWindowFromHWnd.Call(uintptr(hWnd))

	return HWINDOW(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 18:24:36
// @Function: XCIsSViewExtend
// @Description: 判断元素是否从滚动视图元素扩展的新元素,包含滚动视图元素.
// @Calls: XC_IsSViewExtend
// @Input: hEle 元素句柄.
// @Return: 如果是返回TRUE,否则相反.
// *******************************************
func XCIsSViewExtend(hEle HELE) bool {
	ret, _, _ := xC_IsSViewExtend.Call(uintptr(hEle))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 18:26:37
// @Function: XCGetObjectType
// @Description: 获取句柄类型.
// @Calls: XC_GetObjectType
// @Input: hXCGUI 炫彩对象句柄.
// @Return: 返回句柄类型.
// *******************************************
func XCGetObjectType(hXCGUI HXCGUI) XC_OBJECT_TYPE {
	ret, _, _ := xC_GetObjectType.Call(uintptr(hXCGUI))

	return XC_OBJECT_TYPE(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 18:31:03
// @Function: XCGetObjectByID
// @Description: 通过ID获取对象句柄.
// @Calls: XC_GetObjectByID
// @Input: nID ID值.
// @Return: 成功返回句柄,否则返回NULL.
// *******************************************
func XCGetObjectByID(nID int) HXCGUI {
	ret, _, _ := xC_GetObjectByID.Call(uintptr(nID))

	if ret == NULL {
		panic("XCGetObjectByID get HXCGUI failed.")
	}

	return HXCGUI(ret)
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 18:36:38
// @Function: XCRectInRect
// @Description: 判断两个矩形是否相交及重叠.
// @Calls: XC_RectInRect
// @Input: pRect1 矩形1. pRect2 矩形2.
// @Return: 如果两个矩形相交返回TRUE,否则相反.
// *******************************************
func XCRectInRect(pRect1, pRect2 *RECT) bool {
	ret, _, _ := xC_RectInRect.Call(
		uintptr(unsafe.Pointer(pRect1)),
		uintptr(unsafe.Pointer(pRect2)))

	if ret != TRUE {
		return false
	}

	return true
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 18:47:04
// @Function: XCCombineRect
// @Description: 组合两个矩形区域.
// @Calls: XC_CombineRect
// @Input: pDest 新的矩形区域. pSrc1 源矩形1. pSrc2 源矩形2.
// @Return:
// *******************************************
func XCCombineRect(pDest, pSrc1, pSrc2 *RECT) {
	xC_CombineRect.Call(
		uintptr(unsafe.Pointer(pDest)),
		uintptr(unsafe.Pointer(pSrc1)),
		uintptr(unsafe.Pointer(pSrc2)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 18:49:09
// @Function: XCShowLayoutFrame
// @Description: 显示布局对象边界.
// @Calls: XC_ShowLayoutFrame
// @Input: bShow 是否显示.
// @Return:
// *******************************************
func XCShowLayoutFrame(bShow bool) {
	xC_ShowLayoutFrame.Call(uintptr(BoolToBOOL(bShow)))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 18:51:12
// @Function: XCSetLayoutFrameColor
// @Description: 设置布局边框颜色.
// @Calls: XC_SetLayoutFrameColor
// @Input: color RGB颜色值.
// @Return:
// *******************************************
func XCSetLayoutFrameColor(color COLORREF) {
	xC_SetLayoutFrameColor.Call(uintptr(color))
}

// *******************************************
// @Author: cody.guo
// @Date: 2015-11-7 18:55:21
// @Function: XCEnableErrorMessageBox
// @Description: 启用错误弹出,通过该接口可以设置遇到严重错误时不弹出消息提示框.
// @Calls: XC_EnableErrorMessageBox
// @Input: bEnabel 是否启用.
// @Return:
// *******************************************
func XCEnableErrorMessageBox(bEnabel bool) {
	xC_EnableErrorMessageBox.Call(uintptr(BoolToBOOL(bEnabel)))
}

// 运行消息循环,当炫彩窗口数量为0时退出.
func XRunXCGUI() {
	xRunXCGUI.Call()
}

// 退出界面库释放资源.
func XExitXCGUI() {
	xExitXCGUI.Call()
}

func StringToUintPtr(str string) uintptr {
	// fmt.Println("正在转换..", str)
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}

func FullPath(path string) (p string) {
	p, _ = filepath.Abs(path)
	return
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	} else {
		return true
	}

}

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}

	return 0
}

func UTF16PtrToString(s *uint16) string {
	if s == nil {
		return ""
	}
	return syscall.UTF16ToString((*[1 << 29]uint16)(unsafe.Pointer(s))[0:])
}
