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

// BACnetLogDataLogDataEntryBooleanValue is the corresponding interface of BACnetLogDataLogDataEntryBooleanValue
type BACnetLogDataLogDataEntryBooleanValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogDataLogDataEntry
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetContextTagBoolean
}

// BACnetLogDataLogDataEntryBooleanValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataEntryBooleanValue.
// This is useful for switch cases.
type BACnetLogDataLogDataEntryBooleanValueExactly interface {
	BACnetLogDataLogDataEntryBooleanValue
	isBACnetLogDataLogDataEntryBooleanValue() bool
}

// _BACnetLogDataLogDataEntryBooleanValue is the data-structure of this message
type _BACnetLogDataLogDataEntryBooleanValue struct {
	*_BACnetLogDataLogDataEntry
	BooleanValue BACnetContextTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataEntryBooleanValue) InitializeParent(parent BACnetLogDataLogDataEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetParent() BACnetLogDataLogDataEntry {
	return m._BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetBooleanValue() BACnetContextTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryBooleanValue factory function for _BACnetLogDataLogDataEntryBooleanValue
func NewBACnetLogDataLogDataEntryBooleanValue(booleanValue BACnetContextTagBoolean, peekedTagHeader BACnetTagHeader) *_BACnetLogDataLogDataEntryBooleanValue {
	_result := &_BACnetLogDataLogDataEntryBooleanValue{
		BooleanValue:               booleanValue,
		_BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryBooleanValue(structType interface{}) BACnetLogDataLogDataEntryBooleanValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryBooleanValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryBooleanValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryBooleanValue"
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogDataLogDataEntryBooleanValueParse(readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntryBooleanValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryBooleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryBooleanValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for booleanValue")
	}
	_booleanValue, _booleanValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field of BACnetLogDataLogDataEntryBooleanValue")
	}
	booleanValue := _booleanValue.(BACnetContextTagBoolean)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for booleanValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryBooleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryBooleanValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogDataEntryBooleanValue{
		_BACnetLogDataLogDataEntry: &_BACnetLogDataLogDataEntry{},
		BooleanValue:               booleanValue,
	}
	_child._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryBooleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryBooleanValue")
		}

		// Simple Field (booleanValue)
		if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for booleanValue")
		}
		_booleanValueErr := writeBuffer.WriteSerializable(m.GetBooleanValue())
		if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for booleanValue")
		}
		if _booleanValueErr != nil {
			return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryBooleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryBooleanValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) isBACnetLogDataLogDataEntryBooleanValue() bool {
	return true
}

func (m *_BACnetLogDataLogDataEntryBooleanValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
