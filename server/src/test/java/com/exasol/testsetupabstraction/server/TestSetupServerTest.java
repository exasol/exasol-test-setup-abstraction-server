package com.exasol.testsetupabstraction.server;

import static io.restassured.RestAssured.get;
import static io.restassured.RestAssured.given;
import static org.hamcrest.Matchers.equalTo;
import static org.mockito.ArgumentMatchers.anyInt;
import static org.mockito.Mockito.*;

import java.io.FileNotFoundException;
import java.nio.file.Path;
import java.util.List;
import java.util.concurrent.TimeoutException;

import org.junit.jupiter.api.*;

import com.exasol.bucketfs.Bucket;
import com.exasol.bucketfs.BucketAccessException;
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

    @Test
    void testGET_listFiles() throws BucketAccessException {
        final Bucket bucket = mock(Bucket.class);
        doReturn(List.of("myFile")).when(bucket).listContents("/my-path");
        doReturn(bucket).when(testSetup).getDefaultBucket();
        given().queryParam("path", "/my-path").get("/bfs/listFiles").then().body(equalTo("{\"files\":[\"myFile\"]}"));
    }

    @Test
    void testPOST_uploadStringContent() throws BucketAccessException, InterruptedException, TimeoutException {
        final Bucket bucket = mock(Bucket.class);
        doReturn(bucket).when(testSetup).getDefaultBucket();
        given().formParam("stringContent", "myString").formParam("remoteName", "myFile.txt")
                .post("/bfs/uploadStringContent").then().body(equalTo("{\"ok\":true}"));
        verify(bucket).uploadStringContent("myString", "myFile.txt");
    }

    @Test
    void testPOST_uploadFile() throws BucketAccessException, TimeoutException, FileNotFoundException {
        final Bucket bucket = mock(Bucket.class);
        doReturn(bucket).when(testSetup).getDefaultBucket();
        given().formParam("localPath", "/tmp/myFile.txt").formParam("remoteName", "myFile.txt").post("/bfs/uploadFile")
                .then().body(equalTo("{\"ok\":true}"));
        verify(bucket).uploadFile(Path.of("/tmp/myFile.txt"), "myFile.txt");
    }

    @Test
    void testDELETE_deleteFile() throws BucketAccessException, TimeoutException, FileNotFoundException {
        final Bucket bucket = mock(Bucket.class);
        doReturn(bucket).when(testSetup).getDefaultBucket();
        given().formParam("path", "myFile.txt").delete("/bfs/deleteFile").then().body(equalTo("{\"ok\":true}"));
        verify(bucket).deleteFileNonBlocking("myFile.txt");
    }
}