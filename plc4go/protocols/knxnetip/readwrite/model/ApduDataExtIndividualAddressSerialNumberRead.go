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

// ApduDataExtIndividualAddressSerialNumberRead is the corresponding interface of ApduDataExtIndividualAddressSerialNumberRead
type ApduDataExtIndividualAddressSerialNumberRead interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtIndividualAddressSerialNumberReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtIndividualAddressSerialNumberRead.
// This is useful for switch cases.
type ApduDataExtIndividualAddressSerialNumberReadExactly interface {
	ApduDataExtIndividualAddressSerialNumberRead
	isApduDataExtIndividualAddressSerialNumberRead() bool
}

// _ApduDataExtIndividualAddressSerialNumberRead is the data-structure of this message
type _ApduDataExtIndividualAddressSerialNumberRead struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetExtApciType() uint8 {
	return 0x1C
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtIndividualAddressSerialNumberRead) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtIndividualAddressSerialNumberRead factory function for _ApduDataExtIndividualAddressSerialNumberRead
func NewApduDataExtIndividualAddressSerialNumberRead(length uint8) *_ApduDataExtIndividualAddressSerialNumberRead {
	_result := &_ApduDataExtIndividualAddressSerialNumberRead{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtIndividualAddressSerialNumberRead(structType interface{}) ApduDataExtIndividualAddressSerialNumberRead {
	if casted, ok := structType.(ApduDataExtIndividualAddressSerialNumberRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtIndividualAddressSerialNumberRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetTypeName() string {
	return "ApduDataExtIndividualAddressSerialNumberRead"
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtIndividualAddressSerialNumberReadParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtIndividualAddressSerialNumberRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtIndividualAddressSerialNumberRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtIndividualAddressSerialNumberRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtIndividualAddressSerialNumberRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtIndividualAddressSerialNumberRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtIndividualAddressSerialNumberRead{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtIndividualAddressSerialNumberRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtIndividualAddressSerialNumberRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtIndividualAddressSerialNumberRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtIndividualAddressSerialNumberRead")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) isApduDataExtIndividualAddressSerialNumberRead() bool {
	return true
}

func (m *_ApduDataExtIndividualAddressSerialNumberRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
