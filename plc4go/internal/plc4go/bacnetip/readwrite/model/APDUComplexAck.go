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
type APDUComplexAck struct {
	SegmentedMessage   bool
	MoreFollows        bool
	OriginalInvokeId   uint8
	SequenceNumber     *uint8
	ProposedWindowSize *uint8
	ServiceAck         *BACnetServiceAck
	Parent             *APDU
}

// The corresponding interface
type IAPDUComplexAck interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUComplexAck) ApduType() uint8 {
	return 0x3
}

func (m *APDUComplexAck) InitializeParent(parent *APDU) {
}

func NewAPDUComplexAck(segmentedMessage bool, moreFollows bool, originalInvokeId uint8, sequenceNumber *uint8, proposedWindowSize *uint8, serviceAck *BACnetServiceAck) *APDU {
	child := &APDUComplexAck{
		SegmentedMessage:   segmentedMessage,
		MoreFollows:        moreFollows,
		OriginalInvokeId:   originalInvokeId,
		SequenceNumber:     sequenceNumber,
		ProposedWindowSize: proposedWindowSize,
		ServiceAck:         serviceAck,
		Parent:             NewAPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAPDUComplexAck(structType interface{}) *APDUComplexAck {
	castFunc := func(typ interface{}) *APDUComplexAck {
		if casted, ok := typ.(APDUComplexAck); ok {
			return &casted
		}
		if casted, ok := typ.(*APDUComplexAck); ok {
			return casted
		}
		if casted, ok := typ.(APDU); ok {
			return CastAPDUComplexAck(casted.Child)
		}
		if casted, ok := typ.(*APDU); ok {
			return CastAPDUComplexAck(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *APDUComplexAck) GetTypeName() string {
	return "APDUComplexAck"
}

func (m *APDUComplexAck) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *APDUComplexAck) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (segmentedMessage)
	lengthInBits += 1

	// Simple field (moreFollows)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Optional Field (sequenceNumber)
	if m.SequenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (proposedWindowSize)
	if m.ProposedWindowSize != nil {
		lengthInBits += 8
	}

	// Simple field (serviceAck)
	lengthInBits += m.ServiceAck.LengthInBits()

	return lengthInBits
}

func (m *APDUComplexAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func APDUComplexAckParse(readBuffer utils.ReadBuffer) (*APDU, error) {
	if pullErr := readBuffer.PullContext("APDUComplexAck"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (segmentedMessage)
	segmentedMessage, _segmentedMessageErr := readBuffer.ReadBit("segmentedMessage")
	if _segmentedMessageErr != nil {
		return nil, errors.Wrap(_segmentedMessageErr, "Error parsing 'segmentedMessage' field")
	}

	// Simple Field (moreFollows)
	moreFollows, _moreFollowsErr := readBuffer.ReadBit("moreFollows")
	if _moreFollowsErr != nil {
		return nil, errors.Wrap(_moreFollowsErr, "Error parsing 'moreFollows' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (originalInvokeId)
	originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}

	// Optional Field (sequenceNumber) (Can be skipped, if a given expression evaluates to false)
	var sequenceNumber *uint8 = nil
	if segmentedMessage {
		_val, _err := readBuffer.ReadUint8("sequenceNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'sequenceNumber' field")
		}
		sequenceNumber = &_val
	}

	// Optional Field (proposedWindowSize) (Can be skipped, if a given expression evaluates to false)
	var proposedWindowSize *uint8 = nil
	if segmentedMessage {
		_val, _err := readBuffer.ReadUint8("proposedWindowSize", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'proposedWindowSize' field")
		}
		proposedWindowSize = &_val
	}

	if pullErr := readBuffer.PullContext("serviceAck"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (serviceAck)
	serviceAck, _serviceAckErr := BACnetServiceAckParse(readBuffer)
	if _serviceAckErr != nil {
		return nil, errors.Wrap(_serviceAckErr, "Error parsing 'serviceAck' field")
	}
	if closeErr := readBuffer.CloseContext("serviceAck"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("APDUComplexAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &APDUComplexAck{
		SegmentedMessage:   segmentedMessage,
		MoreFollows:        moreFollows,
		OriginalInvokeId:   originalInvokeId,
		SequenceNumber:     sequenceNumber,
		ProposedWindowSize: proposedWindowSize,
		ServiceAck:         serviceAck,
		Parent:             &APDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *APDUComplexAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUComplexAck"); pushErr != nil {
			return pushErr
		}

		// Simple Field (segmentedMessage)
		segmentedMessage := bool(m.SegmentedMessage)
		_segmentedMessageErr := writeBuffer.WriteBit("segmentedMessage", (segmentedMessage))
		if _segmentedMessageErr != nil {
			return errors.Wrap(_segmentedMessageErr, "Error serializing 'segmentedMessage' field")
		}

		// Simple Field (moreFollows)
		moreFollows := bool(m.MoreFollows)
		_moreFollowsErr := writeBuffer.WriteBit("moreFollows", (moreFollows))
		if _moreFollowsErr != nil {
			return errors.Wrap(_moreFollowsErr, "Error serializing 'moreFollows' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 2, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Optional Field (sequenceNumber) (Can be skipped, if the value is null)
		var sequenceNumber *uint8 = nil
		if m.SequenceNumber != nil {
			sequenceNumber = m.SequenceNumber
			_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 8, *(sequenceNumber))
			if _sequenceNumberErr != nil {
				return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
			}
		}

		// Optional Field (proposedWindowSize) (Can be skipped, if the value is null)
		var proposedWindowSize *uint8 = nil
		if m.ProposedWindowSize != nil {
			proposedWindowSize = m.ProposedWindowSize
			_proposedWindowSizeErr := writeBuffer.WriteUint8("proposedWindowSize", 8, *(proposedWindowSize))
			if _proposedWindowSizeErr != nil {
				return errors.Wrap(_proposedWindowSizeErr, "Error serializing 'proposedWindowSize' field")
			}
		}

		// Simple Field (serviceAck)
		if pushErr := writeBuffer.PushContext("serviceAck"); pushErr != nil {
			return pushErr
		}
		_serviceAckErr := m.ServiceAck.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("serviceAck"); popErr != nil {
			return popErr
		}
		if _serviceAckErr != nil {
			return errors.Wrap(_serviceAckErr, "Error serializing 'serviceAck' field")
		}

		if popErr := writeBuffer.PopContext("APDUComplexAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
