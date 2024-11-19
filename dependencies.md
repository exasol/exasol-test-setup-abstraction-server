<!-- @formatter:off -->
# Dependencies

## Server for the Exasol Test-setup Abstraction

### Compile Dependencies

| Dependency                                                       | License                                                                               |
| ---------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| [exasol-test-setup-abstraction-java][0]                          | [MIT License][1]                                                                      |
| [javalin][2]                                                     | [The Apache Software License, Version 2.0][3]                                         |
| [Core :: Server][4]                                              | [Eclipse Public License - Version 2.0][5]; [Apache Software License - Version 2.0][6] |
| [Jetty :: Websocket :: org.eclipse.jetty.websocket :: Server][7] | [Eclipse Public License - Version 2.0][5]; [Apache Software License - Version 2.0][6] |

### Test Dependencies

| Dependency                     | License                           |
| ------------------------------ | --------------------------------- |
| [JUnit Jupiter Engine][8]      | [Eclipse Public License v2.0][9]  |
| [JUnit Jupiter Params][8]      | [Eclipse Public License v2.0][9]  |
| [Hamcrest][10]                 | [BSD-3-Clause][11]                |
| [mockito-junit-jupiter][12]    | [MIT][13]                         |
| [REST Assured][14]             | [Apache 2.0][15]                  |
| [JUnit5 System Extensions][16] | [Eclipse Public License v2.0][17] |

### Runtime Dependencies

| Dependency                 | License           |
| -------------------------- | ----------------- |
| [SLF4J JDK14 Provider][18] | [MIT License][19] |

### Plugin Dependencies

| Dependency                                              | License                           |
| ------------------------------------------------------- | --------------------------------- |
| [Apache Maven Clean Plugin][20]                         | [Apache-2.0][21]                  |
| [Apache Maven Install Plugin][22]                       | [Apache-2.0][21]                  |
| [Apache Maven Resources Plugin][23]                     | [Apache-2.0][21]                  |
| [Apache Maven Site Plugin][24]                          | [Apache License, Version 2.0][21] |
| [SonarQube Scanner for Maven][25]                       | [GNU LGPL 3][26]                  |
| [Apache Maven Toolchains Plugin][27]                    | [Apache-2.0][21]                  |
| [Apache Maven Compiler Plugin][28]                      | [Apache-2.0][21]                  |
| [Apache Maven Enforcer Plugin][29]                      | [Apache-2.0][21]                  |
| [Maven Flatten Plugin][30]                              | [Apache Software Licenese][21]    |
| [Exec Maven Plugin][31]                                 | [Apache License 2][21]            |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][32] | [ASL2][3]                         |
| [Maven Surefire Plugin][33]                             | [Apache-2.0][21]                  |
| [Versions Maven Plugin][34]                             | [Apache License, Version 2.0][21] |
| [duplicate-finder-maven-plugin Maven Mojo][35]          | [Apache License 2.0][36]          |
| [Apache Maven Assembly Plugin][37]                      | [Apache-2.0][21]                  |
| [Apache Maven JAR Plugin][38]                           | [Apache-2.0][21]                  |
| [Artifact reference checker and unifier][39]            | [MIT License][40]                 |
| [Maven Failsafe Plugin][41]                             | [Apache-2.0][21]                  |
| [JaCoCo :: Maven Plugin][42]                            | [EPL-2.0][5]                      |
| [Quality Summarizer Maven Plugin][43]                   | [MIT License][44]                 |
| [error-code-crawler-maven-plugin][45]                   | [MIT License][46]                 |
| [Reproducible Build Maven Plugin][47]                   | [Apache 2.0][3]                   |

## Go-client

### Compile Dependencies

| Dependency                         | License   |
| ---------------------------------- | --------- |
| github.com/exasol/exasol-driver-go | [MIT][48] |
| github.com/stretchr/testify        | [MIT][49] |
| gopkg.in/yaml.v3                   | [MIT][50] |

### Test Dependencies

| Dependency                  | License   |
| --------------------------- | --------- |
| github.com/antchfx/xmlquery | [MIT][51] |

[0]: https://github.com/exasol/exasol-test-setup-abstraction-java/
[1]: https://github.com/exasol/exasol-test-setup-abstraction-java/blob/main/LICENSE
[2]: https://javalin.io/
[3]: http://www.apache.org/licenses/LICENSE-2.0.txt
[4]: https://jetty.org/jetty-core/jetty-server
[5]: https://www.eclipse.org/legal/epl-2.0/
[6]: https://www.apache.org/licenses/LICENSE-2.0
[7]: https://jetty.org/websocket-parent/websocket-jetty-server
[8]: https://junit.org/junit5/
[9]: https://www.eclipse.org/legal/epl-v20.html
[10]: http://hamcrest.org/JavaHamcrest/
[11]: https://raw.githubusercontent.com/hamcrest/JavaHamcrest/master/LICENSE
[12]: https://github.com/mockito/mockito
[13]: https://opensource.org/licenses/MIT
[14]: https://rest-assured.io/
[15]: https://www.apache.org/licenses/LICENSE-2.0.html
[16]: https://github.com/itsallcode/junit5-system-extensions
[17]: http://www.eclipse.org/legal/epl-v20.html
[18]: http://www.slf4j.org
[19]: http://www.opensource.org/licenses/mit-license.php
[20]: https://maven.apache.org/plugins/maven-clean-plugin/
[21]: https://www.apache.org/licenses/LICENSE-2.0.txt
[22]: https://maven.apache.org/plugins/maven-install-plugin/
[23]: https://maven.apache.org/plugins/maven-resources-plugin/
[24]: https://maven.apache.org/plugins/maven-site-plugin/
[25]: http://sonarsource.github.io/sonar-scanner-maven/
[26]: http://www.gnu.org/licenses/lgpl.txt
[27]: https://maven.apache.org/plugins/maven-toolchains-plugin/
[28]: https://maven.apache.org/plugins/maven-compiler-plugin/
[29]: https://maven.apache.org/enforcer/maven-enforcer-plugin/
[30]: https://www.mojohaus.org/flatten-maven-plugin/
[31]: https://www.mojohaus.org/exec-maven-plugin
[32]: https://sonatype.github.io/ossindex-maven/maven-plugin/
[33]: https://maven.apache.org/surefire/maven-surefire-plugin/
[34]: https://www.mojohaus.org/versions/versions-maven-plugin/
[35]: https://basepom.github.io/duplicate-finder-maven-plugin
[36]: http://www.apache.org/licenses/LICENSE-2.0.html
[37]: https://maven.apache.org/plugins/maven-assembly-plugin/
[38]: https://maven.apache.org/plugins/maven-jar-plugin/
[39]: https://github.com/exasol/artifact-reference-checker-maven-plugin/
[40]: https://github.com/exasol/artifact-reference-checker-maven-plugin/blob/main/LICENSE
[41]: https://maven.apache.org/surefire/maven-failsafe-plugin/
[42]: https://www.jacoco.org/jacoco/trunk/doc/maven.html
[43]: https://github.com/exasol/quality-summarizer-maven-plugin/
[44]: https://github.com/exasol/quality-summarizer-maven-plugin/blob/main/LICENSE
[45]: https://github.com/exasol/error-code-crawler-maven-plugin/
[46]: https://github.com/exasol/error-code-crawler-maven-plugin/blob/main/LICENSE
[47]: http://zlika.github.io/reproducible-build-maven-plugin
[48]: https://github.com/exasol/exasol-driver-go/blob/v1.0.10/LICENSE
[49]: https://github.com/stretchr/testify/blob/v1.9.0/LICENSE
[50]: https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE
[51]: https://github.com/antchfx/xmlquery/blob/HEAD/LICENSE
