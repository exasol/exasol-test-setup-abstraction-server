package com.exasol.testsetupabstraction.server;

import java.nio.file.Path;
import java.util.*;

import com.exasol.errorreporting.ExaError;
import com.exasol.exasoltestsetup.*;

import io.javalin.Javalin;
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
            config.showJavalinBanner = false;
        });
        configureRequests();
        this.server.exception(Exception.class, (exception, ctx) -> {
            ctx.result(exception.getMessage());
            ctx.status(HttpStatus.INTERNAL_SERVER_ERROR);
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

    private void configureRequests() {
        this.server.post("/makeDatabaseTcpServiceAccessibleFromLocalhost",
                this::handleMakeDatabaseTcpServiceAccessibleFromLocalhost);
        this.server.post("/makeLocalTcpServiceAccessibleFromDatabase", ctx -> {
            final int localPort = Integer.parseInt(Objects.requireNonNull(ctx.formParam("localPort")));
            validatePort(localPort);
            final ServiceAddress serviceAddress = this.testSetup.makeLocalTcpServiceAccessibleFromDatabase(localPort);
            ctx.json(serviceAddress);
        });
        this.server.post("/makeTcpServiceAccessibleFromDatabase", ctx -> {
            final String hostName = Objects.requireNonNull(ctx.formParam("hostName"));
            final int portNumber = Integer.parseInt(Objects.requireNonNull(ctx.formParam("port")));
            validatePort(portNumber);
            final ServiceAddress serviceAddress = this.testSetup
                    .makeTcpServiceAccessibleFromDatabase(new ServiceAddress(hostName, portNumber));
            ctx.json(serviceAddress);
        });
        this.server.get("/connectionInfo", ctx -> {
            final SqlConnectionInfo connectionInfo = this.testSetup.getConnectionInfo();
            ctx.json(connectionInfo);
        });
        this.server.post("/bfs/uploadFile", ctx -> {
            final String localPath = Objects.requireNonNull(ctx.formParam("localPath"));
            final String remoteName = Objects.requireNonNull(ctx.formParam("remoteName"));
            this.testSetup.getDefaultBucket().uploadFile(Path.of(localPath), remoteName);
            ctx.json(Map.of("ok", true));
        });
        this.server.post("/bfs/uploadStringContent", ctx -> {
            final String stringContent = Objects.requireNonNull(ctx.formParam("stringContent"));
            final String remoteName = Objects.requireNonNull(ctx.formParam("remoteName"));
            this.testSetup.getDefaultBucket().uploadStringContent(stringContent, remoteName);
            ctx.json(Map.of("ok", true));
        });
        this.server.delete("/bfs/deleteFile", ctx -> {
            final String path = Objects.requireNonNull(ctx.formParam("path"));
            this.testSetup.getDefaultBucket().deleteFileNonBlocking(path);
            ctx.json(Map.of("ok", true));
        });
        this.server.get("/bfs/listFiles", ctx -> {
            final String path = Objects.requireNonNull(ctx.queryParam("path"));
            final List<String> result = this.testSetup.getDefaultBucket().listContents(path);
            ctx.json(Map.of("files", result));
        });
        this.server.get("/bfs/downloadFileAsString", ctx -> {
            final String path = Objects.requireNonNull(ctx.queryParam("path"));
            final String result = this.testSetup.getDefaultBucket().downloadFileAsString(path);
            ctx.json(Map.of("content", result));
        });
        this.server.get("/bfs/downloadFile", ctx -> {
            final String path = Objects.requireNonNull(ctx.queryParam("remotePath"));
            final String localPath = Objects.requireNonNull(ctx.queryParam("localPath"));
            this.testSetup.getDefaultBucket().downloadFile(path, Path.of(localPath));
            ctx.json(Map.of("ok", true));
        });
    }

    public void join() throws InterruptedException {
        this.server.jettyServer().server().join();
    }

    private void handleMakeDatabaseTcpServiceAccessibleFromLocalhost(final Context ctx) {
        final int databasePort = Integer.parseInt(Objects.requireNonNull(ctx.formParam("databasePort")));
        validatePort(databasePort);
        final List<Integer> ports = this.testSetup.makeDatabaseTcpServiceAccessibleFromLocalhost(databasePort);
        ctx.json(ports);
    }

    @Override
    public void close() {
        this.server.stop();
        this.server.close();
    }
}
