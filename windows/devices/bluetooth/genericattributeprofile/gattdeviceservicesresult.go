// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint
package genericattributeprofile

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation/collections"
)

const SignatureGattDeviceServicesResult string = "rc(Windows.Devices.Bluetooth.GenericAttributeProfile.GattDeviceServicesResult;{171dd3ee-016d-419d-838a-576cf475a3d8})"

type GattDeviceServicesResult struct {
	ole.IUnknown
}

func (impl *GattDeviceServicesResult) GetStatus() (GattCommunicationStatus, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattDeviceServicesResult))
	defer itf.Release()
	v := (*iGattDeviceServicesResult)(unsafe.Pointer(itf))
	return v.GetStatus()
}

func (impl *GattDeviceServicesResult) GetServices() (*collections.IVectorView, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattDeviceServicesResult))
	defer itf.Release()
	v := (*iGattDeviceServicesResult)(unsafe.Pointer(itf))
	return v.GetServices()
}

const GUIDiGattDeviceServicesResult string = "171dd3ee-016d-419d-838a-576cf475a3d8"
const SignatureiGattDeviceServicesResult string = "{171dd3ee-016d-419d-838a-576cf475a3d8}"

type iGattDeviceServicesResult struct {
	ole.IInspectable
}

type iGattDeviceServicesResultVtbl struct {
	ole.IInspectableVtbl

	GetStatus        uintptr
	GetProtocolError uintptr
	GetServices      uintptr
}

func (v *iGattDeviceServicesResult) VTable() *iGattDeviceServicesResultVtbl {
	return (*iGattDeviceServicesResultVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iGattDeviceServicesResult) GetStatus() (GattCommunicationStatus, error) {
	var out GattCommunicationStatus
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetStatus,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out GattCommunicationStatus
	)

	if hr != 0 {
		return GattCommunicationStatusSuccess, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattDeviceServicesResult) GetServices() (*collections.IVectorView, error) {
	var out *collections.IVectorView
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetServices,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVectorView
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
