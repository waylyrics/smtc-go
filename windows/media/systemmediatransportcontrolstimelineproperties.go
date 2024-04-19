// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package media

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/waylyrics/winrt-go/windows/foundation"
)

const SignatureSystemMediaTransportControlsTimelineProperties string = "rc(Windows.Media.SystemMediaTransportControlsTimelineProperties;{5125316a-c3a2-475b-8507-93534dc88f15})"

type SystemMediaTransportControlsTimelineProperties struct {
	ole.IUnknown
}

func NewSystemMediaTransportControlsTimelineProperties() (*SystemMediaTransportControlsTimelineProperties, error) {
	inspectable, err := ole.RoActivateInstance("Windows.Media.SystemMediaTransportControlsTimelineProperties")
	if err != nil {
		return nil, err
	}
	return (*SystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(inspectable)), nil
}

func (impl *SystemMediaTransportControlsTimelineProperties) GetStartTime() (foundation.TimeSpan, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControlsTimelineProperties))
	defer itf.Release()
	v := (*iSystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(itf))
	return v.GetStartTime()
}

func (impl *SystemMediaTransportControlsTimelineProperties) SetStartTime(value foundation.TimeSpan) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControlsTimelineProperties))
	defer itf.Release()
	v := (*iSystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(itf))
	return v.SetStartTime(value)
}

func (impl *SystemMediaTransportControlsTimelineProperties) GetEndTime() (foundation.TimeSpan, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControlsTimelineProperties))
	defer itf.Release()
	v := (*iSystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(itf))
	return v.GetEndTime()
}

func (impl *SystemMediaTransportControlsTimelineProperties) SetEndTime(value foundation.TimeSpan) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControlsTimelineProperties))
	defer itf.Release()
	v := (*iSystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(itf))
	return v.SetEndTime(value)
}

func (impl *SystemMediaTransportControlsTimelineProperties) GetMinSeekTime() (foundation.TimeSpan, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControlsTimelineProperties))
	defer itf.Release()
	v := (*iSystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(itf))
	return v.GetMinSeekTime()
}

func (impl *SystemMediaTransportControlsTimelineProperties) SetMinSeekTime(value foundation.TimeSpan) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControlsTimelineProperties))
	defer itf.Release()
	v := (*iSystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(itf))
	return v.SetMinSeekTime(value)
}

func (impl *SystemMediaTransportControlsTimelineProperties) GetMaxSeekTime() (foundation.TimeSpan, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControlsTimelineProperties))
	defer itf.Release()
	v := (*iSystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(itf))
	return v.GetMaxSeekTime()
}

func (impl *SystemMediaTransportControlsTimelineProperties) SetMaxSeekTime(value foundation.TimeSpan) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControlsTimelineProperties))
	defer itf.Release()
	v := (*iSystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(itf))
	return v.SetMaxSeekTime(value)
}

func (impl *SystemMediaTransportControlsTimelineProperties) GetPosition() (foundation.TimeSpan, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControlsTimelineProperties))
	defer itf.Release()
	v := (*iSystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(itf))
	return v.GetPosition()
}

func (impl *SystemMediaTransportControlsTimelineProperties) SetPosition(value foundation.TimeSpan) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiSystemMediaTransportControlsTimelineProperties))
	defer itf.Release()
	v := (*iSystemMediaTransportControlsTimelineProperties)(unsafe.Pointer(itf))
	return v.SetPosition(value)
}

const GUIDiSystemMediaTransportControlsTimelineProperties string = "5125316a-c3a2-475b-8507-93534dc88f15"
const SignatureiSystemMediaTransportControlsTimelineProperties string = "{5125316a-c3a2-475b-8507-93534dc88f15}"

type iSystemMediaTransportControlsTimelineProperties struct {
	ole.IInspectable
}

type iSystemMediaTransportControlsTimelinePropertiesVtbl struct {
	ole.IInspectableVtbl

	GetStartTime   uintptr
	SetStartTime   uintptr
	GetEndTime     uintptr
	SetEndTime     uintptr
	GetMinSeekTime uintptr
	SetMinSeekTime uintptr
	GetMaxSeekTime uintptr
	SetMaxSeekTime uintptr
	GetPosition    uintptr
	SetPosition    uintptr
}

func (v *iSystemMediaTransportControlsTimelineProperties) VTable() *iSystemMediaTransportControlsTimelinePropertiesVtbl {
	return (*iSystemMediaTransportControlsTimelinePropertiesVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iSystemMediaTransportControlsTimelineProperties) GetStartTime() (foundation.TimeSpan, error) {
	var out foundation.TimeSpan
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetStartTime,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.TimeSpan
	)

	if hr != 0 {
		return foundation.TimeSpan{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iSystemMediaTransportControlsTimelineProperties) SetStartTime(value foundation.TimeSpan) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetStartTime,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&value)), // in foundation.TimeSpan
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iSystemMediaTransportControlsTimelineProperties) GetEndTime() (foundation.TimeSpan, error) {
	var out foundation.TimeSpan
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetEndTime,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.TimeSpan
	)

	if hr != 0 {
		return foundation.TimeSpan{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iSystemMediaTransportControlsTimelineProperties) SetEndTime(value foundation.TimeSpan) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetEndTime,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&value)), // in foundation.TimeSpan
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iSystemMediaTransportControlsTimelineProperties) GetMinSeekTime() (foundation.TimeSpan, error) {
	var out foundation.TimeSpan
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetMinSeekTime,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.TimeSpan
	)

	if hr != 0 {
		return foundation.TimeSpan{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iSystemMediaTransportControlsTimelineProperties) SetMinSeekTime(value foundation.TimeSpan) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetMinSeekTime,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&value)), // in foundation.TimeSpan
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iSystemMediaTransportControlsTimelineProperties) GetMaxSeekTime() (foundation.TimeSpan, error) {
	var out foundation.TimeSpan
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetMaxSeekTime,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.TimeSpan
	)

	if hr != 0 {
		return foundation.TimeSpan{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iSystemMediaTransportControlsTimelineProperties) SetMaxSeekTime(value foundation.TimeSpan) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetMaxSeekTime,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&value)), // in foundation.TimeSpan
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iSystemMediaTransportControlsTimelineProperties) GetPosition() (foundation.TimeSpan, error) {
	var out foundation.TimeSpan
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetPosition,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.TimeSpan
	)

	if hr != 0 {
		return foundation.TimeSpan{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iSystemMediaTransportControlsTimelineProperties) SetPosition(value foundation.TimeSpan) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetPosition,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&value)), // in foundation.TimeSpan
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
