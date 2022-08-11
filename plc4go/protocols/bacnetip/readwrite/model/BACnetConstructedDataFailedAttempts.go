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

// BACnetConstructedDataFailedAttempts is the corresponding interface of BACnetConstructedDataFailedAttempts
type BACnetConstructedDataFailedAttempts interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFailedAttempts returns FailedAttempts (property field)
	GetFailedAttempts() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataFailedAttemptsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataFailedAttempts.
// This is useful for switch cases.
type BACnetConstructedDataFailedAttemptsExactly interface {
	BACnetConstructedDataFailedAttempts
	isBACnetConstructedDataFailedAttempts() bool
}

// _BACnetConstructedDataFailedAttempts is the data-structure of this message
type _BACnetConstructedDataFailedAttempts struct {
	*_BACnetConstructedData
	FailedAttempts BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFailedAttempts) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFailedAttempts) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAILED_ATTEMPTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFailedAttempts) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataFailedAttempts) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFailedAttempts) GetFailedAttempts() BACnetApplicationTagUnsignedInteger {
	return m.FailedAttempts
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFailedAttempts) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetFailedAttempts())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFailedAttempts factory function for _BACnetConstructedDataFailedAttempts
func NewBACnetConstructedDataFailedAttempts(failedAttempts BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFailedAttempts {
	_result := &_BACnetConstructedDataFailedAttempts{
		FailedAttempts:         failedAttempts,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFailedAttempts(structType interface{}) BACnetConstructedDataFailedAttempts {
	if casted, ok := structType.(BACnetConstructedDataFailedAttempts); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFailedAttempts); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFailedAttempts) GetTypeName() string {
	return "BACnetConstructedDataFailedAttempts"
}

func (m *_BACnetConstructedDataFailedAttempts) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataFailedAttempts) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (failedAttempts)
	lengthInBits += m.FailedAttempts.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFailedAttempts) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataFailedAttemptsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFailedAttempts, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFailedAttempts"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFailedAttempts")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (failedAttempts)
	if pullErr := readBuffer.PullContext("failedAttempts"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for failedAttempts")
	}
	_failedAttempts, _failedAttemptsErr := BACnetApplicationTagParse(readBuffer)
	if _failedAttemptsErr != nil {
		return nil, errors.Wrap(_failedAttemptsErr, "Error parsing 'failedAttempts' field of BACnetConstructedDataFailedAttempts")
	}
	failedAttempts := _failedAttempts.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("failedAttempts"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for failedAttempts")
	}

	// Virtual field
	_actualValue := failedAttempts
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFailedAttempts"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFailedAttempts")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataFailedAttempts{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FailedAttempts: failedAttempts,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataFailedAttempts) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFailedAttempts"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFailedAttempts")
		}

		// Simple Field (failedAttempts)
		if pushErr := writeBuffer.PushContext("failedAttempts"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for failedAttempts")
		}
		_failedAttemptsErr := writeBuffer.WriteSerializable(m.GetFailedAttempts())
		if popErr := writeBuffer.PopContext("failedAttempts"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for failedAttempts")
		}
		if _failedAttemptsErr != nil {
			return errors.Wrap(_failedAttemptsErr, "Error serializing 'failedAttempts' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFailedAttempts"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFailedAttempts")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFailedAttempts) isBACnetConstructedDataFailedAttempts() bool {
	return true
}

func (m *_BACnetConstructedDataFailedAttempts) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
