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

// BACnetBackupStateTagged is the corresponding interface of BACnetBackupStateTagged
type BACnetBackupStateTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetBackupState
}

// BACnetBackupStateTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetBackupStateTagged.
// This is useful for switch cases.
type BACnetBackupStateTaggedExactly interface {
	BACnetBackupStateTagged
	isBACnetBackupStateTagged() bool
}

// _BACnetBackupStateTagged is the data-structure of this message
type _BACnetBackupStateTagged struct {
	Header BACnetTagHeader
	Value  BACnetBackupState

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetBackupStateTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetBackupStateTagged) GetValue() BACnetBackupState {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetBackupStateTagged factory function for _BACnetBackupStateTagged
func NewBACnetBackupStateTagged(header BACnetTagHeader, value BACnetBackupState, tagNumber uint8, tagClass TagClass) *_BACnetBackupStateTagged {
	return &_BACnetBackupStateTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetBackupStateTagged(structType interface{}) BACnetBackupStateTagged {
	if casted, ok := structType.(BACnetBackupStateTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetBackupStateTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetBackupStateTagged) GetTypeName() string {
	return "BACnetBackupStateTagged"
}

func (m *_BACnetBackupStateTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetBackupStateTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetBackupStateTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetBackupStateTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetBackupStateTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetBackupStateTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetBackupStateTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetBackupStateTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetBackupState_IDLE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetBackupStateTagged")
	}
	var value BACnetBackupState
	if _value != nil {
		value = _value.(BACnetBackupState)
	}

	if closeErr := readBuffer.CloseContext("BACnetBackupStateTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetBackupStateTagged")
	}

	// Create the instance
	return &_BACnetBackupStateTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetBackupStateTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetBackupStateTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetBackupStateTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetBackupStateTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetBackupStateTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetBackupStateTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetBackupStateTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetBackupStateTagged) isBACnetBackupStateTagged() bool {
	return true
}

func (m *_BACnetBackupStateTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
