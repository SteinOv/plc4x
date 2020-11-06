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
)

// The data-structure of this message
type TunnelingResponse struct {
    TunnelingResponseDataBlock *TunnelingResponseDataBlock
    Parent *KNXNetIPMessage
    ITunnelingResponse
}

// The corresponding interface
type ITunnelingResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *TunnelingResponse) MsgType() uint16 {
    return 0x0421
}


func (m *TunnelingResponse) InitializeParent(parent *KNXNetIPMessage) {
}

func NewTunnelingResponse(tunnelingResponseDataBlock *TunnelingResponseDataBlock, ) *KNXNetIPMessage {
    child := &TunnelingResponse{
        TunnelingResponseDataBlock: tunnelingResponseDataBlock,
        Parent: NewKNXNetIPMessage(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastTunnelingResponse(structType interface{}) TunnelingResponse {
    castFunc := func(typ interface{}) TunnelingResponse {
        if casted, ok := typ.(TunnelingResponse); ok {
            return casted
        }
        if casted, ok := typ.(*TunnelingResponse); ok {
            return *casted
        }
        if casted, ok := typ.(KNXNetIPMessage); ok {
            return CastTunnelingResponse(casted.Child)
        }
        if casted, ok := typ.(*KNXNetIPMessage); ok {
            return CastTunnelingResponse(casted.Child)
        }
        return TunnelingResponse{}
    }
    return castFunc(structType)
}

func (m *TunnelingResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (tunnelingResponseDataBlock)
    lengthInBits += m.TunnelingResponseDataBlock.LengthInBits()

    return lengthInBits
}

func (m *TunnelingResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func TunnelingResponseParse(io *utils.ReadBuffer) (*KNXNetIPMessage, error) {

    // Simple Field (tunnelingResponseDataBlock)
    tunnelingResponseDataBlock, _tunnelingResponseDataBlockErr := TunnelingResponseDataBlockParse(io)
    if _tunnelingResponseDataBlockErr != nil {
        return nil, errors.New("Error parsing 'tunnelingResponseDataBlock' field " + _tunnelingResponseDataBlockErr.Error())
    }

    // Create a partially initialized instance
    _child := &TunnelingResponse{
        TunnelingResponseDataBlock: tunnelingResponseDataBlock,
        Parent: &KNXNetIPMessage{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *TunnelingResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (tunnelingResponseDataBlock)
    _tunnelingResponseDataBlockErr := m.TunnelingResponseDataBlock.Serialize(io)
    if _tunnelingResponseDataBlockErr != nil {
        return errors.New("Error serializing 'tunnelingResponseDataBlock' field " + _tunnelingResponseDataBlockErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *TunnelingResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "tunnelingResponseDataBlock":
                var data *TunnelingResponseDataBlock
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.TunnelingResponseDataBlock = data
            }
        }
    }
}

func (m *TunnelingResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.TunnelingResponse"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.TunnelingResponseDataBlock, xml.StartElement{Name: xml.Name{Local: "tunnelingResponseDataBlock"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

