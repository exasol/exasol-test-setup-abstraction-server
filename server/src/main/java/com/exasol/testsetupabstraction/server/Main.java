package com.exasol.testsetupabstraction.server;

import java.io.IOException;
import java.io.InputStream;
import java.net.ServerSocket;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.logging.*;

import com.exasol.errorreporting.ExaError;
import com.exasol.exasoltestsetup.ExasolTestSetup;
import com.exasol.exasoltestsetup.ExasolTestSetupFactory;

public class Main {

    static {
        configureLogging();
    }

    private static final Logger LOGGER = Logger.getLogger(Main.class.getName());

    @SuppressWarnings("java:S4792") // Logger configuration is safe, we don't log sensitive information.
    private static void configureLogging() {
        try (final InputStream is = Main.class.getClassLoader().getResourceAsStream("logging.properties")) {
            LogManager.getLogManager().readConfiguration(is);
        } catch (final IOException exception) {
            LOGGER.log(Level.WARNING, ExaError.messageBuilder("W-ETSAS-6")
                    .message("Failed to load logging configuration.").ticketMitigation().toString(), exception);
        }
    }

    @SuppressWarnings({ "java:S106" }) // Using System.out by intention to log server port
    public static void main(final String[] args) {
        final int port = findFreePort();
        try (final ExasolTestSetup exasol = new ExasolTestSetupFactory(getConfigPath(args)).getTestSetup();
                final TestSetupServer server = new TestSetupServer(exasol, port)) {
            System.out.println("Server running on port: " + port);
            server.join();
        } catch (final InterruptedException exception) {
            Thread.currentThread().interrupt();
        } catch (final Exception exception) {
            LOGGER.log(Level.SEVERE,
                    ExaError.messageBuilder("E-ETSAS-7")
                            .message("Failed to start server: {{error|q}}", exception.getMessage()).toString(),
                    exception);
            System.exit(100); // Exit to kill all daemon-threads of javalin, otherwise main terminates but not the app
        }
    }

    private static Path getConfigPath(final String[] args) {
        if (args.length > 0) {
            final Path configPath = Path.of(args[0]).toAbsolutePath();
            LOGGER.info(() -> "Starting exasol test setup using config file '" + configPath + "'");
            return configPath;
        } else {
            LOGGER.info(() -> "Starting exasol test setup using dummy config file");
            return Paths.get("non-existing-config-file-" + System.currentTimeMillis() + ".json");
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
