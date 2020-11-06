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
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

type DataTransportErrorCode uint8

type IDataTransportErrorCode interface {
    Serialize(io utils.WriteBuffer) error
}

const(
    DataTransportErrorCode_RESERVED DataTransportErrorCode = 0x00
    DataTransportErrorCode_OK DataTransportErrorCode = 0xFF
    DataTransportErrorCode_ACCESS_DENIED DataTransportErrorCode = 0x03
    DataTransportErrorCode_INVALID_ADDRESS DataTransportErrorCode = 0x05
    DataTransportErrorCode_DATA_TYPE_NOT_SUPPORTED DataTransportErrorCode = 0x06
    DataTransportErrorCode_NOT_FOUND DataTransportErrorCode = 0x0A
)

func DataTransportErrorCodeValueOf(value uint8) DataTransportErrorCode {
    switch value {
        case 0x00:
            return DataTransportErrorCode_RESERVED
        case 0x03:
            return DataTransportErrorCode_ACCESS_DENIED
        case 0x05:
            return DataTransportErrorCode_INVALID_ADDRESS
        case 0x06:
            return DataTransportErrorCode_DATA_TYPE_NOT_SUPPORTED
        case 0x0A:
            return DataTransportErrorCode_NOT_FOUND
        case 0xFF:
            return DataTransportErrorCode_OK
    }
    return 0
}

func CastDataTransportErrorCode(structType interface{}) DataTransportErrorCode {
    castFunc := func(typ interface{}) DataTransportErrorCode {
        if sDataTransportErrorCode, ok := typ.(DataTransportErrorCode); ok {
            return sDataTransportErrorCode
        }
        return 0
    }
    return castFunc(structType)
}

func (m DataTransportErrorCode) LengthInBits() uint16 {
    return 8
}

func (m DataTransportErrorCode) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func DataTransportErrorCodeParse(io *utils.ReadBuffer) (DataTransportErrorCode, error) {
    // TODO: Implement ...
    return 0, nil
}

func (e DataTransportErrorCode) Serialize(io utils.WriteBuffer) error {
    // TODO: Implement ...
    return nil
}
