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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const S7Message_PROTOCOLID uint8 = 0x32

// The data-structure of this message
type S7Message struct {
	TpduReference uint16
	Parameter     *S7Parameter
	Payload       *S7Payload
	Child         IS7MessageChild
}

// The corresponding interface
type IS7Message interface {
	MessageType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IS7MessageParent interface {
	SerializeParent(io utils.WriteBuffer, child IS7Message, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7MessageChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *S7Message, tpduReference uint16, parameter *S7Parameter, payload *S7Payload)
	GetTypeName() string
	IS7Message
	utils.AsciiBoxer
}

func NewS7Message(tpduReference uint16, parameter *S7Parameter, payload *S7Payload) *S7Message {
	return &S7Message{TpduReference: tpduReference, Parameter: parameter, Payload: payload}
}

func CastS7Message(structType interface{}) *S7Message {
	castFunc := func(typ interface{}) *S7Message {
		if casted, ok := typ.(S7Message); ok {
			return &casted
		}
		if casted, ok := typ.(*S7Message); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7Message) GetTypeName() string {
	return "S7Message"
}

func (m *S7Message) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7Message) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *S7Message) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Const Field (protocolId)
	lengthInBits += 8
	// Discriminator Field (messageType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (tpduReference)
	lengthInBits += 16

	// Implicit Field (parameterLength)
	lengthInBits += 16

	// Implicit Field (payloadLength)
	lengthInBits += 16

	// Optional Field (parameter)
	if m.Parameter != nil {
		lengthInBits += (*m.Parameter).LengthInBits()
	}

	// Optional Field (payload)
	if m.Payload != nil {
		lengthInBits += (*m.Payload).LengthInBits()
	}

	return lengthInBits
}

func (m *S7Message) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7MessageParse(io utils.ReadBuffer) (*S7Message, error) {
	if pullErr := io.PullContext("S7Message"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (protocolId)
	protocolId, _protocolIdErr := io.ReadUint8("protocolId", 8)
	if _protocolIdErr != nil {
		return nil, errors.Wrap(_protocolIdErr, "Error parsing 'protocolId' field")
	}
	if protocolId != S7Message_PROTOCOLID {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7Message_PROTOCOLID) + " but got " + fmt.Sprintf("%d", protocolId))
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType, _messageTypeErr := io.ReadUint8("messageType", 8)
	if _messageTypeErr != nil {
		return nil, errors.Wrap(_messageTypeErr, "Error parsing 'messageType' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := io.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (tpduReference)
	tpduReference, _tpduReferenceErr := io.ReadUint16("tpduReference", 16)
	if _tpduReferenceErr != nil {
		return nil, errors.Wrap(_tpduReferenceErr, "Error parsing 'tpduReference' field")
	}

	// Implicit Field (parameterLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	parameterLength, _parameterLengthErr := io.ReadUint16("parameterLength", 16)
	_ = parameterLength
	if _parameterLengthErr != nil {
		return nil, errors.Wrap(_parameterLengthErr, "Error parsing 'parameterLength' field")
	}

	// Implicit Field (payloadLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	payloadLength, _payloadLengthErr := io.ReadUint16("payloadLength", 16)
	_ = payloadLength
	if _payloadLengthErr != nil {
		return nil, errors.Wrap(_payloadLengthErr, "Error parsing 'payloadLength' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *S7Message
	var typeSwitchError error
	switch {
	case messageType == 0x01: // S7MessageRequest
		_parent, typeSwitchError = S7MessageRequestParse(io)
	case messageType == 0x02: // S7MessageResponse
		_parent, typeSwitchError = S7MessageResponseParse(io)
	case messageType == 0x03: // S7MessageResponseData
		_parent, typeSwitchError = S7MessageResponseDataParse(io)
	case messageType == 0x07: // S7MessageUserData
		_parent, typeSwitchError = S7MessageUserDataParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Optional Field (parameter) (Can be skipped, if a given expression evaluates to false)
	var parameter *S7Parameter = nil
	if bool((parameterLength) > (0)) {
		_val, _err := S7ParameterParse(io, messageType)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'parameter' field")
		}
		parameter = _val
	}

	// Optional Field (payload) (Can be skipped, if a given expression evaluates to false)
	var payload *S7Payload = nil
	if bool((payloadLength) > (0)) {
		_val, _err := S7PayloadParse(io, messageType, (parameter))
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'payload' field")
		}
		payload = _val
	}

	if closeErr := io.CloseContext("S7Message"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, tpduReference, parameter, payload)
	return _parent, nil
}

func (m *S7Message) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *S7Message) SerializeParent(io utils.WriteBuffer, child IS7Message, serializeChildFunction func() error) error {
	if pushErr := io.PushContext("S7Message"); pushErr != nil {
		return pushErr
	}

	// Const Field (protocolId)
	_protocolIdErr := io.WriteUint8("protocolId", 8, 0x32)
	if _protocolIdErr != nil {
		return errors.Wrap(_protocolIdErr, "Error serializing 'protocolId' field")
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := uint8(child.MessageType())
	_messageTypeErr := io.WriteUint8("messageType", 8, (messageType))

	if _messageTypeErr != nil {
		return errors.Wrap(_messageTypeErr, "Error serializing 'messageType' field")
	}

	// Reserved Field (reserved)
	{
		_err := io.WriteUint16("reserved", 16, uint16(0x0000))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (tpduReference)
	tpduReference := uint16(m.TpduReference)
	_tpduReferenceErr := io.WriteUint16("tpduReference", 16, (tpduReference))
	if _tpduReferenceErr != nil {
		return errors.Wrap(_tpduReferenceErr, "Error serializing 'tpduReference' field")
	}

	// Implicit Field (parameterLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	parameterLength := uint16(utils.InlineIf(bool((m.Parameter) != (nil)), func() uint16 { return uint16(m.Parameter.LengthInBytes()) }, func() uint16 { return uint16(uint16(0)) }))
	_parameterLengthErr := io.WriteUint16("parameterLength", 16, (parameterLength))
	if _parameterLengthErr != nil {
		return errors.Wrap(_parameterLengthErr, "Error serializing 'parameterLength' field")
	}

	// Implicit Field (payloadLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	payloadLength := uint16(utils.InlineIf(bool((m.Payload) != (nil)), func() uint16 { return uint16(m.Payload.LengthInBytes()) }, func() uint16 { return uint16(uint16(0)) }))
	_payloadLengthErr := io.WriteUint16("payloadLength", 16, (payloadLength))
	if _payloadLengthErr != nil {
		return errors.Wrap(_payloadLengthErr, "Error serializing 'payloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Optional Field (parameter) (Can be skipped, if the value is null)
	var parameter *S7Parameter = nil
	if m.Parameter != nil {
		parameter = m.Parameter
		_parameterErr := parameter.Serialize(io)
		if _parameterErr != nil {
			return errors.Wrap(_parameterErr, "Error serializing 'parameter' field")
		}
	}

	// Optional Field (payload) (Can be skipped, if the value is null)
	var payload *S7Payload = nil
	if m.Payload != nil {
		payload = m.Payload
		_payloadErr := payload.Serialize(io)
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
	}

	if popErr := io.PopContext("S7Message"); popErr != nil {
		return popErr
	}
	return nil
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *S7Message) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// S7MessageRequest needs special treatment as it has no fields
		case "org.apache.plc4x.java.s7.readwrite.S7MessageRequest":
			if m.Child == nil {
				m.Child = &S7MessageRequest{
					Parent: m,
				}
			}
		// S7MessageUserData needs special treatment as it has no fields
		case "org.apache.plc4x.java.s7.readwrite.S7MessageUserData":
			if m.Child == nil {
				m.Child = &S7MessageUserData{
					Parent: m,
				}
			}
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
			case "tpduReference":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TpduReference = data
			case "parameter":
				var dt *S7Parameter
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.Parameter = dt
			case "payload":
				var dt *S7Payload
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.Payload = dt
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of S7Message")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.s7.readwrite.S7MessageRequest":
					var dt *S7MessageRequest
					if m.Child != nil {
						dt = m.Child.(*S7MessageRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.s7.readwrite.S7MessageResponse":
					var dt *S7MessageResponse
					if m.Child != nil {
						dt = m.Child.(*S7MessageResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.s7.readwrite.S7MessageResponseData":
					var dt *S7MessageResponseData
					if m.Child != nil {
						dt = m.Child.(*S7MessageResponseData)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.s7.readwrite.S7MessageUserData":
					var dt *S7MessageUserData
					if m.Child != nil {
						dt = m.Child.(*S7MessageUserData)
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
func (m *S7Message) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.s7.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TpduReference, xml.StartElement{Name: xml.Name{Local: "tpduReference"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Parameter, xml.StartElement{Name: xml.Name{Local: "parameter"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Payload, xml.StartElement{Name: xml.Name{Local: "payload"}}); err != nil {
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

func (m S7Message) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m *S7Message) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m *S7Message) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "S7Message"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Const Field (protocolId)
	boxes = append(boxes, utils.BoxAnything("ProtocolId", uint8(0x32), -1))
	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := uint8(m.Child.MessageType())
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("MessageType", messageType, -1))
	// Reserved Field (reserved)
	// reserved field can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("reserved", uint16(0x0000), -1))
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("TpduReference", m.TpduReference, -1))
	// Implicit Field (parameterLength)
	parameterLength := uint16(utils.InlineIf(bool((m.Parameter) != (nil)), func() uint16 { return uint16(m.Parameter.LengthInBytes()) }, func() uint16 { return uint16(uint16(0)) }))
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ParameterLength", parameterLength, -1))
	// Implicit Field (payloadLength)
	payloadLength := uint16(utils.InlineIf(bool((m.Payload) != (nil)), func() uint16 { return uint16(m.Payload.LengthInBytes()) }, func() uint16 { return uint16(uint16(0)) }))
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("PayloadLength", payloadLength, -1))
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	// Optional Field (parameter) (Can be skipped, if the value is null)
	var parameter *S7Parameter = nil
	if m.Parameter != nil {
		parameter = m.Parameter
		boxes = append(boxes, parameter.Box("parameter", width-2))
	}
	// Optional Field (payload) (Can be skipped, if the value is null)
	var payload *S7Payload = nil
	if m.Payload != nil {
		payload = m.Payload
		boxes = append(boxes, payload.Box("payload", width-2))
	}
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
