<!-- @formatter:off -->
# Dependencies

## Server for the Exasol Test-setup Abstraction

### Compile Dependencies

| Dependency                              | License                                                                               |
| --------------------------------------- | ------------------------------------------------------------------------------------- |
| [exasol-test-setup-abstraction-java][0] | [MIT License][1]                                                                      |
| [javalin][2]                            | [The Apache Software License, Version 2.0][3]                                         |
| [Jetty :: Http Utility][4]              | [Eclipse Public License - Version 2.0][5]; [Apache Software License - Version 2.0][6] |

### Test Dependencies

| Dependency                     | License                           |
| ------------------------------ | --------------------------------- |
| [JUnit Jupiter Engine][7]      | [Eclipse Public License v2.0][8]  |
| [JUnit Jupiter Params][7]      | [Eclipse Public License v2.0][8]  |
| [Hamcrest][9]                  | [BSD License 3][10]               |
| [mockito-junit-jupiter][11]    | [MIT][12]                         |
| [REST Assured][13]             | [Apache 2.0][14]                  |
| [JUnit5 System Extensions][15] | [Eclipse Public License v2.0][16] |

### Runtime Dependencies

| Dependency                 | License           |
| -------------------------- | ----------------- |
| [SLF4J JDK14 Provider][17] | [MIT License][18] |

### Plugin Dependencies

| Dependency                                              | License                                       |
| ------------------------------------------------------- | --------------------------------------------- |
| [SonarQube Scanner for Maven][19]                       | [GNU LGPL 3][20]                              |
| [Apache Maven Compiler Plugin][21]                      | [Apache-2.0][22]                              |
| [Apache Maven Enforcer Plugin][23]                      | [Apache-2.0][22]                              |
| [Maven Flatten Plugin][24]                              | [Apache Software Licenese][22]                |
| [Exec Maven Plugin][25]                                 | [Apache License 2][22]                        |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][26] | [ASL2][3]                                     |
| [Maven Surefire Plugin][27]                             | [Apache-2.0][22]                              |
| [Versions Maven Plugin][28]                             | [Apache License, Version 2.0][22]             |
| [duplicate-finder-maven-plugin Maven Mojo][29]          | [Apache License 2.0][14]                      |
| [Apache Maven Assembly Plugin][30]                      | [Apache-2.0][22]                              |
| [Apache Maven JAR Plugin][31]                           | [Apache License, Version 2.0][22]             |
| [Artifact reference checker and unifier][32]            | [MIT License][33]                             |
| [Maven Failsafe Plugin][34]                             | [Apache-2.0][22]                              |
| [JaCoCo :: Maven Plugin][35]                            | [Eclipse Public License 2.0][5]               |
| [error-code-crawler-maven-plugin][36]                   | [MIT License][37]                             |
| [Reproducible Build Maven Plugin][38]                   | [Apache 2.0][3]                               |
| [Maven Clean Plugin][39]                                | [The Apache Software License, Version 2.0][3] |
| [Maven Resources Plugin][40]                            | [The Apache Software License, Version 2.0][3] |
| [Maven Install Plugin][41]                              | [The Apache Software License, Version 2.0][3] |
| [Maven Deploy Plugin][42]                               | [The Apache Software License, Version 2.0][3] |
| [Maven Site Plugin 3][43]                               | [The Apache Software License, Version 2.0][3] |

## Go-client

### Compile Dependencies

| Dependency                         | License   |
| ---------------------------------- | --------- |
| github.com/exasol/exasol-driver-go | [MIT][44] |

### Test Dependencies

| Dependency                  | License   |
| --------------------------- | --------- |
| github.com/antchfx/xmlquery | [MIT][45] |
| github.com/stretchr/testify | [MIT][46] |
| gopkg.in/yaml.v3            | [MIT][47] |

[0]: https://github.com/exasol/exasol-test-setup-abstraction-java/
[1]: https://github.com/exasol/exasol-test-setup-abstraction-java/blob/main/LICENSE
[2]: https://javalin.io/
[3]: http://www.apache.org/licenses/LICENSE-2.0.txt
[4]: https://eclipse.dev/jetty/jetty-http
[5]: https://www.eclipse.org/legal/epl-2.0/
[6]: https://www.apache.org/licenses/LICENSE-2.0
[7]: https://junit.org/junit5/
[8]: https://www.eclipse.org/legal/epl-v20.html
[9]: http://hamcrest.org/JavaHamcrest/
[10]: http://opensource.org/licenses/BSD-3-Clause
[11]: https://github.com/mockito/mockito
[12]: https://github.com/mockito/mockito/blob/main/LICENSE
[13]: http://code.google.com/p/rest-assured
[14]: http://www.apache.org/licenses/LICENSE-2.0.html
[15]: https://github.com/itsallcode/junit5-system-extensions
[16]: http://www.eclipse.org/legal/epl-v20.html
[17]: http://www.slf4j.org
[18]: http://www.opensource.org/licenses/mit-license.php
[19]: http://sonarsource.github.io/sonar-scanner-maven/
[20]: http://www.gnu.org/licenses/lgpl.txt
[21]: https://maven.apache.org/plugins/maven-compiler-plugin/
[22]: https://www.apache.org/licenses/LICENSE-2.0.txt
[23]: https://maven.apache.org/enforcer/maven-enforcer-plugin/
[24]: https://www.mojohaus.org/flatten-maven-plugin/
[25]: https://www.mojohaus.org/exec-maven-plugin
[26]: https://sonatype.github.io/ossindex-maven/maven-plugin/
[27]: https://maven.apache.org/surefire/maven-surefire-plugin/
[28]: https://www.mojohaus.org/versions/versions-maven-plugin/
[29]: https://basepom.github.io/duplicate-finder-maven-plugin
[30]: https://maven.apache.org/plugins/maven-assembly-plugin/
[31]: https://maven.apache.org/plugins/maven-jar-plugin/
[32]: https://github.com/exasol/artifact-reference-checker-maven-plugin/
[33]: https://github.com/exasol/artifact-reference-checker-maven-plugin/blob/main/LICENSE
[34]: https://maven.apache.org/surefire/maven-failsafe-plugin/
[35]: https://www.jacoco.org/jacoco/trunk/doc/maven.html
[36]: https://github.com/exasol/error-code-crawler-maven-plugin/
[37]: https://github.com/exasol/error-code-crawler-maven-plugin/blob/main/LICENSE
[38]: http://zlika.github.io/reproducible-build-maven-plugin
[39]: http://maven.apache.org/plugins/maven-clean-plugin/
[40]: http://maven.apache.org/plugins/maven-resources-plugin/
[41]: http://maven.apache.org/plugins/maven-install-plugin/
[42]: http://maven.apache.org/plugins/maven-deploy-plugin/
[43]: http://maven.apache.org/plugins/maven-site-plugin/
[44]: https://github.com/exasol/exasol-driver-go/blob/v1.0.0/LICENSE
[45]: https://github.com/antchfx/xmlquery/blob/HEAD/LICENSE
[46]: https://github.com/stretchr/testify/blob/HEAD/LICENSE
[47]: https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE
