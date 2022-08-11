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

// BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement is the corresponding interface of BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
type BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameterChangeOfValueCivCriteria
	// GetReferencedPropertyIncrement returns ReferencedPropertyIncrement (property field)
	GetReferencedPropertyIncrement() BACnetContextTagReal
}

// BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement.
// This is useful for switch cases.
type BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementExactly interface {
	BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
	isBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement() bool
}

// _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement is the data-structure of this message
type _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement struct {
	*_BACnetEventParameterChangeOfValueCivCriteria
	ReferencedPropertyIncrement BACnetContextTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) InitializeParent(parent BACnetEventParameterChangeOfValueCivCriteria, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetParent() BACnetEventParameterChangeOfValueCivCriteria {
	return m._BACnetEventParameterChangeOfValueCivCriteria
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetReferencedPropertyIncrement() BACnetContextTagReal {
	return m.ReferencedPropertyIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement factory function for _BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
func NewBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement(referencedPropertyIncrement BACnetContextTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement {
	_result := &_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement{
		ReferencedPropertyIncrement:                   referencedPropertyIncrement,
		_BACnetEventParameterChangeOfValueCivCriteria: NewBACnetEventParameterChangeOfValueCivCriteria(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetEventParameterChangeOfValueCivCriteria._BACnetEventParameterChangeOfValueCivCriteriaChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement(structType interface{}) BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement {
	if casted, ok := structType.(BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetTypeName() string {
	return "BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (referencedPropertyIncrement)
	lengthInBits += m.ReferencedPropertyIncrement.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referencedPropertyIncrement)
	if pullErr := readBuffer.PullContext("referencedPropertyIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referencedPropertyIncrement")
	}
	_referencedPropertyIncrement, _referencedPropertyIncrementErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_REAL))
	if _referencedPropertyIncrementErr != nil {
		return nil, errors.Wrap(_referencedPropertyIncrementErr, "Error parsing 'referencedPropertyIncrement' field of BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
	}
	referencedPropertyIncrement := _referencedPropertyIncrement.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("referencedPropertyIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referencedPropertyIncrement")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement{
		_BACnetEventParameterChangeOfValueCivCriteria: &_BACnetEventParameterChangeOfValueCivCriteria{
			TagNumber: tagNumber,
		},
		ReferencedPropertyIncrement: referencedPropertyIncrement,
	}
	_child._BACnetEventParameterChangeOfValueCivCriteria._BACnetEventParameterChangeOfValueCivCriteriaChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
		}

		// Simple Field (referencedPropertyIncrement)
		if pushErr := writeBuffer.PushContext("referencedPropertyIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referencedPropertyIncrement")
		}
		_referencedPropertyIncrementErr := writeBuffer.WriteSerializable(m.GetReferencedPropertyIncrement())
		if popErr := writeBuffer.PopContext("referencedPropertyIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referencedPropertyIncrement")
		}
		if _referencedPropertyIncrementErr != nil {
			return errors.Wrap(_referencedPropertyIncrementErr, "Error serializing 'referencedPropertyIncrement' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) isBACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
