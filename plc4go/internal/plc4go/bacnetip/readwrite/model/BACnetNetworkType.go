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

type BACnetNetworkType uint8

type IBACnetNetworkType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetNetworkType_ETHERNET           BACnetNetworkType = 0x0
	BACnetNetworkType_ARCNET             BACnetNetworkType = 0x1
	BACnetNetworkType_MSTP               BACnetNetworkType = 0x2
	BACnetNetworkType_PTP                BACnetNetworkType = 0x3
	BACnetNetworkType_LONTALK            BACnetNetworkType = 0x4
	BACnetNetworkType_IPV4               BACnetNetworkType = 0x5
	BACnetNetworkType_ZIGBEE             BACnetNetworkType = 0x6
	BACnetNetworkType_VIRTUAL            BACnetNetworkType = 0x7
	BACnetNetworkType_REMOVED_NON_BACNET BACnetNetworkType = 0x8
	BACnetNetworkType_IPV6               BACnetNetworkType = 0x9
	BACnetNetworkType_SERIAL             BACnetNetworkType = 0xA
)

var BACnetNetworkTypeValues []BACnetNetworkType

func init() {
	BACnetNetworkTypeValues = []BACnetNetworkType{
		BACnetNetworkType_ETHERNET,
		BACnetNetworkType_ARCNET,
		BACnetNetworkType_MSTP,
		BACnetNetworkType_PTP,
		BACnetNetworkType_LONTALK,
		BACnetNetworkType_IPV4,
		BACnetNetworkType_ZIGBEE,
		BACnetNetworkType_VIRTUAL,
		BACnetNetworkType_REMOVED_NON_BACNET,
		BACnetNetworkType_IPV6,
		BACnetNetworkType_SERIAL,
	}
}

func BACnetNetworkTypeByValue(value uint8) BACnetNetworkType {
	switch value {
	case 0x0:
		return BACnetNetworkType_ETHERNET
	case 0x1:
		return BACnetNetworkType_ARCNET
	case 0x2:
		return BACnetNetworkType_MSTP
	case 0x3:
		return BACnetNetworkType_PTP
	case 0x4:
		return BACnetNetworkType_LONTALK
	case 0x5:
		return BACnetNetworkType_IPV4
	case 0x6:
		return BACnetNetworkType_ZIGBEE
	case 0x7:
		return BACnetNetworkType_VIRTUAL
	case 0x8:
		return BACnetNetworkType_REMOVED_NON_BACNET
	case 0x9:
		return BACnetNetworkType_IPV6
	case 0xA:
		return BACnetNetworkType_SERIAL
	}
	return 0
}

func BACnetNetworkTypeByName(value string) BACnetNetworkType {
	switch value {
	case "ETHERNET":
		return BACnetNetworkType_ETHERNET
	case "ARCNET":
		return BACnetNetworkType_ARCNET
	case "MSTP":
		return BACnetNetworkType_MSTP
	case "PTP":
		return BACnetNetworkType_PTP
	case "LONTALK":
		return BACnetNetworkType_LONTALK
	case "IPV4":
		return BACnetNetworkType_IPV4
	case "ZIGBEE":
		return BACnetNetworkType_ZIGBEE
	case "VIRTUAL":
		return BACnetNetworkType_VIRTUAL
	case "REMOVED_NON_BACNET":
		return BACnetNetworkType_REMOVED_NON_BACNET
	case "IPV6":
		return BACnetNetworkType_IPV6
	case "SERIAL":
		return BACnetNetworkType_SERIAL
	}
	return 0
}

func CastBACnetNetworkType(structType interface{}) BACnetNetworkType {
	castFunc := func(typ interface{}) BACnetNetworkType {
		if sBACnetNetworkType, ok := typ.(BACnetNetworkType); ok {
			return sBACnetNetworkType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetNetworkType) LengthInBits() uint16 {
	return 4
}

func (m BACnetNetworkType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetNetworkTypeParse(readBuffer utils.ReadBuffer) (BACnetNetworkType, error) {
	val, err := readBuffer.ReadUint8("BACnetNetworkType", 4)
	if err != nil {
		return 0, nil
	}
	return BACnetNetworkTypeByValue(val), nil
}

func (e BACnetNetworkType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetNetworkType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetNetworkType) name() string {
	switch e {
	case BACnetNetworkType_ETHERNET:
		return "ETHERNET"
	case BACnetNetworkType_ARCNET:
		return "ARCNET"
	case BACnetNetworkType_MSTP:
		return "MSTP"
	case BACnetNetworkType_PTP:
		return "PTP"
	case BACnetNetworkType_LONTALK:
		return "LONTALK"
	case BACnetNetworkType_IPV4:
		return "IPV4"
	case BACnetNetworkType_ZIGBEE:
		return "ZIGBEE"
	case BACnetNetworkType_VIRTUAL:
		return "VIRTUAL"
	case BACnetNetworkType_REMOVED_NON_BACNET:
		return "REMOVED_NON_BACNET"
	case BACnetNetworkType_IPV6:
		return "IPV6"
	case BACnetNetworkType_SERIAL:
		return "SERIAL"
	}
	return ""
}

func (e BACnetNetworkType) String() string {
	return e.name()
}
