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
type EipPacket struct {
	SessionHandle uint32
	Status        uint32
	SenderContext []uint8
	Options       uint32
	Child         IEipPacketChild
}

// The corresponding interface
type IEipPacket interface {
	Command() uint16
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IEipPacketParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IEipPacket, serializeChildFunction func() error) error
	GetTypeName() string
}

type IEipPacketChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *EipPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32)
	GetTypeName() string
	IEipPacket
}

func NewEipPacket(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *EipPacket {
	return &EipPacket{SessionHandle: sessionHandle, Status: status, SenderContext: senderContext, Options: options}
}

func CastEipPacket(structType interface{}) *EipPacket {
	castFunc := func(typ interface{}) *EipPacket {
		if casted, ok := typ.(EipPacket); ok {
			return &casted
		}
		if casted, ok := typ.(*EipPacket); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *EipPacket) GetTypeName() string {
	return "EipPacket"
}

func (m *EipPacket) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *EipPacket) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *EipPacket) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (command)
	lengthInBits += 16

	// Implicit Field (len)
	lengthInBits += 16

	// Simple field (sessionHandle)
	lengthInBits += 32

	// Simple field (status)
	lengthInBits += 32

	// Array field
	if len(m.SenderContext) > 0 {
		lengthInBits += 8 * uint16(len(m.SenderContext))
	}

	// Simple field (options)
	lengthInBits += 32

	return lengthInBits
}

func (m *EipPacket) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func EipPacketParse(readBuffer utils.ReadBuffer) (*EipPacket, error) {
	if pullErr := readBuffer.PullContext("EipPacket"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (command) (Used as input to a switch field)
	command, _commandErr := readBuffer.ReadUint16("command", 16)
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field")
	}

	// Implicit Field (len) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	len, _lenErr := readBuffer.ReadUint16("len", 16)
	_ = len
	if _lenErr != nil {
		return nil, errors.Wrap(_lenErr, "Error parsing 'len' field")
	}

	// Simple Field (sessionHandle)
	sessionHandle, _sessionHandleErr := readBuffer.ReadUint32("sessionHandle", 32)
	if _sessionHandleErr != nil {
		return nil, errors.Wrap(_sessionHandleErr, "Error parsing 'sessionHandle' field")
	}

	// Simple Field (status)
	status, _statusErr := readBuffer.ReadUint32("status", 32)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}

	// Array field (senderContext)
	if pullErr := readBuffer.PullContext("senderContext", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	senderContext := make([]uint8, uint16(8))
	for curItem := uint16(0); curItem < uint16(uint16(8)); curItem++ {
		_item, _err := readBuffer.ReadUint8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'senderContext' field")
		}
		senderContext[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("senderContext", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (options)
	options, _optionsErr := readBuffer.ReadUint32("options", 32)
	if _optionsErr != nil {
		return nil, errors.Wrap(_optionsErr, "Error parsing 'options' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *EipPacket
	var typeSwitchError error
	switch {
	case command == 0x0065: // EipConnectionRequest
		_parent, typeSwitchError = EipConnectionRequestParse(readBuffer)
	case command == 0x0066: // EipDisconnectRequest
		_parent, typeSwitchError = EipDisconnectRequestParse(readBuffer)
	case command == 0x006F: // CipRRData
		_parent, typeSwitchError = CipRRDataParse(readBuffer, len)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("EipPacket"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, sessionHandle, status, senderContext, options)
	return _parent, nil
}

func (m *EipPacket) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *EipPacket) SerializeParent(writeBuffer utils.WriteBuffer, child IEipPacket, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("EipPacket"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (command) (Used as input to a switch field)
	command := uint16(child.Command())
	_commandErr := writeBuffer.WriteUint16("command", 16, (command))

	if _commandErr != nil {
		return errors.Wrap(_commandErr, "Error serializing 'command' field")
	}

	// Implicit Field (len) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	len := uint16(uint16(uint16(m.LengthInBytes())) - uint16(uint16(24)))
	_lenErr := writeBuffer.WriteUint16("len", 16, (len))
	if _lenErr != nil {
		return errors.Wrap(_lenErr, "Error serializing 'len' field")
	}

	// Simple Field (sessionHandle)
	sessionHandle := uint32(m.SessionHandle)
	_sessionHandleErr := writeBuffer.WriteUint32("sessionHandle", 32, (sessionHandle))
	if _sessionHandleErr != nil {
		return errors.Wrap(_sessionHandleErr, "Error serializing 'sessionHandle' field")
	}

	// Simple Field (status)
	status := uint32(m.Status)
	_statusErr := writeBuffer.WriteUint32("status", 32, (status))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	// Array Field (senderContext)
	if m.SenderContext != nil {
		if pushErr := writeBuffer.PushContext("senderContext", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.SenderContext {
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'senderContext' field")
			}
		}
		if popErr := writeBuffer.PopContext("senderContext", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	// Simple Field (options)
	options := uint32(m.Options)
	_optionsErr := writeBuffer.WriteUint32("options", 32, (options))
	if _optionsErr != nil {
		return errors.Wrap(_optionsErr, "Error serializing 'options' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("EipPacket"); popErr != nil {
		return popErr
	}
	return nil
}
