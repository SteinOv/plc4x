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

package org.apache.plc4x.java.profinet.gsdml;

import org.apache.commons.lang3.NotImplementedException;
import org.apache.plc4x.java.profinet.readwrite.PnIoCm_Block_ExpectedSubmoduleReq;
import org.apache.plc4x.java.profinet.readwrite.PnIoCm_IoCs;
import org.apache.plc4x.java.profinet.readwrite.PnIoCm_IoDataObject;

import java.util.Collection;

public class ProfinetModule {

    public Collection<PnIoCm_IoDataObject> getInputIoDataApiBlocks(Integer startingOffset) {
        throw new NotImplementedException("Not yet Implemented");
    }

    public Collection<PnIoCm_IoCs> getOutputIoCsApiBlocks(Integer startingOffset) {
        throw new NotImplementedException("Not yet Implemented");
    }

    public Collection<PnIoCm_Block_ExpectedSubmoduleReq> getExpectedSubmoduleReq() {
        throw new NotImplementedException("Not yet Implemented");
    }

    public Collection<PnIoCm_IoCs> getInputIoCsApiBlocks(Integer startingOffset) {
        throw new NotImplementedException("Not yet Implemented");
    }

    public Collection<PnIoCm_IoDataObject> getOutputIoDataApiBlocks(Integer startingOffset) {
        throw new NotImplementedException("Not yet Implemented");
    }

}
