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

// SALDataSecurity is the corresponding interface of SALDataSecurity
type SALDataSecurity interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetSecurityData returns SecurityData (property field)
	GetSecurityData() SecurityData
}

// SALDataSecurityExactly can be used when we want exactly this type and not a type which fulfills SALDataSecurity.
// This is useful for switch cases.
type SALDataSecurityExactly interface {
	SALDataSecurity
	isSALDataSecurity() bool
}

// _SALDataSecurity is the data-structure of this message
type _SALDataSecurity struct {
	*_SALData
	SecurityData SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataSecurity) GetApplicationId() ApplicationId {
	return ApplicationId_SECURITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataSecurity) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataSecurity) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataSecurity) GetSecurityData() SecurityData {
	return m.SecurityData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataSecurity factory function for _SALDataSecurity
func NewSALDataSecurity(securityData SecurityData, salData SALData) *_SALDataSecurity {
	_result := &_SALDataSecurity{
		SecurityData: securityData,
		_SALData:     NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataSecurity(structType interface{}) SALDataSecurity {
	if casted, ok := structType.(SALDataSecurity); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataSecurity); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataSecurity) GetTypeName() string {
	return "SALDataSecurity"
}

func (m *_SALDataSecurity) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataSecurity) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (securityData)
	lengthInBits += m.SecurityData.GetLengthInBits()

	return lengthInBits
}

func (m *_SALDataSecurity) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataSecurityParse(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataSecurity, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataSecurity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataSecurity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (securityData)
	if pullErr := readBuffer.PullContext("securityData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityData")
	}
	_securityData, _securityDataErr := SecurityDataParse(readBuffer)
	if _securityDataErr != nil {
		return nil, errors.Wrap(_securityDataErr, "Error parsing 'securityData' field of SALDataSecurity")
	}
	securityData := _securityData.(SecurityData)
	if closeErr := readBuffer.CloseContext("securityData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityData")
	}

	if closeErr := readBuffer.CloseContext("SALDataSecurity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataSecurity")
	}

	// Create a partially initialized instance
	_child := &_SALDataSecurity{
		_SALData:     &_SALData{},
		SecurityData: securityData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataSecurity) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataSecurity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataSecurity")
		}

		// Simple Field (securityData)
		if pushErr := writeBuffer.PushContext("securityData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityData")
		}
		_securityDataErr := writeBuffer.WriteSerializable(m.GetSecurityData())
		if popErr := writeBuffer.PopContext("securityData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityData")
		}
		if _securityDataErr != nil {
			return errors.Wrap(_securityDataErr, "Error serializing 'securityData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataSecurity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataSecurity")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SALDataSecurity) isSALDataSecurity() bool {
	return true
}

func (m *_SALDataSecurity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
