package com.exasol.testsetupabstraction.server;

import org.itsallcode.io.Capturable;
import org.itsallcode.junit.sysextensions.SystemOutGuard;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;

@ExtendWith(SystemOutGuard.class)
class MainIT {
    @Test
    void testWithArgument(final Capturable stream) throws InterruptedException {
        stream.capture();
        final Thread threadRunningMainMethod = startMainMethod("unknown");
        try {
            assertServerLogsStartupMessage(stream);
        } finally {
            threadRunningMainMethod.interrupt();
        }
    }

    @Test
    void testWithoutArgument(final Capturable stream) throws InterruptedException {
        stream.capture();
        final Thread threadRunningMainMethod = startMainMethod();
        try {
            assertServerLogsStartupMessage(stream);
        } finally {
            threadRunningMainMethod.interrupt();
        }
    }

    private Thread startMainMethod(String... args) {
        final Thread threadRunningMainMethod = new Thread(() -> Main.main(args));
        threadRunningMainMethod.start();
        return threadRunningMainMethod;
    }

    @SuppressWarnings("java:S2925") // Sleep is required for waiting for output
    private void assertServerLogsStartupMessage(final Capturable stream) throws InterruptedException, AssertionError {
        int counter = 0;
        while (true) {
            Thread.sleep(500);
            final String captured = stream.getCapturedData();
            if (captured.contains("Server running on port:")) {
                break;
            }
            if (counter++ > 10 * 2 * 60) {// 10 mins
                throw new AssertionError("timeout waiting for server to start");
            }
        }
    }
}