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
| [Apache Maven Clean Plugin][16]                         | [Apache-2.0][17]                  |
| [Apache Maven Install Plugin][18]                       | [Apache-2.0][17]                  |
| [Apache Maven Resources Plugin][19]                     | [Apache-2.0][17]                  |
| [Apache Maven Site Plugin][20]                          | [Apache License, Version 2.0][17] |
| [SonarQube Scanner for Maven][21]                       | [GNU LGPL 3][22]                  |
| [Apache Maven Toolchains Plugin][23]                    | [Apache-2.0][17]                  |
| [Apache Maven Compiler Plugin][24]                      | [Apache-2.0][17]                  |
| [Apache Maven Enforcer Plugin][25]                      | [Apache-2.0][17]                  |
| [Maven Flatten Plugin][26]                              | [Apache Software Licenese][17]    |
| [Exec Maven Plugin][27]                                 | [Apache License 2][17]            |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][28] | [ASL2][3]                         |
| [Maven Surefire Plugin][29]                             | [Apache-2.0][17]                  |
| [Versions Maven Plugin][30]                             | [Apache License, Version 2.0][17] |
| [duplicate-finder-maven-plugin Maven Mojo][31]          | [Apache License 2.0][11]          |
| [Apache Maven Assembly Plugin][32]                      | [Apache-2.0][17]                  |
| [Apache Maven JAR Plugin][33]                           | [Apache-2.0][17]                  |
| [Artifact reference checker and unifier][34]            | [MIT License][35]                 |
| [Maven Failsafe Plugin][36]                             | [Apache-2.0][17]                  |
| [JaCoCo :: Maven Plugin][37]                            | [EPL-2.0][38]                     |
| [Quality Summarizer Maven Plugin][39]                   | [MIT License][40]                 |
| [error-code-crawler-maven-plugin][41]                   | [MIT License][42]                 |
| [Reproducible Build Maven Plugin][43]                   | [Apache 2.0][3]                   |

## Go-client

### Compile Dependencies

| Dependency                         | License   |
| ---------------------------------- | --------- |
| github.com/exasol/exasol-driver-go | [MIT][44] |
| github.com/stretchr/testify        | [MIT][45] |
| gopkg.in/yaml.v3                   | [MIT][46] |

### Test Dependencies

| Dependency                  | License   |
| --------------------------- | --------- |
| github.com/antchfx/xmlquery | [MIT][47] |

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
[16]: https://maven.apache.org/plugins/maven-clean-plugin/
[17]: https://www.apache.org/licenses/LICENSE-2.0.txt
[18]: https://maven.apache.org/plugins/maven-install-plugin/
[19]: https://maven.apache.org/plugins/maven-resources-plugin/
[20]: https://maven.apache.org/plugins/maven-site-plugin/
[21]: http://sonarsource.github.io/sonar-scanner-maven/
[22]: http://www.gnu.org/licenses/lgpl.txt
[23]: https://maven.apache.org/plugins/maven-toolchains-plugin/
[24]: https://maven.apache.org/plugins/maven-compiler-plugin/
[25]: https://maven.apache.org/enforcer/maven-enforcer-plugin/
[26]: https://www.mojohaus.org/flatten-maven-plugin/
[27]: https://www.mojohaus.org/exec-maven-plugin
[28]: https://sonatype.github.io/ossindex-maven/maven-plugin/
[29]: https://maven.apache.org/surefire/maven-surefire-plugin/
[30]: https://www.mojohaus.org/versions/versions-maven-plugin/
[31]: https://basepom.github.io/duplicate-finder-maven-plugin
[32]: https://maven.apache.org/plugins/maven-assembly-plugin/
[33]: https://maven.apache.org/plugins/maven-jar-plugin/
[34]: https://github.com/exasol/artifact-reference-checker-maven-plugin/
[35]: https://github.com/exasol/artifact-reference-checker-maven-plugin/blob/main/LICENSE
[36]: https://maven.apache.org/surefire/maven-failsafe-plugin/
[37]: https://www.jacoco.org/jacoco/trunk/doc/maven.html
[38]: https://www.eclipse.org/legal/epl-2.0/
[39]: https://github.com/exasol/quality-summarizer-maven-plugin/
[40]: https://github.com/exasol/quality-summarizer-maven-plugin/blob/main/LICENSE
[41]: https://github.com/exasol/error-code-crawler-maven-plugin/
[42]: https://github.com/exasol/error-code-crawler-maven-plugin/blob/main/LICENSE
[43]: http://zlika.github.io/reproducible-build-maven-plugin
[44]: https://github.com/exasol/exasol-driver-go/blob/v1.0.10/LICENSE
[45]: https://github.com/stretchr/testify/blob/v1.9.0/LICENSE
[46]: https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE
[47]: https://github.com/antchfx/xmlquery/blob/HEAD/LICENSE
