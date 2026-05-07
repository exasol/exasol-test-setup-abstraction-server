<!-- @formatter:off -->
# Dependencies

## Server for the Exasol Test-setup Abstraction

### Compile Dependencies

| Dependency                              | License                                       |
| --------------------------------------- | --------------------------------------------- |
| [exasol-test-setup-abstraction-java][0] | [MIT License][1]                              |
| [Javalin][2]                            | [The Apache Software License, Version 2.0][3] |

### Test Dependencies

| Dependency                     | License                           |
| ------------------------------ | --------------------------------- |
| [JUnit Jupiter API][4]         | [Eclipse Public License v2.0][5]  |
| [JUnit Jupiter Params][4]      | [Eclipse Public License v2.0][5]  |
| [Hamcrest][6]                  | [BSD-3-Clause][7]                 |
| [mockito-junit-jupiter][8]     | [MIT][9]                          |
| [REST Assured][10]             | [Apache 2.0][11]                  |
| [JUnit5 System Extensions][12] | [Eclipse Public License v2.0][13] |

### Runtime Dependencies

| Dependency                 | License   |
| -------------------------- | --------- |
| [SLF4J JDK14 Provider][14] | [MIT][15] |

### Plugin Dependencies

| Dependency                                              | License                                     |
| ------------------------------------------------------- | ------------------------------------------- |
| [SonarQube Scanner for Maven][16]                       | [GNU LGPL 3][17]                            |
| [Apache Maven Toolchains Plugin][18]                    | [Apache-2.0][19]                            |
| [Apache Maven Compiler Plugin][20]                      | [Apache-2.0][19]                            |
| [Apache Maven Enforcer Plugin][21]                      | [Apache-2.0][19]                            |
| [Maven Flatten Plugin][22]                              | [Apache Software License][19]               |
| [Exec Maven Plugin][23]                                 | [Apache License 2][19]                      |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][24] | [ASL2][3]                                   |
| [Maven Surefire Plugin][25]                             | [Apache-2.0][19]                            |
| [Versions Maven Plugin][26]                             | [Apache License, Version 2.0][19]           |
| [duplicate-finder-maven-plugin Maven Mojo][27]          | [Apache License 2.0][28]                    |
| [Apache Maven Artifact Plugin][29]                      | [Apache-2.0][19]                            |
| [Apache Maven Assembly Plugin][30]                      | [Apache-2.0][19]                            |
| [Apache Maven JAR Plugin][31]                           | [Apache-2.0][19]                            |
| [Artifact reference checker and unifier][32]            | [MIT License][33]                           |
| [Maven Failsafe Plugin][34]                             | [Apache-2.0][19]                            |
| [JaCoCo :: Maven Plugin][35]                            | [EPL-2.0][36]                               |
| [Quality Summarizer Maven Plugin][37]                   | [MIT License][38]                           |
| [error-code-crawler-maven-plugin][39]                   | [MIT License][40]                           |
| [Git Commit Id Maven Plugin][41]                        | [GNU Lesser General Public License 3.0][42] |
| [Apache Maven Clean Plugin][43]                         | [Apache-2.0][19]                            |
| [Apache Maven Resources Plugin][44]                     | [Apache-2.0][19]                            |
| [Apache Maven Install Plugin][45]                       | [Apache-2.0][19]                            |
| [Apache Maven Site Plugin][46]                          | [Apache-2.0][19]                            |

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
[4]: https://junit.org/
[5]: https://www.eclipse.org/legal/epl-v20.html
[6]: http://hamcrest.org/JavaHamcrest/
[7]: https://raw.githubusercontent.com/hamcrest/JavaHamcrest/master/LICENSE
[8]: https://github.com/mockito/mockito
[9]: https://opensource.org/licenses/MIT
[10]: https://rest-assured.io/
[11]: https://www.apache.org/licenses/LICENSE-2.0.html
[12]: https://github.com/itsallcode/junit5-system-extensions
[13]: http://www.eclipse.org/legal/epl-v20.html
[14]: http://www.slf4j.org
[15]: https://opensource.org/license/mit
[16]: https://docs.sonarsource.com/sonarqube-server/latest/extension-guide/developing-a-plugin/plugin-basics/sonar-scanner-maven/sonar-maven-plugin/
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
[28]: http://www.apache.org/licenses/LICENSE-2.0.html
[29]: https://maven.apache.org/plugins/maven-artifact-plugin/
[30]: https://maven.apache.org/plugins/maven-assembly-plugin/
[31]: https://maven.apache.org/plugins/maven-jar-plugin/
[32]: https://github.com/exasol/artifact-reference-checker-maven-plugin/
[33]: https://github.com/exasol/artifact-reference-checker-maven-plugin/blob/main/LICENSE
[34]: https://maven.apache.org/surefire/maven-failsafe-plugin/
[35]: https://www.jacoco.org/jacoco/trunk/doc/maven.html
[36]: https://www.eclipse.org/legal/epl-2.0/
[37]: https://github.com/exasol/quality-summarizer-maven-plugin/
[38]: https://github.com/exasol/quality-summarizer-maven-plugin/blob/main/LICENSE
[39]: https://github.com/exasol/error-code-crawler-maven-plugin/
[40]: https://github.com/exasol/error-code-crawler-maven-plugin/blob/main/LICENSE
[41]: https://github.com/git-commit-id/git-commit-id-maven-plugin
[42]: http://www.gnu.org/licenses/lgpl-3.0.txt
[43]: https://maven.apache.org/plugins/maven-clean-plugin/
[44]: https://maven.apache.org/plugins/maven-resources-plugin/
[45]: https://maven.apache.org/plugins/maven-install-plugin/
[46]: https://maven.apache.org/plugins/maven-site-plugin/
[47]: https://github.com/exasol/exasol-driver-go/blob/v1.0.16/LICENSE
[48]: https://github.com/stretchr/testify/blob/v1.11.1/LICENSE
[49]: https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE
[50]: https://github.com/antchfx/xmlquery/blob/HEAD/LICENSE
