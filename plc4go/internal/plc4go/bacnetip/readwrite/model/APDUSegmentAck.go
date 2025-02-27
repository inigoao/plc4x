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
	"github.com/rs/zerolog/log"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type APDUSegmentAck struct {
	NegativeAck        bool
	Server             bool
	OriginalInvokeId   uint8
	SequenceNumber     uint8
	ProposedWindowSize uint8
	Parent             *APDU
}

// The corresponding interface
type IAPDUSegmentAck interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUSegmentAck) ApduType() uint8 {
	return 0x4
}

func (m *APDUSegmentAck) InitializeParent(parent *APDU) {
}

func NewAPDUSegmentAck(negativeAck bool, server bool, originalInvokeId uint8, sequenceNumber uint8, proposedWindowSize uint8) *APDU {
	child := &APDUSegmentAck{
		NegativeAck:        negativeAck,
		Server:             server,
		OriginalInvokeId:   originalInvokeId,
		SequenceNumber:     sequenceNumber,
		ProposedWindowSize: proposedWindowSize,
		Parent:             NewAPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAPDUSegmentAck(structType interface{}) *APDUSegmentAck {
	castFunc := func(typ interface{}) *APDUSegmentAck {
		if casted, ok := typ.(APDUSegmentAck); ok {
			return &casted
		}
		if casted, ok := typ.(*APDUSegmentAck); ok {
			return casted
		}
		if casted, ok := typ.(APDU); ok {
			return CastAPDUSegmentAck(casted.Child)
		}
		if casted, ok := typ.(*APDU); ok {
			return CastAPDUSegmentAck(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *APDUSegmentAck) GetTypeName() string {
	return "APDUSegmentAck"
}

func (m *APDUSegmentAck) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *APDUSegmentAck) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (negativeAck)
	lengthInBits += 1

	// Simple field (server)
	lengthInBits += 1

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	// Simple field (proposedWindowSize)
	lengthInBits += 8

	return lengthInBits
}

func (m *APDUSegmentAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func APDUSegmentAckParse(readBuffer utils.ReadBuffer) (*APDU, error) {
	if pullErr := readBuffer.PullContext("APDUSegmentAck"); pullErr != nil {
		return nil, pullErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (negativeAck)
	negativeAck, _negativeAckErr := readBuffer.ReadBit("negativeAck")
	if _negativeAckErr != nil {
		return nil, errors.Wrap(_negativeAckErr, "Error parsing 'negativeAck' field")
	}

	// Simple Field (server)
	server, _serverErr := readBuffer.ReadBit("server")
	if _serverErr != nil {
		return nil, errors.Wrap(_serverErr, "Error parsing 'server' field")
	}

	// Simple Field (originalInvokeId)
	originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}

	// Simple Field (sequenceNumber)
	sequenceNumber, _sequenceNumberErr := readBuffer.ReadUint8("sequenceNumber", 8)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field")
	}

	// Simple Field (proposedWindowSize)
	proposedWindowSize, _proposedWindowSizeErr := readBuffer.ReadUint8("proposedWindowSize", 8)
	if _proposedWindowSizeErr != nil {
		return nil, errors.Wrap(_proposedWindowSizeErr, "Error parsing 'proposedWindowSize' field")
	}

	if closeErr := readBuffer.CloseContext("APDUSegmentAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &APDUSegmentAck{
		NegativeAck:        negativeAck,
		Server:             server,
		OriginalInvokeId:   originalInvokeId,
		SequenceNumber:     sequenceNumber,
		ProposedWindowSize: proposedWindowSize,
		Parent:             &APDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *APDUSegmentAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUSegmentAck"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 2, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (negativeAck)
		negativeAck := bool(m.NegativeAck)
		_negativeAckErr := writeBuffer.WriteBit("negativeAck", (negativeAck))
		if _negativeAckErr != nil {
			return errors.Wrap(_negativeAckErr, "Error serializing 'negativeAck' field")
		}

		// Simple Field (server)
		server := bool(m.Server)
		_serverErr := writeBuffer.WriteBit("server", (server))
		if _serverErr != nil {
			return errors.Wrap(_serverErr, "Error serializing 'server' field")
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint8(m.SequenceNumber)
		_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 8, (sequenceNumber))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		// Simple Field (proposedWindowSize)
		proposedWindowSize := uint8(m.ProposedWindowSize)
		_proposedWindowSizeErr := writeBuffer.WriteUint8("proposedWindowSize", 8, (proposedWindowSize))
		if _proposedWindowSizeErr != nil {
			return errors.Wrap(_proposedWindowSizeErr, "Error serializing 'proposedWindowSize' field")
		}

		if popErr := writeBuffer.PopContext("APDUSegmentAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
