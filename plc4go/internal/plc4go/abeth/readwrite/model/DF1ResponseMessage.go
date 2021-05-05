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
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type DF1ResponseMessage struct {
	DestinationAddress uint8
	SourceAddress      uint8
	Status             uint8
	TransactionCounter uint16
	Child              IDF1ResponseMessageChild
}

// The corresponding interface
type IDF1ResponseMessage interface {
	CommandCode() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IDF1ResponseMessageParent interface {
	SerializeParent(io utils.WriteBuffer, child IDF1ResponseMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type IDF1ResponseMessageChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *DF1ResponseMessage, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16)
	GetTypeName() string
	IDF1ResponseMessage
	utils.AsciiBoxer
}

func NewDF1ResponseMessage(destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) *DF1ResponseMessage {
	return &DF1ResponseMessage{DestinationAddress: destinationAddress, SourceAddress: sourceAddress, Status: status, TransactionCounter: transactionCounter}
}

func CastDF1ResponseMessage(structType interface{}) *DF1ResponseMessage {
	castFunc := func(typ interface{}) *DF1ResponseMessage {
		if casted, ok := typ.(DF1ResponseMessage); ok {
			return &casted
		}
		if casted, ok := typ.(*DF1ResponseMessage); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DF1ResponseMessage) GetTypeName() string {
	return "DF1ResponseMessage"
}

func (m *DF1ResponseMessage) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DF1ResponseMessage) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *DF1ResponseMessage) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (destinationAddress)
	lengthInBits += 8

	// Simple field (sourceAddress)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8
	// Discriminator Field (commandCode)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (transactionCounter)
	lengthInBits += 16

	return lengthInBits
}

func (m *DF1ResponseMessage) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DF1ResponseMessageParse(io utils.ReadBuffer, payloadLength uint16) (*DF1ResponseMessage, error) {
	if pullErr := io.PullContext("DF1ResponseMessage"); pullErr != nil {
		return nil, pullErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8("reserved", 8)
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

	// Simple Field (destinationAddress)
	destinationAddress, _destinationAddressErr := io.ReadUint8("destinationAddress", 8)
	if _destinationAddressErr != nil {
		return nil, errors.Wrap(_destinationAddressErr, "Error parsing 'destinationAddress' field")
	}

	// Simple Field (sourceAddress)
	sourceAddress, _sourceAddressErr := io.ReadUint8("sourceAddress", 8)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint8("reserved", 8)
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

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode, _commandCodeErr := io.ReadUint8("commandCode", 8)
	if _commandCodeErr != nil {
		return nil, errors.Wrap(_commandCodeErr, "Error parsing 'commandCode' field")
	}

	// Simple Field (status)
	status, _statusErr := io.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}

	// Simple Field (transactionCounter)
	transactionCounter, _transactionCounterErr := io.ReadUint16("transactionCounter", 16)
	if _transactionCounterErr != nil {
		return nil, errors.Wrap(_transactionCounterErr, "Error parsing 'transactionCounter' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *DF1ResponseMessage
	var typeSwitchError error
	switch {
	case commandCode == 0x4F: // DF1CommandResponseMessageProtectedTypedLogicalRead
		_parent, typeSwitchError = DF1CommandResponseMessageProtectedTypedLogicalReadParse(io, payloadLength, status)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := io.CloseContext("DF1ResponseMessage"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, destinationAddress, sourceAddress, status, transactionCounter)
	return _parent, nil
}

func (m *DF1ResponseMessage) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *DF1ResponseMessage) SerializeParent(io utils.WriteBuffer, child IDF1ResponseMessage, serializeChildFunction func() error) error {
	if pushErr := io.PushContext("DF1ResponseMessage"); pushErr != nil {
		return pushErr
	}

	// Reserved Field (reserved)
	{
		_err := io.WriteUint8("reserved", 8, uint8(0x00))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (destinationAddress)
	destinationAddress := uint8(m.DestinationAddress)
	_destinationAddressErr := io.WriteUint8("destinationAddress", 8, (destinationAddress))
	if _destinationAddressErr != nil {
		return errors.Wrap(_destinationAddressErr, "Error serializing 'destinationAddress' field")
	}

	// Simple Field (sourceAddress)
	sourceAddress := uint8(m.SourceAddress)
	_sourceAddressErr := io.WriteUint8("sourceAddress", 8, (sourceAddress))
	if _sourceAddressErr != nil {
		return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
	}

	// Reserved Field (reserved)
	{
		_err := io.WriteUint8("reserved", 8, uint8(0x00))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode := uint8(child.CommandCode())
	_commandCodeErr := io.WriteUint8("commandCode", 8, (commandCode))

	if _commandCodeErr != nil {
		return errors.Wrap(_commandCodeErr, "Error serializing 'commandCode' field")
	}

	// Simple Field (status)
	status := uint8(m.Status)
	_statusErr := io.WriteUint8("status", 8, (status))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	// Simple Field (transactionCounter)
	transactionCounter := uint16(m.TransactionCounter)
	_transactionCounterErr := io.WriteUint16("transactionCounter", 16, (transactionCounter))
	if _transactionCounterErr != nil {
		return errors.Wrap(_transactionCounterErr, "Error serializing 'transactionCounter' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := io.PopContext("DF1ResponseMessage"); popErr != nil {
		return popErr
	}
	return nil
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *DF1ResponseMessage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		}
	}
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "destinationAddress":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DestinationAddress = data
			case "sourceAddress":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.SourceAddress = data
			case "status":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Status = data
			case "transactionCounter":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TransactionCounter = data
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of DF1ResponseMessage")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.ab-eth.readwrite.DF1CommandResponseMessageProtectedTypedLogicalRead":
					var dt *DF1CommandResponseMessageProtectedTypedLogicalRead
					if m.Child != nil {
						dt = m.Child.(*DF1CommandResponseMessageProtectedTypedLogicalRead)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				}
			}
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *DF1ResponseMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.ab-eth.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DestinationAddress, xml.StartElement{Name: xml.Name{Local: "destinationAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.SourceAddress, xml.StartElement{Name: xml.Name{Local: "sourceAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Status, xml.StartElement{Name: xml.Name{Local: "status"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TransactionCounter, xml.StartElement{Name: xml.Name{Local: "transactionCounter"}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m DF1ResponseMessage) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m *DF1ResponseMessage) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m *DF1ResponseMessage) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "DF1ResponseMessage"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Reserved Field (reserved)
	// reserved field can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("reserved", uint8(0x00), -1))
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("DestinationAddress", m.DestinationAddress, -1))
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("SourceAddress", m.SourceAddress, -1))
	// Reserved Field (reserved)
	// reserved field can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("reserved", uint8(0x00), -1))
	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode := uint8(m.Child.CommandCode())
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("CommandCode", commandCode, -1))
	// Simple field (case simple)
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("Status", m.Status, -1))
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("TransactionCounter", m.TransactionCounter, -1))
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
