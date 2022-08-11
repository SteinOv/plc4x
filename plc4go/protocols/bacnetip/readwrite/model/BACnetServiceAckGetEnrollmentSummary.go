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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetServiceAckGetEnrollmentSummary is the corresponding interface of BACnetServiceAckGetEnrollmentSummary
type BACnetServiceAckGetEnrollmentSummary interface {
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventTypeTagged
	// GetEventState returns EventState (property field)
	GetEventState() BACnetEventStateTagged
	// GetPriority returns Priority (property field)
	GetPriority() BACnetApplicationTagUnsignedInteger
	// GetNotificationClass returns NotificationClass (property field)
	GetNotificationClass() BACnetApplicationTagUnsignedInteger
}

// BACnetServiceAckGetEnrollmentSummaryExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckGetEnrollmentSummary.
// This is useful for switch cases.
type BACnetServiceAckGetEnrollmentSummaryExactly interface {
	BACnetServiceAckGetEnrollmentSummary
	isBACnetServiceAckGetEnrollmentSummary() bool
}

// _BACnetServiceAckGetEnrollmentSummary is the data-structure of this message
type _BACnetServiceAckGetEnrollmentSummary struct {
	*_BACnetServiceAck
	ObjectIdentifier  BACnetApplicationTagObjectIdentifier
	EventType         BACnetEventTypeTagged
	EventState        BACnetEventStateTagged
	Priority          BACnetApplicationTagUnsignedInteger
	NotificationClass BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckGetEnrollmentSummary) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckGetEnrollmentSummary) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckGetEnrollmentSummary) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetEventType() BACnetEventTypeTagged {
	return m.EventType
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetEventState() BACnetEventStateTagged {
	return m.EventState
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetPriority() BACnetApplicationTagUnsignedInteger {
	return m.Priority
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetNotificationClass() BACnetApplicationTagUnsignedInteger {
	return m.NotificationClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckGetEnrollmentSummary factory function for _BACnetServiceAckGetEnrollmentSummary
func NewBACnetServiceAckGetEnrollmentSummary(objectIdentifier BACnetApplicationTagObjectIdentifier, eventType BACnetEventTypeTagged, eventState BACnetEventStateTagged, priority BACnetApplicationTagUnsignedInteger, notificationClass BACnetApplicationTagUnsignedInteger, serviceAckLength uint16) *_BACnetServiceAckGetEnrollmentSummary {
	_result := &_BACnetServiceAckGetEnrollmentSummary{
		ObjectIdentifier:  objectIdentifier,
		EventType:         eventType,
		EventState:        eventState,
		Priority:          priority,
		NotificationClass: notificationClass,
		_BACnetServiceAck: NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckGetEnrollmentSummary(structType interface{}) BACnetServiceAckGetEnrollmentSummary {
	if casted, ok := structType.(BACnetServiceAckGetEnrollmentSummary); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckGetEnrollmentSummary); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetTypeName() string {
	return "BACnetServiceAckGetEnrollmentSummary"
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits()

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits()

	// Simple field (priority)
	lengthInBits += m.Priority.GetLengthInBits()

	// Optional Field (notificationClass)
	if m.NotificationClass != nil {
		lengthInBits += m.NotificationClass.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckGetEnrollmentSummaryParse(readBuffer utils.ReadBuffer, serviceAckLength uint16) (BACnetServiceAckGetEnrollmentSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckGetEnrollmentSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckGetEnrollmentSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetServiceAckGetEnrollmentSummary")
	}
	objectIdentifier := _objectIdentifier.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (eventType)
	if pullErr := readBuffer.PullContext("eventType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventType")
	}
	_eventType, _eventTypeErr := BACnetEventTypeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _eventTypeErr != nil {
		return nil, errors.Wrap(_eventTypeErr, "Error parsing 'eventType' field of BACnetServiceAckGetEnrollmentSummary")
	}
	eventType := _eventType.(BACnetEventTypeTagged)
	if closeErr := readBuffer.CloseContext("eventType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventType")
	}

	// Simple Field (eventState)
	if pullErr := readBuffer.PullContext("eventState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventState")
	}
	_eventState, _eventStateErr := BACnetEventStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field of BACnetServiceAckGetEnrollmentSummary")
	}
	eventState := _eventState.(BACnetEventStateTagged)
	if closeErr := readBuffer.CloseContext("eventState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventState")
	}

	// Simple Field (priority)
	if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priority")
	}
	_priority, _priorityErr := BACnetApplicationTagParse(readBuffer)
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field of BACnetServiceAckGetEnrollmentSummary")
	}
	priority := _priority.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priority")
	}

	// Optional Field (notificationClass) (Can be skipped, if a given expression evaluates to false)
	var notificationClass BACnetApplicationTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("notificationClass"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for notificationClass")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'notificationClass' field of BACnetServiceAckGetEnrollmentSummary")
		default:
			notificationClass = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("notificationClass"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for notificationClass")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckGetEnrollmentSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckGetEnrollmentSummary")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckGetEnrollmentSummary{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		ObjectIdentifier:  objectIdentifier,
		EventType:         eventType,
		EventState:        eventState,
		Priority:          priority,
		NotificationClass: notificationClass,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckGetEnrollmentSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckGetEnrollmentSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckGetEnrollmentSummary")
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		_objectIdentifierErr := writeBuffer.WriteSerializable(m.GetObjectIdentifier())
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
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

		// Simple Field (eventState)
		if pushErr := writeBuffer.PushContext("eventState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventState")
		}
		_eventStateErr := writeBuffer.WriteSerializable(m.GetEventState())
		if popErr := writeBuffer.PopContext("eventState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventState")
		}
		if _eventStateErr != nil {
			return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
		}

		// Simple Field (priority)
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priority")
		}
		_priorityErr := writeBuffer.WriteSerializable(m.GetPriority())
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priority")
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}

		// Optional Field (notificationClass) (Can be skipped, if the value is null)
		var notificationClass BACnetApplicationTagUnsignedInteger = nil
		if m.GetNotificationClass() != nil {
			if pushErr := writeBuffer.PushContext("notificationClass"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for notificationClass")
			}
			notificationClass = m.GetNotificationClass()
			_notificationClassErr := writeBuffer.WriteSerializable(notificationClass)
			if popErr := writeBuffer.PopContext("notificationClass"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for notificationClass")
			}
			if _notificationClassErr != nil {
				return errors.Wrap(_notificationClassErr, "Error serializing 'notificationClass' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckGetEnrollmentSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckGetEnrollmentSummary")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetServiceAckGetEnrollmentSummary) isBACnetServiceAckGetEnrollmentSummary() bool {
	return true
}

func (m *_BACnetServiceAckGetEnrollmentSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
