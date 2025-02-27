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
type S7Address struct {
	Child IS7AddressChild
}

// The corresponding interface
type IS7Address interface {
	AddressType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7AddressParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7Address, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7AddressChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7Address)
	GetTypeName() string
	IS7Address
}

func NewS7Address() *S7Address {
	return &S7Address{}
}

func CastS7Address(structType interface{}) *S7Address {
	castFunc := func(typ interface{}) *S7Address {
		if casted, ok := typ.(S7Address); ok {
			return &casted
		}
		if casted, ok := typ.(*S7Address); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7Address) GetTypeName() string {
	return "S7Address"
}

func (m *S7Address) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7Address) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *S7Address) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (addressType)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7Address) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7AddressParse(readBuffer utils.ReadBuffer) (*S7Address, error) {
	if pullErr := readBuffer.PullContext("S7Address"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (addressType) (Used as input to a switch field)
	addressType, _addressTypeErr := readBuffer.ReadUint8("addressType", 8)
	if _addressTypeErr != nil {
		return nil, errors.Wrap(_addressTypeErr, "Error parsing 'addressType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *S7Address
	var typeSwitchError error
	switch {
	case addressType == 0x10: // S7AddressAny
		_parent, typeSwitchError = S7AddressAnyParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7Address"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *S7Address) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7Address) SerializeParent(writeBuffer utils.WriteBuffer, child IS7Address, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("S7Address"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (addressType) (Used as input to a switch field)
	addressType := uint8(child.AddressType())
	_addressTypeErr := writeBuffer.WriteUint8("addressType", 8, (addressType))

	if _addressTypeErr != nil {
		return errors.Wrap(_addressTypeErr, "Error serializing 'addressType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Address"); popErr != nil {
		return popErr
	}
	return nil
}
