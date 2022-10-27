package com.exasol.testsetupabstraction.server;

import java.io.*;
import java.lang.reflect.Type;
import java.nio.charset.StandardCharsets;

import org.jetbrains.annotations.NotNull;

import com.exasol.errorreporting.ExaError;

import io.javalin.json.JsonMapper;
import jakarta.json.bind.Jsonb;
import jakarta.json.bind.JsonbBuilder;
import jakarta.json.bind.spi.JsonbProvider;

class JsonBJsonMapper implements JsonMapper {
    private static final JsonbProvider JSONB_PROVIDER = JsonbProvider.provider();

    @NotNull
    @Override
    public String toJsonString(@NotNull final Object obj, Type type) {
        try (final Jsonb jsonb = JsonbBuilder.create()) {
            return jsonb.toJson(obj, type);
        } catch (final Exception exception) {
            throw new IllegalStateException(ExaError.messageBuilder("F-ETSAS-1")
                    .message("Failed to JSON serialize response.").ticketMitigation().toString());
        }
    }

    @NotNull
    @Override
    public InputStream toJsonStream(@NotNull final Object obj, Type type) {
        return new ByteArrayInputStream(toJsonString(obj, type).getBytes(StandardCharsets.UTF_8));
    }

    @NotNull
    @Override
    public <T> T fromJsonString(@NotNull final String json, @NotNull final Type type) {
        try (final Jsonb jsonb = JSONB_PROVIDER.create().build()) {
            return jsonb.fromJson(json, type);
        } catch (final Exception exception) {
            throw new IllegalStateException(ExaError.messageBuilder("F-ETSAS-2")
                    .message("Failed to JSON deserialize message.").ticketMitigation().toString());
        }
    }

    @NotNull
    @Override
    public <T> T fromJsonStream(@NotNull final InputStream json, @NotNull final Type type) {
        try {
            final String jsonString = new String(json.readAllBytes(), StandardCharsets.UTF_8);
            return this.fromJsonString(jsonString, type);
        } catch (final IOException e) {
            throw new IllegalStateException(ExaError.messageBuilder("F-ETSAS-3").message("Failed to read JSON message.")
                    .ticketMitigation().toString());
        }
    }
}
