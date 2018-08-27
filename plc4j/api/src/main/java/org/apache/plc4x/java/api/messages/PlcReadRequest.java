/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.api.messages;

import org.apache.plc4x.java.api.messages.items.PlcReadRequestItem;
import org.apache.plc4x.java.api.messages.specific.TypeSafePlcReadRequest;
import org.apache.plc4x.java.api.model.PlcField;

import java.util.List;
import java.util.Objects;

/**
 * Request to read one or more values from a plc.
 * The {@link PlcReadRequest} is a container for one or more {@link PlcReadRequestItem}s that are contained.
 * <p>
 * By default this is NOT typesafe as it can contain {@link PlcReadRequestItem}s for different types.
 * If there are only {@link PlcReadRequestItem}s of the same type one can use the {@link TypeSafePlcReadRequest} to enforce
 * type safety.
 * <p>
 * Provides a builder for object construction through {@link PlcReadRequest#builder()}.
 *
 * @see TypeSafePlcReadRequest
 */
public class PlcReadRequest extends PlcRequest<PlcReadRequestItem<?>> {

    public PlcReadRequest() {
        // Exists for construction of inherited classes
    }

    public PlcReadRequest(PlcReadRequestItem<?> requestItem) {
        addItem(requestItem);
    }

    public PlcReadRequest(Class<?> dataType, PlcField field) {
        addItem(new PlcReadRequestItem<>(dataType, field));
    }

    public PlcReadRequest(Class<?> dataType, PlcField field, int size) {
        Objects.requireNonNull(dataType, "Data type must not be null");
        Objects.requireNonNull(field, "PlcField must not be null");
        addItem(new PlcReadRequestItem<>(dataType, field, size));
    }

    public PlcReadRequest(List<PlcReadRequestItem<?>> requestItems) {
        super(requestItems);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static class Builder extends PlcRequest.Builder<PlcReadRequestItem<?>> {

        public final Builder addItem(Class<?> dataType, PlcField field) {
            checkType(dataType);
            requests.add(new PlcReadRequestItem<>(dataType, field));
            return this;
        }

        public final Builder addItem(Class<?> dataType, PlcField field, int size) {
            checkType(dataType);
            requests.add(new PlcReadRequestItem<>(dataType, field, size));
            return this;
        }

        public final Builder addItem(PlcReadRequestItem readRequestItem) {
            checkType(readRequestItem.getDatatype());
            requests.add(readRequestItem);
            return this;
        }

        public final PlcReadRequest build() {
            if (requests.isEmpty()) {
                throw new IllegalStateException("No requests added");
            }
            PlcReadRequest plcReadRequest;
            if (mixed) {
                plcReadRequest = new PlcReadRequest();
            } else {
                plcReadRequest = new TypeSafePlcReadRequest<>(firstType);
            }
            for (PlcReadRequestItem request : requests) {
                plcReadRequest.addItem(request);
            }
            return plcReadRequest;
        }

        @SuppressWarnings("unchecked")
        public final <T> TypeSafePlcReadRequest<T> build(Class<T> type) {
            if (firstType != type) {
                throw new ClassCastException("Incompatible type " + type + ". Required " + firstType);
            }
            if (mixed) {
                throw new IllegalStateException("Mixed types contained");
            }
            return (TypeSafePlcReadRequest<T>) build();
        }

    }

    @Override
    public String toString() {
        return "PlcReadRequest{} " + super.toString();
    }
}

