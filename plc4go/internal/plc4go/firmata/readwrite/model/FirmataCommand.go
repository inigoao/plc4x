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
	"io"
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type FirmataCommand struct {
	Child IFirmataCommandChild
}

// The corresponding interface
type IFirmataCommand interface {
	CommandCode() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IFirmataCommandParent interface {
	SerializeParent(io utils.WriteBuffer, child IFirmataCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type IFirmataCommandChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *FirmataCommand)
	GetTypeName() string
	IFirmataCommand
	utils.AsciiBoxer
}

func NewFirmataCommand() *FirmataCommand {
	return &FirmataCommand{}
}

func CastFirmataCommand(structType interface{}) *FirmataCommand {
	castFunc := func(typ interface{}) *FirmataCommand {
		if casted, ok := typ.(FirmataCommand); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataCommand); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataCommand) GetTypeName() string {
	return "FirmataCommand"
}

func (m *FirmataCommand) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *FirmataCommand) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *FirmataCommand) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandCode)
	lengthInBits += 4

	return lengthInBits
}

func (m *FirmataCommand) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func FirmataCommandParse(io utils.ReadBuffer, response bool) (*FirmataCommand, error) {
	if pullErr := io.PullContext("FirmataCommand"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode, _commandCodeErr := io.ReadUint8("commandCode", 4)
	if _commandCodeErr != nil {
		return nil, errors.Wrap(_commandCodeErr, "Error parsing 'commandCode' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *FirmataCommand
	var typeSwitchError error
	switch {
	case commandCode == 0x0: // FirmataCommandSysex
		_parent, typeSwitchError = FirmataCommandSysexParse(io, response)
	case commandCode == 0x4: // FirmataCommandSetPinMode
		_parent, typeSwitchError = FirmataCommandSetPinModeParse(io)
	case commandCode == 0x5: // FirmataCommandSetDigitalPinValue
		_parent, typeSwitchError = FirmataCommandSetDigitalPinValueParse(io)
	case commandCode == 0x9: // FirmataCommandProtocolVersion
		_parent, typeSwitchError = FirmataCommandProtocolVersionParse(io)
	case commandCode == 0xF: // FirmataCommandSystemReset
		_parent, typeSwitchError = FirmataCommandSystemResetParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := io.CloseContext("FirmataCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *FirmataCommand) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *FirmataCommand) SerializeParent(io utils.WriteBuffer, child IFirmataCommand, serializeChildFunction func() error) error {
	if pushErr := io.PushContext("FirmataCommand"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode := uint8(child.CommandCode())
	_commandCodeErr := io.WriteUint8("commandCode", 4, (commandCode))

	if _commandCodeErr != nil {
		return errors.Wrap(_commandCodeErr, "Error serializing 'commandCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := io.PopContext("FirmataCommand"); popErr != nil {
		return popErr
	}
	return nil
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *FirmataCommand) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// FirmataCommandSystemReset needs special treatment as it has no fields
		case "org.apache.plc4x.java.firmata.readwrite.FirmataCommandSystemReset":
			if m.Child == nil {
				m.Child = &FirmataCommandSystemReset{
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
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of FirmataCommand")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.firmata.readwrite.FirmataCommandSysex":
					var dt *FirmataCommandSysex
					if m.Child != nil {
						dt = m.Child.(*FirmataCommandSysex)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.firmata.readwrite.FirmataCommandSetPinMode":
					var dt *FirmataCommandSetPinMode
					if m.Child != nil {
						dt = m.Child.(*FirmataCommandSetPinMode)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.firmata.readwrite.FirmataCommandSetDigitalPinValue":
					var dt *FirmataCommandSetDigitalPinValue
					if m.Child != nil {
						dt = m.Child.(*FirmataCommandSetDigitalPinValue)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.firmata.readwrite.FirmataCommandProtocolVersion":
					var dt *FirmataCommandProtocolVersion
					if m.Child != nil {
						dt = m.Child.(*FirmataCommandProtocolVersion)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.firmata.readwrite.FirmataCommandSystemReset":
					var dt *FirmataCommandSystemReset
					if m.Child != nil {
						dt = m.Child.(*FirmataCommandSystemReset)
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
func (m *FirmataCommand) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.firmata.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
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

func (m FirmataCommand) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m *FirmataCommand) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m *FirmataCommand) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "FirmataCommand"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode := uint8(m.Child.CommandCode())
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("CommandCode", commandCode, -1))
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
