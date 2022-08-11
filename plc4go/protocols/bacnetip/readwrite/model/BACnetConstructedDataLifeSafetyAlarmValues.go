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

// BACnetConstructedDataLifeSafetyAlarmValues is the corresponding interface of BACnetConstructedDataLifeSafetyAlarmValues
type BACnetConstructedDataLifeSafetyAlarmValues interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetLifeSafetyStateTagged
}

// BACnetConstructedDataLifeSafetyAlarmValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLifeSafetyAlarmValues.
// This is useful for switch cases.
type BACnetConstructedDataLifeSafetyAlarmValuesExactly interface {
	BACnetConstructedDataLifeSafetyAlarmValues
	isBACnetConstructedDataLifeSafetyAlarmValues() bool
}

// _BACnetConstructedDataLifeSafetyAlarmValues is the data-structure of this message
type _BACnetConstructedDataLifeSafetyAlarmValues struct {
	*_BACnetConstructedData
	AlarmValues []BACnetLifeSafetyStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIFE_SAFETY_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetAlarmValues() []BACnetLifeSafetyStateTagged {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLifeSafetyAlarmValues factory function for _BACnetConstructedDataLifeSafetyAlarmValues
func NewBACnetConstructedDataLifeSafetyAlarmValues(alarmValues []BACnetLifeSafetyStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyAlarmValues {
	_result := &_BACnetConstructedDataLifeSafetyAlarmValues{
		AlarmValues:            alarmValues,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyAlarmValues(structType interface{}) BACnetConstructedDataLifeSafetyAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyAlarmValues"
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLifeSafetyAlarmValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLifeSafetyAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (alarmValues)
	if pullErr := readBuffer.PullContext("alarmValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmValues")
	}
	// Terminated array
	var alarmValues []BACnetLifeSafetyStateTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLifeSafetyStateTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'alarmValues' field of BACnetConstructedDataLifeSafetyAlarmValues")
			}
			alarmValues = append(alarmValues, _item.(BACnetLifeSafetyStateTagged))

		}
	}
	if closeErr := readBuffer.CloseContext("alarmValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyAlarmValues")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLifeSafetyAlarmValues{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AlarmValues: alarmValues,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyAlarmValues")
		}

		// Array Field (alarmValues)
		if pushErr := writeBuffer.PushContext("alarmValues", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alarmValues")
		}
		for _, _element := range m.GetAlarmValues() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'alarmValues' field")
			}
		}
		if popErr := writeBuffer.PopContext("alarmValues", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alarmValues")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyAlarmValues")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) isBACnetConstructedDataLifeSafetyAlarmValues() bool {
	return true
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
