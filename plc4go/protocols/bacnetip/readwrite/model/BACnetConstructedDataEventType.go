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

// BACnetConstructedDataEventType is the corresponding interface of BACnetConstructedDataEventType
type BACnetConstructedDataEventType interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventTypeTagged
}

// BACnetConstructedDataEventTypeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventType.
// This is useful for switch cases.
type BACnetConstructedDataEventTypeExactly interface {
	BACnetConstructedDataEventType
	isBACnetConstructedDataEventType() bool
}

// _BACnetConstructedDataEventType is the data-structure of this message
type _BACnetConstructedDataEventType struct {
	*_BACnetConstructedData
	EventType BACnetEventTypeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventType) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEventType) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventType) GetEventType() BACnetEventTypeTagged {
	return m.EventType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventType) GetActualValue() BACnetEventTypeTagged {
	return CastBACnetEventTypeTagged(m.GetEventType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventType factory function for _BACnetConstructedDataEventType
func NewBACnetConstructedDataEventType(eventType BACnetEventTypeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventType {
	_result := &_BACnetConstructedDataEventType{
		EventType:              eventType,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventType(structType interface{}) BACnetConstructedDataEventType {
	if casted, ok := structType.(BACnetConstructedDataEventType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventType) GetTypeName() string {
	return "BACnetConstructedDataEventType"
}

func (m *_BACnetConstructedDataEventType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEventType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventTypeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventType)
	if pullErr := readBuffer.PullContext("eventType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventType")
	}
	_eventType, _eventTypeErr := BACnetEventTypeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _eventTypeErr != nil {
		return nil, errors.Wrap(_eventTypeErr, "Error parsing 'eventType' field of BACnetConstructedDataEventType")
	}
	eventType := _eventType.(BACnetEventTypeTagged)
	if closeErr := readBuffer.CloseContext("eventType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventType")
	}

	// Virtual field
	_actualValue := eventType
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventType")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEventType{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		EventType: eventType,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEventType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventType")
		}

		// Simple Field (eventType)
		if pushErr := writeBuffer.PushContext("eventType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventType")
		}
		_eventTypeErr := writeBuffer.WriteSerializable(m.GetEventType())
		if popErr := writeBuffer.PopContext("eventType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventType")
		}
		if _eventTypeErr != nil {
			return errors.Wrap(_eventTypeErr, "Error serializing 'eventType' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventType")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventType) isBACnetConstructedDataEventType() bool {
	return true
}

func (m *_BACnetConstructedDataEventType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
