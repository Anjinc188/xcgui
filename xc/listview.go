package xc

import (
	"syscall"
	"unsafe"
)

var (
	// Functions
	xListView_Create                       *syscall.Proc
	xListView_BindAdapter                  *syscall.Proc
	xListView_GetAdapter                   *syscall.Proc
	xListView_SetItemTemplateXML           *syscall.Proc
	xListView_SetItemTemplateXMLFromString *syscall.Proc
	xListView_GetTemplateObject            *syscall.Proc
	xListView_GetTemplateObjectGroup       *syscall.Proc
	xListView_GetItemIDFromHXCGUI          *syscall.Proc
	xListView_HitTest                      *syscall.Proc
	xListView_HitTestOffset                *syscall.Proc
	xListView_EnableMultiSel               *syscall.Proc
	xListView_SetDrawItemBkFlags           *syscall.Proc
	xListView_SetSelectItem                *syscall.Proc
	xListView_GetSelectItem                *syscall.Proc
	xListView_GetSelectItemCount           *syscall.Proc
	xListView_GetSelectItemAll             *syscall.Proc
	xListView_SetSelectItemAll             *syscall.Proc
	xListView_CancelSelectItemAll          *syscall.Proc
	xListView_SetColumnSpace               *syscall.Proc
	xListView_SetRowSpace                  *syscall.Proc
	xListView_SetAlignSizeLeft             *syscall.Proc
	xListView_SetAlignSizeTop              *syscall.Proc
	xListView_SetItemSize                  *syscall.Proc
	xListView_GetItemSize                  *syscall.Proc
	xListView_SetGroupHeight               *syscall.Proc
	xListView_GetGroupHeight               *syscall.Proc
	xListView_AddItemBkBorder              *syscall.Proc
	xListView_AddItemBkFill                *syscall.Proc
	xListView_AddItemBkImage               *syscall.Proc
	xListView_GetItemBkInfoCount           *syscall.Proc
	xListView_ClearItemBkInfo              *syscall.Proc
	xListView_GetItemBkInfoManager         *syscall.Proc
	xListView_ExpandGroup                  *syscall.Proc
)

func init() {
	// Functions
	xListView_Create = XCDLL.MustFindProc("XListView_Create")
	xListView_BindAdapter = XCDLL.MustFindProc("XListView_BindAdapter")
	xListView_GetAdapter = XCDLL.MustFindProc("XListView_GetAdapter")
	xListView_SetItemTemplateXML = XCDLL.MustFindProc("XListView_SetItemTemplateXML")
	xListView_SetItemTemplateXMLFromString = XCDLL.MustFindProc("XListView_SetItemTemplateXMLFromString")
	xListView_GetTemplateObject = XCDLL.MustFindProc("XListView_GetTemplateObject")
	xListView_GetTemplateObjectGroup = XCDLL.MustFindProc("XListView_GetTemplateObjectGroup")
	xListView_GetItemIDFromHXCGUI = XCDLL.MustFindProc("XListView_GetItemIDFromHXCGUI")
	xListView_HitTest = XCDLL.MustFindProc("XListView_HitTest")
	xListView_HitTestOffset = XCDLL.MustFindProc("XListView_HitTestOffset")
	xListView_EnableMultiSel = XCDLL.MustFindProc("XListView_EnableMultiSel")
	xListView_SetDrawItemBkFlags = XCDLL.MustFindProc("XListView_SetDrawItemBkFlags")
	xListView_SetSelectItem = XCDLL.MustFindProc("XListView_SetSelectItem")
	xListView_GetSelectItem = XCDLL.MustFindProc("XListView_GetSelectItem")
	xListView_GetSelectItemCount = XCDLL.MustFindProc("XListView_GetSelectItemCount")
	xListView_GetSelectItemAll = XCDLL.MustFindProc("XListView_GetSelectItemAll")
	xListView_SetSelectItemAll = XCDLL.MustFindProc("XListView_SetSelectItemAll")
	xListView_CancelSelectItemAll = XCDLL.MustFindProc("XListView_CancelSelectItemAll")
	xListView_SetColumnSpace = XCDLL.MustFindProc("XListView_SetColumnSpace")
	xListView_SetRowSpace = XCDLL.MustFindProc("XListView_SetRowSpace")
	xListView_SetAlignSizeLeft = XCDLL.MustFindProc("XListView_SetAlignSizeLeft")
	xListView_SetAlignSizeTop = XCDLL.MustFindProc("XListView_SetAlignSizeTop")
	xListView_SetItemSize = XCDLL.MustFindProc("XListView_SetItemSize")
	xListView_GetItemSize = XCDLL.MustFindProc("XListView_GetItemSize")
	xListView_SetGroupHeight = XCDLL.MustFindProc("XListView_SetGroupHeight")
	xListView_GetGroupHeight = XCDLL.MustFindProc("XListView_GetGroupHeight")
	xListView_AddItemBkBorder = XCDLL.MustFindProc("XListView_AddItemBkBorder")
	xListView_AddItemBkFill = XCDLL.MustFindProc("XListView_AddItemBkFill")
	xListView_AddItemBkImage = XCDLL.MustFindProc("XListView_AddItemBkImage")
	xListView_GetItemBkInfoCount = XCDLL.MustFindProc("XListView_GetItemBkInfoCount")
	xListView_ClearItemBkInfo = XCDLL.MustFindProc("XListView_ClearItemBkInfo")
	xListView_GetItemBkInfoManager = XCDLL.MustFindProc("XListView_GetItemBkInfoManager")
	xListView_ExpandGroup = XCDLL.MustFindProc("XListView_ExpandGroup")
}

func XListViewCreate(x, y, cx, cy int, hParent HXCGUI) HELE {
	ret, _, _ := xListView_Create.Call(
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hParent))

	return HELE(ret)
}

func XListViewBindAdapter(hEle HELE, hAdapter HXCGUI) {
	xListView_BindAdapter.Call(
		uintptr(hEle),
		uintptr(hAdapter))
}

func XListViewGetAdapter(hEle HELE) HXCGUI {
	ret, _, _ := xListView_GetAdapter.Call(uintptr(hEle))

	return HXCGUI(ret)
}

func XListViewSetItemTemplateXML(hEle HELE, pXmlFile string) bool {
	ret, _, _ := xListView_SetItemTemplateXML.Call(
		uintptr(hEle),
		StringToUintPtr(pXmlFile))

	if ret != TRUE {
		return false
	}

	return true
}

func XListViewSetItemTemplateXMLFromString(hEle HELE, pStringXML string) bool {
	ret, _, _ := xListView_SetItemTemplateXMLFromString.Call(
		uintptr(hEle),
		StringToUintPtr(pStringXML))

	if ret != TRUE {
		return false
	}

	return true
}

func XListViewGetTemplateObject(hEle HELE, iGroup, iItem, nTempItemID int) HXCGUI {
	ret, _, _ := xListView_GetTemplateObject.Call(
		uintptr(hEle),
		uintptr(iGroup),
		uintptr(iItem),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

func XListViewGetTemplateObjectGroup(hEle HELE, iGroup, nTempItemID int) HXCGUI {
	ret, _, _ := xListView_GetTemplateObjectGroup.Call(
		uintptr(hEle),
		uintptr(iGroup),
		uintptr(nTempItemID))

	return HXCGUI(ret)
}

func XListViewGetItemIDFromHXCGUI(hEle HELE, hXCGUI HXCGUI, piGroup, piItem *uint16) bool {
	ret, _, _ := xListView_GetItemIDFromHXCGUI.Call(
		uintptr(hEle),
		uintptr(hXCGUI),
		uintptr(unsafe.Pointer(piGroup)),
		uintptr(unsafe.Pointer(piItem)))

	if ret != TRUE {
		return false
	}

	return true
}

func XListViewHitTest(hEle HELE, pPt *POINT, pOutGroup, pOutItem *uint16) bool {
	ret, _, _ := xListView_HitTest.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(unsafe.Pointer(pOutGroup)),
		uintptr(unsafe.Pointer(pOutItem)))

	if ret != TRUE {
		return false
	}

	return true
}

func XListViewHitTestOffset(hEle HELE, pPt *POINT, pOutGroup, pOutItem *uint16) int {
	ret, _, _ := xListView_HitTestOffset.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pPt)),
		uintptr(unsafe.Pointer(pOutGroup)),
		uintptr(unsafe.Pointer(pOutItem)))

	return int(ret)
}

func XListViewEnableMultiSel(hEle HELE, bEnable bool) {
	xListView_EnableMultiSel.Call(
		uintptr(hEle),
		uintptr(BoolToBOOL(bEnable)))
}

func XListViewSetDrawItemBkFlags(hEle HELE, nFlags LIST_DRAWITEMBK_FLAGS_) {
	xListView_SetDrawItemBkFlags.Call(
		uintptr(hEle),
		uintptr(nFlags))
}

func XListViewSetSelectItem(hEle HELE, iGroup, iItem int) bool {
	ret, _, _ := xListView_SetSelectItem.Call(
		uintptr(hEle),
		uintptr(iGroup),
		uintptr(iItem))

	if ret != TRUE {
		return false
	}

	return true
}

func XListViewGetSelectItem(hEle HELE, piGroup, piItem *uint16) bool {
	ret, _, _ := xListView_GetSelectItem.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(piGroup)),
		uintptr(unsafe.Pointer(piItem)))

	if ret != TRUE {
		return false
	}

	return true

}

func XListViewGetSelectItemCount(hEle HELE) int {
	ret, _, _ := xListView_GetSelectItemCount.Call(uintptr(hEle))

	return int(ret)
}

func XListViewGetSelectItemAll(hEle HELE, pArray *LISTVIEW_ITEM_ID_I, nArraySize int) int {
	ret, _, _ := xListView_GetSelectItemAll.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pArray)),
		uintptr(nArraySize))

	return int(ret)
}

func XListViewSetSelectItemAll(hEle HELE) {
	xListView_SetSelectItemAll.Call(uintptr(hEle))

}

func XListViewCancelSelectItemAll(hEle HELE) {
	xListView_CancelSelectItemAll.Call(uintptr(hEle))
}

func XListViewSetColumnSpace(hEle HELE, space int) {
	xListView_SetColumnSpace.Call(
		uintptr(hEle),
		uintptr(space))
}

func XListViewSetRowSpace(hEle HELE, space int) {
	xListView_SetRowSpace.Call(
		uintptr(hEle),
		uintptr(space))
}

func XListViewSetAlignSizeLeft(hEle HELE, size int) {
	xListView_SetAlignSizeLeft.Call(
		uintptr(hEle),
		uintptr(size))
}

func XListViewSetAlignSizeTop(hEle HELE, size int) {
	xListView_SetAlignSizeTop.Call(
		uintptr(hEle),
		uintptr(size))
}

func XListViewSetItemSize(hEle HELE, width, height int) {
	xListView_SetItemSize.Call(
		uintptr(hEle),
		uintptr(width),
		uintptr(height))
}

func XListViewGetItemSize(hEle HELE, pSize *SIZE) {
	xListView_GetItemSize.Call(
		uintptr(hEle),
		uintptr(unsafe.Pointer(pSize)))
}

func XListViewSetGroupHeight(hEle HELE, height int) {
	xListView_SetGroupHeight.Call(
		uintptr(hEle),
		uintptr(height))
}

func XListViewGetGroupHeight(hEle HELE) int {
	ret, _, _ := xListView_GetGroupHeight.Call(uintptr(hEle))

	return int(ret)
}

func XListViewAddItemBkBorder(hEle HELE, nState LIST_ITEM_STATE_, color COLORREF, alpha byte, width int) {
	xListView_AddItemBkBorder.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha),
		uintptr(width))
}

func XListViewAddItemBkFill(hEle HELE, nState LIST_ITEM_STATE_, color COLORREF, alpha byte) {
	xListView_AddItemBkFill.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(color),
		uintptr(alpha))
}

func XListViewAddItemBkImage(hEle HELE, nState LIST_ITEM_STATE_, hImage HIMAGE) {
	xListView_AddItemBkImage.Call(
		uintptr(hEle),
		uintptr(nState),
		uintptr(hImage))
}

func XListViewGetItemBkInfoCount(hEle HELE, nState LIST_ITEM_STATE_) int {
	ret, _, _ := xListView_GetItemBkInfoCount.Call(
		uintptr(hEle),
		uintptr(nState))

	return int(ret)
}

func XListViewClearItemBkInfo(hEle HELE, nState LIST_ITEM_STATE_) {
	xListView_ClearItemBkInfo.Call(
		uintptr(hEle),
		uintptr(nState))
}

func XListViewGetItemBkInfoManager(hEle HELE, nState LIST_ITEM_STATE_) HBKINFOM {
	ret, _, _ := xListView_GetItemBkInfoManager.Call(
		uintptr(hEle),
		uintptr(nState))

	return HBKINFOM(ret)
}

func XListViewExpandGroup(hEle HELE, iGroup int, bExpand bool) bool {
	ret, _, _ := xListView_ExpandGroup.Call(
		uintptr(hEle),
		uintptr(iGroup),
		uintptr(BoolToBOOL(bExpand)))

	if ret != TRUE {
		return false
	}

	return true
}
