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

// TriggerControlLabelOptions is the corresponding interface of TriggerControlLabelOptions
type TriggerControlLabelOptions interface {
	utils.LengthAware
	utils.Serializable
	// GetLabelFlavour returns LabelFlavour (property field)
	GetLabelFlavour() TriggerControlLabelFlavour
	// GetLabelType returns LabelType (property field)
	GetLabelType() TriggerControlLabelType
}

// TriggerControlLabelOptionsExactly can be used when we want exactly this type and not a type which fulfills TriggerControlLabelOptions.
// This is useful for switch cases.
type TriggerControlLabelOptionsExactly interface {
	TriggerControlLabelOptions
	isTriggerControlLabelOptions() bool
}

// _TriggerControlLabelOptions is the data-structure of this message
type _TriggerControlLabelOptions struct {
	LabelFlavour TriggerControlLabelFlavour
	LabelType    TriggerControlLabelType
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
	reservedField2 *bool
	reservedField3 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TriggerControlLabelOptions) GetLabelFlavour() TriggerControlLabelFlavour {
	return m.LabelFlavour
}

func (m *_TriggerControlLabelOptions) GetLabelType() TriggerControlLabelType {
	return m.LabelType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTriggerControlLabelOptions factory function for _TriggerControlLabelOptions
func NewTriggerControlLabelOptions(labelFlavour TriggerControlLabelFlavour, labelType TriggerControlLabelType) *_TriggerControlLabelOptions {
	return &_TriggerControlLabelOptions{LabelFlavour: labelFlavour, LabelType: labelType}
}

// Deprecated: use the interface for direct cast
func CastTriggerControlLabelOptions(structType interface{}) TriggerControlLabelOptions {
	if casted, ok := structType.(TriggerControlLabelOptions); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlLabelOptions); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlLabelOptions) GetTypeName() string {
	return "TriggerControlLabelOptions"
}

func (m *_TriggerControlLabelOptions) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TriggerControlLabelOptions) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (labelFlavour)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (labelType)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	return lengthInBits
}

func (m *_TriggerControlLabelOptions) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TriggerControlLabelOptionsParse(readBuffer utils.ReadBuffer) (TriggerControlLabelOptions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TriggerControlLabelOptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlLabelOptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of TriggerControlLabelOptions")
		}
		if reserved != bool(false) {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (labelFlavour)
	if pullErr := readBuffer.PullContext("labelFlavour"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for labelFlavour")
	}
	_labelFlavour, _labelFlavourErr := TriggerControlLabelFlavourParse(readBuffer)
	if _labelFlavourErr != nil {
		return nil, errors.Wrap(_labelFlavourErr, "Error parsing 'labelFlavour' field of TriggerControlLabelOptions")
	}
	labelFlavour := _labelFlavour
	if closeErr := readBuffer.CloseContext("labelFlavour"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for labelFlavour")
	}

	var reservedField1 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of TriggerControlLabelOptions")
		}
		if reserved != bool(false) {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	var reservedField2 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of TriggerControlLabelOptions")
		}
		if reserved != bool(false) {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField2 = &reserved
		}
	}

	// Simple Field (labelType)
	if pullErr := readBuffer.PullContext("labelType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for labelType")
	}
	_labelType, _labelTypeErr := TriggerControlLabelTypeParse(readBuffer)
	if _labelTypeErr != nil {
		return nil, errors.Wrap(_labelTypeErr, "Error parsing 'labelType' field of TriggerControlLabelOptions")
	}
	labelType := _labelType
	if closeErr := readBuffer.CloseContext("labelType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for labelType")
	}

	var reservedField3 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of TriggerControlLabelOptions")
		}
		if reserved != bool(false) {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField3 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("TriggerControlLabelOptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlLabelOptions")
	}

	// Create the instance
	return &_TriggerControlLabelOptions{
		LabelFlavour:   labelFlavour,
		LabelType:      labelType,
		reservedField0: reservedField0,
		reservedField1: reservedField1,
		reservedField2: reservedField2,
		reservedField3: reservedField3,
	}, nil
}

func (m *_TriggerControlLabelOptions) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("TriggerControlLabelOptions"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TriggerControlLabelOptions")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (labelFlavour)
	if pushErr := writeBuffer.PushContext("labelFlavour"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for labelFlavour")
	}
	_labelFlavourErr := writeBuffer.WriteSerializable(m.GetLabelFlavour())
	if popErr := writeBuffer.PopContext("labelFlavour"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for labelFlavour")
	}
	if _labelFlavourErr != nil {
		return errors.Wrap(_labelFlavourErr, "Error serializing 'labelFlavour' field")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField1 != nil {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField1
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField2 != nil {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField2
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (labelType)
	if pushErr := writeBuffer.PushContext("labelType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for labelType")
	}
	_labelTypeErr := writeBuffer.WriteSerializable(m.GetLabelType())
	if popErr := writeBuffer.PopContext("labelType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for labelType")
	}
	if _labelTypeErr != nil {
		return errors.Wrap(_labelTypeErr, "Error serializing 'labelType' field")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField3 != nil {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField3
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	if popErr := writeBuffer.PopContext("TriggerControlLabelOptions"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TriggerControlLabelOptions")
	}
	return nil
}

func (m *_TriggerControlLabelOptions) isTriggerControlLabelOptions() bool {
	return true
}

func (m *_TriggerControlLabelOptions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
