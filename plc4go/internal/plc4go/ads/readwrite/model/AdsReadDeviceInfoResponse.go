//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AdsReadDeviceInfoResponse struct {
	Result       ReturnCode
	MajorVersion uint8
	MinorVersion uint8
	Version      uint16
	Device       []int8
	Parent       *AdsData
}

// The corresponding interface
type IAdsReadDeviceInfoResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsReadDeviceInfoResponse) CommandId() CommandId {
	return CommandId_ADS_READ_DEVICE_INFO
}

func (m *AdsReadDeviceInfoResponse) Response() bool {
	return true
}

func (m *AdsReadDeviceInfoResponse) InitializeParent(parent *AdsData) {
}

func NewAdsReadDeviceInfoResponse(result ReturnCode, majorVersion uint8, minorVersion uint8, version uint16, device []int8) *AdsData {
	child := &AdsReadDeviceInfoResponse{
		Result:       result,
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Version:      version,
		Device:       device,
		Parent:       NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsReadDeviceInfoResponse(structType interface{}) *AdsReadDeviceInfoResponse {
	castFunc := func(typ interface{}) *AdsReadDeviceInfoResponse {
		if casted, ok := typ.(AdsReadDeviceInfoResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsReadDeviceInfoResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsReadDeviceInfoResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsReadDeviceInfoResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsReadDeviceInfoResponse) GetTypeName() string {
	return "AdsReadDeviceInfoResponse"
}

func (m *AdsReadDeviceInfoResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsReadDeviceInfoResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	// Simple field (version)
	lengthInBits += 16

	// Array field
	if len(m.Device) > 0 {
		lengthInBits += 8 * uint16(len(m.Device))
	}

	return lengthInBits
}

func (m *AdsReadDeviceInfoResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsReadDeviceInfoResponseParse(readBuffer utils.ReadBuffer) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsReadDeviceInfoResponse"); pullErr != nil {
		return nil, pullErr
	}

	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (result)
	result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (majorVersion)
	majorVersion, _majorVersionErr := readBuffer.ReadUint8("majorVersion", 8)
	if _majorVersionErr != nil {
		return nil, errors.Wrap(_majorVersionErr, "Error parsing 'majorVersion' field")
	}

	// Simple Field (minorVersion)
	minorVersion, _minorVersionErr := readBuffer.ReadUint8("minorVersion", 8)
	if _minorVersionErr != nil {
		return nil, errors.Wrap(_minorVersionErr, "Error parsing 'minorVersion' field")
	}

	// Simple Field (version)
	version, _versionErr := readBuffer.ReadUint16("version", 16)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}

	// Array field (device)
	if pullErr := readBuffer.PullContext("device", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	device := make([]int8, uint16(16))
	for curItem := uint16(0); curItem < uint16(uint16(16)); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'device' field")
		}
		device[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("device", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AdsReadDeviceInfoResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsReadDeviceInfoResponse{
		Result:       result,
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Version:      version,
		Device:       device,
		Parent:       &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsReadDeviceInfoResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadDeviceInfoResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return pushErr
		}
		_resultErr := m.Result.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return popErr
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (majorVersion)
		majorVersion := uint8(m.MajorVersion)
		_majorVersionErr := writeBuffer.WriteUint8("majorVersion", 8, (majorVersion))
		if _majorVersionErr != nil {
			return errors.Wrap(_majorVersionErr, "Error serializing 'majorVersion' field")
		}

		// Simple Field (minorVersion)
		minorVersion := uint8(m.MinorVersion)
		_minorVersionErr := writeBuffer.WriteUint8("minorVersion", 8, (minorVersion))
		if _minorVersionErr != nil {
			return errors.Wrap(_minorVersionErr, "Error serializing 'minorVersion' field")
		}

		// Simple Field (version)
		version := uint16(m.Version)
		_versionErr := writeBuffer.WriteUint16("version", 16, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		// Array Field (device)
		if m.Device != nil {
			if pushErr := writeBuffer.PushContext("device", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Device {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'device' field")
				}
			}
			if popErr := writeBuffer.PopContext("device", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("AdsReadDeviceInfoResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
