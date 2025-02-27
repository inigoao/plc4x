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
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduControlNack struct {
	Parent *ApduControl
}

// The corresponding interface
type IApduControlNack interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduControlNack) ControlType() uint8 {
	return 0x3
}

func (m *ApduControlNack) InitializeParent(parent *ApduControl) {
}

func NewApduControlNack() *ApduControl {
	child := &ApduControlNack{
		Parent: NewApduControl(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduControlNack(structType interface{}) *ApduControlNack {
	castFunc := func(typ interface{}) *ApduControlNack {
		if casted, ok := typ.(ApduControlNack); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduControlNack); ok {
			return casted
		}
		if casted, ok := typ.(ApduControl); ok {
			return CastApduControlNack(casted.Child)
		}
		if casted, ok := typ.(*ApduControl); ok {
			return CastApduControlNack(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduControlNack) GetTypeName() string {
	return "ApduControlNack"
}

func (m *ApduControlNack) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduControlNack) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduControlNack) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduControlNackParse(readBuffer utils.ReadBuffer) (*ApduControl, error) {
	if pullErr := readBuffer.PullContext("ApduControlNack"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduControlNack"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduControlNack{
		Parent: &ApduControl{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduControlNack) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlNack"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduControlNack"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
