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

// BACnetPropertyStatesShedState is the corresponding interface of BACnetPropertyStatesShedState
type BACnetPropertyStatesShedState interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetShedState returns ShedState (property field)
	GetShedState() BACnetShedStateTagged
}

// BACnetPropertyStatesShedStateExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesShedState.
// This is useful for switch cases.
type BACnetPropertyStatesShedStateExactly interface {
	BACnetPropertyStatesShedState
	isBACnetPropertyStatesShedState() bool
}

// _BACnetPropertyStatesShedState is the data-structure of this message
type _BACnetPropertyStatesShedState struct {
	*_BACnetPropertyStates
	ShedState BACnetShedStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesShedState) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesShedState) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesShedState) GetShedState() BACnetShedStateTagged {
	return m.ShedState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesShedState factory function for _BACnetPropertyStatesShedState
func NewBACnetPropertyStatesShedState(shedState BACnetShedStateTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesShedState {
	_result := &_BACnetPropertyStatesShedState{
		ShedState:             shedState,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesShedState(structType interface{}) BACnetPropertyStatesShedState {
	if casted, ok := structType.(BACnetPropertyStatesShedState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesShedState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesShedState) GetTypeName() string {
	return "BACnetPropertyStatesShedState"
}

func (m *_BACnetPropertyStatesShedState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesShedState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (shedState)
	lengthInBits += m.ShedState.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesShedState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesShedStateParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesShedState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesShedState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesShedState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (shedState)
	if pullErr := readBuffer.PullContext("shedState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for shedState")
	}
	_shedState, _shedStateErr := BACnetShedStateTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _shedStateErr != nil {
		return nil, errors.Wrap(_shedStateErr, "Error parsing 'shedState' field of BACnetPropertyStatesShedState")
	}
	shedState := _shedState.(BACnetShedStateTagged)
	if closeErr := readBuffer.CloseContext("shedState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for shedState")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesShedState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesShedState")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesShedState{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		ShedState:             shedState,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesShedState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesShedState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesShedState")
		}

		// Simple Field (shedState)
		if pushErr := writeBuffer.PushContext("shedState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for shedState")
		}
		_shedStateErr := writeBuffer.WriteSerializable(m.GetShedState())
		if popErr := writeBuffer.PopContext("shedState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for shedState")
		}
		if _shedStateErr != nil {
			return errors.Wrap(_shedStateErr, "Error serializing 'shedState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesShedState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesShedState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesShedState) isBACnetPropertyStatesShedState() bool {
	return true
}

func (m *_BACnetPropertyStatesShedState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
