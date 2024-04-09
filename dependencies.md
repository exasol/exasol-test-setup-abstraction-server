<!-- @formatter:off -->
# Dependencies

## Server for the Exasol Test-setup Abstraction

### Compile Dependencies

| Dependency                              | License                                       |
| --------------------------------------- | --------------------------------------------- |
| [exasol-test-setup-abstraction-java][0] | [MIT License][1]                              |
| [javalin][2]                            | [The Apache Software License, Version 2.0][3] |

### Test Dependencies

| Dependency                     | License                           |
| ------------------------------ | --------------------------------- |
| [JUnit Jupiter Engine][4]      | [Eclipse Public License v2.0][5]  |
| [JUnit Jupiter Params][4]      | [Eclipse Public License v2.0][5]  |
| [Hamcrest][6]                  | [BSD License 3][7]                |
| [mockito-junit-jupiter][8]     | [MIT][9]                          |
| [REST Assured][10]             | [Apache 2.0][11]                  |
| [JUnit5 System Extensions][12] | [Eclipse Public License v2.0][13] |

### Runtime Dependencies

| Dependency                 | License           |
| -------------------------- | ----------------- |
| [SLF4J JDK14 Provider][14] | [MIT License][15] |

### Plugin Dependencies

| Dependency                                              | License                           |
| ------------------------------------------------------- | --------------------------------- |
| [SonarQube Scanner for Maven][16]                       | [GNU LGPL 3][17]                  |
| [Apache Maven Toolchains Plugin][18]                    | [Apache License, Version 2.0][19] |
| [Apache Maven Compiler Plugin][20]                      | [Apache-2.0][19]                  |
| [Apache Maven Enforcer Plugin][21]                      | [Apache-2.0][19]                  |
| [Maven Flatten Plugin][22]                              | [Apache Software Licenese][19]    |
| [Exec Maven Plugin][23]                                 | [Apache License 2][19]            |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][24] | [ASL2][3]                         |
| [Maven Surefire Plugin][25]                             | [Apache-2.0][19]                  |
| [Versions Maven Plugin][26]                             | [Apache License, Version 2.0][19] |
| [duplicate-finder-maven-plugin Maven Mojo][27]          | [Apache License 2.0][11]          |
| [Apache Maven Assembly Plugin][28]                      | [Apache-2.0][19]                  |
| [Apache Maven JAR Plugin][29]                           | [Apache License, Version 2.0][19] |
| [Artifact reference checker and unifier][30]            | [MIT License][31]                 |
| [Maven Failsafe Plugin][32]                             | [Apache-2.0][19]                  |
| [JaCoCo :: Maven Plugin][33]                            | [EPL-2.0][34]                     |
| [error-code-crawler-maven-plugin][35]                   | [MIT License][36]                 |
| [Reproducible Build Maven Plugin][37]                   | [Apache 2.0][3]                   |

## Go-client

### Compile Dependencies

| Dependency                         | License   |
| ---------------------------------- | --------- |
| github.com/exasol/exasol-driver-go | [MIT][38] |
| github.com/stretchr/testify        | [MIT][39] |
| gopkg.in/yaml.v3                   | [MIT][40] |

### Test Dependencies

| Dependency                  | License   |
| --------------------------- | --------- |
| github.com/antchfx/xmlquery | [MIT][41] |

[0]: https://github.com/exasol/exasol-test-setup-abstraction-java/
[1]: https://github.com/exasol/exasol-test-setup-abstraction-java/blob/main/LICENSE
[2]: https://javalin.io/
[3]: http://www.apache.org/licenses/LICENSE-2.0.txt
[4]: https://junit.org/junit5/
[5]: https://www.eclipse.org/legal/epl-v20.html
[6]: http://hamcrest.org/JavaHamcrest/
[7]: http://opensource.org/licenses/BSD-3-Clause
[8]: https://github.com/mockito/mockito
[9]: https://opensource.org/licenses/MIT
[10]: http://code.google.com/p/rest-assured
[11]: http://www.apache.org/licenses/LICENSE-2.0.html
[12]: https://github.com/itsallcode/junit5-system-extensions
[13]: http://www.eclipse.org/legal/epl-v20.html
[14]: http://www.slf4j.org
[15]: http://www.opensource.org/licenses/mit-license.php
[16]: http://sonarsource.github.io/sonar-scanner-maven/
[17]: http://www.gnu.org/licenses/lgpl.txt
[18]: https://maven.apache.org/plugins/maven-toolchains-plugin/
[19]: https://www.apache.org/licenses/LICENSE-2.0.txt
[20]: https://maven.apache.org/plugins/maven-compiler-plugin/
[21]: https://maven.apache.org/enforcer/maven-enforcer-plugin/
[22]: https://www.mojohaus.org/flatten-maven-plugin/
[23]: https://www.mojohaus.org/exec-maven-plugin
[24]: https://sonatype.github.io/ossindex-maven/maven-plugin/
[25]: https://maven.apache.org/surefire/maven-surefire-plugin/
[26]: https://www.mojohaus.org/versions/versions-maven-plugin/
[27]: https://basepom.github.io/duplicate-finder-maven-plugin
[28]: https://maven.apache.org/plugins/maven-assembly-plugin/
[29]: https://maven.apache.org/plugins/maven-jar-plugin/
[30]: https://github.com/exasol/artifact-reference-checker-maven-plugin/
[31]: https://github.com/exasol/artifact-reference-checker-maven-plugin/blob/main/LICENSE
[32]: https://maven.apache.org/surefire/maven-failsafe-plugin/
[33]: https://www.jacoco.org/jacoco/trunk/doc/maven.html
[34]: https://www.eclipse.org/legal/epl-2.0/
[35]: https://github.com/exasol/error-code-crawler-maven-plugin/
[36]: https://github.com/exasol/error-code-crawler-maven-plugin/blob/main/LICENSE
[37]: http://zlika.github.io/reproducible-build-maven-plugin
[38]: https://github.com/exasol/exasol-driver-go/blob/v1.0.6/LICENSE
[39]: https://github.com/stretchr/testify/blob/v1.9.0/LICENSE
[40]: https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE
[41]: https://github.com/antchfx/xmlquery/blob/HEAD/LICENSE
