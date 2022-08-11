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

// BACnetWeekNDayTagged is the corresponding interface of BACnetWeekNDayTagged
type BACnetWeekNDayTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetMonth returns Month (property field)
	GetMonth() uint8
	// GetWeekOfMonth returns WeekOfMonth (property field)
	GetWeekOfMonth() uint8
	// GetDayOfWeek returns DayOfWeek (property field)
	GetDayOfWeek() uint8
	// GetOddMonths returns OddMonths (virtual field)
	GetOddMonths() bool
	// GetEvenMonths returns EvenMonths (virtual field)
	GetEvenMonths() bool
	// GetAnyMonth returns AnyMonth (virtual field)
	GetAnyMonth() bool
	// GetDays1to7 returns Days1to7 (virtual field)
	GetDays1to7() bool
	// GetDays8to14 returns Days8to14 (virtual field)
	GetDays8to14() bool
	// GetDays15to21 returns Days15to21 (virtual field)
	GetDays15to21() bool
	// GetDays22to28 returns Days22to28 (virtual field)
	GetDays22to28() bool
	// GetDays29to31 returns Days29to31 (virtual field)
	GetDays29to31() bool
	// GetLast7DaysOfThisMonth returns Last7DaysOfThisMonth (virtual field)
	GetLast7DaysOfThisMonth() bool
	// GetAny7DaysPriorToLast7DaysOfThisMonth returns Any7DaysPriorToLast7DaysOfThisMonth (virtual field)
	GetAny7DaysPriorToLast7DaysOfThisMonth() bool
	// GetAny7DaysPriorToLast14DaysOfThisMonth returns Any7DaysPriorToLast14DaysOfThisMonth (virtual field)
	GetAny7DaysPriorToLast14DaysOfThisMonth() bool
	// GetAny7DaysPriorToLast21DaysOfThisMonth returns Any7DaysPriorToLast21DaysOfThisMonth (virtual field)
	GetAny7DaysPriorToLast21DaysOfThisMonth() bool
	// GetAnyWeekOfthisMonth returns AnyWeekOfthisMonth (virtual field)
	GetAnyWeekOfthisMonth() bool
	// GetAnyDayOfWeek returns AnyDayOfWeek (virtual field)
	GetAnyDayOfWeek() bool
}

// BACnetWeekNDayTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetWeekNDayTagged.
// This is useful for switch cases.
type BACnetWeekNDayTaggedExactly interface {
	BACnetWeekNDayTagged
	isBACnetWeekNDayTagged() bool
}

// _BACnetWeekNDayTagged is the data-structure of this message
type _BACnetWeekNDayTagged struct {
	Header      BACnetTagHeader
	Month       uint8
	WeekOfMonth uint8
	DayOfWeek   uint8

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetWeekNDayTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetWeekNDayTagged) GetMonth() uint8 {
	return m.Month
}

func (m *_BACnetWeekNDayTagged) GetWeekOfMonth() uint8 {
	return m.WeekOfMonth
}

func (m *_BACnetWeekNDayTagged) GetDayOfWeek() uint8 {
	return m.DayOfWeek
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetWeekNDayTagged) GetOddMonths() bool {
	return bool(bool((m.GetMonth()) == (13)))
}

func (m *_BACnetWeekNDayTagged) GetEvenMonths() bool {
	return bool(bool((m.GetMonth()) == (14)))
}

func (m *_BACnetWeekNDayTagged) GetAnyMonth() bool {
	return bool(bool((m.GetMonth()) == (0xFF)))
}

func (m *_BACnetWeekNDayTagged) GetDays1to7() bool {
	return bool(bool((m.GetWeekOfMonth()) == (1)))
}

func (m *_BACnetWeekNDayTagged) GetDays8to14() bool {
	return bool(bool((m.GetWeekOfMonth()) == (2)))
}

func (m *_BACnetWeekNDayTagged) GetDays15to21() bool {
	return bool(bool((m.GetWeekOfMonth()) == (3)))
}

func (m *_BACnetWeekNDayTagged) GetDays22to28() bool {
	return bool(bool((m.GetWeekOfMonth()) == (4)))
}

func (m *_BACnetWeekNDayTagged) GetDays29to31() bool {
	return bool(bool((m.GetWeekOfMonth()) == (5)))
}

func (m *_BACnetWeekNDayTagged) GetLast7DaysOfThisMonth() bool {
	return bool(bool((m.GetWeekOfMonth()) == (6)))
}

func (m *_BACnetWeekNDayTagged) GetAny7DaysPriorToLast7DaysOfThisMonth() bool {
	return bool(bool((m.GetWeekOfMonth()) == (7)))
}

func (m *_BACnetWeekNDayTagged) GetAny7DaysPriorToLast14DaysOfThisMonth() bool {
	return bool(bool((m.GetWeekOfMonth()) == (8)))
}

func (m *_BACnetWeekNDayTagged) GetAny7DaysPriorToLast21DaysOfThisMonth() bool {
	return bool(bool((m.GetWeekOfMonth()) == (9)))
}

func (m *_BACnetWeekNDayTagged) GetAnyWeekOfthisMonth() bool {
	return bool(bool((m.GetWeekOfMonth()) == (0xFF)))
}

func (m *_BACnetWeekNDayTagged) GetAnyDayOfWeek() bool {
	return bool(bool((m.GetDayOfWeek()) == (0xFF)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetWeekNDayTagged factory function for _BACnetWeekNDayTagged
func NewBACnetWeekNDayTagged(header BACnetTagHeader, month uint8, weekOfMonth uint8, dayOfWeek uint8, tagNumber uint8, tagClass TagClass) *_BACnetWeekNDayTagged {
	return &_BACnetWeekNDayTagged{Header: header, Month: month, WeekOfMonth: weekOfMonth, DayOfWeek: dayOfWeek, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetWeekNDayTagged(structType interface{}) BACnetWeekNDayTagged {
	if casted, ok := structType.(BACnetWeekNDayTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetWeekNDayTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetWeekNDayTagged) GetTypeName() string {
	return "BACnetWeekNDayTagged"
}

func (m *_BACnetWeekNDayTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetWeekNDayTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Simple field (month)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (weekOfMonth)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (dayOfWeek)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetWeekNDayTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetWeekNDayTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetWeekNDayTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetWeekNDayTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetWeekNDayTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetWeekNDayTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Validation
	if !(bool((header.GetActualLength()) == (3))) {
		return nil, errors.WithStack(utils.ParseValidationError{"We should have at least 3 octets"})
	}

	// Simple Field (month)
	_month, _monthErr := readBuffer.ReadUint8("month", 8)
	if _monthErr != nil {
		return nil, errors.Wrap(_monthErr, "Error parsing 'month' field of BACnetWeekNDayTagged")
	}
	month := _month

	// Virtual field
	_oddMonths := bool((month) == (13))
	oddMonths := bool(_oddMonths)
	_ = oddMonths

	// Virtual field
	_evenMonths := bool((month) == (14))
	evenMonths := bool(_evenMonths)
	_ = evenMonths

	// Virtual field
	_anyMonth := bool((month) == (0xFF))
	anyMonth := bool(_anyMonth)
	_ = anyMonth

	// Simple Field (weekOfMonth)
	_weekOfMonth, _weekOfMonthErr := readBuffer.ReadUint8("weekOfMonth", 8)
	if _weekOfMonthErr != nil {
		return nil, errors.Wrap(_weekOfMonthErr, "Error parsing 'weekOfMonth' field of BACnetWeekNDayTagged")
	}
	weekOfMonth := _weekOfMonth

	// Virtual field
	_days1to7 := bool((weekOfMonth) == (1))
	days1to7 := bool(_days1to7)
	_ = days1to7

	// Virtual field
	_days8to14 := bool((weekOfMonth) == (2))
	days8to14 := bool(_days8to14)
	_ = days8to14

	// Virtual field
	_days15to21 := bool((weekOfMonth) == (3))
	days15to21 := bool(_days15to21)
	_ = days15to21

	// Virtual field
	_days22to28 := bool((weekOfMonth) == (4))
	days22to28 := bool(_days22to28)
	_ = days22to28

	// Virtual field
	_days29to31 := bool((weekOfMonth) == (5))
	days29to31 := bool(_days29to31)
	_ = days29to31

	// Virtual field
	_last7DaysOfThisMonth := bool((weekOfMonth) == (6))
	last7DaysOfThisMonth := bool(_last7DaysOfThisMonth)
	_ = last7DaysOfThisMonth

	// Virtual field
	_any7DaysPriorToLast7DaysOfThisMonth := bool((weekOfMonth) == (7))
	any7DaysPriorToLast7DaysOfThisMonth := bool(_any7DaysPriorToLast7DaysOfThisMonth)
	_ = any7DaysPriorToLast7DaysOfThisMonth

	// Virtual field
	_any7DaysPriorToLast14DaysOfThisMonth := bool((weekOfMonth) == (8))
	any7DaysPriorToLast14DaysOfThisMonth := bool(_any7DaysPriorToLast14DaysOfThisMonth)
	_ = any7DaysPriorToLast14DaysOfThisMonth

	// Virtual field
	_any7DaysPriorToLast21DaysOfThisMonth := bool((weekOfMonth) == (9))
	any7DaysPriorToLast21DaysOfThisMonth := bool(_any7DaysPriorToLast21DaysOfThisMonth)
	_ = any7DaysPriorToLast21DaysOfThisMonth

	// Virtual field
	_anyWeekOfthisMonth := bool((weekOfMonth) == (0xFF))
	anyWeekOfthisMonth := bool(_anyWeekOfthisMonth)
	_ = anyWeekOfthisMonth

	// Simple Field (dayOfWeek)
	_dayOfWeek, _dayOfWeekErr := readBuffer.ReadUint8("dayOfWeek", 8)
	if _dayOfWeekErr != nil {
		return nil, errors.Wrap(_dayOfWeekErr, "Error parsing 'dayOfWeek' field of BACnetWeekNDayTagged")
	}
	dayOfWeek := _dayOfWeek

	// Virtual field
	_anyDayOfWeek := bool((dayOfWeek) == (0xFF))
	anyDayOfWeek := bool(_anyDayOfWeek)
	_ = anyDayOfWeek

	if closeErr := readBuffer.CloseContext("BACnetWeekNDayTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetWeekNDayTagged")
	}

	// Create the instance
	return &_BACnetWeekNDayTagged{
		TagNumber:   tagNumber,
		TagClass:    tagClass,
		Header:      header,
		Month:       month,
		WeekOfMonth: weekOfMonth,
		DayOfWeek:   dayOfWeek,
	}, nil
}

func (m *_BACnetWeekNDayTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetWeekNDayTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetWeekNDayTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Simple Field (month)
	month := uint8(m.GetMonth())
	_monthErr := writeBuffer.WriteUint8("month", 8, (month))
	if _monthErr != nil {
		return errors.Wrap(_monthErr, "Error serializing 'month' field")
	}
	// Virtual field
	if _oddMonthsErr := writeBuffer.WriteVirtual("oddMonths", m.GetOddMonths()); _oddMonthsErr != nil {
		return errors.Wrap(_oddMonthsErr, "Error serializing 'oddMonths' field")
	}
	// Virtual field
	if _evenMonthsErr := writeBuffer.WriteVirtual("evenMonths", m.GetEvenMonths()); _evenMonthsErr != nil {
		return errors.Wrap(_evenMonthsErr, "Error serializing 'evenMonths' field")
	}
	// Virtual field
	if _anyMonthErr := writeBuffer.WriteVirtual("anyMonth", m.GetAnyMonth()); _anyMonthErr != nil {
		return errors.Wrap(_anyMonthErr, "Error serializing 'anyMonth' field")
	}

	// Simple Field (weekOfMonth)
	weekOfMonth := uint8(m.GetWeekOfMonth())
	_weekOfMonthErr := writeBuffer.WriteUint8("weekOfMonth", 8, (weekOfMonth))
	if _weekOfMonthErr != nil {
		return errors.Wrap(_weekOfMonthErr, "Error serializing 'weekOfMonth' field")
	}
	// Virtual field
	if _days1to7Err := writeBuffer.WriteVirtual("days1to7", m.GetDays1to7()); _days1to7Err != nil {
		return errors.Wrap(_days1to7Err, "Error serializing 'days1to7' field")
	}
	// Virtual field
	if _days8to14Err := writeBuffer.WriteVirtual("days8to14", m.GetDays8to14()); _days8to14Err != nil {
		return errors.Wrap(_days8to14Err, "Error serializing 'days8to14' field")
	}
	// Virtual field
	if _days15to21Err := writeBuffer.WriteVirtual("days15to21", m.GetDays15to21()); _days15to21Err != nil {
		return errors.Wrap(_days15to21Err, "Error serializing 'days15to21' field")
	}
	// Virtual field
	if _days22to28Err := writeBuffer.WriteVirtual("days22to28", m.GetDays22to28()); _days22to28Err != nil {
		return errors.Wrap(_days22to28Err, "Error serializing 'days22to28' field")
	}
	// Virtual field
	if _days29to31Err := writeBuffer.WriteVirtual("days29to31", m.GetDays29to31()); _days29to31Err != nil {
		return errors.Wrap(_days29to31Err, "Error serializing 'days29to31' field")
	}
	// Virtual field
	if _last7DaysOfThisMonthErr := writeBuffer.WriteVirtual("last7DaysOfThisMonth", m.GetLast7DaysOfThisMonth()); _last7DaysOfThisMonthErr != nil {
		return errors.Wrap(_last7DaysOfThisMonthErr, "Error serializing 'last7DaysOfThisMonth' field")
	}
	// Virtual field
	if _any7DaysPriorToLast7DaysOfThisMonthErr := writeBuffer.WriteVirtual("any7DaysPriorToLast7DaysOfThisMonth", m.GetAny7DaysPriorToLast7DaysOfThisMonth()); _any7DaysPriorToLast7DaysOfThisMonthErr != nil {
		return errors.Wrap(_any7DaysPriorToLast7DaysOfThisMonthErr, "Error serializing 'any7DaysPriorToLast7DaysOfThisMonth' field")
	}
	// Virtual field
	if _any7DaysPriorToLast14DaysOfThisMonthErr := writeBuffer.WriteVirtual("any7DaysPriorToLast14DaysOfThisMonth", m.GetAny7DaysPriorToLast14DaysOfThisMonth()); _any7DaysPriorToLast14DaysOfThisMonthErr != nil {
		return errors.Wrap(_any7DaysPriorToLast14DaysOfThisMonthErr, "Error serializing 'any7DaysPriorToLast14DaysOfThisMonth' field")
	}
	// Virtual field
	if _any7DaysPriorToLast21DaysOfThisMonthErr := writeBuffer.WriteVirtual("any7DaysPriorToLast21DaysOfThisMonth", m.GetAny7DaysPriorToLast21DaysOfThisMonth()); _any7DaysPriorToLast21DaysOfThisMonthErr != nil {
		return errors.Wrap(_any7DaysPriorToLast21DaysOfThisMonthErr, "Error serializing 'any7DaysPriorToLast21DaysOfThisMonth' field")
	}
	// Virtual field
	if _anyWeekOfthisMonthErr := writeBuffer.WriteVirtual("anyWeekOfthisMonth", m.GetAnyWeekOfthisMonth()); _anyWeekOfthisMonthErr != nil {
		return errors.Wrap(_anyWeekOfthisMonthErr, "Error serializing 'anyWeekOfthisMonth' field")
	}

	// Simple Field (dayOfWeek)
	dayOfWeek := uint8(m.GetDayOfWeek())
	_dayOfWeekErr := writeBuffer.WriteUint8("dayOfWeek", 8, (dayOfWeek))
	if _dayOfWeekErr != nil {
		return errors.Wrap(_dayOfWeekErr, "Error serializing 'dayOfWeek' field")
	}
	// Virtual field
	if _anyDayOfWeekErr := writeBuffer.WriteVirtual("anyDayOfWeek", m.GetAnyDayOfWeek()); _anyDayOfWeekErr != nil {
		return errors.Wrap(_anyDayOfWeekErr, "Error serializing 'anyDayOfWeek' field")
	}

	if popErr := writeBuffer.PopContext("BACnetWeekNDayTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetWeekNDayTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetWeekNDayTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetWeekNDayTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetWeekNDayTagged) isBACnetWeekNDayTagged() bool {
	return true
}

func (m *_BACnetWeekNDayTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
