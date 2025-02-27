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
type FirmataCommandProtocolVersion struct {
	MajorVersion uint8
	MinorVersion uint8
	Parent       *FirmataCommand
}

// The corresponding interface
type IFirmataCommandProtocolVersion interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *FirmataCommandProtocolVersion) CommandCode() uint8 {
	return 0x9
}

func (m *FirmataCommandProtocolVersion) InitializeParent(parent *FirmataCommand) {
}

func NewFirmataCommandProtocolVersion(majorVersion uint8, minorVersion uint8) *FirmataCommand {
	child := &FirmataCommandProtocolVersion{
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Parent:       NewFirmataCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastFirmataCommandProtocolVersion(structType interface{}) *FirmataCommandProtocolVersion {
	castFunc := func(typ interface{}) *FirmataCommandProtocolVersion {
		if casted, ok := typ.(FirmataCommandProtocolVersion); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataCommandProtocolVersion); ok {
			return casted
		}
		if casted, ok := typ.(FirmataCommand); ok {
			return CastFirmataCommandProtocolVersion(casted.Child)
		}
		if casted, ok := typ.(*FirmataCommand); ok {
			return CastFirmataCommandProtocolVersion(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataCommandProtocolVersion) GetTypeName() string {
	return "FirmataCommandProtocolVersion"
}

func (m *FirmataCommandProtocolVersion) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *FirmataCommandProtocolVersion) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	return lengthInBits
}

func (m *FirmataCommandProtocolVersion) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func FirmataCommandProtocolVersionParse(readBuffer utils.ReadBuffer) (*FirmataCommand, error) {
	if pullErr := readBuffer.PullContext("FirmataCommandProtocolVersion"); pullErr != nil {
		return nil, pullErr
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

	if closeErr := readBuffer.CloseContext("FirmataCommandProtocolVersion"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataCommandProtocolVersion{
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Parent:       &FirmataCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *FirmataCommandProtocolVersion) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandProtocolVersion"); pushErr != nil {
			return pushErr
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

		if popErr := writeBuffer.PopContext("FirmataCommandProtocolVersion"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
