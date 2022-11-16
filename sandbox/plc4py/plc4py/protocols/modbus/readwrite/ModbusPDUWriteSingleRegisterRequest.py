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
import math


@dataclass
class ModbusPDUWriteSingleRegisterRequest(PlcMessage, ModbusPDU):
    address: c_uint16
    value: c_uint16

    # Accessors for discriminator values.
    def getErrorFlag(self) -> c_bool:
        return c_bool(False)

    def getFunctionFlag(self) -> c_uint8:
        return c_uint8(0x06)

    def getResponse(self) -> c_bool:
        return c_bool(False)

    def __post_init__(self):
        super().__init__()

    def getAddress(self) -> c_uint16:
        return self.address

    def getValue(self) -> c_uint16:
        return self.value

    def serializeModbusPDUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
        startPos: int = positionAware.getPos()
        writeBuffer.pushContext("ModbusPDUWriteSingleRegisterRequest")

        # Simple Field (address)
        writeSimpleField("address", address, writeUnsignedInt(writeBuffer, 16))

        # Simple Field (value)
        writeSimpleField("value", value, writeUnsignedInt(writeBuffer, 16))

        writeBuffer.popContext("ModbusPDUWriteSingleRegisterRequest")

    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(self.getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusPDUWriteSingleRegisterRequest = self

        # Simple field (address)
        lengthInBits += 16

        # Simple field (value)
        lengthInBits += 16

        return lengthInBits

    @staticmethod
    def staticParseBuilder(
        readBuffer: ReadBuffer, response: c_bool
    ) -> ModbusPDUWriteSingleRegisterRequestBuilder:
        readBuffer.pullContext("ModbusPDUWriteSingleRegisterRequest")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

        address: c_uint16 = readSimpleField("address", readUnsignedInt(readBuffer, 16))

        value: c_uint16 = readSimpleField("value", readUnsignedInt(readBuffer, 16))

        readBuffer.closeContext("ModbusPDUWriteSingleRegisterRequest")
        # Create the instance
        return ModbusPDUWriteSingleRegisterRequestBuilder(address, value)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUWriteSingleRegisterRequest):
            return False

        that: ModbusPDUWriteSingleRegisterRequest = ModbusPDUWriteSingleRegisterRequest(
            o
        )
        return (
            (self.getAddress() == that.getAddress())
            and (self.getValue() == that.getValue())
            and super().equals(that)
            and True
        )

    def hashCode(self) -> int:
        return hash(super().hashCode(), self.getAddress(), self.getValue())

    def __str__(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            writeBufferBoxBased.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(writeBufferBoxBased.getBox()) + "\n"


@dataclass
class ModbusPDUWriteSingleRegisterRequestBuilder(ModbusPDUModbusPDUBuilder):
    address: c_uint16
    value: c_uint16

    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUWriteSingleRegisterRequest:
        modbusPDUWriteSingleRegisterRequest: ModbusPDUWriteSingleRegisterRequest = (
            ModbusPDUWriteSingleRegisterRequest(self.address, self.value)
        )
        return modbusPDUWriteSingleRegisterRequest
