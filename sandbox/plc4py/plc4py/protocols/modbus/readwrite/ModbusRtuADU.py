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
from ctypes import c_uint16
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.modbus.readwrite.DriverType import DriverType
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDU
import math


@dataclass
class ModbusRtuADU(PlcMessage, ModbusADU):
    address: c_uint8
    pdu: ModbusPDU
    # Arguments.
    response: c_bool

    # Accessors for discriminator values.
    def getDriverType(self) -> DriverType:
        return DriverType.MODBUS_RTU

    def __post_init__(self):
        super().__init__(self.response)

    def getAddress(self) -> c_uint8:
        return self.address

    def getPdu(self) -> ModbusPDU:
        return self.pdu

    def serializeModbusADUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
        startPos: int = positionAware.getPos()
        writeBuffer.pushContext("ModbusRtuADU")

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
            (c_uint16)(modbus.readwrite.utils.StaticHelper.rtuCrcCheck(address, pdu)),
            writeUnsignedInt(writeBuffer, 16),
        )

        writeBuffer.popContext("ModbusRtuADU")

    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(self.getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusRtuADU = self

        # Simple field (address)
        lengthInBits += 8

        # Simple field (pdu)
        lengthInBits += pdu.getLengthInBits()

        # Checksum Field (checksum)
        lengthInBits += 16

        return lengthInBits

    @staticmethod
    def staticParseBuilder(
        readBuffer: ReadBuffer, driverType: DriverType, response: c_bool
    ) -> ModbusRtuADUBuilder:
        readBuffer.pullContext("ModbusRtuADU")
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

        crc: c_uint16 = readChecksumField(
            "crc",
            readUnsignedInt(readBuffer, 16),
            (c_uint16)(modbus.readwrite.utils.StaticHelper.rtuCrcCheck(address, pdu)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN),
        )

        readBuffer.closeContext("ModbusRtuADU")
        # Create the instance
        return ModbusRtuADUBuilder(address, pdu, response)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusRtuADU):
            return False

        that: ModbusRtuADU = ModbusRtuADU(o)
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
class ModbusRtuADUBuilder(ModbusADUModbusADUBuilder):
    address: c_uint8
    pdu: ModbusPDU
    response: c_bool

    def __post_init__(self):
        pass

    def build(self, response: c_bool) -> ModbusRtuADU:
        modbusRtuADU: ModbusRtuADU = ModbusRtuADU(self.address, self.pdu, response)
        return modbusRtuADU
