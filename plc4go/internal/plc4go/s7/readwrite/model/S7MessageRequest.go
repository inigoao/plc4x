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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type S7MessageRequest struct {
	Parent *S7Message
}

// The corresponding interface
type IS7MessageRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7MessageRequest) MessageType() uint8 {
	return 0x01
}

func (m *S7MessageRequest) InitializeParent(parent *S7Message, tpduReference uint16, parameter *S7Parameter, payload *S7Payload) {
	m.Parent.TpduReference = tpduReference
	m.Parent.Parameter = parameter
	m.Parent.Payload = payload
}

func NewS7MessageRequest(tpduReference uint16, parameter *S7Parameter, payload *S7Payload) *S7Message {
	child := &S7MessageRequest{
		Parent: NewS7Message(tpduReference, parameter, payload),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7MessageRequest(structType interface{}) *S7MessageRequest {
	castFunc := func(typ interface{}) *S7MessageRequest {
		if casted, ok := typ.(S7MessageRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*S7MessageRequest); ok {
			return casted
		}
		if casted, ok := typ.(S7Message); ok {
			return CastS7MessageRequest(casted.Child)
		}
		if casted, ok := typ.(*S7Message); ok {
			return CastS7MessageRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7MessageRequest) GetTypeName() string {
	return "S7MessageRequest"
}

func (m *S7MessageRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7MessageRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *S7MessageRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7MessageRequestParse(io utils.ReadBuffer) (*S7Message, error) {
	if pullErr := io.PullContext("S7MessageRequest"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := io.CloseContext("S7MessageRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7MessageRequest{
		Parent: &S7Message{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7MessageRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("S7MessageRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := io.PopContext("S7MessageRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *S7MessageRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
func (m *S7MessageRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m S7MessageRequest) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m S7MessageRequest) Box(name string, width int) utils.AsciiBox {
	boxName := "S7MessageRequest"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
