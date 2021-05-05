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
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AmsTCPPacket struct {
	Userdata *AmsPacket
}

// The corresponding interface
type IAmsTCPPacket interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewAmsTCPPacket(userdata *AmsPacket) *AmsTCPPacket {
	return &AmsTCPPacket{Userdata: userdata}
}

func CastAmsTCPPacket(structType interface{}) *AmsTCPPacket {
	castFunc := func(typ interface{}) *AmsTCPPacket {
		if casted, ok := typ.(AmsTCPPacket); ok {
			return &casted
		}
		if casted, ok := typ.(*AmsTCPPacket); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AmsTCPPacket) GetTypeName() string {
	return "AmsTCPPacket"
}

func (m *AmsTCPPacket) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AmsTCPPacket) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 32

	// Simple field (userdata)
	lengthInBits += m.Userdata.LengthInBits()

	return lengthInBits
}

func (m *AmsTCPPacket) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AmsTCPPacketParse(io utils.ReadBuffer) (*AmsTCPPacket, error) {
	if pullErr := io.PullContext("AmsTCPPacket"); pullErr != nil {
		return nil, pullErr
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

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := io.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	if pullErr := io.PullContext("userdata"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (userdata)
	userdata, _userdataErr := AmsPacketParse(io)
	if _userdataErr != nil {
		return nil, errors.Wrap(_userdataErr, "Error parsing 'userdata' field")
	}
	if closeErr := io.CloseContext("userdata"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("AmsTCPPacket"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAmsTCPPacket(userdata), nil
}

func (m *AmsTCPPacket) Serialize(io utils.WriteBuffer) error {
	if pushErr := io.PushContext("AmsTCPPacket"); pushErr != nil {
		return pushErr
	}

	// Reserved Field (reserved)
	{
		_err := io.WriteUint16("reserved", 16, uint16(0x0000))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length := uint32(m.Userdata.LengthInBytes())
	_lengthErr := io.WriteUint32("length", 32, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (userdata)
	if pushErr := io.PushContext("userdata"); pushErr != nil {
		return pushErr
	}
	_userdataErr := m.Userdata.Serialize(io)
	if popErr := io.PopContext("userdata"); popErr != nil {
		return popErr
	}
	if _userdataErr != nil {
		return errors.Wrap(_userdataErr, "Error serializing 'userdata' field")
	}

	if popErr := io.PopContext("AmsTCPPacket"); popErr != nil {
		return popErr
	}
	return nil
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *AmsTCPPacket) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
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
			case "userdata":
				var data AmsPacket
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Userdata = &data
			}
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *AmsTCPPacket) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.ads.readwrite.AmsTCPPacket"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Userdata, xml.StartElement{Name: xml.Name{Local: "userdata"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m AmsTCPPacket) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m AmsTCPPacket) Box(name string, width int) utils.AsciiBox {
	boxName := "AmsTCPPacket"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Reserved Field (reserved)
	// reserved field can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("reserved", uint16(0x0000), -1))
	// Implicit Field (length)
	length := uint32(m.Userdata.LengthInBytes())
	// uint32 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("Length", length, -1))
	// Complex field (case complex)
	boxes = append(boxes, m.Userdata.Box("userdata", width-2))
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
