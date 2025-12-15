# Exasol Test Setup Abstraction Server 1.0.0, released 2025-12-15

Code name: Fixes for vulnerability CVE-2025-48924 CVE-2024-55551, CVE-2025-22872, CVE-2025-22870

## Summary

This release updates Java and requires Java 17.

The release also fixes the following vulnerability:

### CVE-2025-48924 (CWE-674) in dependency `org.apache.commons:commons-lang3:jar:3.11:compile`

Uncontrolled Recursion vulnerability in Apache Commons Lang.

This issue affects Apache Commons Lang: Starting with commons-lang:commons-lang 2.0 to 2.6, and, from org.apache.commons:commons-lang3 3.0 before 3.18.0.

The methods ClassUtils.getClass(...) can throw StackOverflowError on very long inputs. Because an Error is usually not handled by applications and libraries, a 
StackOverflowError could cause an application to stop.

Users are recommended to upgrade to version 3.18.0, which fixes the issue.

CVE: CVE-2025-48924
CWE: CWE-674

#### References

- https://ossindex.sonatype.org/vulnerability/CVE-2025-48924?component-type=maven&component-name=org.apache.commons%2Fcommons-lang3&utm_source=ossindex-client&utm_medium=integration&utm_content=1.8.1
- http://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2025-48924
- https://github.com/advisories/GHSA-j288-q9x7-2f5v

### CVE-2024-55551 (CWE-94) in dependency `com.exasol:exasol-jdbc:jar:24.2.1:runtime`

An issue was discovered in Exasol jdbc driver 24.2.0. Attackers can inject malicious parameters into the JDBC URL, triggering JNDI injection during the process when the JDBC Driver uses this URL to connect to the database. This can further lead to remote code execution vulnerability.

CVE: CVE-2024-55551
CWE: CWE-94

#### References

- https://ossindex.sonatype.org/vulnerability/CVE-2024-55551?component-type=maven&component-name=com.exasol%2Fexasol-jdbc&utm_source=ossindex-client&utm_medium=integration&utm_content=1.8.1
- http://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2024-55551
- https://gist.github.com/azraelxuemo/9565ec9219e0c3e9afd5474904c39d0f

## CVE-2025-22872
golang.org/x/net vulnerable to Cross-site Scripting

## CVE-2025-22870
HTTP Proxy bypass using IPv6 Zone IDs in golang.org/x/net

## Security

* #61: CVE-2025-48924 in `org.apache.commons:commons-lang3:jar:3.11:compile`
* #60: CVE-2024-55551 in `com.exasol:exasol-jdbc:jar:24.2.1:runtime`
* #63: CVE-2025-22872, CVE-2025-22870 in `golang.org/x/net`
* #64: Fix CVE-2025-58178 in `sonarqube-scan-action`

## Dependency Updates

### Server for the Exasol Test-Setup Abstraction

#### Compile Dependency Updates

* Updated `com.exasol:exasol-test-setup-abstraction-java:2.1.7` to `2.1.10`
* Updated `io.javalin:javalin:6.4.0` to `6.7.0`

#### Runtime Dependency Updates

* Updated `org.slf4j:slf4j-jdk14:2.0.16` to `2.0.17`

#### Test Dependency Updates

* Updated `io.rest-assured:rest-assured:5.5.0` to `6.0.0`
* Updated `org.junit.jupiter:junit-jupiter-engine:5.11.4` to `6.0.1`
* Updated `org.junit.jupiter:junit-jupiter-params:5.11.4` to `6.0.1`
* Updated `org.mockito:mockito-junit-jupiter:5.15.2` to `5.21.0`

#### Plugin Dependency Updates

* Updated `com.exasol:artifact-reference-checker-maven-plugin:0.4.2` to `0.4.4`
* Updated `com.exasol:error-code-crawler-maven-plugin:2.0.3` to `2.0.5`
* Updated `com.exasol:quality-summarizer-maven-plugin:0.2.0` to `0.2.1`
* Added `io.github.git-commit-id:git-commit-id-maven-plugin:9.0.2`
* Removed `io.github.zlika:reproducible-build-maven-plugin:0.17`
* Added `org.apache.maven.plugins:maven-artifact-plugin:3.6.1`
* Updated `org.apache.maven.plugins:maven-clean-plugin:3.4.0` to `3.5.0`
* Updated `org.apache.maven.plugins:maven-compiler-plugin:3.13.0` to `3.14.1`
* Updated `org.apache.maven.plugins:maven-enforcer-plugin:3.5.0` to `3.6.2`
* Updated `org.apache.maven.plugins:maven-failsafe-plugin:3.5.2` to `3.5.4`
* Updated `org.apache.maven.plugins:maven-install-plugin:3.1.3` to `3.1.4`
* Updated `org.apache.maven.plugins:maven-surefire-plugin:3.5.2` to `3.5.4`
* Updated `org.codehaus.mojo:exec-maven-plugin:3.5.0` to `3.5.1`
* Updated `org.codehaus.mojo:flatten-maven-plugin:1.6.0` to `1.7.3`
* Updated `org.codehaus.mojo:versions-maven-plugin:2.18.0` to `2.19.1`
* Updated `org.jacoco:jacoco-maven-plugin:0.8.12` to `0.8.14`
* Updated `org.sonarsource.scanner.maven:sonar-maven-plugin:5.0.0.4389` to `5.2.0.4988`

### Go-client

#### Compile Dependency Updates

* Updated `golang:1.23` to `1.24.0`
* Updated `github.com/stretchr/testify:v1.10.0` to `v1.11.1`
* Updated `github.com/exasol/exasol-driver-go:v1.0.12` to `v1.0.14`

#### Test Dependency Updates

* Updated `github.com/antchfx/xmlquery:v1.4.4` to `v1.5.0`

#### Other Dependency Updates

* Added `toolchain:go1.25.5`
