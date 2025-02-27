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
type ModbusPDUReadWriteMultipleHoldingRegistersRequest struct {
	ReadStartingAddress  uint16
	ReadQuantity         uint16
	WriteStartingAddress uint16
	WriteQuantity        uint16
	Value                []int8
	Parent               *ModbusPDU
}

// The corresponding interface
type IModbusPDUReadWriteMultipleHoldingRegistersRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) ErrorFlag() bool {
	return false
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) FunctionFlag() uint8 {
	return 0x17
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) Response() bool {
	return false
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUReadWriteMultipleHoldingRegistersRequest(readStartingAddress uint16, readQuantity uint16, writeStartingAddress uint16, writeQuantity uint16, value []int8) *ModbusPDU {
	child := &ModbusPDUReadWriteMultipleHoldingRegistersRequest{
		ReadStartingAddress:  readStartingAddress,
		ReadQuantity:         readQuantity,
		WriteStartingAddress: writeStartingAddress,
		WriteQuantity:        writeQuantity,
		Value:                value,
		Parent:               NewModbusPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastModbusPDUReadWriteMultipleHoldingRegistersRequest(structType interface{}) *ModbusPDUReadWriteMultipleHoldingRegistersRequest {
	castFunc := func(typ interface{}) *ModbusPDUReadWriteMultipleHoldingRegistersRequest {
		if casted, ok := typ.(ModbusPDUReadWriteMultipleHoldingRegistersRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadWriteMultipleHoldingRegistersRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadWriteMultipleHoldingRegistersRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadWriteMultipleHoldingRegistersRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetTypeName() string {
	return "ModbusPDUReadWriteMultipleHoldingRegistersRequest"
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (readStartingAddress)
	lengthInBits += 16

	// Simple field (readQuantity)
	lengthInBits += 16

	// Simple field (writeStartingAddress)
	lengthInBits += 16

	// Simple field (writeQuantity)
	lengthInBits += 16

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusPDUReadWriteMultipleHoldingRegistersRequestParse(readBuffer utils.ReadBuffer) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (readStartingAddress)
	readStartingAddress, _readStartingAddressErr := readBuffer.ReadUint16("readStartingAddress", 16)
	if _readStartingAddressErr != nil {
		return nil, errors.Wrap(_readStartingAddressErr, "Error parsing 'readStartingAddress' field")
	}

	// Simple Field (readQuantity)
	readQuantity, _readQuantityErr := readBuffer.ReadUint16("readQuantity", 16)
	if _readQuantityErr != nil {
		return nil, errors.Wrap(_readQuantityErr, "Error parsing 'readQuantity' field")
	}

	// Simple Field (writeStartingAddress)
	writeStartingAddress, _writeStartingAddressErr := readBuffer.ReadUint16("writeStartingAddress", 16)
	if _writeStartingAddressErr != nil {
		return nil, errors.Wrap(_writeStartingAddressErr, "Error parsing 'writeStartingAddress' field")
	}

	// Simple Field (writeQuantity)
	writeQuantity, _writeQuantityErr := readBuffer.ReadUint16("writeQuantity", 16)
	if _writeQuantityErr != nil {
		return nil, errors.Wrap(_writeQuantityErr, "Error parsing 'writeQuantity' field")
	}

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Array field (value)
	if pullErr := readBuffer.PullContext("value", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	value := make([]int8, byteCount)
	for curItem := uint16(0); curItem < uint16(byteCount); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'value' field")
		}
		value[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("value", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadWriteMultipleHoldingRegistersRequest{
		ReadStartingAddress:  readStartingAddress,
		ReadQuantity:         readQuantity,
		WriteStartingAddress: writeStartingAddress,
		WriteQuantity:        writeQuantity,
		Value:                value,
		Parent:               &ModbusPDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ModbusPDUReadWriteMultipleHoldingRegistersRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (readStartingAddress)
		readStartingAddress := uint16(m.ReadStartingAddress)
		_readStartingAddressErr := writeBuffer.WriteUint16("readStartingAddress", 16, (readStartingAddress))
		if _readStartingAddressErr != nil {
			return errors.Wrap(_readStartingAddressErr, "Error serializing 'readStartingAddress' field")
		}

		// Simple Field (readQuantity)
		readQuantity := uint16(m.ReadQuantity)
		_readQuantityErr := writeBuffer.WriteUint16("readQuantity", 16, (readQuantity))
		if _readQuantityErr != nil {
			return errors.Wrap(_readQuantityErr, "Error serializing 'readQuantity' field")
		}

		// Simple Field (writeStartingAddress)
		writeStartingAddress := uint16(m.WriteStartingAddress)
		_writeStartingAddressErr := writeBuffer.WriteUint16("writeStartingAddress", 16, (writeStartingAddress))
		if _writeStartingAddressErr != nil {
			return errors.Wrap(_writeStartingAddressErr, "Error serializing 'writeStartingAddress' field")
		}

		// Simple Field (writeQuantity)
		writeQuantity := uint16(m.WriteQuantity)
		_writeQuantityErr := writeBuffer.WriteUint16("writeQuantity", 16, (writeQuantity))
		if _writeQuantityErr != nil {
			return errors.Wrap(_writeQuantityErr, "Error serializing 'writeQuantity' field")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(len(m.Value)))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (value)
		if m.Value != nil {
			if pushErr := writeBuffer.PushContext("value", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Value {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'value' field")
				}
			}
			if popErr := writeBuffer.PopContext("value", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
