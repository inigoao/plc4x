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
type ServiceId struct {
	Child IServiceIdChild
}

// The corresponding interface
type IServiceId interface {
	ServiceType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IServiceIdParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IServiceId, serializeChildFunction func() error) error
	GetTypeName() string
}

type IServiceIdChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *ServiceId)
	GetTypeName() string
	IServiceId
}

func NewServiceId() *ServiceId {
	return &ServiceId{}
}

func CastServiceId(structType interface{}) *ServiceId {
	castFunc := func(typ interface{}) *ServiceId {
		if casted, ok := typ.(ServiceId); ok {
			return &casted
		}
		if casted, ok := typ.(*ServiceId); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ServiceId) GetTypeName() string {
	return "ServiceId"
}

func (m *ServiceId) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ServiceId) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *ServiceId) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceType)
	lengthInBits += 8

	return lengthInBits
}

func (m *ServiceId) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ServiceIdParse(readBuffer utils.ReadBuffer) (*ServiceId, error) {
	if pullErr := readBuffer.PullContext("ServiceId"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (serviceType) (Used as input to a switch field)
	serviceType, _serviceTypeErr := readBuffer.ReadUint8("serviceType", 8)
	if _serviceTypeErr != nil {
		return nil, errors.Wrap(_serviceTypeErr, "Error parsing 'serviceType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ServiceId
	var typeSwitchError error
	switch {
	case serviceType == 0x02: // KnxNetIpCore
		_parent, typeSwitchError = KnxNetIpCoreParse(readBuffer)
	case serviceType == 0x03: // KnxNetIpDeviceManagement
		_parent, typeSwitchError = KnxNetIpDeviceManagementParse(readBuffer)
	case serviceType == 0x04: // KnxNetIpTunneling
		_parent, typeSwitchError = KnxNetIpTunnelingParse(readBuffer)
	case serviceType == 0x05: // KnxNetIpRouting
		_parent, typeSwitchError = KnxNetIpRoutingParse(readBuffer)
	case serviceType == 0x06: // KnxNetRemoteLogging
		_parent, typeSwitchError = KnxNetRemoteLoggingParse(readBuffer)
	case serviceType == 0x07: // KnxNetRemoteConfigurationAndDiagnosis
		_parent, typeSwitchError = KnxNetRemoteConfigurationAndDiagnosisParse(readBuffer)
	case serviceType == 0x08: // KnxNetObjectServer
		_parent, typeSwitchError = KnxNetObjectServerParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("ServiceId"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ServiceId) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *ServiceId) SerializeParent(writeBuffer utils.WriteBuffer, child IServiceId, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("ServiceId"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (serviceType) (Used as input to a switch field)
	serviceType := uint8(child.ServiceType())
	_serviceTypeErr := writeBuffer.WriteUint8("serviceType", 8, (serviceType))

	if _serviceTypeErr != nil {
		return errors.Wrap(_serviceTypeErr, "Error serializing 'serviceType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ServiceId"); popErr != nil {
		return popErr
	}
	return nil
}
