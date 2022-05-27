package com.exasol.testsetupabstraction.server;

import java.io.IOException;
import java.net.ServerSocket;
import java.nio.file.Path;

import com.exasol.errorreporting.ExaError;
import com.exasol.exasoltestsetup.ExasolTestSetup;
import com.exasol.exasoltestsetup.ExasolTestSetupFactory;

public class Main {
    @SuppressWarnings({ "java:S106", "java:S4507" }) // we don't want to use a logger, printStackTrace is ok here
                                                     // since it's a tool for testing
    public static void main(final String[] args) {
        final int port = findFreePort();
        try (final ExasolTestSetup exasol = new ExasolTestSetupFactory(Path.of(args[0])).getTestSetup();
                final TestSetupServer server = new TestSetupServer(exasol, port)) {
            System.out.println("Server running on port: " + port);
            server.join();
        } catch (final InterruptedException exception) {
            Thread.currentThread().interrupt();
        } catch (final Exception exception) {
            exception.printStackTrace();
            System.exit(100); // Exit to kill all daemon-threads of javalin, otherwise main terminates but not the app
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
