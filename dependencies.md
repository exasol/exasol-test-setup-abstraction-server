<!-- @formatter:off -->
# Dependencies

## Server for the Exasol Test-setup Abstraction

### Compile Dependencies

| Dependency                                | License                                       |
| ----------------------------------------- | --------------------------------------------- |
| [exasol-test-setup-abstraction-java][0]   | [MIT License][1]                              |
| [Test containers for Exasol on Docker][2] | [MIT License][3]                              |
| [javalin][4]                              | [The Apache Software License, Version 2.0][5] |

### Test Dependencies

| Dependency                     | License                           |
| ------------------------------ | --------------------------------- |
| [JUnit Jupiter Engine][6]      | [Eclipse Public License v2.0][7]  |
| [JUnit Jupiter Params][6]      | [Eclipse Public License v2.0][7]  |
| [Hamcrest][8]                  | [BSD License 3][9]                |
| [mockito-junit-jupiter][10]    | [The MIT License][11]             |
| [REST Assured][12]             | [Apache 2.0][13]                  |
| [JUnit5 System Extensions][14] | [Eclipse Public License v2.0][15] |

### Plugin Dependencies

| Dependency                                              | License                                       |
| ------------------------------------------------------- | --------------------------------------------- |
| [SonarQube Scanner for Maven][16]                       | [GNU LGPL 3][17]                              |
| [Apache Maven Compiler Plugin][18]                      | [Apache License, Version 2.0][19]             |
| [Apache Maven Enforcer Plugin][20]                      | [Apache License, Version 2.0][19]             |
| [Maven Flatten Plugin][21]                              | [Apache Software Licenese][19]                |
| [Exec Maven Plugin][22]                                 | [Apache License 2][19]                        |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][23] | [ASL2][5]                                     |
| [Maven Surefire Plugin][24]                             | [Apache License, Version 2.0][19]             |
| [Versions Maven Plugin][25]                             | [Apache License, Version 2.0][19]             |
| [Apache Maven Assembly Plugin][26]                      | [Apache License, Version 2.0][19]             |
| [Apache Maven JAR Plugin][27]                           | [Apache License, Version 2.0][19]             |
| [Artifact reference checker and unifier][28]            | [MIT License][29]                             |
| [Maven Failsafe Plugin][30]                             | [Apache License, Version 2.0][19]             |
| [JaCoCo :: Maven Plugin][31]                            | [Eclipse Public License 2.0][32]              |
| [error-code-crawler-maven-plugin][33]                   | [MIT License][34]                             |
| [Reproducible Build Maven Plugin][35]                   | [Apache 2.0][5]                               |
| [Maven Clean Plugin][36]                                | [The Apache Software License, Version 2.0][5] |
| [Maven Resources Plugin][37]                            | [The Apache Software License, Version 2.0][5] |
| [Maven Install Plugin][38]                              | [The Apache Software License, Version 2.0][5] |
| [Maven Deploy Plugin][39]                               | [The Apache Software License, Version 2.0][5] |
| [Maven Site Plugin 3][40]                               | [The Apache Software License, Version 2.0][5] |

## Go-client

### Compile Dependencies

| Dependency                         | License   |
| ---------------------------------- | --------- |
| github.com/exasol/exasol-driver-go | [MIT][41] |

### Test Dependencies

| Dependency                  | License   |
| --------------------------- | --------- |
| github.com/antchfx/xmlquery | [MIT][42] |
| github.com/stretchr/testify | [MIT][43] |
| gopkg.in/yaml.v3            | [MIT][44] |

[0]: https://github.com/exasol/exasol-test-setup-abstraction-java/
[1]: https://github.com/exasol/exasol-test-setup-abstraction-java/blob/main/LICENSE
[2]: https://github.com/exasol/exasol-testcontainers/
[3]: https://github.com/exasol/exasol-testcontainers/blob/main/LICENSE
[4]: https://javalin.io/
[5]: http://www.apache.org/licenses/LICENSE-2.0.txt
[6]: https://junit.org/junit5/
[7]: https://www.eclipse.org/legal/epl-v20.html
[8]: http://hamcrest.org/JavaHamcrest/
[9]: http://opensource.org/licenses/BSD-3-Clause
[10]: https://github.com/mockito/mockito
[11]: https://github.com/mockito/mockito/blob/main/LICENSE
[12]: http://code.google.com/p/rest-assured
[13]: http://www.apache.org/licenses/LICENSE-2.0.html
[14]: https://github.com/itsallcode/junit5-system-extensions
[15]: http://www.eclipse.org/legal/epl-v20.html
[16]: http://sonarsource.github.io/sonar-scanner-maven/
[17]: http://www.gnu.org/licenses/lgpl.txt
[18]: https://maven.apache.org/plugins/maven-compiler-plugin/
[19]: https://www.apache.org/licenses/LICENSE-2.0.txt
[20]: https://maven.apache.org/enforcer/maven-enforcer-plugin/
[21]: https://www.mojohaus.org/flatten-maven-plugin/
[22]: https://www.mojohaus.org/exec-maven-plugin
[23]: https://sonatype.github.io/ossindex-maven/maven-plugin/
[24]: https://maven.apache.org/surefire/maven-surefire-plugin/
[25]: https://www.mojohaus.org/versions-maven-plugin/
[26]: https://maven.apache.org/plugins/maven-assembly-plugin/
[27]: https://maven.apache.org/plugins/maven-jar-plugin/
[28]: https://github.com/exasol/artifact-reference-checker-maven-plugin/
[29]: https://github.com/exasol/artifact-reference-checker-maven-plugin/blob/main/LICENSE
[30]: https://maven.apache.org/surefire/maven-failsafe-plugin/
[31]: https://www.jacoco.org/jacoco/trunk/doc/maven.html
[32]: https://www.eclipse.org/legal/epl-2.0/
[33]: https://github.com/exasol/error-code-crawler-maven-plugin/
[34]: https://github.com/exasol/error-code-crawler-maven-plugin/blob/main/LICENSE
[35]: http://zlika.github.io/reproducible-build-maven-plugin
[36]: http://maven.apache.org/plugins/maven-clean-plugin/
[37]: http://maven.apache.org/plugins/maven-resources-plugin/
[38]: http://maven.apache.org/plugins/maven-install-plugin/
[39]: http://maven.apache.org/plugins/maven-deploy-plugin/
[40]: http://maven.apache.org/plugins/maven-site-plugin/
[41]: https://github.com/exasol/exasol-driver-go/blob/v0.4.6/LICENSE
[42]: https://github.com/antchfx/xmlquery/blob/HEAD/LICENSE
[43]: https://github.com/stretchr/testify/blob/HEAD/LICENSE
[44]: https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE
