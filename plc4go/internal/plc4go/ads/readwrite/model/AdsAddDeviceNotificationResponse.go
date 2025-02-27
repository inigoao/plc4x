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
type AdsAddDeviceNotificationResponse struct {
	Result             ReturnCode
	NotificationHandle uint32
	Parent             *AdsData
}

// The corresponding interface
type IAdsAddDeviceNotificationResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsAddDeviceNotificationResponse) CommandId() CommandId {
	return CommandId_ADS_ADD_DEVICE_NOTIFICATION
}

func (m *AdsAddDeviceNotificationResponse) Response() bool {
	return true
}

func (m *AdsAddDeviceNotificationResponse) InitializeParent(parent *AdsData) {
}

func NewAdsAddDeviceNotificationResponse(result ReturnCode, notificationHandle uint32) *AdsData {
	child := &AdsAddDeviceNotificationResponse{
		Result:             result,
		NotificationHandle: notificationHandle,
		Parent:             NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsAddDeviceNotificationResponse(structType interface{}) *AdsAddDeviceNotificationResponse {
	castFunc := func(typ interface{}) *AdsAddDeviceNotificationResponse {
		if casted, ok := typ.(AdsAddDeviceNotificationResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsAddDeviceNotificationResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsAddDeviceNotificationResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsAddDeviceNotificationResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsAddDeviceNotificationResponse) GetTypeName() string {
	return "AdsAddDeviceNotificationResponse"
}

func (m *AdsAddDeviceNotificationResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsAddDeviceNotificationResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	// Simple field (notificationHandle)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsAddDeviceNotificationResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsAddDeviceNotificationResponseParse(readBuffer utils.ReadBuffer) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsAddDeviceNotificationResponse"); pullErr != nil {
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

	// Simple Field (notificationHandle)
	notificationHandle, _notificationHandleErr := readBuffer.ReadUint32("notificationHandle", 32)
	if _notificationHandleErr != nil {
		return nil, errors.Wrap(_notificationHandleErr, "Error parsing 'notificationHandle' field")
	}

	if closeErr := readBuffer.CloseContext("AdsAddDeviceNotificationResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsAddDeviceNotificationResponse{
		Result:             result,
		NotificationHandle: notificationHandle,
		Parent:             &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsAddDeviceNotificationResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsAddDeviceNotificationResponse"); pushErr != nil {
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

		// Simple Field (notificationHandle)
		notificationHandle := uint32(m.NotificationHandle)
		_notificationHandleErr := writeBuffer.WriteUint32("notificationHandle", 32, (notificationHandle))
		if _notificationHandleErr != nil {
			return errors.Wrap(_notificationHandleErr, "Error serializing 'notificationHandle' field")
		}

		if popErr := writeBuffer.PopContext("AdsAddDeviceNotificationResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
