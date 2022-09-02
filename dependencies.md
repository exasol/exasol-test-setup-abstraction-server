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
| [Apache Commons Codec][12]     | [Apache License, Version 2.0][13] |
| [JUnit5 System Extensions][14] | [Eclipse Public License v2.0][15] |

### Plugin Dependencies

| Dependency                                              | License                                       |
| ------------------------------------------------------- | --------------------------------------------- |
| [SonarQube Scanner for Maven][16]                       | [GNU LGPL 3][17]                              |
| [Apache Maven Compiler Plugin][18]                      | [Apache License, Version 2.0][13]             |
| [Apache Maven Enforcer Plugin][19]                      | [Apache License, Version 2.0][13]             |
| [Maven Flatten Plugin][20]                              | [Apache Software Licenese][3]                 |
| [Exec Maven Plugin][21]                                 | [Apache License 2][3]                         |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][22] | [ASL2][3]                                     |
| [Maven Surefire Plugin][23]                             | [Apache License, Version 2.0][13]             |
| [Versions Maven Plugin][24]                             | [Apache License, Version 2.0][13]             |
| [Apache Maven Assembly Plugin][25]                      | [Apache License, Version 2.0][13]             |
| [Apache Maven JAR Plugin][26]                           | [Apache License, Version 2.0][13]             |
| [Artifact reference checker and unifier][27]            | [MIT][28]                                     |
| [Maven Failsafe Plugin][29]                             | [Apache License, Version 2.0][13]             |
| [JaCoCo :: Maven Plugin][30]                            | [Eclipse Public License 2.0][31]              |
| [error-code-crawler-maven-plugin][32]                   | [MIT License][33]                             |
| [Reproducible Build Maven Plugin][34]                   | [Apache 2.0][3]                               |
| [Maven Clean Plugin][35]                                | [The Apache Software License, Version 2.0][3] |
| [Maven Resources Plugin][36]                            | [The Apache Software License, Version 2.0][3] |
| [Maven Install Plugin][37]                              | [The Apache Software License, Version 2.0][3] |
| [Maven Deploy Plugin][38]                               | [The Apache Software License, Version 2.0][3] |
| [Maven Site Plugin 3][39]                               | [The Apache Software License, Version 2.0][3] |

## Go-client

### Compile Dependencies

| Dependency                         | License   |
| ---------------------------------- | --------- |
| github.com/exasol/exasol-driver-go | [MIT][40] |

### Test Dependencies

| Dependency                  | License   |
| --------------------------- | --------- |
| github.com/antchfx/xmlquery | [MIT][41] |
| github.com/stretchr/testify | [MIT][42] |
| gopkg.in/yaml.v3            | [MIT][43] |

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
[12]: https://commons.apache.org/proper/commons-codec/
[13]: https://www.apache.org/licenses/LICENSE-2.0.txt
[14]: https://github.com/itsallcode/junit5-system-extensions
[15]: http://www.eclipse.org/legal/epl-v20.html
[16]: http://sonarsource.github.io/sonar-scanner-maven/
[17]: http://www.gnu.org/licenses/lgpl.txt
[18]: https://maven.apache.org/plugins/maven-compiler-plugin/
[19]: https://maven.apache.org/enforcer/maven-enforcer-plugin/
[20]: https://www.mojohaus.org/flatten-maven-plugin/
[21]: http://www.mojohaus.org/exec-maven-plugin
[22]: https://sonatype.github.io/ossindex-maven/maven-plugin/
[23]: https://maven.apache.org/surefire/maven-surefire-plugin/
[24]: http://www.mojohaus.org/versions-maven-plugin/
[25]: https://maven.apache.org/plugins/maven-assembly-plugin/
[26]: https://maven.apache.org/plugins/maven-jar-plugin/
[27]: https://github.com/exasol/artifact-reference-checker-maven-plugin
[28]: https://opensource.org/licenses/MIT
[29]: https://maven.apache.org/surefire/maven-failsafe-plugin/
[30]: https://www.jacoco.org/jacoco/trunk/doc/maven.html
[31]: https://www.eclipse.org/legal/epl-2.0/
[32]: https://github.com/exasol/error-code-crawler-maven-plugin/
[33]: https://github.com/exasol/error-code-crawler-maven-plugin/blob/main/LICENSE
[34]: http://zlika.github.io/reproducible-build-maven-plugin
[35]: http://maven.apache.org/plugins/maven-clean-plugin/
[36]: http://maven.apache.org/plugins/maven-resources-plugin/
[37]: http://maven.apache.org/plugins/maven-install-plugin/
[38]: http://maven.apache.org/plugins/maven-deploy-plugin/
[39]: http://maven.apache.org/plugins/maven-site-plugin/
[40]: https://github.com/exasol/exasol-driver-go/blob/v0.4.5/LICENSE
[41]: https://github.com/antchfx/xmlquery/blob/HEAD/LICENSE
[42]: https://github.com/stretchr/testify/blob/HEAD/LICENSE
[43]: https://github.com/go-yaml/yaml/blob/v3.0.1/LICENSE
