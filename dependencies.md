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
| [mockito-junit-jupiter][8]     | [The MIT License][9]              |
| [REST Assured][10]             | [Apache 2.0][11]                  |
| [JUnit5 System Extensions][12] | [Eclipse Public License v2.0][13] |

### Plugin Dependencies

| Dependency                                              | License                                       |
| ------------------------------------------------------- | --------------------------------------------- |
| [SonarQube Scanner for Maven][14]                       | [GNU LGPL 3][15]                              |
| [Apache Maven Compiler Plugin][16]                      | [Apache License, Version 2.0][17]             |
| [Apache Maven Enforcer Plugin][18]                      | [Apache License, Version 2.0][17]             |
| [Maven Flatten Plugin][19]                              | [Apache Software Licenese][3]                 |
| [Exec Maven Plugin][20]                                 | [Apache License 2][3]                         |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][21] | [ASL2][3]                                     |
| [Maven Surefire Plugin][22]                             | [Apache License, Version 2.0][17]             |
| [Versions Maven Plugin][23]                             | [Apache License, Version 2.0][17]             |
| [Apache Maven Assembly Plugin][24]                      | [Apache License, Version 2.0][17]             |
| [Apache Maven JAR Plugin][25]                           | [Apache License, Version 2.0][17]             |
| [Artifact reference checker and unifier][26]            | [MIT][27]                                     |
| [Maven Failsafe Plugin][28]                             | [Apache License, Version 2.0][17]             |
| [JaCoCo :: Maven Plugin][29]                            | [Eclipse Public License 2.0][30]              |
| [error-code-crawler-maven-plugin][31]                   | [MIT License][32]                             |
| [Reproducible Build Maven Plugin][33]                   | [Apache 2.0][3]                               |
| [Maven Clean Plugin][34]                                | [The Apache Software License, Version 2.0][3] |
| [Maven Resources Plugin][35]                            | [The Apache Software License, Version 2.0][3] |
| [Maven Install Plugin][36]                              | [The Apache Software License, Version 2.0][3] |
| [Maven Deploy Plugin][37]                               | [The Apache Software License, Version 2.0][3] |
| [Maven Site Plugin 3][38]                               | [The Apache Software License, Version 2.0][3] |

## Go-client

### Compile Dependencies

| Dependency                         | License   |
| ---------------------------------- | --------- |
| github.com/exasol/exasol-driver-go | [MIT][39] |

### Test Dependencies

| Dependency                  | License   |
| --------------------------- | --------- |
| github.com/antchfx/xmlquery | [MIT][40] |
| github.com/stretchr/testify | [MIT][41] |
| gopkg.in/yaml.v3            | [MIT][42] |

[0]: https://github.com/exasol/exasol-test-setup-abstraction-java/
[1]: https://github.com/exasol/exasol-test-setup-abstraction-java/blob/main/LICENSE
[2]: https://javalin.io/
[3]: http://www.apache.org/licenses/LICENSE-2.0.txt
[4]: https://junit.org/junit5/
[5]: https://www.eclipse.org/legal/epl-v20.html
[6]: http://hamcrest.org/JavaHamcrest/
[7]: http://opensource.org/licenses/BSD-3-Clause
[8]: https://github.com/mockito/mockito
[9]: https://github.com/mockito/mockito/blob/main/LICENSE
[10]: http://code.google.com/p/rest-assured
[11]: http://www.apache.org/licenses/LICENSE-2.0.html
[12]: https://github.com/itsallcode/junit5-system-extensions
[13]: http://www.eclipse.org/legal/epl-v20.html
[14]: http://sonarsource.github.io/sonar-scanner-maven/
[15]: http://www.gnu.org/licenses/lgpl.txt
[16]: https://maven.apache.org/plugins/maven-compiler-plugin/
[17]: https://www.apache.org/licenses/LICENSE-2.0.txt
[18]: https://maven.apache.org/enforcer/maven-enforcer-plugin/
[19]: https://www.mojohaus.org/flatten-maven-plugin/
[20]: http://www.mojohaus.org/exec-maven-plugin
[21]: https://sonatype.github.io/ossindex-maven/maven-plugin/
[22]: https://maven.apache.org/surefire/maven-surefire-plugin/
[23]: http://www.mojohaus.org/versions-maven-plugin/
[24]: https://maven.apache.org/plugins/maven-assembly-plugin/
[25]: https://maven.apache.org/plugins/maven-jar-plugin/
[26]: https://github.com/exasol/artifact-reference-checker-maven-plugin
[27]: https://opensource.org/licenses/MIT
[28]: https://maven.apache.org/surefire/maven-failsafe-plugin/
[29]: https://www.jacoco.org/jacoco/trunk/doc/maven.html
[30]: https://www.eclipse.org/legal/epl-2.0/
[31]: https://github.com/exasol/error-code-crawler-maven-plugin/
[32]: https://github.com/exasol/error-code-crawler-maven-plugin/blob/main/LICENSE
[33]: http://zlika.github.io/reproducible-build-maven-plugin
[34]: http://maven.apache.org/plugins/maven-clean-plugin/
[35]: http://maven.apache.org/plugins/maven-resources-plugin/
[36]: http://maven.apache.org/plugins/maven-install-plugin/
[37]: http://maven.apache.org/plugins/maven-deploy-plugin/
[38]: http://maven.apache.org/plugins/maven-site-plugin/
[39]: https://github.com/exasol/exasol-driver-go/blob/v0.4.6/LICENSE
[40]: https://github.com/antchfx/xmlquery/blob/HEAD/LICENSE
[41]: https://github.com/stretchr/testify/blob/HEAD/LICENSE
[42]: https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE
