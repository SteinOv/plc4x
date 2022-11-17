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
from dataclasses import dataclass

from ctypes import c_uint16
from plc4py.api.messages.PlcMessage import PlcMessage
import math


@dataclass
class ModbusConstants(PlcMessage):
    MODBUSTCPDEFAULTPORT: c_uint16 = 502

    def __post_init__(self):
        super().__init__()

    def serialize(self, write_buffer: WriteBuffer):
        position_aware: PositionAware = write_buffer
        start_pos: int = position_aware.get_pos()
        write_buffer.push_context("ModbusConstants")

        # Const Field (modbusTcpDefaultPort)
        write_const_field(
            "modbusTcpDefaultPort",
            self.modbus_tcp_default_port,
            write_unsigned_int(write_buffer, 16),
        )

        write_buffer.pop_context("ModbusConstants")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.get_length_in_bits() / 8.0)))

    def get_length_in_bits(self) -> int:
        length_in_bits: int = 0
        _value: ModbusConstants = self

        # Const Field (modbusTcpDefaultPort)
        length_in_bits += 16

        return length_in_bits

    def static_parse(read_buffer: ReadBuffer, args):
        position_aware: PositionAware = read_buffer
        return staticParse(read_buffer)

    @staticmethod
    def static_parse_context(read_buffer: ReadBuffer):
        read_buffer.pull_context("ModbusConstants")
        position_aware: PositionAware = read_buffer
        start_pos: int = position_aware.get_pos()
        cur_pos: int = 0

        modbus_tcp_default_port: c_uint16 = read_const_field(
            "modbusTcpDefaultPort",
            read_unsigned_int(read_buffer, 16),
            ModbusConstants.MODBUSTCPDEFAULTPORT,
        )

        read_buffer.close_context("ModbusConstants")
        # Create the instance
        _modbus_constants: ModbusConstants = ModbusConstants()
        return _modbus_constants

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusConstants):
            return False

        that: ModbusConstants = ModbusConstants(o)
        return True

    def hash_code(self) -> int:
        return hash(self)

    def __str__(self) -> str:
        write_buffer_box_based: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            write_buffer_box_based.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(write_buffer_box_based.get_box()) + "\n"
