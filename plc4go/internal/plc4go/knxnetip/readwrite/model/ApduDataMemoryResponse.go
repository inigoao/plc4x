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
type ApduDataMemoryResponse struct {
	Address uint16
	Data    []uint8
	Parent  *ApduData
}

// The corresponding interface
type IApduDataMemoryResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataMemoryResponse) ApciType() uint8 {
	return 0x9
}

func (m *ApduDataMemoryResponse) InitializeParent(parent *ApduData) {
}

func NewApduDataMemoryResponse(address uint16, data []uint8) *ApduData {
	child := &ApduDataMemoryResponse{
		Address: address,
		Data:    data,
		Parent:  NewApduData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataMemoryResponse(structType interface{}) *ApduDataMemoryResponse {
	castFunc := func(typ interface{}) *ApduDataMemoryResponse {
		if casted, ok := typ.(ApduDataMemoryResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataMemoryResponse); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataMemoryResponse(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataMemoryResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataMemoryResponse) GetTypeName() string {
	return "ApduDataMemoryResponse"
}

func (m *ApduDataMemoryResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataMemoryResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Implicit Field (numBytes)
	lengthInBits += 6

	// Simple field (address)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataMemoryResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataMemoryResponseParse(readBuffer utils.ReadBuffer) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataMemoryResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (numBytes) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	numBytes, _numBytesErr := readBuffer.ReadUint8("numBytes", 6)
	_ = numBytes
	if _numBytesErr != nil {
		return nil, errors.Wrap(_numBytesErr, "Error parsing 'numBytes' field")
	}

	// Simple Field (address)
	address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	data := make([]uint8, numBytes)
	for curItem := uint16(0); curItem < uint16(numBytes); curItem++ {
		_item, _err := readBuffer.ReadUint8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataMemoryResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataMemoryResponse{
		Address: address,
		Data:    data,
		Parent:  &ApduData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataMemoryResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryResponse"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (numBytes) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numBytes := uint8(uint8(len(m.Data)))
		_numBytesErr := writeBuffer.WriteUint8("numBytes", 6, (numBytes))
		if _numBytesErr != nil {
			return errors.Wrap(_numBytesErr, "Error serializing 'numBytes' field")
		}

		// Simple Field (address)
		address := uint16(m.Address)
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := writeBuffer.WriteUint8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
