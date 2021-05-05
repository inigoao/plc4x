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
type BVLCRegisterForeignDevice struct {
	Parent *BVLC
}

// The corresponding interface
type IBVLCRegisterForeignDevice interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCRegisterForeignDevice) BvlcFunction() uint8 {
	return 0x05
}

func (m *BVLCRegisterForeignDevice) InitializeParent(parent *BVLC) {
}

func NewBVLCRegisterForeignDevice() *BVLC {
	child := &BVLCRegisterForeignDevice{
		Parent: NewBVLC(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBVLCRegisterForeignDevice(structType interface{}) *BVLCRegisterForeignDevice {
	castFunc := func(typ interface{}) *BVLCRegisterForeignDevice {
		if casted, ok := typ.(BVLCRegisterForeignDevice); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCRegisterForeignDevice); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCRegisterForeignDevice(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCRegisterForeignDevice(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCRegisterForeignDevice) GetTypeName() string {
	return "BVLCRegisterForeignDevice"
}

func (m *BVLCRegisterForeignDevice) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCRegisterForeignDevice) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BVLCRegisterForeignDevice) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCRegisterForeignDeviceParse(io utils.ReadBuffer) (*BVLC, error) {
	if pullErr := io.PullContext("BVLCRegisterForeignDevice"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := io.CloseContext("BVLCRegisterForeignDevice"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCRegisterForeignDevice{
		Parent: &BVLC{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BVLCRegisterForeignDevice) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("BVLCRegisterForeignDevice"); pushErr != nil {
			return pushErr
		}

		if popErr := io.PopContext("BVLCRegisterForeignDevice"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *BVLCRegisterForeignDevice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
func (m *BVLCRegisterForeignDevice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m BVLCRegisterForeignDevice) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m BVLCRegisterForeignDevice) Box(name string, width int) utils.AsciiBox {
	boxName := "BVLCRegisterForeignDevice"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
