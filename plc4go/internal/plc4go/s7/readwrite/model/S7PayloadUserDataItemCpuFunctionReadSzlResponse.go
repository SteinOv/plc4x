//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "encoding/xml"
    "errors"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
    "strconv"
)

// Constant values.
const S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH uint16 = 28

// The data-structure of this message
type S7PayloadUserDataItemCpuFunctionReadSzlResponse struct {
    Items []*SzlDataTreeItem
    Parent *S7PayloadUserDataItem
    IS7PayloadUserDataItemCpuFunctionReadSzlResponse
}

// The corresponding interface
type IS7PayloadUserDataItemCpuFunctionReadSzlResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) CpuFunctionType() uint8 {
    return 0x08
}


func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, szlId *SzlId, szlIndex uint16) {
    m.Parent.ReturnCode = returnCode
    m.Parent.TransportSize = transportSize
    m.Parent.SzlId = szlId
    m.Parent.SzlIndex = szlIndex
}

func NewS7PayloadUserDataItemCpuFunctionReadSzlResponse(items []*SzlDataTreeItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, szlId *SzlId, szlIndex uint16) *S7PayloadUserDataItem {
    child := &S7PayloadUserDataItemCpuFunctionReadSzlResponse{
        Items: items,
        Parent: NewS7PayloadUserDataItem(returnCode, transportSize, szlId, szlIndex),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(structType interface{}) S7PayloadUserDataItemCpuFunctionReadSzlResponse {
    castFunc := func(typ interface{}) S7PayloadUserDataItemCpuFunctionReadSzlResponse {
        if casted, ok := typ.(S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
            return casted
        }
        if casted, ok := typ.(*S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
            return *casted
        }
        if casted, ok := typ.(S7PayloadUserDataItem); ok {
            return CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(casted.Child)
        }
        if casted, ok := typ.(*S7PayloadUserDataItem); ok {
            return CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(casted.Child)
        }
        return S7PayloadUserDataItemCpuFunctionReadSzlResponse{}
    }
    return castFunc(structType)
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Const Field (szlItemLength)
    lengthInBits += 16

    // Implicit Field (szlItemCount)
    lengthInBits += 16

    // Array field
    if len(m.Items) > 0 {
        for _, element := range m.Items {
            lengthInBits += element.LengthInBits()
        }
    }

    return lengthInBits
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionReadSzlResponseParse(io *utils.ReadBuffer) (*S7PayloadUserDataItem, error) {

    // Const Field (szlItemLength)
    szlItemLength, _szlItemLengthErr := io.ReadUint16(16)
    if _szlItemLengthErr != nil {
        return nil, errors.New("Error parsing 'szlItemLength' field " + _szlItemLengthErr.Error())
    }
    if szlItemLength != S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH)) + " but got " + strconv.Itoa(int(szlItemLength)))
    }

    // Implicit Field (szlItemCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    szlItemCount, _szlItemCountErr := io.ReadUint16(16)
    if _szlItemCountErr != nil {
        return nil, errors.New("Error parsing 'szlItemCount' field " + _szlItemCountErr.Error())
    }

    // Array field (items)
    // Count array
    items := make([]*SzlDataTreeItem, szlItemCount)
    for curItem := uint16(0); curItem < uint16(szlItemCount); curItem++ {
        _item, _err := SzlDataTreeItemParse(io)
        if _err != nil {
            return nil, errors.New("Error parsing 'items' field " + _err.Error())
        }
        items[curItem] = _item
    }

    // Create a partially initialized instance
    _child := &S7PayloadUserDataItemCpuFunctionReadSzlResponse{
        Items: items,
        Parent: &S7PayloadUserDataItem{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Const Field (szlItemLength)
    _szlItemLengthErr := io.WriteUint16(16, 28)
    if _szlItemLengthErr != nil {
        return errors.New("Error serializing 'szlItemLength' field " + _szlItemLengthErr.Error())
    }

    // Implicit Field (szlItemCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    szlItemCount := uint16(uint16(len(m.Items)))
    _szlItemCountErr := io.WriteUint16(16, (szlItemCount))
    if _szlItemCountErr != nil {
        return errors.New("Error serializing 'szlItemCount' field " + _szlItemCountErr.Error())
    }

    // Array Field (items)
    if m.Items != nil {
        for _, _element := range m.Items {
            _elementErr := _element.Serialize(io)
            if _elementErr != nil {
                return errors.New("Error serializing 'items' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "items":
                var data []*SzlDataTreeItem
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Items = data
            }
        }
    }
}

func (m *S7PayloadUserDataItemCpuFunctionReadSzlResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.s7.readwrite.S7PayloadUserDataItemCpuFunctionReadSzlResponse"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Items, xml.StartElement{Name: xml.Name{Local: "items"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "items"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

