package xc

import (
	"syscall"
	"unsafe"
)

// Font weight constants
const (
	FW_DONTCARE   = 0
	FW_THIN       = 100
	FW_EXTRALIGHT = 200
	FW_ULTRALIGHT = FW_EXTRALIGHT
	FW_LIGHT      = 300
	FW_NORMAL     = 400
	FW_REGULAR    = 400
	FW_MEDIUM     = 500
	FW_SEMIBOLD   = 600
	FW_DEMIBOLD   = FW_SEMIBOLD
	FW_BOLD       = 700
	FW_EXTRABOLD  = 800
	FW_ULTRABOLD  = FW_EXTRABOLD
	FW_HEAVY      = 900
	FW_BLACK      = FW_HEAVY
)

type LOGFONT struct {
	LfHeight         int32
	LfWidth          int32
	LfEscapement     int32
	LfOrientation    int32
	LfWeight         int32
	LfItalic         byte
	LfUnderline      byte
	LfStrikeOut      byte
	LfCharSet        byte
	LfOutPrecision   byte
	LfClipPrecision  byte
	LfQuality        byte
	LfPitchAndFamily byte
	LfFaceName       [LF_FACESIZE]uint16
}

var (
	// Functions
	xDraw_Create                 *syscall.Proc
	xDraw_Destroy                *syscall.Proc
	xDraw_SetOffset              *syscall.Proc
	xDraw_GetOffset              *syscall.Proc
	xDraw_RestoreGDIOBJ          *syscall.Proc
	xDraw_GetHDC                 *syscall.Proc
	xDraw_SetBrushColor          *syscall.Proc
	xDraw_SetTextVertical        *syscall.Proc
	xDraw_SetTextAlign           *syscall.Proc
	xDraw_SetFontX               *syscall.Proc
	xDraw_SetFont                *syscall.Proc
	xDraw_SetFont2               *syscall.Proc
	xDraw_SetFont3               *syscall.Proc
	xDraw_SetLineWidth           *syscall.Proc
	xDraw_SetBkMode              *syscall.Proc
	xDraw_CreateSolidBrush       *syscall.Proc
	xDraw_CreatePen              *syscall.Proc
	xDraw_CreateRectRgn          *syscall.Proc
	xDraw_CreateRoundRectRgn     *syscall.Proc
	xDraw_CreatePolygonRgn       *syscall.Proc
	xDraw_SelectClipRgn          *syscall.Proc
	xDraw_FillRect               *syscall.Proc
	xDraw_FillRectColor          *syscall.Proc
	xDraw_FillRgn                *syscall.Proc
	xDraw_FillEllipse            *syscall.Proc
	xDraw_DrawEllipse            *syscall.Proc
	xDraw_FillRoundRect          *syscall.Proc
	xDraw_DrawRoundRect          *syscall.Proc
	xDraw_Rectangle              *syscall.Proc
	xDraw_DrawGroupBox_Rect      *syscall.Proc
	xDraw_DrawGroupBox_RoundRect *syscall.Proc
	xDraw_GradientFill2          *syscall.Proc
	xDraw_GradientFill4          *syscall.Proc
	xDraw_FrameRgn               *syscall.Proc
	xDraw_FrameRect              *syscall.Proc
	xDraw_DrawLine               *syscall.Proc
	xDraw_FocusRect              *syscall.Proc
	xDraw_MoveToEx               *syscall.Proc
	xDraw_LineTo                 *syscall.Proc
	xDraw_Polyline               *syscall.Proc
	xDraw_Dottedline             *syscall.Proc
	xDraw_SetPixel               *syscall.Proc
	xDraw_Check                  *syscall.Proc
	xDraw_DrawIconEx             *syscall.Proc
	xDraw_BitBlt                 *syscall.Proc
	xDraw_BitBlt2                *syscall.Proc
	xDraw_AlphaBlend             *syscall.Proc
	xDraw_TriangularArrow        *syscall.Proc
	xDraw_Image                  *syscall.Proc
	xDraw_Image2                 *syscall.Proc
	xDraw_ImageStretch           *syscall.Proc
	xDraw_ImageAdaptive          *syscall.Proc
	xDraw_ImageExTile            *syscall.Proc
	xDraw_ImageSuper             *syscall.Proc
	xDraw_ImageSuper2            *syscall.Proc
	xDraw_DrawText               *syscall.Proc
	xDraw_DrawTextUnderline      *syscall.Proc
	xDraw_TextOut                *syscall.Proc
	xDraw_TextOutA               *syscall.Proc
)

func init() {
	// Functions
	xDraw_Create = xcDLL.MustFindProc("XDraw_Create")
	xDraw_Destroy = xcDLL.MustFindProc("XDraw_Destroy")
	xDraw_SetOffset = xcDLL.MustFindProc("XDraw_SetOffset")
	xDraw_GetOffset = xcDLL.MustFindProc("XDraw_GetOffset")
	xDraw_RestoreGDIOBJ = xcDLL.MustFindProc("XDraw_RestoreGDIOBJ")
	xDraw_GetHDC = xcDLL.MustFindProc("XDraw_GetHDC")
	xDraw_SetBrushColor = xcDLL.MustFindProc("XDraw_SetBrushColor")
	xDraw_SetTextVertical = xcDLL.MustFindProc("XDraw_SetTextVertical")
	xDraw_SetTextAlign = xcDLL.MustFindProc("XDraw_SetTextAlign")
	xDraw_SetFontX = xcDLL.MustFindProc("XDraw_SetFontX")
	xDraw_SetFont = xcDLL.MustFindProc("XDraw_SetFont")
	xDraw_SetFont2 = xcDLL.MustFindProc("XDraw_SetFont2")
	xDraw_SetFont3 = xcDLL.MustFindProc("XDraw_SetFont3")
	xDraw_SetLineWidth = xcDLL.MustFindProc("XDraw_SetLineWidth")
	xDraw_SetBkMode = xcDLL.MustFindProc("XDraw_SetBkMode")
	xDraw_CreateSolidBrush = xcDLL.MustFindProc("XDraw_CreateSolidBrush")
	xDraw_CreatePen = xcDLL.MustFindProc("XDraw_CreatePen")
	xDraw_CreateRectRgn = xcDLL.MustFindProc("XDraw_CreateRectRgn")
	xDraw_CreateRoundRectRgn = xcDLL.MustFindProc("XDraw_CreateRoundRectRgn")
	xDraw_CreatePolygonRgn = xcDLL.MustFindProc("XDraw_CreatePolygonRgn")
	xDraw_SelectClipRgn = xcDLL.MustFindProc("XDraw_SelectClipRgn")
	xDraw_FillRect = xcDLL.MustFindProc("XDraw_FillRect")
	xDraw_FillRectColor = xcDLL.MustFindProc("XDraw_FillRectColor")
	xDraw_FillRgn = xcDLL.MustFindProc("XDraw_FillRgn")
	xDraw_FillEllipse = xcDLL.MustFindProc("XDraw_FillEllipse")
	xDraw_DrawEllipse = xcDLL.MustFindProc("XDraw_DrawEllipse")
	xDraw_FillRoundRect = xcDLL.MustFindProc("XDraw_FillRoundRect")
	xDraw_DrawRoundRect = xcDLL.MustFindProc("XDraw_DrawRoundRect")
	xDraw_Rectangle = xcDLL.MustFindProc("XDraw_Rectangle")
	xDraw_DrawGroupBox_Rect = xcDLL.MustFindProc("XDraw_DrawGroupBox_Rect")
	xDraw_DrawGroupBox_RoundRect = xcDLL.MustFindProc("XDraw_DrawGroupBox_RoundRect")
	xDraw_GradientFill2 = xcDLL.MustFindProc("XDraw_GradientFill2")
	xDraw_GradientFill4 = xcDLL.MustFindProc("XDraw_GradientFill4")
	xDraw_FrameRgn = xcDLL.MustFindProc("XDraw_FrameRgn")
	xDraw_FrameRect = xcDLL.MustFindProc("XDraw_FrameRect")
	xDraw_DrawLine = xcDLL.MustFindProc("XDraw_DrawLine")
	xDraw_FocusRect = xcDLL.MustFindProc("XDraw_FocusRect")
	xDraw_MoveToEx = xcDLL.MustFindProc("XDraw_MoveToEx")
	xDraw_LineTo = xcDLL.MustFindProc("XDraw_LineTo")
	xDraw_Polyline = xcDLL.MustFindProc("XDraw_Polyline")
	xDraw_Dottedline = xcDLL.MustFindProc("XDraw_Dottedline")
	xDraw_SetPixel = xcDLL.MustFindProc("XDraw_SetPixel")
	xDraw_Check = xcDLL.MustFindProc("XDraw_Check")
	xDraw_DrawIconEx = xcDLL.MustFindProc("XDraw_DrawIconEx")
	xDraw_BitBlt = xcDLL.MustFindProc("XDraw_BitBlt")
	xDraw_BitBlt2 = xcDLL.MustFindProc("XDraw_BitBlt2")
	xDraw_AlphaBlend = xcDLL.MustFindProc("XDraw_AlphaBlend")
	xDraw_TriangularArrow = xcDLL.MustFindProc("XDraw_TriangularArrow")
	xDraw_Image = xcDLL.MustFindProc("XDraw_Image")
	xDraw_Image2 = xcDLL.MustFindProc("XDraw_Image2")
	xDraw_ImageStretch = xcDLL.MustFindProc("XDraw_ImageStretch")
	xDraw_ImageAdaptive = xcDLL.MustFindProc("XDraw_ImageAdaptive")
	xDraw_ImageExTile = xcDLL.MustFindProc("XDraw_ImageExTile")
	xDraw_ImageSuper = xcDLL.MustFindProc("XDraw_ImageSuper")
	xDraw_ImageSuper2 = xcDLL.MustFindProc("XDraw_ImageSuper2")
	xDraw_DrawText = xcDLL.MustFindProc("XDraw_DrawText")
	xDraw_DrawTextUnderline = xcDLL.MustFindProc("XDraw_DrawTextUnderline")
	xDraw_TextOut = xcDLL.MustFindProc("XDraw_TextOut")
	xDraw_TextOutA = xcDLL.MustFindProc("XDraw_TextOutA")
}

/*
创建图形绘制模块实例.

参数:
	hdc 设备上下文句柄HDC.
返回:
	图形绘制模块实例句柄.
*/
func XDrawCreate(hdc HDC) HDRAW {
	ret, _, _ := xDraw_Create.Call(uintptr(hdc))

	return HDRAW(ret)
}

/*
销毁图形绘制模块实例句柄.

参数:
	hDraw 图形绘制句柄.
*/
func XDrawDestroy(hDraw HDRAW) {
	xDraw_Destroy.Call(uintptr(hDraw))
}

/*
设置坐标偏移量,X向左偏移为负数,向右偏移为正数.

参数:
	hDraw 图形绘制句柄.
	x X轴偏移量.
	y Y轴偏移量.
*/
func XDrawSetOffset(hDraw HDRAW, x, y int) {
	xDraw_SetOffset.Call(
		uintptr(hDraw),
		uintptr(x),
		uintptr(y))
}

/*
获取坐标偏移量,X向左偏移为负数,向右偏移为正数.

参数:
	hDraw 图形绘制句柄.
	pX 接收X轴偏移量.
	pY 接收Y轴偏移量.
*/
func XDrawGetOffset(hDraw HDRAW, pX, pY *uint16) {
	xDraw_GetOffset.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pX)),
		uintptr(unsafe.Pointer(pY)))
}

/*
还原状态,释放用户绑定的GDI对象,例如画刷,画笔.

参数:
	hDraw 图形绘制句柄.
*/
func XDrawRestoreGDIOBJ(hDraw HDRAW) {
	xDraw_RestoreGDIOBJ.Call(uintptr(hDraw))
}

/*
获取绑定的设备上下文HDC.

参数:
	hDraw 图形绘制句柄.
返回:
	返回HDC句柄.
*/
func XDrawGetHDC(hDraw HDRAW) HDC {
	ret, _, _ := xDraw_GetHDC.Call(uintptr(hDraw))

	return HDC(ret)
}

/*
设置画刷颜色.

参数:
	hDraw 图形绘制句柄.
	color RGB颜色值
	alpha 透明度,0-255,0完全透明,255不透明.
*/
func XDrawSetBrushColor(hDraw HDRAW, color COLORREF, alpha byte) {
	xDraw_SetBrushColor.Call(
		uintptr(hDraw),
		uintptr(color),
		uintptr(alpha))
}

/*
设置文本垂直显示.

参数:
	hDraw 图形绘制句柄.
	bVertical 是否垂直显示文本.
*/
func XDrawSetTextVertical(hDraw HDRAW, bVertical bool) {
	xDraw_SetTextVertical.Call(
		uintptr(hDraw),
		uintptr(BoolToBOOL(bVertical)))
}

/*
设置文本对齐.

参数:
	hDraw 图形绘制句柄.
	nFlag 对齐标识.
*/
func XDrawSetTextAlign(hDraw HDRAW, nFlag int) {
	xDraw_SetTextAlign.Call(
		uintptr(hDraw),
		uintptr(nFlag))
}

/*
设置字体.

参数:
	hDraw 图形绘制句柄.
	hFontx 炫彩字体.
*/
func XDrawSetFontX(hDraw HDRAW, hFontx HFONTX) {
	xDraw_SetFontX.Call(
		uintptr(hDraw),
		uintptr(hFontx))
}

/*
设置字体.

参数:
	hDraw 图形绘制句柄.
	size 大小.
*/
func XDrawSetFont(hDraw HDRAW, size int) {
	xDraw_SetFont.Call(
		uintptr(hDraw),
		uintptr(size))
}

/*
设置字体.

参数:
	hDraw 图形绘制句柄.
	pName 字体名称.
	size 大小.
*/
func XDrawSetFont2(hDraw HDRAW, pName *uint16, size int) {
	xDraw_SetFont2.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pName)),
		uintptr(size))
}

/*
设置字体.

参数:
	hDraw 图形绘制句柄.
	pLogFont 字体信息.
*/
func XDrawSetFont3(hDraw HDRAW, pLogFont *LOGFONT) {
	xDraw_SetFont3.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pLogFont)))
}

/*
设置线宽.

参数:
	hDraw 图形绘制句柄.
	nWidth 宽度.
*/
func XDrawSetLineWidth(hDraw HDRAW, nWidth int) {
	xDraw_SetLineWidth.Call(
		uintptr(hDraw),
		uintptr(nWidth))
}

/*
SetBkMode() 参见MSDN.

参数:
	hDraw 图形绘制句柄.
	bTransparent 参见MSDN.
返回:
	参见MSDN.
*/
func XDrawSetBkMode(hDraw HDRAW, bTransparent bool) int {
	ret, _, _ := xDraw_SetBkMode.Call(
		uintptr(hDraw),
		uintptr(BoolToBOOL(bTransparent)))

	return int(ret)
}

/*
创建具有指定的纯色逻辑刷.

参数:
	hDraw 图形绘制句柄.
	crColor 画刷颜色.
返回:
	如果函数成功,返回值标识一个逻辑刷,如果函数失败,返回值是NULL.
*/
func XDrawCreateSolidBrush(hDraw HDRAW, crColor COLORREF) HBRUSH {
	ret, _, _ := xDraw_CreateSolidBrush.Call(
		uintptr(hDraw),
		uintptr(crColor))

	if ret == NULL {
		return 0
	}

	return HBRUSH(ret)
}

/*
创建一个逻辑笔,指定的样式,宽度和颜色,随后的笔可以选择到设备上下文,用于绘制线条和曲线.

参数:
	hDraw 图形绘制句柄.
	fnPenStyle 画笔颜色.
	nWidth 画笔宽度.
	crColor 颜色.
返回:
	如果函数成功,返回值是一个句柄,标识一个逻辑笔,如果函数失败,返回值是NULL.
*/
func XDrawCreatePen(hDraw HDRAW, fnPenStyle, nWidth int, crColor COLORREF) HPEN {
	ret, _, _ := xDraw_CreatePen.Call(
		uintptr(hDraw),
		uintptr(fnPenStyle),
		uintptr(nWidth),
		uintptr(crColor))

	if ret == NULL {
		return 0
	}
	return HPEN(ret)
}

/*
创建矩形区域.

参数:
	hDraw 图形绘制句柄.
	nLeftRect 左上角X坐标.
	nTopRect 左上角Y坐标.
	nRightRect 右下角X坐标.
	nBottomRect 右下角Y坐标.
返回:
	成功返回区域句柄,失败返回NULL.
*/
func XDrawCreateRectRgn(hDraw HDRAW, nLeftRect, nTopRect, nRightRect, nBottomRect int) HRGN {
	ret, _, _ := xDraw_CreateRectRgn.Call(
		uintptr(hDraw),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect))

	if ret == NULL {
		return 0
	}

	return HRGN(ret)
}

/*
创建一个圆角的矩形区域.

参数:
	hDraw 图形绘制句柄.
	nLeftRect X-坐标的左上角.
	nTopRect Y-坐标左上角坐标
	nRightRect X-坐标右下角
	nBottomRect Y-坐标右下角
	nWidthEllipse 椭圆的宽度.
	nHeightEllipse 椭圆的高度.
返回:
	如果函数成功,返回值是该区域的句柄,如果函数失败,返回值是NULL.
*/
func XDrawCreateRoundRectRgn(hDraw HDRAW, nLeftRect, nTopRect, nRightRect, nBottomRect, nWidthEllipse, nHeightEllipse int) HRGN {
	ret, _, _ := xDraw_CreateRoundRectRgn.Call(
		uintptr(hDraw),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect),
		uintptr(nWidthEllipse),
		uintptr(nHeightEllipse))

	if ret == NULL {
		return 0
	}

	return HRGN(ret)
}

/*
创建一个多边形区域.

参数:
	hDraw 图形绘制句柄.
	pPt POINT数组.
	cPoints 数组大小.
	fnPolyFillMode 多边形填充模式,指定用于确定在该地区的像素填充模式,这个参数可以是下列值之一.
	ALTERNATE Selects alternate mode (fills area between odd-numbered and even-numbered polygon sides on each scan line).
	 WINDING Selects winding mode (fills any region with a nonzero winding value).
返回:
	如果函数成功,返回值是该区域的句柄,如果函数失败,返回值是NULL.
*/
func XDrawCreatePolygonRgn(hDraw HDRAW, pPt *POINT, cPoints, fnPolyFillMode int) HRGN {
	ret, _, _ := xDraw_CreatePolygonRgn.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(cPoints),
		uintptr(fnPolyFillMode))

	if ret == NULL {
		return 0
	}

	return HRGN(ret)
}

/*
选择一个区域作为当前裁剪区域.

参数:
	hDraw 图形绘制句柄.
	hRgn 区域句柄.
返回:
	返回值指定地区的复杂性，可以是下列值之一.
	NULLREGION Region is empty.
	 SIMPLEREGION Region is a single rectangle.
	 COMPLEXREGION Region is more than one rectangle.
	 ERROR An error occurred. (The previous clipping region is unaffected).
*/
func XDrawSelectClipRgn(hDraw HDRAW, hRgn HRGN) int {
	ret, _, _ := xDraw_SelectClipRgn.Call(
		uintptr(hDraw),
		uintptr(hRgn))

	return int(ret)
}

/*
通过使用指定的刷子填充一个矩形,此功能包括左侧和顶部的边界,但不包括矩形的右边和底部边界.

参数:
	hDraw 图形绘制句柄.
	pRect 矩形区域.
*/
func XDrawFillRect(hDraw HDRAW, pRect *RECT) {
	xDraw_FillRect.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)))
}

/*
填充矩形.

参数:
	hDraw 图形绘制句柄.
	pRect XX.
	color 矩形区域.
	alpha 透明度.
*/
func XDrawFillRectColor(hDraw HDRAW, pRect *RECT, color COLORREF, alpha byte) {
	xDraw_FillRectColor.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(color),
		uintptr(alpha))
}

/*
通过使用指定的画刷填充一个区域.

参数:
	hDraw 图形绘制句柄.
	hrgn 区域句柄.
	hbr 画刷句柄.
返回:
	如果函数成功,返回非零值,如果函数失败,返回值是零.
*/
func XDrawFillRgn(hDraw HDRAW, hrgn HRGN, hbr HBRUSH) bool {
	ret, _, _ := xDraw_FillRgn.Call(
		uintptr(hDraw),
		uintptr(hrgn),
		uintptr(hbr))

	return ret != FALSE
}

/*
填充圆.

参数:
	hDraw 图形绘制句柄.
	pRect 矩形区域
*/
func XDrawFillEllipse(hDraw HDRAW, pRect *RECT) {
	xDraw_FillEllipse.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)))
}

/*
绘制圆边框.

参数:
	hDraw 图形绘制句柄.
	pRect 矩形区域.
*/
func XDrawDrawEllipse(hDraw HDRAW, pRect *RECT) {
	xDraw_DrawEllipse.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)))
}

/*
填充圆角矩形.

参数:
	hDraw 图形绘制句柄.
	pRect 矩形坐标.
	nWidth 圆角宽度.
	nHeight 圆角高度.
*/
func XDrawFillRoundRect(hDraw HDRAW, pRect *RECT, nWidth, nHeight int) {
	xDraw_FillRoundRect.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(nWidth),
		uintptr(nHeight))
}

/*
绘制圆角矩形边框.

参数:
	hDraw 图形绘制句柄.
	pRect 矩形坐标.
	nWidth 圆角宽度.
	nHeight 圆角高度.
*/
func XDrawDrawRoundRect(hDraw HDRAW, pRect *RECT, nWidth, nHeight int) {
	xDraw_DrawRoundRect.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(nWidth),
		uintptr(nHeight))
}

/*
绘制矩形,使用当前的画刷和画笔.

参数:
	hDraw 图形绘制句柄.
	nLeftRect 左上角X坐标.
	nTopRect 左上角Y坐标.
	nRightRect 右下角X坐标.
	nBottomRect 右下角Y坐标.
返回:
	如果函数成功,返回非零值,如果函数失败,返回值是零.
*/
func XDrawRectangle(hDraw HDRAW, nLeftRect, nTopRect, nRightRect, nBottomRect int) bool {
	ret, _, _ := xDraw_Rectangle.Call(
		uintptr(hDraw),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect))

	return ret != FALSE
}

/*
绘制组框,矩形边框.

参数:
	hDraw 图形绘制句柄.
	pRect 矩形坐标.
	pName 标题名称.
	textColor 文本颜色.
	textAlpha 文本透明度.
	pOffset 文本偏移
*/
func XDrawDrawGroupBoxRect(hDraw HDRAW, pRect *RECT, pName *uint16, textColor COLORREF, textAlpha byte, pOffset *POINT) {
	xDraw_DrawGroupBox_Rect.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(unsafe.Pointer(pName)),
		uintptr(textColor),
		uintptr(textAlpha),
		uintptr(unsafe.Pointer(pOffset)))
}

/*
绘制组框,圆角边框.

参数:
	hDraw 图形绘制句柄.
	pRect 矩形坐标.
	pName 标题名称.
	textColor 文本颜色.
	textAlpha 文本透明度.
	pOffset 文本偏移
	nWidth 圆角宽度.
	nHeight 圆角高度.
*/
func XDrawDrawGroupBoxRoundRect(hDraw HDRAW, pRect *RECT, pName *uint16, textColor COLORREF, textAlpha byte, pOffset *POINT, nWidth, nHeight int) {
	xDraw_DrawGroupBox_RoundRect.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(unsafe.Pointer(pName)),
		uintptr(textColor),
		uintptr(textAlpha),
		uintptr(unsafe.Pointer(pOffset)),
		uintptr(nWidth),
		uintptr(nHeight))
}

/*
渐变填充,从一种颜色过渡到另一种颜色.

参数:
	hDraw 图形绘制句柄.
	color1 开始颜色.
	alpha1 透明度.
	color2 结束颜色.
	alpha2 透明度.
	pRect 矩形坐标.
	mode 模式. GRADIENT_FILL_RECT_H 水平填充 . GRADIENT_FILL_RECT_V 垂直填充. GRADIENT_FILL_TRIANGLE 三角形.
*/
func XDrawGradientFill2(hDraw HDRAW, color1 COLORREF, alpha1 byte, color2 COLORREF, alpha2 byte, pRect *RECT, mode int) {
	xDraw_GradientFill2.Call(
		uintptr(hDraw),
		uintptr(color1),
		uintptr(alpha1),
		uintptr(color2),
		uintptr(alpha2),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(mode))
}

/*
渐变填充,从一种颜色过渡到另一种颜色.

参数:
	hDraw 图形绘制句柄.
	color1 开始颜色.
	color2 结束颜色,中间.
	color3 开始颜色,中间.
	color4 结束颜色.
	pRect 矩形坐标.
	mode 模式. GRADIENT_FILL_RECT_H 水平填充. GRADIENT_FILL_RECT_V 垂直填充. GRADIENT_FILL_TRIANGLE 三角形.
返回:
	如果函数成功，返回值为TRUE,如果函数失败,返回值是FALSE.
*/
func XDrawGradientFill4(hDraw HDRAW, color1, color2, color3, color4 COLORREF, pRect *RECT, mode int) bool {
	ret, _, _ := xDraw_GradientFill4.Call(
		uintptr(hDraw),
		uintptr(color1),
		uintptr(color2),
		uintptr(color3),
		uintptr(color4),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(mode))

	return ret == TRUE
}

/*
绘制边框,使用指定的画刷绘制指定的区域的边框.

参数:
	hDraw 图形绘制句柄.
	hrgn 区域句柄.
	hbr 画刷句柄.
	nWidth 边框宽度,垂直边.
	nHeight 边框高度,水平边.
返回:
	如果函数成功,返回非零值,如果函数失败,返回值是零.
*/
func XDrawFrameRgn(hDraw HDRAW, hrgn HRGN, hbr HBRUSH, nWidth, nHeight int) bool {
	ret, _, _ := xDraw_FrameRgn.Call(
		uintptr(hDraw),
		uintptr(hrgn),
		uintptr(hbr),
		uintptr(nWidth),
		uintptr(nHeight))

	return ret != FALSE
}

/*
绘制矩形边框,使用指定的画刷.

参数:
	hDraw 图形绘制句柄.
	pRect 矩形坐标
*/
func XDrawFrameRect(hDraw HDRAW, pRect *RECT) {
	xDraw_FrameRect.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)))
}

/*
绘制线条.

参数:
	hDraw 图形绘制句柄.
	x1 坐标.
	y1 坐标.
	x2 坐标.
	y2 坐标.
*/
func XDrawDrawLine(hDraw HDRAW, x1, y1, x2, y2 int) {
	xDraw_DrawLine.Call(
		uintptr(hDraw),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2))
}

/*
绘制焦点矩形.

参数:
	hDraw 图形绘制句柄.
	pRect 矩形坐标.
*/
func XDrawFocusRect(hDraw HDRAW, pRect *RECT) {
	xDraw_FocusRect.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pRect)))
}

/*
更新当前位置到指定点，并返回以前的位置。

参数:
	hDraw 图形绘制句柄.
	X 坐标.
	Y 坐标.
	lpPoint 接收以前的当前位置到一个POINT结构的指针,如果这个参数是NULL指针,没有返回原来的位置.
返回:
	如果函数成功,返回非零值,如果函数失败,返回值是零.
*/
func XDrawMoveToEx(hDraw HDRAW, X, Y int, lpPoint *POINT) bool {
	ret, _, _ := xDraw_MoveToEx.Call(
		uintptr(hDraw),
		uintptr(X),
		uintptr(Y),
		uintptr(unsafe.Pointer(lpPoint)))

	return ret != FALSE
}

/*
函数绘制一条线从当前位置到,但不包括指定点.

参数:
	hDraw 图形绘制句柄.
	nXEnd X坐标,线结束点.
	nYEnd Y坐标,线结束点.
返回:
	如果函数成功,返回非零值,如果函数失败,返回值是零.
*/
func XDrawLineTo(hDraw HDRAW, nXEnd, nYEnd int) bool {
	ret, _, _ := xDraw_LineTo.Call(
		uintptr(hDraw),
		uintptr(nXEnd),
		uintptr(nYEnd))

	return ret != FALSE
}

/*
Polyline() 参见MSDN.用当前画笔描绘一系列线段。

参数:
	hDraw 图形绘制句柄.
	pArrayPt 参见MSDN.
	arrayPtSize 参见MSDN.
返回:
	参见MSDN.
	bool，非零表示成功，零表示失败
*/
func XDrawPolyline(hDraw HDRAW, pArrayPt *POINT, arrayPtSize int) bool {
	ret, _, _ := xDraw_Polyline.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(pArrayPt)),
		uintptr(arrayPtSize))

	return ret != FALSE
}

/*
绘制水平或垂直虚线.

参数:
	hDraw 图形绘制句柄.
	x1 起点x坐标.
	y1 起点y坐标.
	x2 结束点x坐标.
	y2 结束点y坐标.
*/
func XDrawDottedline(hDraw HDRAW, x1, y1, x2, y2 int) {
	xDraw_Dottedline.Call(
		uintptr(hDraw),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2))
}

/*
函数设置在指定的坐标到指定的颜色的像素.

参数:
	hDraw 图形绘制句柄.
	X 坐标
	Y 坐标
	crColor RGB颜色值
返回:
	如果函数成功返回RGB值,如果失败返回-1.
*/
func XDrawSetPixel(hDraw HDRAW, X, Y int, crColor COLORREF) COLORREF {
	ret, _, _ := xDraw_SetPixel.Call(
		uintptr(hDraw),
		uintptr(X),
		uintptr(Y),
		uintptr(crColor))

	return COLORREF(ret)
}

/*
绘制复选框.

参数:
	hDraw 图形绘制句柄.
	x 坐标.
	y 坐标.
	color 边框颜色.
	bCheck 是否选中状态.
*/
func XDrawCheck(hDraw HDRAW, x, y int, color COLORREF, bCheck bool) {
	xDraw_Check.Call(
		uintptr(hDraw),
		uintptr(x),
		uintptr(y),
		uintptr(color),
		uintptr(BoolToBOOL(bCheck)))
}

/*
绘制图标,DrawIconEx()参见MSDN.
	http://baike.baidu.com/link?url=OAc-hbHFn9QC_Ay-OgcUUhcTAwDKb3CylCdNF5C_TlXsZy1N_sYFaBZ7DmqVfTDbkhxDqcfuYPV2BEeP-JR_zq

参数:
	hDraw .
	xLeft .
	yTop .
	hIcon .
	cxWidth .
	cyWidth .
	istepIfAniCur .
	hbrFlickerFreeDraw .
	diFlags .
返回:
*/
func XDrawDrawIconEx(hDraw HDRAW, xLeft, yTop int, hIcon HICON, cxWidth, cyWidth int, istepIfAniCur uint32, hbrFlickerFreeDraw HBRUSH, diFlags uint32) bool {
	ret, _, _ := xDraw_DrawIconEx.Call(
		uintptr(hDraw),
		uintptr(xLeft),
		uintptr(yTop),
		uintptr(hIcon),
		uintptr(cxWidth),
		uintptr(cyWidth),
		uintptr(istepIfAniCur),
		uintptr(hbrFlickerFreeDraw),
		uintptr(diFlags))

	return ret != FALSE
}

/*
BitBlt() 参见MSDN.
	http://baike.baidu.com/link?url=1qMIDG4Yd4KkVVoa1ritf_wyGbPR_JbiaecQVsz_qJfFWeK2vD6qPOO5ztuR8JQQVJa73epTX0Qh99qGFqXkQa
参数:
	hDrawDest XX.
	nXDest XX.
	nYDest XX.
	nWidth XX.
	nHeight XX.
	hdcSrc XX.
	nXSrc XX.
	nYSrc XX.
	dwRop XX. DWORD 现在表示 32bit 无符号整数
返回:
*/
func XDrawBitBlt(hDrawDest HDRAW, nXDest, nYDest, nWidth, nHeight int, hdcSrc HDC, nXSrc, nYSrc int, dwRop uint32) bool {
	ret, _, _ := xDraw_BitBlt.Call(
		uintptr(hDrawDest),
		uintptr(nXDest),
		uintptr(nYDest),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hdcSrc),
		uintptr(nXSrc),
		uintptr(nYSrc),
		uintptr(dwRop))

	return ret != FALSE
}

/*
BitBlt() 参见MSDN.
	http://baike.baidu.com/link?url=1qMIDG4Yd4KkVVoa1ritf_wyGbPR_JbiaecQVsz_qJfFWeK2vD6qPOO5ztuR8JQQVJa73epTX0Qh99qGFqXkQa

参数:
	hDrawDest XX.
	nXDest XX.
	nYDest XX.
	nWidth XX.
	nHeight XX.
	hDrawSrc XX.
	nXSrc XX.
	nYSrc XX.
	dwRop XX.
返回:
*/
func XDrawBitBlt2(hDrawDest HDRAW, nXDest, nYDest, nWidth, nHeight int, hDrawSrc HDRAW, nXSrc, nYSrc int, dwRop uint32) bool {
	ret, _, _ := xDraw_BitBlt2.Call(
		uintptr(hDrawDest),
		uintptr(nXDest),
		uintptr(nYDest),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hDrawSrc),
		uintptr(nXSrc),
		uintptr(nYSrc),
		uintptr(dwRop))

	return ret != FALSE
}

/*
AlphaBlend() 参见MSDN.
	http://baike.baidu.com/link?url=gRaZjsaQDo5L4crox6slDq_ncVV_rAe5mX3v96P089Lp4OfDNCiHIJw9CdVeuqXywqdAi728e1OiRMw40ba5Yq

参数:
	hDraw XX.
	nXOriginDest XX.
	nYOriginDest XX.
	nWidthDest XX.
	nHeightDest XX.
	hdcSrc XX.
	nXOriginSrc XX.
	nYOriginSrc XX.
	nWidthSrc XX.
	nHeightSrc XX.
	alpha XX.
返回:
	成功返回TRUE否则返回FALSE.
*/
func XDrawAlphaBlend(hDraw HDRAW, nXOriginDest, nYOriginDest, nWidthDest, nHeightDest int, hdcSrc HDC, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc, alpha int) bool {
	ret, _, _ := xDraw_AlphaBlend.Call(
		uintptr(hDraw),
		uintptr(nXOriginDest),
		uintptr(nYOriginDest),
		uintptr(nWidthDest),
		uintptr(nHeightDest),
		uintptr(hdcSrc),
		uintptr(nXOriginSrc),
		uintptr(nYOriginSrc),
		uintptr(nWidthSrc),
		uintptr(nHeightSrc),
		uintptr(alpha))

	return ret == TRUE
}

/*
绘制三角型箭头.

参数:
	hDraw 图形绘制句柄.
	align 箭头方向,左1,上2,右3,下4.
	x 中心点X坐标.
	y 中心点Y坐标.
	width 三角形宽度.
	height 三角形高度.
*/
func XDrawTriangularArrow(hDraw HDRAW, align, x, y, width, height int) {
	xDraw_TriangularArrow.Call(
		uintptr(hDraw),
		uintptr(align),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
}

/*
绘制图片.

参数:
	hDraw 图形绘制句柄.
	hImage 图片句柄.
	x x坐标.
	y y坐标.
*/
func XDrawImage(hDraw HDRAW, hImage HIMAGE, x, y int) {
	xDraw_Image.Call(
		uintptr(hDraw),
		uintptr(hImage),
		uintptr(x),
		uintptr(y))
}

/*
绘制图片.

参数:
	hDraw 图形绘制句柄.
	hImage 图片句柄.
	x x坐标.
	y y坐标.
	width 宽度.
	height 高度.
*/
func XDrawImage2(hDraw HDRAW, hImage HIMAGE, x, y, width, height int) {
	xDraw_Image2.Call(
		uintptr(hDraw),
		uintptr(hImage),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
}

/*
绘制图片.

参数:
	hDraw 图形绘制句柄.
	hImage 图片句柄.
	x x坐标.
	y y坐标.
	width 宽度.
	height 高度.
*/
func XDrawImageStretch(hDraw HDRAW, hImage HIMAGE, x, y, width, height int) {
	xDraw_ImageStretch.Call(
		uintptr(hDraw),
		uintptr(hImage),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height))
}

/*
绘制图片.

参数:
	hDraw 图形绘制句柄.
	hImage 图片句柄.
	pRect 坐标.
	bOnlyBorder 是否只绘制边缘区域.
*/
func XDrawImageAdaptive(hDraw HDRAW, hImage HIMAGE, pRect *RECT, bOnlyBorder bool) {
	xDraw_ImageAdaptive.Call(
		uintptr(hDraw),
		uintptr(hImage),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(BoolToBOOL(bOnlyBorder)))
}

/*
绘制图片.

参数:
	hDraw 图形绘制句柄.
	hImage 图片句柄.
	pRect 坐标.
	flag 标识.
*/
func XDrawImageExTile(hDraw HDRAW, hImage HIMAGE, pRect *RECT, flag int) {
	xDraw_ImageExTile.Call(
		uintptr(hDraw),
		uintptr(hImage),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(flag))
}

/*
绘制图片.

参数:
	hDraw 图形绘制句柄.
	hImage 图片句柄.
	pRect 坐标.
	bClip 是否裁剪区域.
*/
func XDrawImageSuper(hDraw HDRAW, hImage HIMAGE, pRect *RECT, bClip bool) {
	xDraw_ImageSuper.Call(
		uintptr(hDraw),
		uintptr(hImage),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(BoolToBOOL(bClip)))
}

/*
绘制图片.

参数:
	hDraw 图形绘制句柄.
	hImage 图片句柄.
	pRcDest 目标坐标.
	pSrcRect 源坐标.
*/
func XDrawImageSuper2(hDraw HDRAW, hImage HIMAGE, pRcDest, pSrcRect *RECT) {
	xDraw_ImageSuper2.Call(
		uintptr(hDraw),
		uintptr(hImage),
		uintptr(unsafe.Pointer(pRcDest)),
		uintptr(unsafe.Pointer(pSrcRect)))
}

/*
DrawText() 参见MSDN.
	http://baike.baidu.com/link?url=OB4-3_d1OSNh1tHWftQjq2TQ1n-njLtquQ-_UAfENKVepRM1zIhMHOO-qF1KwiS_fUqKzA2wTFZ0jmFbXUtFta

参数:
	hDraw 图形绘制句柄.
	lpString 字符串.
	nCount 字符串长度.
	lpRect 坐标.
*/
func XDrawDrawText(hDraw HDRAW, lpString *uint16, nCount int, lpRect *RECT) {
	xDraw_DrawText.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(nCount),
		uintptr(unsafe.Pointer(lpRect)))
}

/*
参见MSDN.

参数:
	hDraw 图形绘制句柄.
	lpString 字符串.
	nCount 字符串长度.
	lpRect 坐标.
	colorLine 下划线颜色.
	alphaLine 下划线透明度.
*/
func XDrawDrawTextUnderline(hDraw HDRAW, lpString *uint16, nCount int, lpRect *RECT, colorLine COLORREF, alphaLine byte) {
	xDraw_DrawTextUnderline.Call(
		uintptr(hDraw),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(nCount),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(colorLine),
		uintptr(alphaLine))
}

/*
TextOut() 参见MSDN.
	http://baike.baidu.com/link?url=KLWCibEvLOSaHGZdqNDfcQspO6uNPJENmSNKmS7bW0g8ilzSAMWY8PcFLa9kuMT91b8vJ22V04THVBHEjyyj7a

参数:
	hDraw 图形绘制句柄.
	nXStart XX.
	nYStart XX.
	lpString XX.
	cbString XX.
*/
func XDrawTextOut(hDraw HDRAW, nXStart, nYStart int, lpString *uint16, cbString int) {
	xDraw_TextOut.Call(
		uintptr(hDraw),
		uintptr(nXStart),
		uintptr(nYStart),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(cbString))
}

/*
TextOut() 参见MSDN.
	http://baike.baidu.com/link?url=KLWCibEvLOSaHGZdqNDfcQspO6uNPJENmSNKmS7bW0g8ilzSAMWY8PcFLa9kuMT91b8vJ22V04THVBHEjyyj7a

参数:
	hDraw 图形绘制句柄.
	nXStart XX.
	nYStart XX.
	lpString XX.
返回:
*/
func XDrawTextOutA(hDraw HDRAW, nXStart, nYStart int, lpString *uint16) {
	xDraw_TextOutA.Call(
		uintptr(hDraw),
		uintptr(nXStart),
		uintptr(nYStart),
		uintptr(unsafe.Pointer(lpString)))
}
