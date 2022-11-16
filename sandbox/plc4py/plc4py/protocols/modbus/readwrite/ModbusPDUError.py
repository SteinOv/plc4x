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
from plc4py.protocols.modbus.readwrite.ModbusErrorCode import ModbusErrorCode
import math


@dataclass
class ModbusPDUError(PlcMessage, ModbusPDU):
    exceptionCode: ModbusErrorCode

    # Accessors for discriminator values.
    def getErrorFlag(self) -> c_bool:
        return c_bool(True)

    def getFunctionFlag(self) -> c_uint8:
        return 0

    def getResponse(self) -> c_bool:
        return False

    def __post_init__(self):
        super().__init__()

    def getExceptionCode(self) -> ModbusErrorCode:
        return self.exceptionCode

    def serializeModbusPDUChild(self, writeBuffer: WriteBuffer):
        positionAware: PositionAware = writeBuffer
        startPos: int = positionAware.getPos()
        writeBuffer.pushContext("ModbusPDUError")

        # Simple Field (exceptionCode)
        writeSimpleEnumField(
            "exceptionCode",
            "ModbusErrorCode",
            exceptionCode,
            DataWriterEnumDefault(
                ModbusErrorCode.getValue,
                ModbusErrorCode.name,
                writeUnsignedShort(writeBuffer, 8),
            ),
        )

        writeBuffer.popContext("ModbusPDUError")

    def getLengthInBytes(self) -> int:
        return int(math.ceil(float(self.getLengthInBits() / 8.0)))

    def getLengthInBits(self) -> int:
        lengthInBits: int = super().getLengthInBits()
        _value: ModbusPDUError = self

        # Simple field (exceptionCode)
        lengthInBits += 8

        return lengthInBits

    @staticmethod
    def staticParseBuilder(
        readBuffer: ReadBuffer, response: c_bool
    ) -> ModbusPDUErrorBuilder:
        readBuffer.pullContext("ModbusPDUError")
        positionAware: PositionAware = readBuffer
        startPos: int = positionAware.getPos()
        curPos: int = 0

        exceptionCode: ModbusErrorCode = readEnumField(
            "exceptionCode",
            "ModbusErrorCode",
            DataReaderEnumDefault(
                ModbusErrorCode.enumForValue, readUnsignedShort(readBuffer, 8)
            ),
        )

        readBuffer.closeContext("ModbusPDUError")
        # Create the instance
        return ModbusPDUErrorBuilder(exceptionCode)

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUError):
            return False

        that: ModbusPDUError = ModbusPDUError(o)
        return (
            (self.getExceptionCode() == that.getExceptionCode())
            and super().equals(that)
            and True
        )

    def hashCode(self) -> int:
        return hash(super().hashCode(), self.getExceptionCode())

    def __str__(self) -> str:
        writeBufferBoxBased: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            writeBufferBoxBased.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(writeBufferBoxBased.getBox()) + "\n"


@dataclass
class ModbusPDUErrorBuilder(ModbusPDUModbusPDUBuilder):
    exceptionCode: ModbusErrorCode

    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUError:
        modbusPDUError: ModbusPDUError = ModbusPDUError(self.exceptionCode)
        return modbusPDUError
