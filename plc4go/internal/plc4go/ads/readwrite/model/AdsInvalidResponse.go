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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type AdsInvalidResponse struct {
	Parent *AdsData
}

// The corresponding interface
type IAdsInvalidResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsInvalidResponse) CommandId() CommandId {
	return CommandId_INVALID
}

func (m *AdsInvalidResponse) Response() bool {
	return true
}

func (m *AdsInvalidResponse) InitializeParent(parent *AdsData) {
}

func NewAdsInvalidResponse() *AdsData {
	child := &AdsInvalidResponse{
		Parent: NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsInvalidResponse(structType interface{}) *AdsInvalidResponse {
	castFunc := func(typ interface{}) *AdsInvalidResponse {
		if casted, ok := typ.(AdsInvalidResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsInvalidResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsInvalidResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsInvalidResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsInvalidResponse) GetTypeName() string {
	return "AdsInvalidResponse"
}

func (m *AdsInvalidResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsInvalidResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *AdsInvalidResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsInvalidResponseParse(readBuffer utils.ReadBuffer) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsInvalidResponse"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("AdsInvalidResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsInvalidResponse{
		Parent: &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsInvalidResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsInvalidResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("AdsInvalidResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
