package com.exasol.testsetupabstraction.server;

import java.io.IOException;
import java.net.ServerSocket;
import java.nio.file.Path;

import com.exasol.errorreporting.ExaError;
import com.exasol.exasoltestsetup.ExasolTestSetup;
import com.exasol.exasoltestsetup.ExasolTestSetupFactory;

public class Main {
    @SuppressWarnings({ "java:S106", "java:S2189" }) // we don't want to use a logger here, endless loop is intended
    public static void main(final String[] args) {
        final ExasolTestSetup exasol = new ExasolTestSetupFactory(Path.of(args[0])).getTestSetup();
        final int port = findFreePort();
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

    private static int findFreePort() {
        try (final ServerSocket socket = new ServerSocket(0)) {
            return socket.getLocalPort();
        } catch (final IOException exception) {
            throw new IllegalStateException(
                    ExaError.messageBuilder("E-ETSAS-5").message("Failed to find a free port.").toString(), exception);
        }
    }

    private Main() {
        // static class
    }
}
