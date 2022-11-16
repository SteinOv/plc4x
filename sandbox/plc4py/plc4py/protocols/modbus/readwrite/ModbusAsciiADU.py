#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License") you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   https:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

# Code generated by code-generation. DO NOT EDIT.
from abc import ABC, abstractmethod
from dataclasses import dataclass


from ctypes import c_bool
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.modbus.readwrite.DriverType import DriverType
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDU
import math


@dataclass
class ModbusAsciiADU(PlcMessage, ModbusADU):
    address: c_uint8
    pdu: ModbusPDU
    # Arguments.
    response: c_bool

    # Accessors for discriminator values.
    def getDriverType(self) -> DriverType:
        return DriverType.MODBUS_ASCII

    def __post_init__(self):
        super().__init__(self.response)

    def getAddress(self) -> c_uint8:
        return self.address

    def getPdu(self) -> ModbusPDU:
        return self.pdu

    def serializeModbusADUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
        startPos: int = positionAware.getPos()
        writeBuffer.pushContext("ModbusAsciiADU")

        # Simple Field (address)
        writeSimpleField(
            "address",
            address,
            writeUnsignedShort(writeBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN),
        )

        # Simple Field (pdu)
        writeSimpleField(
            "pdu",
            pdu,
            DataWriterComplexDefault(writeBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN),
        )

        # Checksum Field (checksum) (Calculated)
        writeChecksumField(
            "crc",
            (c_uint8)(modbus.readwrite.utils.StaticHelper.asciiLrcCheck(address, pdu)),
            writeUnsignedShort(writeBuffer, 8),
        )

        writeBuffer.popContext("ModbusAsciiADU")

    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(self.getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusAsciiADU = self

        # Simple field (address)
        lengthInBits += 8

        # Simple field (pdu)
        lengthInBits += pdu.getLengthInBits()

        # Checksum Field (checksum)
        lengthInBits += 8

        return lengthInBits

    @staticmethod
    def staticParseBuilder(
        readBuffer: ReadBuffer, driverType: DriverType, response: c_bool
    ) -> ModbusAsciiADUBuilder:
        readBuffer.pullContext("ModbusAsciiADU")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

        address: c_uint8 = readSimpleField(
            "address",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN),
        )

        pdu: ModbusPDU = readSimpleField(
            "pdu",
            DataReaderComplexDefault(
                ModbusPDU.staticParse(readBuffer, c_bool(response)), readBuffer
            ),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN),
        )

        crc: c_uint8 = readChecksumField(
            "crc",
            readUnsignedShort(readBuffer, 8),
            (c_uint8)(modbus.readwrite.utils.StaticHelper.asciiLrcCheck(address, pdu)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN),
        )

        readBuffer.closeContext("ModbusAsciiADU")
        # Create the instance
        return ModbusAsciiADUBuilder(address, pdu, response)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusAsciiADU):
            return False

        that: ModbusAsciiADU = ModbusAsciiADU(o)
        return (
            (self.getAddress() == that.getAddress())
            and (self.getPdu() == that.getPdu())
            and super().equals(that)
            and True
        )

    def hashCode(self) -> int:
        return hash(super().hashCode(), self.getAddress(), self.getPdu())

    def __str__(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            writeBufferBoxBased.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(writeBufferBoxBased.getBox()) + "\n"


@dataclass
class ModbusAsciiADUBuilder(ModbusADUModbusADUBuilder):
    address: c_uint8
    pdu: ModbusPDU
    response: c_bool

    def __post_init__(self):
        pass

    def build(self, response: c_bool) -> ModbusAsciiADU:
        modbusAsciiADU: ModbusAsciiADU = ModbusAsciiADU(
            self.address, self.pdu, response
        )
        return modbusAsciiADU
