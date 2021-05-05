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

package readwrite

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/s7/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

type S7XmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.Atoi
	_ = strings.Join
	_ = utils.Dump
	_ = xml.NewDecoder
}

func (m S7XmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	switch typeName {
	case "SzlId":
		return model.SzlIdParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "S7Message":
		return nil, errors.New("S7Message unmappable")
	case "S7VarPayloadStatusItem":
		return model.S7VarPayloadStatusItemParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "S7Parameter":
		return nil, errors.New("S7Parameter unmappable")
	case "SzlDataTreeItem":
		return model.SzlDataTreeItemParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "COTPPacket":
		return nil, errors.New("COTPPacket unmappable")
	case "S7PayloadUserDataItem":
		return nil, errors.New("S7PayloadUserDataItem unmappable")
	case "COTPParameter":
		return nil, errors.New("COTPParameter unmappable")
	case "TPKTPacket":
		return model.TPKTPacketParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "S7Payload":
		return nil, errors.New("S7Payload unmappable")
	case "S7VarRequestParameterItem":
		return nil, errors.New("S7VarRequestParameterItem unmappable")
	case "S7VarPayloadDataItem":
		//var lastItem bool
		lastItem := parserArguments[0] == "true"
		return model.S7VarPayloadDataItemParse(utils.NewXmlReadBuffer(strings.NewReader(xmlString)), lastItem)
	case "S7Address":
		return nil, errors.New("S7Address unmappable")
	case "S7ParameterUserDataItem":
		return nil, errors.New("S7ParameterUserDataItem unmappable")
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}

// Deprecated: will be removed in favor of Parse soon
func (m S7XmlParserHelper) ParseOld(typeName string, xmlString string) (interface{}, error) {
	switch typeName {
	case "SzlId":
		var obj *model.SzlId
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "S7Message":
		var obj *model.S7Message
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "S7VarPayloadStatusItem":
		var obj *model.S7VarPayloadStatusItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "S7Parameter":
		var obj *model.S7Parameter
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "SzlDataTreeItem":
		var obj *model.SzlDataTreeItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "COTPPacket":
		var obj *model.COTPPacket
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "S7PayloadUserDataItem":
		var obj *model.S7PayloadUserDataItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "COTPParameter":
		var obj *model.COTPParameter
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "TPKTPacket":
		var obj *model.TPKTPacket
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "S7Payload":
		var obj *model.S7Payload
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "S7VarRequestParameterItem":
		var obj *model.S7VarRequestParameterItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "S7VarPayloadDataItem":
		var obj *model.S7VarPayloadDataItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "S7Address":
		var obj *model.S7Address
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	case "S7ParameterUserDataItem":
		var obj *model.S7ParameterUserDataItem
		err := xml.Unmarshal([]byte(xmlString), &obj)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling xml")
		}
		return obj, nil
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
