package com.exasol.testsetupabstraction.server;

import org.itsallcode.io.Capturable;
import org.itsallcode.junit.sysextensions.SystemOutGuard;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;

@ExtendWith(SystemOutGuard.class)
class MainTest {
    @Test
    @SuppressWarnings({ "java:S2699", "2925" }) // no assertions required, sleep is ok here
    void test(final Capturable stream) throws InterruptedException {
        stream.capture();
        final Thread threadRunningMainMethod = new Thread(() -> Main.main(new String[] { "unknown" }));
        threadRunningMainMethod.start();
        int counter = 0;
        while (true) {
            Thread.sleep(500);
            final String captured = stream.getCapturedData();
            if (captured.contains("Server running on port:")) {
                break;
            }
            if (counter++ > 10 * 60) {
                throw new AssertionError("timeout waiting for server to start");
            }
        }
        threadRunningMainMethod.interrupt();
    }
}