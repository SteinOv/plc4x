/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataOther is the corresponding interface of ApduDataOther
type ApduDataOther interface {
	utils.LengthAware
	utils.Serializable
	ApduData
	// GetExtendedApdu returns ExtendedApdu (property field)
	GetExtendedApdu() ApduDataExt
}

// ApduDataOtherExactly can be used when we want exactly this type and not a type which fulfills ApduDataOther.
// This is useful for switch cases.
type ApduDataOtherExactly interface {
	ApduDataOther
	isApduDataOther() bool
}

// _ApduDataOther is the data-structure of this message
type _ApduDataOther struct {
	*_ApduData
	ExtendedApdu ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataOther) GetApciType() uint8 {
	return 0xF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataOther) InitializeParent(parent ApduData) {}

func (m *_ApduDataOther) GetParent() ApduData {
	return m._ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataOther) GetExtendedApdu() ApduDataExt {
	return m.ExtendedApdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataOther factory function for _ApduDataOther
func NewApduDataOther(extendedApdu ApduDataExt, dataLength uint8) *_ApduDataOther {
	_result := &_ApduDataOther{
		ExtendedApdu: extendedApdu,
		_ApduData:    NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataOther(structType interface{}) ApduDataOther {
	if casted, ok := structType.(ApduDataOther); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataOther); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataOther) GetTypeName() string {
	return "ApduDataOther"
}

func (m *_ApduDataOther) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataOther) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (extendedApdu)
	lengthInBits += m.ExtendedApdu.GetLengthInBits()

	return lengthInBits
}

func (m *_ApduDataOther) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataOtherParse(readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataOther, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataOther"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataOther")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (extendedApdu)
	if pullErr := readBuffer.PullContext("extendedApdu"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for extendedApdu")
	}
	_extendedApdu, _extendedApduErr := ApduDataExtParse(readBuffer, uint8(dataLength))
	if _extendedApduErr != nil {
		return nil, errors.Wrap(_extendedApduErr, "Error parsing 'extendedApdu' field of ApduDataOther")
	}
	extendedApdu := _extendedApdu.(ApduDataExt)
	if closeErr := readBuffer.CloseContext("extendedApdu"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for extendedApdu")
	}

	if closeErr := readBuffer.CloseContext("ApduDataOther"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataOther")
	}

	// Create a partially initialized instance
	_child := &_ApduDataOther{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
		ExtendedApdu: extendedApdu,
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataOther) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataOther"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataOther")
		}

		// Simple Field (extendedApdu)
		if pushErr := writeBuffer.PushContext("extendedApdu"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for extendedApdu")
		}
		_extendedApduErr := writeBuffer.WriteSerializable(m.GetExtendedApdu())
		if popErr := writeBuffer.PopContext("extendedApdu"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for extendedApdu")
		}
		if _extendedApduErr != nil {
			return errors.Wrap(_extendedApduErr, "Error serializing 'extendedApdu' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataOther"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataOther")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataOther) isApduDataOther() bool {
	return true
}

func (m *_ApduDataOther) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
