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
type KnxNetIpCore struct {
	Version uint8
	Parent  *ServiceId
}

// The corresponding interface
type IKnxNetIpCore interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *KnxNetIpCore) ServiceType() uint8 {
	return 0x02
}

func (m *KnxNetIpCore) InitializeParent(parent *ServiceId) {
}

func NewKnxNetIpCore(version uint8) *ServiceId {
	child := &KnxNetIpCore{
		Version: version,
		Parent:  NewServiceId(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastKnxNetIpCore(structType interface{}) *KnxNetIpCore {
	castFunc := func(typ interface{}) *KnxNetIpCore {
		if casted, ok := typ.(KnxNetIpCore); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxNetIpCore); ok {
			return casted
		}
		if casted, ok := typ.(ServiceId); ok {
			return CastKnxNetIpCore(casted.Child)
		}
		if casted, ok := typ.(*ServiceId); ok {
			return CastKnxNetIpCore(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxNetIpCore) GetTypeName() string {
	return "KnxNetIpCore"
}

func (m *KnxNetIpCore) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *KnxNetIpCore) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *KnxNetIpCore) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxNetIpCoreParse(readBuffer utils.ReadBuffer) (*ServiceId, error) {
	if pullErr := readBuffer.PullContext("KnxNetIpCore"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (version)
	version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}

	if closeErr := readBuffer.CloseContext("KnxNetIpCore"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxNetIpCore{
		Version: version,
		Parent:  &ServiceId{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *KnxNetIpCore) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpCore"); pushErr != nil {
			return pushErr
		}

		// Simple Field (version)
		version := uint8(m.Version)
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpCore"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
