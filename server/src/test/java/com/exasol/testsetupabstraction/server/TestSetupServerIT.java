package com.exasol.testsetupabstraction.server;

import static io.restassured.RestAssured.get;
import static io.restassured.RestAssured.given;
import static org.hamcrest.Matchers.equalTo;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

import java.io.FileNotFoundException;
import java.net.InetSocketAddress;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.List;
import java.util.concurrent.TimeoutException;

import org.junit.jupiter.api.*;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.jupiter.MockitoExtension;
import org.mockito.stubbing.Stubber;

import com.exasol.bucketfs.Bucket;
import com.exasol.bucketfs.BucketAccessException;
import com.exasol.exasoltestsetup.ExasolTestSetup;
import com.exasol.exasoltestsetup.SqlConnectionInfo;

@ExtendWith(MockitoExtension.class)
class TestSetupServerIT {

    private static final String MOCKED_EXCEPTION = "mockedException";
    @Mock
    private Bucket bucket;
    private static TestSetupServer server;
    private static ExasolTestSetup testSetup;

    @BeforeAll
    static void beforeAll() {
        testSetup = mock(ExasolTestSetup.class);
        server = new TestSetupServer(testSetup, 8080);
    }

    @BeforeEach
    void resetMock() {
        Mockito.reset(testSetup);
    }

    @AfterAll
    static void afterAll() {
        server.close();
    }

    @Test
    void testPOST_makeDatabaseTcpServiceAccessibleFromLocalhost_succeeds() {
        doReturn(List.of(345)).when(testSetup).makeDatabaseTcpServiceAccessibleFromLocalhost(123);
        given().param("databasePort", 123)//
                .post("/makeDatabaseTcpServiceAccessibleFromLocalhost").then().statusCode(200).assertThat()
                .body(equalTo("[345]"));
        verify(testSetup).makeDatabaseTcpServiceAccessibleFromLocalhost(123);
    }

    @Test
    void testPOST_makeDatabaseTcpServiceAccessibleFromLocalhost_validatesPort() {
        given().param("databasePort", -123)//
                .post("/makeDatabaseTcpServiceAccessibleFromLocalhost").then().assertThat().statusCode(500)
                .body(equalTo("E-ETSAS-8: Port number -123 is negative. Please specify a valid port."));
    }

    @Test
    void testPOST_makeDatabaseTcpServiceAccessibleFromLocalhost_fails() {
        doThrowMockException().when(testSetup).makeDatabaseTcpServiceAccessibleFromLocalhost(123);
        given().param("databasePort", 123) //
                .post("/makeDatabaseTcpServiceAccessibleFromLocalhost").then().statusCode(500).assertThat()
                .body(equalTo(MOCKED_EXCEPTION));
        verify(testSetup).makeDatabaseTcpServiceAccessibleFromLocalhost(123);
    }

    @Test
    void testPOST_makeLocalTcpServiceAccessibleFromDatabase_succeeds() {
        doReturn(new InetSocketAddress("localhost", 456)).when(testSetup)
                .makeLocalTcpServiceAccessibleFromDatabase(123);
        given().param("localPort", 123) //
                .post("/makeLocalTcpServiceAccessibleFromDatabase").then().statusCode(200).assertThat()
                .body("hostName", equalTo("localhost")).body("port", equalTo(456));
        verify(testSetup).makeLocalTcpServiceAccessibleFromDatabase(123);
    }

    @Test
    void testPOST_makeLocalTcpServiceAccessibleFromDatabase_fails() {
        doThrowMockException().when(testSetup).makeLocalTcpServiceAccessibleFromDatabase(123);
        given().param("localPort", 123) //
                .post("/makeLocalTcpServiceAccessibleFromDatabase").then().statusCode(500).assertThat()
                .body(equalTo(MOCKED_EXCEPTION));
        verify(testSetup).makeLocalTcpServiceAccessibleFromDatabase(123);
    }

    @Test
    void testPOST_makeTcpServiceAccessibleFromDatabase_succeeds() {
        doReturn(new InetSocketAddress("localhost", 456)).when(testSetup).makeTcpServiceAccessibleFromDatabase(any());
        given().param("port", 123)//
                .param("hostName", "localhost").post("/makeTcpServiceAccessibleFromDatabase").then().statusCode(200)
                .assertThat().body("hostName", equalTo("localhost")).body("port", equalTo(456));
        verify(testSetup).makeTcpServiceAccessibleFromDatabase(new InetSocketAddress("localhost", 123));
    }

    @Test
    void testPOST_makeTcpServiceAccessibleFromDatabase_fails() {
        doThrowMockException().when(testSetup).makeTcpServiceAccessibleFromDatabase(any());
        given().param("port", 123)//
                .param("hostName", "localhost").post("/makeTcpServiceAccessibleFromDatabase").then().statusCode(500)
                .assertThat().body(equalTo(MOCKED_EXCEPTION));
        verify(testSetup).makeTcpServiceAccessibleFromDatabase(new InetSocketAddress("localhost", 123));
    }

    @Test
    void testGET_connectionInfo_succeeds() {
        doReturn(new SqlConnectionInfo("localhost", 123, "myUser", "myPass")).when(testSetup).getConnectionInfo();
        get("/connectionInfo").then().statusCode(200) //
                .body("host", equalTo("localhost")).body("port", equalTo(123)).body("user", equalTo("myUser"))
                .body("password", equalTo("myPass"));
    }

    @Test
    void testGET_connectionInfo_fails() {
        doThrowMockException().when(testSetup).getConnectionInfo();
        get("/connectionInfo").then().statusCode(500).body(equalTo(MOCKED_EXCEPTION));
    }

    @Test
    void testGET_listFiles_succeeds() throws BucketAccessException {
        doReturn(List.of("myFile")).when(bucket).listContents("/my-path");
        doReturn(bucket).when(testSetup).getDefaultBucket();
        given().queryParam("path", "/my-path").get("/bfs/listFiles").then().statusCode(200)
                .body(equalTo("{\"files\":[\"myFile\"]}"));
    }

    @Test
    void testGET_listFiles_fails() throws BucketAccessException {
        doThrowMockException().when(bucket).listContents("/my-path");
        doReturn(bucket).when(testSetup).getDefaultBucket();
        given().queryParam("path", "/my-path").get("/bfs/listFiles").then().statusCode(500)
                .body(equalTo(MOCKED_EXCEPTION));
    }

    @Test
    void testPOST_uploadStringContent_succeeds() throws BucketAccessException, InterruptedException, TimeoutException {
        doReturn(bucket).when(testSetup).getDefaultBucket();
        given().formParam("stringContent", "myString").formParam("remoteName", "myFile.txt")
                .post("/bfs/uploadStringContent").then().statusCode(200).body(equalTo("{\"ok\":true}"));
        verify(bucket).uploadStringContent("myString", "myFile.txt");
    }

    @Test
    void testPOST_uploadStringContent_fails() throws BucketAccessException, InterruptedException, TimeoutException {
        doThrowMockException().when(testSetup).getDefaultBucket();
        given().formParam("stringContent", "myString").formParam("remoteName", "myFile.txt")
                .post("/bfs/uploadStringContent").then().statusCode(500).body(equalTo(MOCKED_EXCEPTION));
    }

    @Test
    void testPOST_uploadFile_succeeds() throws BucketAccessException, TimeoutException, FileNotFoundException {
        doReturn(bucket).when(testSetup).getDefaultBucket();
        given().formParam("localPath", "/tmp/myFile.txt").formParam("remoteName", "myFile.txt").post("/bfs/uploadFile")
                .then().statusCode(200).body(equalTo("{\"ok\":true}"));
        verify(bucket).uploadFile(Path.of("/tmp/myFile.txt"), "myFile.txt");
    }

    @Test
    void testPOST_uploadFile_fails() throws BucketAccessException, TimeoutException, FileNotFoundException {
        doThrowMockException().when(testSetup).getDefaultBucket();
        given().formParam("localPath", "/tmp/myFile.txt").formParam("remoteName", "myFile.txt").post("/bfs/uploadFile")
                .then().statusCode(500).body(equalTo(MOCKED_EXCEPTION));
    }

    @Test
    void testDELETE_deleteFile_succeeds() throws BucketAccessException, TimeoutException, FileNotFoundException {
        doReturn(bucket).when(testSetup).getDefaultBucket();
        given().formParam("path", "myFile.txt").delete("/bfs/deleteFile").then().statusCode(200)
                .body(equalTo("{\"ok\":true}"));
        verify(bucket).deleteFileNonBlocking("myFile.txt");
    }

    @Test
    void testDELETE_deleteFile_fails() throws BucketAccessException, TimeoutException, FileNotFoundException {
        doThrowMockException().when(testSetup).getDefaultBucket();
        given().formParam("path", "myFile.txt").delete("/bfs/deleteFile").then().statusCode(500)
                .body(equalTo(MOCKED_EXCEPTION));
    }

    @Test
    void testGET_downloadFileAsString_succeeds() throws BucketAccessException {
        doReturn(bucket).when(testSetup).getDefaultBucket();
        doReturn("mockContent").when(bucket).downloadFileAsString("/pathInBucket");
        given().formParam("path", "/pathInBucket").get("/bfs/downloadFileAsString").then().statusCode(200)
                .body(equalTo("{\"content\":\"mockContent\"}"));
    }

    @Test
    void testGET_downloadFileAsString_fails() throws BucketAccessException {
        doReturn(bucket).when(testSetup).getDefaultBucket();
        doThrowMockException().when(bucket).downloadFileAsString("/pathInBucket");
        given().formParam("path", "/pathInBucket").get("/bfs/downloadFileAsString").then().statusCode(500)
                .body(equalTo(MOCKED_EXCEPTION));
    }

    @Test
    void testGET_downloadFile_succeeds() throws BucketAccessException {
        doReturn(bucket).when(testSetup).getDefaultBucket();
        given().formParam("remotePath", "/remotePath").formParam("localPath", "/localPath").get("/bfs/downloadFile")
                .then().statusCode(200).body(equalTo("{\"ok\":true}"));
        verify(bucket).downloadFile("/remotePath", Paths.get("/localPath"));
    }

    @Test
    void testGET_downloadFile_fails() throws BucketAccessException {
        doReturn(bucket).when(testSetup).getDefaultBucket();
        doThrowMockException().when(bucket).downloadFile("/remotePath", Paths.get("/localPath"));
        given().formParam("remotePath", "/remotePath").formParam("localPath", "/localPath").get("/bfs/downloadFile")
                .then().statusCode(500).body(equalTo(MOCKED_EXCEPTION));
    }

    private Stubber doThrowMockException() {
        return doThrow(new RuntimeException(MOCKED_EXCEPTION));
    }
}