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

### Runtime Dependencies

| Dependency                | License           |
| ------------------------- | ----------------- |
| [SLF4J JDK14 Binding][16] | [MIT License][17] |

### Plugin Dependencies

| Dependency                                              | License                                       |
| ------------------------------------------------------- | --------------------------------------------- |
| [SonarQube Scanner for Maven][18]                       | [GNU LGPL 3][19]                              |
| [Apache Maven Compiler Plugin][20]                      | [Apache License, Version 2.0][21]             |
| [Apache Maven Enforcer Plugin][22]                      | [Apache-2.0][21]                              |
| [Maven Flatten Plugin][23]                              | [Apache Software Licenese][21]                |
| [Exec Maven Plugin][24]                                 | [Apache License 2][21]                        |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][25] | [ASL2][5]                                     |
| [Maven Surefire Plugin][26]                             | [Apache License, Version 2.0][21]             |
| [Versions Maven Plugin][27]                             | [Apache License, Version 2.0][21]             |
| [Apache Maven Assembly Plugin][28]                      | [Apache License, Version 2.0][21]             |
| [Apache Maven JAR Plugin][29]                           | [Apache License, Version 2.0][21]             |
| [Artifact reference checker and unifier][30]            | [MIT License][31]                             |
| [Maven Failsafe Plugin][32]                             | [Apache License, Version 2.0][21]             |
| [JaCoCo :: Maven Plugin][33]                            | [Eclipse Public License 2.0][34]              |
| [error-code-crawler-maven-plugin][35]                   | [MIT License][36]                             |
| [Reproducible Build Maven Plugin][37]                   | [Apache 2.0][5]                               |
| [Maven Clean Plugin][38]                                | [The Apache Software License, Version 2.0][5] |
| [Maven Resources Plugin][39]                            | [The Apache Software License, Version 2.0][5] |
| [Maven Install Plugin][40]                              | [The Apache Software License, Version 2.0][5] |
| [Maven Deploy Plugin][41]                               | [The Apache Software License, Version 2.0][5] |
| [Maven Site Plugin 3][42]                               | [The Apache Software License, Version 2.0][5] |

## Go-client

### Compile Dependencies

| Dependency                         | License   |
| ---------------------------------- | --------- |
| github.com/exasol/exasol-driver-go | [MIT][43] |

### Test Dependencies

| Dependency                  | License   |
| --------------------------- | --------- |
| github.com/antchfx/xmlquery | [MIT][44] |
| github.com/stretchr/testify | [MIT][45] |
| gopkg.in/yaml.v3            | [MIT][46] |

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
[16]: http://www.slf4j.org
[17]: http://www.opensource.org/licenses/mit-license.php
[18]: http://sonarsource.github.io/sonar-scanner-maven/
[19]: http://www.gnu.org/licenses/lgpl.txt
[20]: https://maven.apache.org/plugins/maven-compiler-plugin/
[21]: https://www.apache.org/licenses/LICENSE-2.0.txt
[22]: https://maven.apache.org/enforcer/maven-enforcer-plugin/
[23]: https://www.mojohaus.org/flatten-maven-plugin/
[24]: https://www.mojohaus.org/exec-maven-plugin
[25]: https://sonatype.github.io/ossindex-maven/maven-plugin/
[26]: https://maven.apache.org/surefire/maven-surefire-plugin/
[27]: https://www.mojohaus.org/versions/versions-maven-plugin/
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
[38]: http://maven.apache.org/plugins/maven-clean-plugin/
[39]: http://maven.apache.org/plugins/maven-resources-plugin/
[40]: http://maven.apache.org/plugins/maven-install-plugin/
[41]: http://maven.apache.org/plugins/maven-deploy-plugin/
[42]: http://maven.apache.org/plugins/maven-site-plugin/
[43]: https://github.com/exasol/exasol-driver-go/blob/v0.4.7/LICENSE
[44]: https://github.com/antchfx/xmlquery/blob/HEAD/LICENSE
[45]: https://github.com/stretchr/testify/blob/HEAD/LICENSE
[46]: https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE
