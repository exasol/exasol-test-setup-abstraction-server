package com.exasol.testsetupabstraction.server;

import static java.util.Objects.requireNonNull;

import java.net.InetSocketAddress;
import java.nio.file.Path;
import java.util.List;
import java.util.Map;

import com.exasol.errorreporting.ExaError;
import com.exasol.exasoltestsetup.ExasolTestSetup;
import com.exasol.exasoltestsetup.SqlConnectionInfo;

import io.javalin.Javalin;
import io.javalin.config.RoutesConfig;
import io.javalin.http.Context;
import io.javalin.http.HttpStatus;

/**
 * Server that exposes the functions of the exasol-test-setup-abstraction-java via a REST API.
 */
public class TestSetupServer implements AutoCloseable {
    private final Javalin server;
    private final ExasolTestSetup testSetup;

    /**
     * Create a new instance of {@link TestSetupServer}.
     *
     * @param testSetup Exasol test setup
     * @param port      HTTP port
     */
    public TestSetupServer(final ExasolTestSetup testSetup, final int port) {
        this.testSetup = testSetup;
        this.server = Javalin.create(config -> {
            config.jsonMapper(new JsonBJsonMapper());
            config.startup.showJavalinBanner = false;
            configureEndpoints(config.routes);
            config.routes.exception(Exception.class, (exception, ctx) -> {
                ctx.result(exception.getMessage());
                ctx.status(HttpStatus.INTERNAL_SERVER_ERROR);
            });
        });
        this.server.start(port);
    }

    private void validatePort(final int port) {
        if (port < 0) {
            throw new IllegalArgumentException(
                    ExaError.messageBuilder("E-ETSAS-8").message("Port number {{port}} is negative.", port)
                            .mitigation("Please specify a valid port.").toString());
        }
    }

    private void configureEndpoints(final RoutesConfig routes) {
        routes.post("/makeDatabaseTcpServiceAccessibleFromLocalhost",
                this::handleMakeDatabaseTcpServiceAccessibleFromLocalhost);
        routes.post("/makeLocalTcpServiceAccessibleFromDatabase", ctx -> {
            final int localPort = Integer.parseInt(requireNonNull(ctx.formParam("localPort")));
            validatePort(localPort);
            final InetSocketAddress serviceAddress = this.testSetup
                    .makeLocalTcpServiceAccessibleFromDatabase(localPort);
            ctx.json(toMap(serviceAddress));
        });
        routes.post("/makeTcpServiceAccessibleFromDatabase", ctx -> {
            final String hostName = requireNonNull(ctx.formParam("hostName"));
            final int portNumber = Integer.parseInt(requireNonNull(ctx.formParam("port")));
            validatePort(portNumber);
            final InetSocketAddress serviceAddress = this.testSetup
                    .makeTcpServiceAccessibleFromDatabase(new InetSocketAddress(hostName, portNumber));
            ctx.json(toMap(serviceAddress));
        });
        routes.get("/connectionInfo", ctx -> {
            final SqlConnectionInfo connectionInfo = this.testSetup.getConnectionInfo();
            ctx.json(connectionInfo);
        });
        routes.post("/bfs/uploadFile", ctx -> {
            final String localPath = requireNonNull(ctx.formParam("localPath"));
            final String remoteName = requireNonNull(ctx.formParam("remoteName"));
            this.testSetup.getDefaultBucket().uploadFile(Path.of(localPath), remoteName);
            ctx.json(Map.of("ok", true));
        });
        routes.post("/bfs/uploadStringContent", ctx -> {
            final String stringContent = requireNonNull(ctx.formParam("stringContent"));
            final String remoteName = requireNonNull(ctx.formParam("remoteName"));
            this.testSetup.getDefaultBucket().uploadStringContent(stringContent, remoteName);
            ctx.json(Map.of("ok", true));
        });
        routes.delete("/bfs/deleteFile", ctx -> {
            final String path = requireNonNull(ctx.formParam("path"));
            this.testSetup.getDefaultBucket().deleteFileNonBlocking(path);
            ctx.json(Map.of("ok", true));
        });
        routes.get("/bfs/listFiles", ctx -> {
            final String path = requireNonNull(ctx.queryParam("path"));
            final List<String> result = this.testSetup.getDefaultBucket().listContents(path);
            ctx.json(Map.of("files", result));
        });
        routes.get("/bfs/downloadFileAsString", ctx -> {
            final String path = requireNonNull(ctx.queryParam("path"));
            final String result = this.testSetup.getDefaultBucket().downloadFileAsString(path);
            ctx.json(Map.of("content", result));
        });
        routes.get("/bfs/downloadFile", ctx -> {
            final String path = requireNonNull(ctx.queryParam("remotePath"));
            final String localPath = requireNonNull(ctx.queryParam("localPath"));
            this.testSetup.getDefaultBucket().downloadFile(path, Path.of(localPath));
            ctx.json(Map.of("ok", true));
        });
    }

    private Map<String, Object> toMap(final InetSocketAddress address) {
        return Map.of("hostName", address.getHostName(), "port", address.getPort());
    }

    public void join() throws InterruptedException {
        this.server.jettyServer().server().join();
    }

    private void handleMakeDatabaseTcpServiceAccessibleFromLocalhost(final Context ctx) {
        final int databasePort = Integer.parseInt(requireNonNull(ctx.formParam("databasePort")));
        validatePort(databasePort);
        final List<Integer> ports = this.testSetup.makeDatabaseTcpServiceAccessibleFromLocalhost(databasePort);
        ctx.json(ports);
    }

    @Override
    public void close() {
        this.server.stop();
    }
}
