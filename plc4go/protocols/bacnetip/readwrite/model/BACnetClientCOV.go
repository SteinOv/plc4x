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

// BACnetClientCOV is the corresponding interface of BACnetClientCOV
type BACnetClientCOV interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetClientCOVExactly can be used when we want exactly this type and not a type which fulfills BACnetClientCOV.
// This is useful for switch cases.
type BACnetClientCOVExactly interface {
	BACnetClientCOV
	isBACnetClientCOV() bool
}

// _BACnetClientCOV is the data-structure of this message
type _BACnetClientCOV struct {
	_BACnetClientCOVChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetClientCOVChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type BACnetClientCOVParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetClientCOV, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetClientCOVChild interface {
	utils.Serializable
	InitializeParent(parent BACnetClientCOV, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetClientCOV

	GetTypeName() string
	BACnetClientCOV
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetClientCOV) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetClientCOV) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetClientCOV factory function for _BACnetClientCOV
func NewBACnetClientCOV(peekedTagHeader BACnetTagHeader) *_BACnetClientCOV {
	return &_BACnetClientCOV{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetClientCOV(structType interface{}) BACnetClientCOV {
	if casted, ok := structType.(BACnetClientCOV); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetClientCOV); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetClientCOV) GetTypeName() string {
	return "BACnetClientCOV"
}

func (m *_BACnetClientCOV) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetClientCOV) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetClientCOVParse(readBuffer utils.ReadBuffer) (BACnetClientCOV, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetClientCOV"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetClientCOV")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetClientCOVChildSerializeRequirement interface {
		BACnetClientCOV
		InitializeParent(BACnetClientCOV, BACnetTagHeader)
		GetParent() BACnetClientCOV
	}
	var _childTemp interface{}
	var _child BACnetClientCOVChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x4: // BACnetClientCOVObject
		_childTemp, typeSwitchError = BACnetClientCOVObjectParse(readBuffer)
	case peekedTagNumber == 0x0: // BACnetClientCOVNone
		_childTemp, typeSwitchError = BACnetClientCOVNoneParse(readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetClientCOV")
	}
	_child = _childTemp.(BACnetClientCOVChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetClientCOV"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetClientCOV")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetClientCOV) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetClientCOV, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetClientCOV"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetClientCOV")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetClientCOV"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetClientCOV")
	}
	return nil
}

func (m *_BACnetClientCOV) isBACnetClientCOV() bool {
	return true
}

func (m *_BACnetClientCOV) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
