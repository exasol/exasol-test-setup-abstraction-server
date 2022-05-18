package com.exasol.testsetupabstraction.server;

import java.nio.file.Path;

import com.exasol.exasoltestsetup.ExasolTestSetup;
import com.exasol.exasoltestsetup.ExasolTestSetupFactory;

public class Main {
    @SuppressWarnings({ "java:S106", "java:S2189" }) // we don't want to use a logger here, endlesslopp is intended
    public static void main(final String[] args) {
        final ExasolTestSetup exasol = new ExasolTestSetupFactory(Path.of(args[0])).getTestSetup();
        final int port = 7070;
        final TestSetupServer server = new TestSetupServer(exasol, port);
        System.out.println("Server running on port: " + port);
        while (true) {
            try {
                Thread.sleep(1000000);
            } catch (final InterruptedException e) {
                Thread.currentThread().interrupt();
                server.close();
                break;
            }
        }
    }

    private Main() {
        // static class
    }
}
