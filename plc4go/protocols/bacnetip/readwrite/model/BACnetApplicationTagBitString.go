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

// BACnetApplicationTagBitString is the corresponding interface of BACnetApplicationTagBitString
type BACnetApplicationTagBitString interface {
	utils.LengthAware
	utils.Serializable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
}

// BACnetApplicationTagBitStringExactly can be used when we want exactly this type and not a type which fulfills BACnetApplicationTagBitString.
// This is useful for switch cases.
type BACnetApplicationTagBitStringExactly interface {
	BACnetApplicationTagBitString
	isBACnetApplicationTagBitString() bool
}

// _BACnetApplicationTagBitString is the data-structure of this message
type _BACnetApplicationTagBitString struct {
	*_BACnetApplicationTag
	Payload BACnetTagPayloadBitString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagBitString) InitializeParent(parent BACnetApplicationTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetApplicationTagBitString) GetParent() BACnetApplicationTag {
	return m._BACnetApplicationTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagBitString) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagBitString factory function for _BACnetApplicationTagBitString
func NewBACnetApplicationTagBitString(payload BACnetTagPayloadBitString, header BACnetTagHeader) *_BACnetApplicationTagBitString {
	_result := &_BACnetApplicationTagBitString{
		Payload:               payload,
		_BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	_result._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagBitString(structType interface{}) BACnetApplicationTagBitString {
	if casted, ok := structType.(BACnetApplicationTagBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagBitString) GetTypeName() string {
	return "BACnetApplicationTagBitString"
}

func (m *_BACnetApplicationTagBitString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetApplicationTagBitString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetApplicationTagBitString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetApplicationTagBitStringParse(readBuffer utils.ReadBuffer, header BACnetTagHeader) (BACnetApplicationTagBitString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadBitStringParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetApplicationTagBitString")
	}
	payload := _payload.(BACnetTagPayloadBitString)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagBitString")
	}

	// Create a partially initialized instance
	_child := &_BACnetApplicationTagBitString{
		_BACnetApplicationTag: &_BACnetApplicationTag{},
		Payload:               payload,
	}
	_child._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetApplicationTagBitString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagBitString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagBitString")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagBitString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagBitString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagBitString) isBACnetApplicationTagBitString() bool {
	return true
}

func (m *_BACnetApplicationTagBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
