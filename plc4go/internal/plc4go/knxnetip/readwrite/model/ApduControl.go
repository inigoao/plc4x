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
type ApduControl struct {
	Child IApduControlChild
}

// The corresponding interface
type IApduControl interface {
	ControlType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

type IApduControlParent interface {
	SerializeParent(io utils.WriteBuffer, child IApduControl, serializeChildFunction func() error) error
	GetTypeName() string
}

type IApduControlChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *ApduControl)
	GetTypeName() string
	IApduControl
	utils.AsciiBoxer
}

func NewApduControl() *ApduControl {
	return &ApduControl{}
}

func CastApduControl(structType interface{}) *ApduControl {
	castFunc := func(typ interface{}) *ApduControl {
		if casted, ok := typ.(ApduControl); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduControl); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduControl) GetTypeName() string {
	return "ApduControl"
}

func (m *ApduControl) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduControl) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *ApduControl) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (controlType)
	lengthInBits += 2

	return lengthInBits
}

func (m *ApduControl) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduControlParse(io utils.ReadBuffer) (*ApduControl, error) {
	if pullErr := io.PullContext("ApduControl"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (controlType) (Used as input to a switch field)
	controlType, _controlTypeErr := io.ReadUint8("controlType", 2)
	if _controlTypeErr != nil {
		return nil, errors.Wrap(_controlTypeErr, "Error parsing 'controlType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ApduControl
	var typeSwitchError error
	switch {
	case controlType == 0x0: // ApduControlConnect
		_parent, typeSwitchError = ApduControlConnectParse(io)
	case controlType == 0x1: // ApduControlDisconnect
		_parent, typeSwitchError = ApduControlDisconnectParse(io)
	case controlType == 0x2: // ApduControlAck
		_parent, typeSwitchError = ApduControlAckParse(io)
	case controlType == 0x3: // ApduControlNack
		_parent, typeSwitchError = ApduControlNackParse(io)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := io.CloseContext("ApduControl"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ApduControl) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *ApduControl) SerializeParent(io utils.WriteBuffer, child IApduControl, serializeChildFunction func() error) error {
	if pushErr := io.PushContext("ApduControl"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (controlType) (Used as input to a switch field)
	controlType := uint8(child.ControlType())
	_controlTypeErr := io.WriteUint8("controlType", 2, (controlType))

	if _controlTypeErr != nil {
		return errors.Wrap(_controlTypeErr, "Error serializing 'controlType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := io.PopContext("ApduControl"); popErr != nil {
		return popErr
	}
	return nil
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *ApduControl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	if start.Attr != nil && len(start.Attr) > 0 {
		switch start.Attr[0].Value {
		// ApduControlConnect needs special treatment as it has no fields
		case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlConnect":
			if m.Child == nil {
				m.Child = &ApduControlConnect{
					Parent: m,
				}
			}
		// ApduControlDisconnect needs special treatment as it has no fields
		case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlDisconnect":
			if m.Child == nil {
				m.Child = &ApduControlDisconnect{
					Parent: m,
				}
			}
		// ApduControlAck needs special treatment as it has no fields
		case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlAck":
			if m.Child == nil {
				m.Child = &ApduControlAck{
					Parent: m,
				}
			}
		// ApduControlNack needs special treatment as it has no fields
		case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlNack":
			if m.Child == nil {
				m.Child = &ApduControlNack{
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
					panic("Couldn't determine class type for childs of ApduControl")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlConnect":
					var dt *ApduControlConnect
					if m.Child != nil {
						dt = m.Child.(*ApduControlConnect)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlDisconnect":
					var dt *ApduControlDisconnect
					if m.Child != nil {
						dt = m.Child.(*ApduControlDisconnect)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlAck":
					var dt *ApduControlAck
					if m.Child != nil {
						dt = m.Child.(*ApduControlAck)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.knxnetip.readwrite.ApduControlNack":
					var dt *ApduControlNack
					if m.Child != nil {
						dt = m.Child.(*ApduControlNack)
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
func (m *ApduControl) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.knxnetip.readwrite." + className[strings.LastIndex(className, ".")+1:]
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

func (m ApduControl) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m *ApduControl) Box(name string, width int) utils.AsciiBox {
	return m.Child.Box(name, width)
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m *ApduControl) BoxParent(name string, width int, childBoxer func() []utils.AsciiBox) utils.AsciiBox {
	boxName := "ApduControl"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Discriminator Field (controlType) (Used as input to a switch field)
	controlType := uint8(m.Child.ControlType())
	// uint8 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("ControlType", controlType, -1))
	// Switch field (Depending on the discriminator values, passes the boxing to a sub-type)
	boxes = append(boxes, childBoxer()...)
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
