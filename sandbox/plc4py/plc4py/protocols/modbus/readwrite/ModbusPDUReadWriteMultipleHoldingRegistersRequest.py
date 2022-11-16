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
from ctypes import c_byte
from ctypes import c_uint16
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
from typing import List
import math


@dataclass
class ModbusPDUReadWriteMultipleHoldingRegistersRequest(PlcMessage, ModbusPDU):
    readStartingAddress: c_uint16
    readQuantity: c_uint16
    writeStartingAddress: c_uint16
    writeQuantity: c_uint16
    value: List[c_byte]

    # Accessors for discriminator values.
    def getErrorFlag(self) -> c_bool:
        return c_bool(False)

    def getFunctionFlag(self) -> c_uint8:
        return c_uint8(0x17)

    def getResponse(self) -> c_bool:
        return c_bool(False)

    def __post_init__(self):
        super().__init__()

    def getReadStartingAddress(self) -> c_uint16:
        return self.readStartingAddress

    def getReadQuantity(self) -> c_uint16:
        return self.readQuantity

    def getWriteStartingAddress(self) -> c_uint16:
        return self.writeStartingAddress

    def getWriteQuantity(self) -> c_uint16:
        return self.writeQuantity

    def getValue(self) -> List[c_byte]:
        return self.value

    def serializeModbusPDUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
        startPos: int = positionAware.getPos()
        writeBuffer.pushContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest")

        # Simple Field (readStartingAddress)
        writeSimpleField(
            "readStartingAddress",
            readStartingAddress,
            writeUnsignedInt(writeBuffer, 16),
        )

        # Simple Field (readQuantity)
        writeSimpleField(
            "readQuantity", readQuantity, writeUnsignedInt(writeBuffer, 16)
        )

        # Simple Field (writeStartingAddress)
        writeSimpleField(
            "writeStartingAddress",
            writeStartingAddress,
            writeUnsignedInt(writeBuffer, 16),
        )

        # Simple Field (writeQuantity)
        writeSimpleField(
            "writeQuantity", writeQuantity, writeUnsignedInt(writeBuffer, 16)
        )

        # Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
        byteCount: c_uint8 = c_uint8((COUNT(self.getValue())))
        writeImplicitField("byteCount", byteCount, writeUnsignedShort(writeBuffer, 8))

        # Array Field (value)
        writeByteArrayField("value", value, writeByteArray(writeBuffer, 8))

        writeBuffer.popContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest")

    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(self.getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusPDUReadWriteMultipleHoldingRegistersRequest = self

        # Simple field (readStartingAddress)
        lengthInBits += 16

        # Simple field (readQuantity)
        lengthInBits += 16

        # Simple field (writeStartingAddress)
        lengthInBits += 16

        # Simple field (writeQuantity)
        lengthInBits += 16

        # Implicit Field (byteCount)
        lengthInBits += 8

        # Array field
        if self.value is not None:
            lengthInBits += 8 * self.value.length

        return lengthInBits

    @staticmethod
    def staticParseBuilder(
        readBuffer: ReadBuffer, response: c_bool
    ) -> ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder:
        readBuffer.pullContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

        readStartingAddress: c_uint16 = readSimpleField(
            "readStartingAddress", readUnsignedInt(readBuffer, 16)
        )

        readQuantity: c_uint16 = readSimpleField(
            "readQuantity", readUnsignedInt(readBuffer, 16)
        )

        writeStartingAddress: c_uint16 = readSimpleField(
            "writeStartingAddress", readUnsignedInt(readBuffer, 16)
        )

        writeQuantity: c_uint16 = readSimpleField(
            "writeQuantity", readUnsignedInt(readBuffer, 16)
        )

        byteCount: c_uint8 = readImplicitField(
            "byteCount", readUnsignedShort(readBuffer, 8)
        )

        value: List[byte] = readBuffer.readByteArray("value", int(byteCount))

        readBuffer.closeContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest")
        # Create the instance
        return ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder(
            readStartingAddress,
            readQuantity,
            writeStartingAddress,
            writeQuantity,
            value,
        )

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUReadWriteMultipleHoldingRegistersRequest):
            return False

        that: ModbusPDUReadWriteMultipleHoldingRegistersRequest = (
            ModbusPDUReadWriteMultipleHoldingRegistersRequest(o)
        )
        return (
            (self.getReadStartingAddress() == that.getReadStartingAddress())
            and (self.getReadQuantity() == that.getReadQuantity())
            and (self.getWriteStartingAddress() == that.getWriteStartingAddress())
            and (self.getWriteQuantity() == that.getWriteQuantity())
            and (self.getValue() == that.getValue())
            and super().equals(that)
            and True
        )

    def hashCode(self) -> int:
        return hash(
            super().hashCode(),
            self.getReadStartingAddress(),
            self.getReadQuantity(),
            self.getWriteStartingAddress(),
            self.getWriteQuantity(),
            self.getValue(),
        )

    def __str__(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            writeBufferBoxBased.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(writeBufferBoxBased.getBox()) + "\n"


@dataclass
class ModbusPDUReadWriteMultipleHoldingRegistersRequestBuilder(
    ModbusPDUModbusPDUBuilder
):
    readStartingAddress: c_uint16
    readQuantity: c_uint16
    writeStartingAddress: c_uint16
    writeQuantity: c_uint16
    value: List[c_byte]

    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUReadWriteMultipleHoldingRegistersRequest:
        modbusPDUReadWriteMultipleHoldingRegistersRequest: ModbusPDUReadWriteMultipleHoldingRegistersRequest = ModbusPDUReadWriteMultipleHoldingRegistersRequest(
            self.readStartingAddress,
            self.readQuantity,
            self.writeStartingAddress,
            self.writeQuantity,
            self.value,
        )
        return modbusPDUReadWriteMultipleHoldingRegistersRequest
