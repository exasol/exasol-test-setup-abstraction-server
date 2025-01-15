<!-- @formatter:off -->
# Dependencies

## Server for the Exasol Test-setup Abstraction

### Compile Dependencies

| Dependency                              | License                                                                               |
| --------------------------------------- | ------------------------------------------------------------------------------------- |
| [exasol-test-setup-abstraction-java][0] | [MIT License][1]                                                                      |
| [javalin][2]                            | [The Apache Software License, Version 2.0][3]                                         |
| [Jetty :: Server Core][4]               | [Eclipse Public License - Version 2.0][5]; [Apache Software License - Version 2.0][6] |

### Test Dependencies

| Dependency                     | License                           |
| ------------------------------ | --------------------------------- |
| [JUnit Jupiter Engine][7]      | [Eclipse Public License v2.0][8]  |
| [JUnit Jupiter Params][7]      | [Eclipse Public License v2.0][8]  |
| [Hamcrest][9]                  | [BSD-3-Clause][10]                |
| [mockito-junit-jupiter][11]    | [MIT][12]                         |
| [REST Assured][13]             | [Apache 2.0][14]                  |
| [JUnit5 System Extensions][15] | [Eclipse Public License v2.0][16] |

### Runtime Dependencies

| Dependency                 | License           |
| -------------------------- | ----------------- |
| [SLF4J JDK14 Provider][17] | [MIT License][18] |

### Plugin Dependencies

| Dependency                                              | License                           |
| ------------------------------------------------------- | --------------------------------- |
| [Apache Maven Clean Plugin][19]                         | [Apache-2.0][20]                  |
| [Apache Maven Install Plugin][21]                       | [Apache-2.0][20]                  |
| [Apache Maven Resources Plugin][22]                     | [Apache-2.0][20]                  |
| [Apache Maven Site Plugin][23]                          | [Apache-2.0][20]                  |
| [SonarQube Scanner for Maven][24]                       | [GNU LGPL 3][25]                  |
| [Apache Maven Toolchains Plugin][26]                    | [Apache-2.0][20]                  |
| [Apache Maven Compiler Plugin][27]                      | [Apache-2.0][20]                  |
| [Apache Maven Enforcer Plugin][28]                      | [Apache-2.0][20]                  |
| [Maven Flatten Plugin][29]                              | [Apache Software Licenese][20]    |
| [Exec Maven Plugin][30]                                 | [Apache License 2][20]            |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][31] | [ASL2][3]                         |
| [Maven Surefire Plugin][32]                             | [Apache-2.0][20]                  |
| [Versions Maven Plugin][33]                             | [Apache License, Version 2.0][20] |
| [duplicate-finder-maven-plugin Maven Mojo][34]          | [Apache License 2.0][35]          |
| [Apache Maven Assembly Plugin][36]                      | [Apache-2.0][20]                  |
| [Apache Maven JAR Plugin][37]                           | [Apache-2.0][20]                  |
| [Artifact reference checker and unifier][38]            | [MIT License][39]                 |
| [Maven Failsafe Plugin][40]                             | [Apache-2.0][20]                  |
| [JaCoCo :: Maven Plugin][41]                            | [EPL-2.0][5]                      |
| [Quality Summarizer Maven Plugin][42]                   | [MIT License][43]                 |
| [error-code-crawler-maven-plugin][44]                   | [MIT License][45]                 |
| [Reproducible Build Maven Plugin][46]                   | [Apache 2.0][3]                   |

## Go-client

### Compile Dependencies

| Dependency                         | License   |
| ---------------------------------- | --------- |
| github.com/exasol/exasol-driver-go | [MIT][47] |
| github.com/stretchr/testify        | [MIT][48] |
| gopkg.in/yaml.v3                   | [MIT][49] |

### Test Dependencies

| Dependency                  | License   |
| --------------------------- | --------- |
| github.com/antchfx/xmlquery | [MIT][50] |

[0]: https://github.com/exasol/exasol-test-setup-abstraction-java/
[1]: https://github.com/exasol/exasol-test-setup-abstraction-java/blob/main/LICENSE
[2]: https://javalin.io/
[3]: http://www.apache.org/licenses/LICENSE-2.0.txt
[4]: https://jetty.org/jetty-server
[5]: https://www.eclipse.org/legal/epl-2.0/
[6]: https://www.apache.org/licenses/LICENSE-2.0
[7]: https://junit.org/junit5/
[8]: https://www.eclipse.org/legal/epl-v20.html
[9]: http://hamcrest.org/JavaHamcrest/
[10]: https://raw.githubusercontent.com/hamcrest/JavaHamcrest/master/LICENSE
[11]: https://github.com/mockito/mockito
[12]: https://opensource.org/licenses/MIT
[13]: https://rest-assured.io/
[14]: https://www.apache.org/licenses/LICENSE-2.0.html
[15]: https://github.com/itsallcode/junit5-system-extensions
[16]: http://www.eclipse.org/legal/epl-v20.html
[17]: http://www.slf4j.org
[18]: http://www.opensource.org/licenses/mit-license.php
[19]: https://maven.apache.org/plugins/maven-clean-plugin/
[20]: https://www.apache.org/licenses/LICENSE-2.0.txt
[21]: https://maven.apache.org/plugins/maven-install-plugin/
[22]: https://maven.apache.org/plugins/maven-resources-plugin/
[23]: https://maven.apache.org/plugins/maven-site-plugin/
[24]: http://docs.sonarqube.org/display/PLUG/Plugin+Library/sonar-maven-plugin
[25]: http://www.gnu.org/licenses/lgpl.txt
[26]: https://maven.apache.org/plugins/maven-toolchains-plugin/
[27]: https://maven.apache.org/plugins/maven-compiler-plugin/
[28]: https://maven.apache.org/enforcer/maven-enforcer-plugin/
[29]: https://www.mojohaus.org/flatten-maven-plugin/
[30]: https://www.mojohaus.org/exec-maven-plugin
[31]: https://sonatype.github.io/ossindex-maven/maven-plugin/
[32]: https://maven.apache.org/surefire/maven-surefire-plugin/
[33]: https://www.mojohaus.org/versions/versions-maven-plugin/
[34]: https://basepom.github.io/duplicate-finder-maven-plugin
[35]: http://www.apache.org/licenses/LICENSE-2.0.html
[36]: https://maven.apache.org/plugins/maven-assembly-plugin/
[37]: https://maven.apache.org/plugins/maven-jar-plugin/
[38]: https://github.com/exasol/artifact-reference-checker-maven-plugin/
[39]: https://github.com/exasol/artifact-reference-checker-maven-plugin/blob/main/LICENSE
[40]: https://maven.apache.org/surefire/maven-failsafe-plugin/
[41]: https://www.jacoco.org/jacoco/trunk/doc/maven.html
[42]: https://github.com/exasol/quality-summarizer-maven-plugin/
[43]: https://github.com/exasol/quality-summarizer-maven-plugin/blob/main/LICENSE
[44]: https://github.com/exasol/error-code-crawler-maven-plugin/
[45]: https://github.com/exasol/error-code-crawler-maven-plugin/blob/main/LICENSE
[46]: http://zlika.github.io/reproducible-build-maven-plugin
[47]: https://github.com/exasol/exasol-driver-go/blob/v1.0.10/LICENSE
[48]: https://github.com/stretchr/testify/blob/v1.9.0/LICENSE
[49]: https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE
[50]: https://github.com/antchfx/xmlquery/blob/HEAD/LICENSE
