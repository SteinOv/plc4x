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

// BACnetFaultParameterFaultExtendedParametersEntryInteger is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryInteger
type BACnetFaultParameterFaultExtendedParametersEntryInteger interface {
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
}

// BACnetFaultParameterFaultExtendedParametersEntryIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParametersEntryInteger.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersEntryIntegerExactly interface {
	BACnetFaultParameterFaultExtendedParametersEntryInteger
	isBACnetFaultParameterFaultExtendedParametersEntryInteger() bool
}

// _BACnetFaultParameterFaultExtendedParametersEntryInteger is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryInteger struct {
	*_BACnetFaultParameterFaultExtendedParametersEntry
	IntegerValue BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) InitializeParent(parent BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetParent() BACnetFaultParameterFaultExtendedParametersEntry {
	return m._BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntryInteger factory function for _BACnetFaultParameterFaultExtendedParametersEntryInteger
func NewBACnetFaultParameterFaultExtendedParametersEntryInteger(integerValue BACnetApplicationTagSignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtendedParametersEntryInteger {
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryInteger{
		IntegerValue: integerValue,
		_BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryInteger(structType interface{}) BACnetFaultParameterFaultExtendedParametersEntryInteger {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryInteger"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryIntegerParse(readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParametersEntryInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerValue)
	if pullErr := readBuffer.PullContext("integerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerValue")
	}
	_integerValue, _integerValueErr := BACnetApplicationTagParse(readBuffer)
	if _integerValueErr != nil {
		return nil, errors.Wrap(_integerValueErr, "Error parsing 'integerValue' field of BACnetFaultParameterFaultExtendedParametersEntryInteger")
	}
	integerValue := _integerValue.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("integerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryInteger")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultExtendedParametersEntryInteger{
		_BACnetFaultParameterFaultExtendedParametersEntry: &_BACnetFaultParameterFaultExtendedParametersEntry{},
		IntegerValue: integerValue,
	}
	_child._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryInteger")
		}

		// Simple Field (integerValue)
		if pushErr := writeBuffer.PushContext("integerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for integerValue")
		}
		_integerValueErr := writeBuffer.WriteSerializable(m.GetIntegerValue())
		if popErr := writeBuffer.PopContext("integerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for integerValue")
		}
		if _integerValueErr != nil {
			return errors.Wrap(_integerValueErr, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryInteger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) isBACnetFaultParameterFaultExtendedParametersEntryInteger() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
