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
type DF1CommandResponseMessageProtectedTypedLogicalRead struct {
	Data   []uint8
	Parent *DF1ResponseMessage
}

// The corresponding interface
type IDF1CommandResponseMessageProtectedTypedLogicalRead interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) CommandCode() uint8 {
	return 0x4F
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) InitializeParent(parent *DF1ResponseMessage, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) {
	m.Parent.DestinationAddress = destinationAddress
	m.Parent.SourceAddress = sourceAddress
	m.Parent.Status = status
	m.Parent.TransactionCounter = transactionCounter
}

func NewDF1CommandResponseMessageProtectedTypedLogicalRead(data []uint8, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) *DF1ResponseMessage {
	child := &DF1CommandResponseMessageProtectedTypedLogicalRead{
		Data:   data,
		Parent: NewDF1ResponseMessage(destinationAddress, sourceAddress, status, transactionCounter),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastDF1CommandResponseMessageProtectedTypedLogicalRead(structType interface{}) *DF1CommandResponseMessageProtectedTypedLogicalRead {
	castFunc := func(typ interface{}) *DF1CommandResponseMessageProtectedTypedLogicalRead {
		if casted, ok := typ.(DF1CommandResponseMessageProtectedTypedLogicalRead); ok {
			return &casted
		}
		if casted, ok := typ.(*DF1CommandResponseMessageProtectedTypedLogicalRead); ok {
			return casted
		}
		if casted, ok := typ.(DF1ResponseMessage); ok {
			return CastDF1CommandResponseMessageProtectedTypedLogicalRead(casted.Child)
		}
		if casted, ok := typ.(*DF1ResponseMessage); ok {
			return CastDF1CommandResponseMessageProtectedTypedLogicalRead(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) GetTypeName() string {
	return "DF1CommandResponseMessageProtectedTypedLogicalRead"
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DF1CommandResponseMessageProtectedTypedLogicalReadParse(io utils.ReadBuffer, payloadLength uint16, status uint8) (*DF1ResponseMessage, error) {
	if pullErr := io.PullContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); pullErr != nil {
		return nil, pullErr
	}

	// Array field (data)
	if pullErr := io.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	data := make([]uint8, 0)
	_dataLength := uint16(payloadLength) - uint16(uint16(8))
	_dataEndPos := io.GetPos() + uint16(_dataLength)
	for io.GetPos() < _dataEndPos {
		_item, _err := io.ReadUint8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data = append(data, _item)
	}
	if closeErr := io.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DF1CommandResponseMessageProtectedTypedLogicalRead{
		Data:   data,
		Parent: &DF1ResponseMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); pushErr != nil {
			return pushErr
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := io.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := io.WriteUint8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := io.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := io.PopContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "data":
				var data []uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Data = data
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
func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m DF1CommandResponseMessageProtectedTypedLogicalRead) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m DF1CommandResponseMessageProtectedTypedLogicalRead) Box(name string, width int) utils.AsciiBox {
	boxName := "DF1CommandResponseMessageProtectedTypedLogicalRead"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Array Field (data)
		if m.Data != nil {
			// Simple array base type uint8 will be hex dumped
			boxes = append(boxes, utils.BoxedDumpAnything("Data", m.Data))
			// Simple array base type uint8 will be rendered one by one
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.Data {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("Data", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
