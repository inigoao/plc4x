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
type AdsReadStateResponse struct {
	Result      ReturnCode
	AdsState    uint16
	DeviceState uint16
	Parent      *AdsData
}

// The corresponding interface
type IAdsReadStateResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsReadStateResponse) CommandId() CommandId {
	return CommandId_ADS_READ_STATE
}

func (m *AdsReadStateResponse) Response() bool {
	return true
}

func (m *AdsReadStateResponse) InitializeParent(parent *AdsData) {
}

func NewAdsReadStateResponse(result ReturnCode, adsState uint16, deviceState uint16) *AdsData {
	child := &AdsReadStateResponse{
		Result:      result,
		AdsState:    adsState,
		DeviceState: deviceState,
		Parent:      NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsReadStateResponse(structType interface{}) *AdsReadStateResponse {
	castFunc := func(typ interface{}) *AdsReadStateResponse {
		if casted, ok := typ.(AdsReadStateResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsReadStateResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsReadStateResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsReadStateResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsReadStateResponse) GetTypeName() string {
	return "AdsReadStateResponse"
}

func (m *AdsReadStateResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsReadStateResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	// Simple field (adsState)
	lengthInBits += 16

	// Simple field (deviceState)
	lengthInBits += 16

	return lengthInBits
}

func (m *AdsReadStateResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsReadStateResponseParse(readBuffer utils.ReadBuffer) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsReadStateResponse"); pullErr != nil {
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

	// Simple Field (adsState)
	adsState, _adsStateErr := readBuffer.ReadUint16("adsState", 16)
	if _adsStateErr != nil {
		return nil, errors.Wrap(_adsStateErr, "Error parsing 'adsState' field")
	}

	// Simple Field (deviceState)
	deviceState, _deviceStateErr := readBuffer.ReadUint16("deviceState", 16)
	if _deviceStateErr != nil {
		return nil, errors.Wrap(_deviceStateErr, "Error parsing 'deviceState' field")
	}

	if closeErr := readBuffer.CloseContext("AdsReadStateResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsReadStateResponse{
		Result:      result,
		AdsState:    adsState,
		DeviceState: deviceState,
		Parent:      &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsReadStateResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadStateResponse"); pushErr != nil {
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

		// Simple Field (adsState)
		adsState := uint16(m.AdsState)
		_adsStateErr := writeBuffer.WriteUint16("adsState", 16, (adsState))
		if _adsStateErr != nil {
			return errors.Wrap(_adsStateErr, "Error serializing 'adsState' field")
		}

		// Simple Field (deviceState)
		deviceState := uint16(m.DeviceState)
		_deviceStateErr := writeBuffer.WriteUint16("deviceState", 16, (deviceState))
		if _deviceStateErr != nil {
			return errors.Wrap(_deviceStateErr, "Error serializing 'deviceState' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadStateResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
