package com.exasol.testsetupabstraction.server;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.Matchers.equalTo;

import java.io.*;
import java.nio.charset.StandardCharsets;
import java.util.Map;

import org.junit.jupiter.api.Test;

class JsonBJsonMapperTest {
    @Test
    void testToJsonString() {
        final String result = new JsonBJsonMapper().toJsonString(Map.of("test", true), Map.class);
        assertThat(result, equalTo("{\"test\":true}"));
    }

    @Test
    void testToJsonStream() throws IOException {
        try (final InputStream stream = new JsonBJsonMapper().toJsonStream(Map.of("test", true), Map.class)) {
            final String result = new String(stream.readAllBytes(), StandardCharsets.UTF_8);
            assertThat(result, equalTo("{\"test\":true}"));
        }
    }

    @Test
    void testFromJsonString() {
        final Map<?, ?> result = new JsonBJsonMapper().fromJsonString("{\"test\":true}", Map.class);
        assertThat(result, equalTo(Map.of("test", true)));
    }

    @Test
    void testFromJsonStream() throws IOException {
        try (final ByteArrayInputStream inputStream = new ByteArrayInputStream(
                "{\"test\":true}".getBytes(StandardCharsets.UTF_8))) {
            final Map<?, ?> result = new JsonBJsonMapper().fromJsonStream(inputStream, Map.class);
            assertThat(result, equalTo(Map.of("test", true)));
        }
    }
}