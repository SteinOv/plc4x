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

// BACnetConstructedDataUpdateInterval is the corresponding interface of BACnetConstructedDataUpdateInterval
type BACnetConstructedDataUpdateInterval interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUpdateInterval returns UpdateInterval (property field)
	GetUpdateInterval() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataUpdateIntervalExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUpdateInterval.
// This is useful for switch cases.
type BACnetConstructedDataUpdateIntervalExactly interface {
	BACnetConstructedDataUpdateInterval
	isBACnetConstructedDataUpdateInterval() bool
}

// _BACnetConstructedDataUpdateInterval is the data-structure of this message
type _BACnetConstructedDataUpdateInterval struct {
	*_BACnetConstructedData
	UpdateInterval BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUpdateInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUpdateInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UPDATE_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUpdateInterval) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUpdateInterval) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUpdateInterval) GetUpdateInterval() BACnetApplicationTagUnsignedInteger {
	return m.UpdateInterval
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUpdateInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetUpdateInterval())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUpdateInterval factory function for _BACnetConstructedDataUpdateInterval
func NewBACnetConstructedDataUpdateInterval(updateInterval BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUpdateInterval {
	_result := &_BACnetConstructedDataUpdateInterval{
		UpdateInterval:         updateInterval,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUpdateInterval(structType interface{}) BACnetConstructedDataUpdateInterval {
	if casted, ok := structType.(BACnetConstructedDataUpdateInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUpdateInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUpdateInterval) GetTypeName() string {
	return "BACnetConstructedDataUpdateInterval"
}

func (m *_BACnetConstructedDataUpdateInterval) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataUpdateInterval) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (updateInterval)
	lengthInBits += m.UpdateInterval.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUpdateInterval) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUpdateIntervalParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUpdateInterval, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUpdateInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUpdateInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (updateInterval)
	if pullErr := readBuffer.PullContext("updateInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for updateInterval")
	}
	_updateInterval, _updateIntervalErr := BACnetApplicationTagParse(readBuffer)
	if _updateIntervalErr != nil {
		return nil, errors.Wrap(_updateIntervalErr, "Error parsing 'updateInterval' field of BACnetConstructedDataUpdateInterval")
	}
	updateInterval := _updateInterval.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("updateInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for updateInterval")
	}

	// Virtual field
	_actualValue := updateInterval
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUpdateInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUpdateInterval")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUpdateInterval{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		UpdateInterval: updateInterval,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUpdateInterval) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUpdateInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUpdateInterval")
		}

		// Simple Field (updateInterval)
		if pushErr := writeBuffer.PushContext("updateInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for updateInterval")
		}
		_updateIntervalErr := writeBuffer.WriteSerializable(m.GetUpdateInterval())
		if popErr := writeBuffer.PopContext("updateInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for updateInterval")
		}
		if _updateIntervalErr != nil {
			return errors.Wrap(_updateIntervalErr, "Error serializing 'updateInterval' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUpdateInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUpdateInterval")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUpdateInterval) isBACnetConstructedDataUpdateInterval() bool {
	return true
}

func (m *_BACnetConstructedDataUpdateInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
