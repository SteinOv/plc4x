/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduType is an enum
type ApduType uint8

type IApduType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	ApduType_CONFIRMED_REQUEST_PDU   ApduType = 0
	ApduType_UNCONFIRMED_REQUEST_PDU ApduType = 1
	ApduType_SIMPLE_ACK_PDU          ApduType = 2
	ApduType_COMPLEX_ACK_PDU         ApduType = 3
	ApduType_SEGMENT_ACK_PDU         ApduType = 4
	ApduType_ERROR_PDU               ApduType = 5
	ApduType_REJECT_PDU              ApduType = 6
	ApduType_ABORT_PDU               ApduType = 7
)

var ApduTypeValues []ApduType

func init() {
	_ = errors.New
	ApduTypeValues = []ApduType{
		ApduType_CONFIRMED_REQUEST_PDU,
		ApduType_UNCONFIRMED_REQUEST_PDU,
		ApduType_SIMPLE_ACK_PDU,
		ApduType_COMPLEX_ACK_PDU,
		ApduType_SEGMENT_ACK_PDU,
		ApduType_ERROR_PDU,
		ApduType_REJECT_PDU,
		ApduType_ABORT_PDU,
	}
}

func ApduTypeByValue(value uint8) ApduType {
	switch value {
	case 0:
		return ApduType_CONFIRMED_REQUEST_PDU
	case 1:
		return ApduType_UNCONFIRMED_REQUEST_PDU
	case 2:
		return ApduType_SIMPLE_ACK_PDU
	case 3:
		return ApduType_COMPLEX_ACK_PDU
	case 4:
		return ApduType_SEGMENT_ACK_PDU
	case 5:
		return ApduType_ERROR_PDU
	case 6:
		return ApduType_REJECT_PDU
	case 7:
		return ApduType_ABORT_PDU
	}
	return 0
}

func ApduTypeByName(value string) ApduType {
	switch value {
	case "CONFIRMED_REQUEST_PDU":
		return ApduType_CONFIRMED_REQUEST_PDU
	case "UNCONFIRMED_REQUEST_PDU":
		return ApduType_UNCONFIRMED_REQUEST_PDU
	case "SIMPLE_ACK_PDU":
		return ApduType_SIMPLE_ACK_PDU
	case "COMPLEX_ACK_PDU":
		return ApduType_COMPLEX_ACK_PDU
	case "SEGMENT_ACK_PDU":
		return ApduType_SEGMENT_ACK_PDU
	case "ERROR_PDU":
		return ApduType_ERROR_PDU
	case "REJECT_PDU":
		return ApduType_REJECT_PDU
	case "ABORT_PDU":
		return ApduType_ABORT_PDU
	}
	return 0
}

func ApduTypeKnows(value uint8) bool {
	for _, typeValue := range ApduTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastApduType(structType interface{}) ApduType {
	castFunc := func(typ interface{}) ApduType {
		if sApduType, ok := typ.(ApduType); ok {
			return sApduType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ApduType) GetLengthInBits() uint16 {
	return 4
}

func (m ApduType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduTypeParse(readBuffer utils.ReadBuffer) (ApduType, error) {
	val, err := readBuffer.ReadUint8("ApduType", 4)
	if err != nil {
		return 0, nil
	}
	return ApduTypeByValue(val), nil
}

func (e ApduType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ApduType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e ApduType) name() string {
	switch e {
	case ApduType_CONFIRMED_REQUEST_PDU:
		return "CONFIRMED_REQUEST_PDU"
	case ApduType_UNCONFIRMED_REQUEST_PDU:
		return "UNCONFIRMED_REQUEST_PDU"
	case ApduType_SIMPLE_ACK_PDU:
		return "SIMPLE_ACK_PDU"
	case ApduType_COMPLEX_ACK_PDU:
		return "COMPLEX_ACK_PDU"
	case ApduType_SEGMENT_ACK_PDU:
		return "SEGMENT_ACK_PDU"
	case ApduType_ERROR_PDU:
		return "ERROR_PDU"
	case ApduType_REJECT_PDU:
		return "REJECT_PDU"
	case ApduType_ABORT_PDU:
		return "ABORT_PDU"
	}
	return ""
}

func (e ApduType) String() string {
	return e.name()
}
