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
)

// Code generated by code-generation. DO NOT EDIT.

// GroupObjectDescriptorRealisationType1 is the corresponding interface of GroupObjectDescriptorRealisationType1
type GroupObjectDescriptorRealisationType1 interface {
	utils.LengthAware
	utils.Serializable
	// GetDataPointer returns DataPointer (property field)
	GetDataPointer() uint8
	// GetTransmitEnable returns TransmitEnable (property field)
	GetTransmitEnable() bool
	// GetSegmentSelectorEnable returns SegmentSelectorEnable (property field)
	GetSegmentSelectorEnable() bool
	// GetWriteEnable returns WriteEnable (property field)
	GetWriteEnable() bool
	// GetReadEnable returns ReadEnable (property field)
	GetReadEnable() bool
	// GetCommunicationEnable returns CommunicationEnable (property field)
	GetCommunicationEnable() bool
	// GetPriority returns Priority (property field)
	GetPriority() CEMIPriority
	// GetValueType returns ValueType (property field)
	GetValueType() ComObjectValueType
}

// GroupObjectDescriptorRealisationType1Exactly can be used when we want exactly this type and not a type which fulfills GroupObjectDescriptorRealisationType1.
// This is useful for switch cases.
type GroupObjectDescriptorRealisationType1Exactly interface {
	GroupObjectDescriptorRealisationType1
	isGroupObjectDescriptorRealisationType1() bool
}

// _GroupObjectDescriptorRealisationType1 is the data-structure of this message
type _GroupObjectDescriptorRealisationType1 struct {
	DataPointer           uint8
	TransmitEnable        bool
	SegmentSelectorEnable bool
	WriteEnable           bool
	ReadEnable            bool
	CommunicationEnable   bool
	Priority              CEMIPriority
	ValueType             ComObjectValueType
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GroupObjectDescriptorRealisationType1) GetDataPointer() uint8 {
	return m.DataPointer
}

func (m *_GroupObjectDescriptorRealisationType1) GetTransmitEnable() bool {
	return m.TransmitEnable
}

func (m *_GroupObjectDescriptorRealisationType1) GetSegmentSelectorEnable() bool {
	return m.SegmentSelectorEnable
}

func (m *_GroupObjectDescriptorRealisationType1) GetWriteEnable() bool {
	return m.WriteEnable
}

func (m *_GroupObjectDescriptorRealisationType1) GetReadEnable() bool {
	return m.ReadEnable
}

func (m *_GroupObjectDescriptorRealisationType1) GetCommunicationEnable() bool {
	return m.CommunicationEnable
}

func (m *_GroupObjectDescriptorRealisationType1) GetPriority() CEMIPriority {
	return m.Priority
}

func (m *_GroupObjectDescriptorRealisationType1) GetValueType() ComObjectValueType {
	return m.ValueType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGroupObjectDescriptorRealisationType1 factory function for _GroupObjectDescriptorRealisationType1
func NewGroupObjectDescriptorRealisationType1(dataPointer uint8, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) *_GroupObjectDescriptorRealisationType1 {
	return &_GroupObjectDescriptorRealisationType1{DataPointer: dataPointer, TransmitEnable: transmitEnable, SegmentSelectorEnable: segmentSelectorEnable, WriteEnable: writeEnable, ReadEnable: readEnable, CommunicationEnable: communicationEnable, Priority: priority, ValueType: valueType}
}

// Deprecated: use the interface for direct cast
func CastGroupObjectDescriptorRealisationType1(structType interface{}) GroupObjectDescriptorRealisationType1 {
	if casted, ok := structType.(GroupObjectDescriptorRealisationType1); ok {
		return casted
	}
	if casted, ok := structType.(*GroupObjectDescriptorRealisationType1); ok {
		return *casted
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType1) GetTypeName() string {
	return "GroupObjectDescriptorRealisationType1"
}

func (m *_GroupObjectDescriptorRealisationType1) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_GroupObjectDescriptorRealisationType1) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dataPointer)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (transmitEnable)
	lengthInBits += 1

	// Simple field (segmentSelectorEnable)
	lengthInBits += 1

	// Simple field (writeEnable)
	lengthInBits += 1

	// Simple field (readEnable)
	lengthInBits += 1

	// Simple field (communicationEnable)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (valueType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_GroupObjectDescriptorRealisationType1) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func GroupObjectDescriptorRealisationType1Parse(readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType1, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GroupObjectDescriptorRealisationType1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GroupObjectDescriptorRealisationType1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dataPointer)
	_dataPointer, _dataPointerErr := readBuffer.ReadUint8("dataPointer", 8)
	if _dataPointerErr != nil {
		return nil, errors.Wrap(_dataPointerErr, "Error parsing 'dataPointer' field of GroupObjectDescriptorRealisationType1")
	}
	dataPointer := _dataPointer

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of GroupObjectDescriptorRealisationType1")
		}
		if reserved != uint8(0x1) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x1),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (transmitEnable)
	_transmitEnable, _transmitEnableErr := readBuffer.ReadBit("transmitEnable")
	if _transmitEnableErr != nil {
		return nil, errors.Wrap(_transmitEnableErr, "Error parsing 'transmitEnable' field of GroupObjectDescriptorRealisationType1")
	}
	transmitEnable := _transmitEnable

	// Simple Field (segmentSelectorEnable)
	_segmentSelectorEnable, _segmentSelectorEnableErr := readBuffer.ReadBit("segmentSelectorEnable")
	if _segmentSelectorEnableErr != nil {
		return nil, errors.Wrap(_segmentSelectorEnableErr, "Error parsing 'segmentSelectorEnable' field of GroupObjectDescriptorRealisationType1")
	}
	segmentSelectorEnable := _segmentSelectorEnable

	// Simple Field (writeEnable)
	_writeEnable, _writeEnableErr := readBuffer.ReadBit("writeEnable")
	if _writeEnableErr != nil {
		return nil, errors.Wrap(_writeEnableErr, "Error parsing 'writeEnable' field of GroupObjectDescriptorRealisationType1")
	}
	writeEnable := _writeEnable

	// Simple Field (readEnable)
	_readEnable, _readEnableErr := readBuffer.ReadBit("readEnable")
	if _readEnableErr != nil {
		return nil, errors.Wrap(_readEnableErr, "Error parsing 'readEnable' field of GroupObjectDescriptorRealisationType1")
	}
	readEnable := _readEnable

	// Simple Field (communicationEnable)
	_communicationEnable, _communicationEnableErr := readBuffer.ReadBit("communicationEnable")
	if _communicationEnableErr != nil {
		return nil, errors.Wrap(_communicationEnableErr, "Error parsing 'communicationEnable' field of GroupObjectDescriptorRealisationType1")
	}
	communicationEnable := _communicationEnable

	// Simple Field (priority)
	if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priority")
	}
	_priority, _priorityErr := CEMIPriorityParse(readBuffer)
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field of GroupObjectDescriptorRealisationType1")
	}
	priority := _priority
	if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priority")
	}

	// Simple Field (valueType)
	if pullErr := readBuffer.PullContext("valueType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for valueType")
	}
	_valueType, _valueTypeErr := ComObjectValueTypeParse(readBuffer)
	if _valueTypeErr != nil {
		return nil, errors.Wrap(_valueTypeErr, "Error parsing 'valueType' field of GroupObjectDescriptorRealisationType1")
	}
	valueType := _valueType
	if closeErr := readBuffer.CloseContext("valueType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for valueType")
	}

	if closeErr := readBuffer.CloseContext("GroupObjectDescriptorRealisationType1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GroupObjectDescriptorRealisationType1")
	}

	// Create the instance
	return &_GroupObjectDescriptorRealisationType1{
		DataPointer:           dataPointer,
		TransmitEnable:        transmitEnable,
		SegmentSelectorEnable: segmentSelectorEnable,
		WriteEnable:           writeEnable,
		ReadEnable:            readEnable,
		CommunicationEnable:   communicationEnable,
		Priority:              priority,
		ValueType:             valueType,
		reservedField0:        reservedField0,
	}, nil
}

func (m *_GroupObjectDescriptorRealisationType1) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("GroupObjectDescriptorRealisationType1"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for GroupObjectDescriptorRealisationType1")
	}

	// Simple Field (dataPointer)
	dataPointer := uint8(m.GetDataPointer())
	_dataPointerErr := writeBuffer.WriteUint8("dataPointer", 8, (dataPointer))
	if _dataPointerErr != nil {
		return errors.Wrap(_dataPointerErr, "Error serializing 'dataPointer' field")
	}

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0x1)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x1),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteUint8("reserved", 1, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (transmitEnable)
	transmitEnable := bool(m.GetTransmitEnable())
	_transmitEnableErr := writeBuffer.WriteBit("transmitEnable", (transmitEnable))
	if _transmitEnableErr != nil {
		return errors.Wrap(_transmitEnableErr, "Error serializing 'transmitEnable' field")
	}

	// Simple Field (segmentSelectorEnable)
	segmentSelectorEnable := bool(m.GetSegmentSelectorEnable())
	_segmentSelectorEnableErr := writeBuffer.WriteBit("segmentSelectorEnable", (segmentSelectorEnable))
	if _segmentSelectorEnableErr != nil {
		return errors.Wrap(_segmentSelectorEnableErr, "Error serializing 'segmentSelectorEnable' field")
	}

	// Simple Field (writeEnable)
	writeEnable := bool(m.GetWriteEnable())
	_writeEnableErr := writeBuffer.WriteBit("writeEnable", (writeEnable))
	if _writeEnableErr != nil {
		return errors.Wrap(_writeEnableErr, "Error serializing 'writeEnable' field")
	}

	// Simple Field (readEnable)
	readEnable := bool(m.GetReadEnable())
	_readEnableErr := writeBuffer.WriteBit("readEnable", (readEnable))
	if _readEnableErr != nil {
		return errors.Wrap(_readEnableErr, "Error serializing 'readEnable' field")
	}

	// Simple Field (communicationEnable)
	communicationEnable := bool(m.GetCommunicationEnable())
	_communicationEnableErr := writeBuffer.WriteBit("communicationEnable", (communicationEnable))
	if _communicationEnableErr != nil {
		return errors.Wrap(_communicationEnableErr, "Error serializing 'communicationEnable' field")
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

	// Simple Field (valueType)
	if pushErr := writeBuffer.PushContext("valueType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for valueType")
	}
	_valueTypeErr := writeBuffer.WriteSerializable(m.GetValueType())
	if popErr := writeBuffer.PopContext("valueType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for valueType")
	}
	if _valueTypeErr != nil {
		return errors.Wrap(_valueTypeErr, "Error serializing 'valueType' field")
	}

	if popErr := writeBuffer.PopContext("GroupObjectDescriptorRealisationType1"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for GroupObjectDescriptorRealisationType1")
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType1) isGroupObjectDescriptorRealisationType1() bool {
	return true
}

func (m *_GroupObjectDescriptorRealisationType1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
