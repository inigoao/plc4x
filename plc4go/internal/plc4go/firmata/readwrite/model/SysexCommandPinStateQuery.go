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
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type SysexCommandPinStateQuery struct {
	Pin    uint8
	Parent *SysexCommand
}

// The corresponding interface
type ISysexCommandPinStateQuery interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SysexCommandPinStateQuery) CommandType() uint8 {
	return 0x6D
}

func (m *SysexCommandPinStateQuery) Response() bool {
	return false
}

func (m *SysexCommandPinStateQuery) InitializeParent(parent *SysexCommand) {
}

func NewSysexCommandPinStateQuery(pin uint8) *SysexCommand {
	child := &SysexCommandPinStateQuery{
		Pin:    pin,
		Parent: NewSysexCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastSysexCommandPinStateQuery(structType interface{}) *SysexCommandPinStateQuery {
	castFunc := func(typ interface{}) *SysexCommandPinStateQuery {
		if casted, ok := typ.(SysexCommandPinStateQuery); ok {
			return &casted
		}
		if casted, ok := typ.(*SysexCommandPinStateQuery); ok {
			return casted
		}
		if casted, ok := typ.(SysexCommand); ok {
			return CastSysexCommandPinStateQuery(casted.Child)
		}
		if casted, ok := typ.(*SysexCommand); ok {
			return CastSysexCommandPinStateQuery(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SysexCommandPinStateQuery) GetTypeName() string {
	return "SysexCommandPinStateQuery"
}

func (m *SysexCommandPinStateQuery) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SysexCommandPinStateQuery) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	return lengthInBits
}

func (m *SysexCommandPinStateQuery) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SysexCommandPinStateQueryParse(io utils.ReadBuffer) (*SysexCommand, error) {
	if pullErr := io.PullContext("SysexCommandPinStateQuery"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (pin)
	pin, _pinErr := io.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field")
	}

	if closeErr := io.CloseContext("SysexCommandPinStateQuery"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandPinStateQuery{
		Pin:    pin,
		Parent: &SysexCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *SysexCommandPinStateQuery) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("SysexCommandPinStateQuery"); pushErr != nil {
			return pushErr
		}

		// Simple Field (pin)
		pin := uint8(m.Pin)
		_pinErr := io.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		if popErr := io.PopContext("SysexCommandPinStateQuery"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *SysexCommandPinStateQuery) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "pin":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Pin = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *SysexCommandPinStateQuery) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Pin, xml.StartElement{Name: xml.Name{Local: "pin"}}); err != nil {
		return err
	}
	return nil
}

func (m SysexCommandPinStateQuery) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m SysexCommandPinStateQuery) Box(name string, width int) utils.AsciiBox {
	boxName := "SysexCommandPinStateQuery"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("Pin", m.Pin, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
