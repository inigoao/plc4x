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
type SysexCommandExendedId struct {
	Id     []int8
	Parent *SysexCommand
}

// The corresponding interface
type ISysexCommandExendedId interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SysexCommandExendedId) CommandType() uint8 {
	return 0x00
}

func (m *SysexCommandExendedId) Response() bool {
	return false
}

func (m *SysexCommandExendedId) InitializeParent(parent *SysexCommand) {
}

func NewSysexCommandExendedId(id []int8) *SysexCommand {
	child := &SysexCommandExendedId{
		Id:     id,
		Parent: NewSysexCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastSysexCommandExendedId(structType interface{}) *SysexCommandExendedId {
	castFunc := func(typ interface{}) *SysexCommandExendedId {
		if casted, ok := typ.(SysexCommandExendedId); ok {
			return &casted
		}
		if casted, ok := typ.(*SysexCommandExendedId); ok {
			return casted
		}
		if casted, ok := typ.(SysexCommand); ok {
			return CastSysexCommandExendedId(casted.Child)
		}
		if casted, ok := typ.(*SysexCommand); ok {
			return CastSysexCommandExendedId(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SysexCommandExendedId) GetTypeName() string {
	return "SysexCommandExendedId"
}

func (m *SysexCommandExendedId) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SysexCommandExendedId) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Array field
	if len(m.Id) > 0 {
		lengthInBits += 8 * uint16(len(m.Id))
	}

	return lengthInBits
}

func (m *SysexCommandExendedId) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SysexCommandExendedIdParse(readBuffer utils.ReadBuffer) (*SysexCommand, error) {
	if pullErr := readBuffer.PullContext("SysexCommandExendedId"); pullErr != nil {
		return nil, pullErr
	}

	// Array field (id)
	if pullErr := readBuffer.PullContext("id", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	id := make([]int8, uint16(2))
	for curItem := uint16(0); curItem < uint16(uint16(2)); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'id' field")
		}
		id[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("id", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("SysexCommandExendedId"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandExendedId{
		Id:     id,
		Parent: &SysexCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *SysexCommandExendedId) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandExendedId"); pushErr != nil {
			return pushErr
		}

		// Array Field (id)
		if m.Id != nil {
			if pushErr := writeBuffer.PushContext("id", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Id {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'id' field")
				}
			}
			if popErr := writeBuffer.PopContext("id", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("SysexCommandExendedId"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
