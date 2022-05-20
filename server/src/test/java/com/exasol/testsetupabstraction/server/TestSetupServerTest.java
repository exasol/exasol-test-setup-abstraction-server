package com.exasol.testsetupabstraction.server;

import static io.restassured.RestAssured.get;
import static io.restassured.RestAssured.given;
import static org.hamcrest.Matchers.equalTo;
import static org.mockito.ArgumentMatchers.anyInt;
import static org.mockito.Mockito.*;

import java.util.List;

import org.junit.jupiter.api.*;

import com.exasol.exasoltestsetup.*;

class TestSetupServerTest {

    private static TestSetupServer server;
    private static ExasolTestSetup testSetup;

    @BeforeAll
    static void beforeAll() {
        testSetup = mock(ExasolTestSetup.class);
        server = new TestSetupServer(testSetup, 8080);
    }

    @AfterAll
    static void afterAll() {
        server.close();
    }

    @Test
    void testPOST_makeDatabaseTcpServiceAccessibleFromLocalhost() {
        doReturn(List.of(345)).when(testSetup).makeDatabaseTcpServiceAccessibleFromLocalhost(anyInt());
        given().param("databasePort", 123)//
                .post("/makeDatabaseTcpServiceAccessibleFromLocalhost").then().statusCode(200).assertThat()
                .body(equalTo("[345]"));
        verify(testSetup).makeDatabaseTcpServiceAccessibleFromLocalhost(123);
    }

    @Test
    void testPOST_makeLocalTcpServiceAccessibleFromDatabase() {
        doReturn(new ServiceAddress("localhost", 456)).when(testSetup)
                .makeLocalTcpServiceAccessibleFromDatabase(anyInt());
        given().param("localPort", 123)//
                .post("/makeLocalTcpServiceAccessibleFromDatabase").then().statusCode(200).assertThat()
                .body(equalTo("{\"hostName\":\"localhost\",\"local\":true,\"port\":456}"));
        verify(testSetup).makeLocalTcpServiceAccessibleFromDatabase(123);
    }

    @Test
    void testPOST_makeTcpServiceAccessibleFromDatabase() {
        doReturn(new ServiceAddress("localhost", 456)).when(testSetup).makeTcpServiceAccessibleFromDatabase(any());
        given().param("port", 123)//
                .param("hostName", "localhost").post("/makeTcpServiceAccessibleFromDatabase").then().statusCode(200)
                .assertThat().body(equalTo("{\"hostName\":\"localhost\",\"local\":true,\"port\":456}"));
        verify(testSetup).makeTcpServiceAccessibleFromDatabase(new ServiceAddress("localhost", 123));
    }

    @Test
    void testGET_connectionInfo() {
        doReturn(new SqlConnectionInfo("localhost", 123, "myUser", "myPass")).when(testSetup).getConnectionInfo();
        get("/connectionInfo").then()
                .body(equalTo("{\"host\":\"localhost\",\"password\":\"myPass\",\"port\":123,\"user\":\"myUser\"}"));
    }
}