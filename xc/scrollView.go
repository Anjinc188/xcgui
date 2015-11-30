package xc

import (
	"syscall"
	"unsafe"
)

type PaddingSize_ int

var (
	xSView_Create                  *syscall.Proc
	xSView_SetTotalSize            *syscall.Proc
	xSView_GetTotalSize            *syscall.Proc
	xSView_SetLineSize             *syscall.Proc
	xSView_GetLineSize             *syscall.Proc
	xSView_GetViewPosH             *syscall.Proc
	xSView_GetViewPosV             *syscall.Proc
	xSView_GetViewWidth            *syscall.Proc
	xSView_GetViewHeight           *syscall.Proc
	xSView_GetViewRect             *syscall.Proc
	xSView_GetScrollBarH           *syscall.Proc
	xSView_GetScrollBarV           *syscall.Proc
	xSView_SetPadding              *syscall.Proc
	xSView_GetPadding              *syscall.Proc
	xSView_ScrollPosH              *syscall.Proc
	xSView_ScrollPosV              *syscall.Proc
	xSView_ScrollPosXH             *syscall.Proc
	xSView_ScrollPosYV             *syscall.Proc
	xSView_ShowSBarH               *syscall.Proc
	xSView_ShowSBarV               *syscall.Proc
	xSView_EnableAutoShowScrollBar *syscall.Proc
	xSView_ScrollLeftLine          *syscall.Proc
	xSView_ScrollRightLine         *syscall.Proc
	xSView_ScrollTopLine           *syscall.Proc
	xSView_ScrollBottomLine        *syscall.Proc
	xSView_ScrollLeft              *syscall.Proc
	xSView_ScrollRight             *syscall.Proc
	xSView_ScrollTop               *syscall.Proc
	xSView_ScrollBottom            *syscall.Proc
)

func init() {
	xSView_Create = xcDLL.MustFindProc("XSView_Create")
	xSView_SetTotalSize = xcDLL.MustFindProc("XSView_SetTotalSize")
	xSView_GetTotalSize = xcDLL.MustFindProc("XSView_GetTotalSize")
	xSView_SetLineSize = xcDLL.MustFindProc("XSView_SetLineSize")
	xSView_GetLineSize = xcDLL.MustFindProc("XSView_GetLineSize")
	xSView_GetViewPosH = xcDLL.MustFindProc("XSView_GetViewPosH")
	xSView_GetViewPosV = xcDLL.MustFindProc("XSView_GetViewPosV")
	xSView_GetViewWidth = xcDLL.MustFindProc("XSView_GetViewWidth")
	xSView_GetViewHeight = xcDLL.MustFindProc("XSView_GetViewHeight")
	xSView_GetViewRect = xcDLL.MustFindProc("XSView_GetViewRect")
	xSView_GetScrollBarH = xcDLL.MustFindProc("XSView_GetScrollBarH")
	xSView_GetScrollBarV = xcDLL.MustFindProc("XSView_GetScrollBarV")
	xSView_SetPadding = xcDLL.MustFindProc("XSView_SetPadding")
	xSView_GetPadding = xcDLL.MustFindProc("XSView_GetPadding")
	xSView_ScrollPosH = xcDLL.MustFindProc("XSView_ScrollPosH")
	xSView_ScrollPosV = xcDLL.MustFindProc("XSView_ScrollPosV")
	xSView_ScrollPosXH = xcDLL.MustFindProc("XSView_ScrollPosXH")
	xSView_ScrollPosYV = xcDLL.MustFindProc("XSView_ScrollPosYV")
	xSView_ShowSBarH = xcDLL.MustFindProc("XSView_ShowSBarH")
	xSView_ShowSBarV = xcDLL.MustFindProc("XSView_ShowSBarV")
	xSView_EnableAutoShowScrollBar = xcDLL.MustFindProc("XSView_EnableAutoShowScrollBar")
	xSView_ScrollLeftLine = xcDLL.MustFindProc("XSView_ScrollLeftLine")
	xSView_ScrollRightLine = xcDLL.MustFindProc("XSView_ScrollRightLine")
	xSView_ScrollTopLine = xcDLL.MustFindProc("XSView_ScrollTopLine")
	xSView_ScrollBottomLine = xcDLL.MustFindProc("XSView_ScrollBottomLine")
	xSView_ScrollLeft = xcDLL.MustFindProc("XSView_ScrollLeft")
	xSView_ScrollRight = xcDLL.MustFindProc("XSView_ScrollRight")
	xSView_ScrollTop = xcDLL.MustFindProc("XSView_ScrollTop")
	xSView_ScrollBottom = xcDLL.MustFindProc("XSView_ScrollBottom")
}

func XSViewCreate(x int, y int, cx int, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xSView_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))
	return HELE(ret)
}

func XSViewSetTotalSize(hEle HELE, cx int, cy int) bool {
	ret, _, _ := xSView_SetTotalSize.Call(
		uintptr(hEle),
		uintptr(cx),
		uintptr(cy))
	return ret == TRUE

}

func XSViewGetTotalSize(hEle HELE, pSize *SIZE) {
	xSView_GetTotalSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

func XSViewSetLineSize(hEle HELE, nHeight int, nWidth int) bool {
	ret, _, _ := xSView_SetLineSize.Call(
		uintptr(hEle),
		uintptr(nHeight),
		uintptr(nWidth))
	return ret == TRUE
}

func XSViewGetLineSize(hEle HELE, pSize *SIZE) {
	xSView_GetLineSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

func XSViewGetViewPosH(hEle HELE) int {
	ret, _, _ := xSView_GetViewPosH.Call(
		uintptr(hEle))
	return int(ret)
}

func XSViewGetViewPosV(hEle HELE) int {
	ret, _, _ := xSView_GetViewPosV.Call(
		uintptr(hEle))
	return int(ret)
}

func XSViewGetViewWidth(hEle HELE) int {
	ret, _, _ := xSView_GetViewWidth.Call(
		uintptr(hEle))
	return int(ret)
}

func XSViewGetViewHeight(hEle HELE) int {
	ret, _, _ := xSView_GetViewHeight.Call(
		uintptr(hEle))
	return int(ret)
}

func XSViewGetViewRect(hEle HELE, pRect *RECT) {
	xSView_GetLineSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pRect)))
}

func XSViewGetScrollBarH(hEle HELE) HELE {
	ret, _, _ := xSView_GetScrollBarH.Call(
		uintptr(hEle))
	return HELE(ret)
}

func XSViewGetScrollBarV(hEle HELE) HELE {
	ret, _, _ := xSView_GetScrollBarV.Call(
		uintptr(hEle))
	return HELE(ret)
}

func XSViewSetPadding(hEle HELE, left int, top int, right int, bottom int) {
	xSView_SetPadding.Call(
		uintptr(hEle),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))
}

func XSViewGetPadding(hEle HELE, pPadding *PaddingSize_) {
	xSView_GetPadding.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPadding)))
}

func XSViewScrollPosH(hEle HELE, pos int) bool {
	ret, _, _ := xSView_ScrollPosH.Call(
		uintptr(hEle),
		uintptr(pos))
	return ret == TRUE
}

func XSViewScrollPosV(hEle HELE, pos int) bool {
	ret, _, _ := xSView_ScrollPosV.Call(
		uintptr(hEle),
		uintptr(pos))
	return ret == TRUE
}

func XSViewScrollPosXH(hEle HELE, posX int) bool {
	ret, _, _ := xSView_ScrollPosXH.Call(
		uintptr(hEle),
		uintptr(posX))
	return ret == TRUE
}

func XSViewScrollPosYV(hEle HELE, posY int) bool {
	ret, _, _ := xSView_ScrollPosYV.Call(
		uintptr(hEle),
		uintptr(posY))
	return ret == TRUE
}

func XSViewShowSBarH(hEle HELE, bShow bool) {
	xSView_ShowSBarH.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bShow)))
}

func XSViewShowSBarV(hEle HELE, bShow bool) {
	xSView_ShowSBarV.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bShow)))
}

func XSViewEnableAutoShowScrollBar(hEle HELE, bEnable bool) {
	xSView_EnableAutoShowScrollBar.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

func XSViewScrollLeftLine(hEle HELE) bool {
	ret, _, _ := xSView_ScrollLeftLine.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSViewScrollRightLine(hEle HELE) bool {
	ret, _, _ := xSView_ScrollRightLine.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSViewScrollTopLine(hEle HELE) bool {
	ret, _, _ := xSView_ScrollTopLine.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSViewScrollBottomLine(hEle HELE) bool {
	ret, _, _ := xSView_ScrollBottomLine.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSViewScrollLeft(hEle HELE) bool {
	ret, _, _ := xSView_ScrollLeft.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSViewScrollRight(hEle HELE) bool {
	ret, _, _ := xSView_ScrollRight.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSViewScrollTop(hEle HELE) bool {
	ret, _, _ := xSView_ScrollTop.Call(
		uintptr(hEle))
	return ret == TRUE
}

func XSViewScrollBottom(hEle HELE) bool {
	ret, _, _ := xSView_ScrollBottom.Call(
		uintptr(hEle))
	return ret == TRUE
}
