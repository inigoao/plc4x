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
type SysexCommandReportFirmwareResponse struct {
	MajorVersion uint8
	MinorVersion uint8
	FileName     []int8
	Parent       *SysexCommand
}

// The corresponding interface
type ISysexCommandReportFirmwareResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *SysexCommandReportFirmwareResponse) CommandType() uint8 {
	return 0x79
}

func (m *SysexCommandReportFirmwareResponse) Response() bool {
	return true
}

func (m *SysexCommandReportFirmwareResponse) InitializeParent(parent *SysexCommand) {
}

func NewSysexCommandReportFirmwareResponse(majorVersion uint8, minorVersion uint8, fileName []int8) *SysexCommand {
	child := &SysexCommandReportFirmwareResponse{
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		FileName:     fileName,
		Parent:       NewSysexCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastSysexCommandReportFirmwareResponse(structType interface{}) *SysexCommandReportFirmwareResponse {
	castFunc := func(typ interface{}) *SysexCommandReportFirmwareResponse {
		if casted, ok := typ.(SysexCommandReportFirmwareResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*SysexCommandReportFirmwareResponse); ok {
			return casted
		}
		if casted, ok := typ.(SysexCommand); ok {
			return CastSysexCommandReportFirmwareResponse(casted.Child)
		}
		if casted, ok := typ.(*SysexCommand); ok {
			return CastSysexCommandReportFirmwareResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *SysexCommandReportFirmwareResponse) GetTypeName() string {
	return "SysexCommandReportFirmwareResponse"
}

func (m *SysexCommandReportFirmwareResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *SysexCommandReportFirmwareResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	// Manual Array Field (fileName)
	// TODO: below expression doesn't work yet
	// lengthInBits += FirmataUtilsLengthSysexString(fileName) * 8

	return lengthInBits
}

func (m *SysexCommandReportFirmwareResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func SysexCommandReportFirmwareResponseParse(io utils.ReadBuffer) (*SysexCommand, error) {
	if pullErr := io.PullContext("SysexCommandReportFirmwareResponse"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (majorVersion)
	majorVersion, _majorVersionErr := io.ReadUint8("majorVersion", 8)
	if _majorVersionErr != nil {
		return nil, errors.Wrap(_majorVersionErr, "Error parsing 'majorVersion' field")
	}

	// Simple Field (minorVersion)
	minorVersion, _minorVersionErr := io.ReadUint8("minorVersion", 8)
	if _minorVersionErr != nil {
		return nil, errors.Wrap(_minorVersionErr, "Error parsing 'minorVersion' field")
	}
	if pullErr := io.PullContext("fileName", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Manual Array Field (fileName)
	// Terminated array
	_fileNameList := make([]int8, 0)
	for !((bool)(FirmataUtilsIsSysexEnd(io))) {
		_fileNameList = append(_fileNameList, ((int8)(FirmataUtilsParseSysexString(io))))

	}
	fileName := make([]int8, len(_fileNameList))
	for i := 0; i < len(_fileNameList); i++ {
		fileName[i] = int8(_fileNameList[i])
	}
	if closeErr := io.CloseContext("fileName", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := io.CloseContext("SysexCommandReportFirmwareResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SysexCommandReportFirmwareResponse{
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		FileName:     fileName,
		Parent:       &SysexCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *SysexCommandReportFirmwareResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("SysexCommandReportFirmwareResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (majorVersion)
		majorVersion := uint8(m.MajorVersion)
		_majorVersionErr := io.WriteUint8("majorVersion", 8, (majorVersion))
		if _majorVersionErr != nil {
			return errors.Wrap(_majorVersionErr, "Error serializing 'majorVersion' field")
		}

		// Simple Field (minorVersion)
		minorVersion := uint8(m.MinorVersion)
		_minorVersionErr := io.WriteUint8("minorVersion", 8, (minorVersion))
		if _minorVersionErr != nil {
			return errors.Wrap(_minorVersionErr, "Error serializing 'minorVersion' field")
		}

		// Manual Array Field (fileName)
		if m.FileName != nil {
			if pushErr := io.PushContext("fileName", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			//for _, element := range m.FileName {
			//FirmataUtilsSerializeSysexString(io, m.Element)
			//}
			if popErr := io.PopContext("fileName", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := io.PopContext("SysexCommandReportFirmwareResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *SysexCommandReportFirmwareResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "majorVersion":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MajorVersion = data
			case "minorVersion":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MinorVersion = data
			case "fileName":
				// TODO: not implemented yet
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
func (m *SysexCommandReportFirmwareResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.MajorVersion, xml.StartElement{Name: xml.Name{Local: "majorVersion"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MinorVersion, xml.StartElement{Name: xml.Name{Local: "minorVersion"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.FileName, xml.StartElement{Name: xml.Name{Local: "fileName"}}); err != nil {
		return err
	}
	return nil
}

func (m SysexCommandReportFirmwareResponse) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m SysexCommandReportFirmwareResponse) Box(name string, width int) utils.AsciiBox {
	boxName := "SysexCommandReportFirmwareResponse"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("MajorVersion", m.MajorVersion, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("MinorVersion", m.MinorVersion, -1))
		// Manual Array Field (fileName)
		//if m.FileName != nil {
		//	for(int8 element : m.FileName) {
		// TODO FirmataUtilsSerializeSysexString(io, m.Element)
		//	}
		//}
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
