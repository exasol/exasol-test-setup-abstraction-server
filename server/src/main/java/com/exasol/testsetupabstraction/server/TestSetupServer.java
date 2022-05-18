package com.exasol.testsetupabstraction.server;

import java.util.List;
import java.util.Objects;

import com.exasol.exasoltestsetup.*;

import io.javalin.Javalin;
import io.javalin.http.Context;

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
     * @param port      http port
     */
    public TestSetupServer(final ExasolTestSetup testSetup, final int port) {
        this.testSetup = testSetup;
        this.server = Javalin.create(config -> {
            config.jsonMapper(new JsonBJsonMapper());
        });
        this.server.start(port);
        this.server.post("/makeDatabaseTcpServiceAccessibleFromLocalhost",
                this::handleMakeDatabaseTcpServiceAccessibleFromLocalhost);
        this.server.post("/makeLocalTcpServiceAccessibleFromDatabase", ctx -> {
            final int localPort = Integer.parseInt(Objects.requireNonNull(ctx.formParam("localPort")));
            final ServiceAddress serviceAddress = testSetup.makeLocalTcpServiceAccessibleFromDatabase(localPort);
            ctx.json(serviceAddress);
        });
        this.server.post("/makeTcpServiceAccessibleFromDatabase", ctx -> {
            final String hostName = Objects.requireNonNull(ctx.formParam("hostName"));
            final int portNumber = Integer.parseInt(Objects.requireNonNull(ctx.formParam("port")));
            final ServiceAddress serviceAddress = testSetup
                    .makeTcpServiceAccessibleFromDatabase(new ServiceAddress(hostName, portNumber));
            ctx.json(serviceAddress);
        });
        this.server.get("/connectionInfo", ctx -> {
            final SqlConnectionInfo connectionInfo = testSetup.getConnectionInfo();
            ctx.json(connectionInfo);
        });
    }

    private void handleMakeDatabaseTcpServiceAccessibleFromLocalhost(final Context ctx) {
        final int databasePort = Integer.parseInt(Objects.requireNonNull(ctx.formParam("databasePort")));
        final List<Integer> ports = this.testSetup.makeDatabaseTcpServiceAccessibleFromLocalhost(databasePort);
        ctx.json(ports);
    }

    @Override
    public void close() {
        this.server.stop();
        this.server.close();
    }
}
