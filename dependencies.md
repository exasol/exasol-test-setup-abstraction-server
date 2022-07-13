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
| [Hamcrest][8]                  | [BSD License 3][9]                |
| [mockito-junit-jupiter][10]    | [The MIT License][11]             |
| [REST Assured][12]             | [Apache 2.0][13]                  |
| [Apache Commons Codec][14]     | [Apache License, Version 2.0][15] |
| [JUnit5 System Extensions][16] | [Eclipse Public License v2.0][17] |

### Plugin Dependencies

| Dependency                                              | License                                       |
| ------------------------------------------------------- | --------------------------------------------- |
| [SonarQube Scanner for Maven][18]                       | [GNU LGPL 3][19]                              |
| [Apache Maven Compiler Plugin][20]                      | [Apache License, Version 2.0][15]             |
| [Apache Maven Enforcer Plugin][22]                      | [Apache License, Version 2.0][15]             |
| [Maven Flatten Plugin][24]                              | [Apache Software Licenese][3]                 |
| [Exec Maven Plugin][26]                                 | [Apache License 2][3]                         |
| [org.sonatype.ossindex.maven:ossindex-maven-plugin][28] | [ASL2][3]                                     |
| [Reproducible Build Maven Plugin][30]                   | [Apache 2.0][3]                               |
| [Maven Surefire Plugin][32]                             | [Apache License, Version 2.0][15]             |
| [Versions Maven Plugin][34]                             | [Apache License, Version 2.0][15]             |
| [Apache Maven Assembly Plugin][36]                      | [Apache License, Version 2.0][15]             |
| [Apache Maven JAR Plugin][38]                           | [Apache License, Version 2.0][15]             |
| [Artifact reference checker and unifier][40]            | [MIT][41]                                     |
| [Maven Failsafe Plugin][42]                             | [Apache License, Version 2.0][15]             |
| [JaCoCo :: Maven Plugin][44]                            | [Eclipse Public License 2.0][45]              |
| [error-code-crawler-maven-plugin][46]                   | [MIT][41]                                     |
| [Maven Clean Plugin][48]                                | [The Apache Software License, Version 2.0][3] |
| [Maven Resources Plugin][50]                            | [The Apache Software License, Version 2.0][3] |
| [Maven Install Plugin][52]                              | [The Apache Software License, Version 2.0][3] |
| [Maven Deploy Plugin][54]                               | [The Apache Software License, Version 2.0][3] |
| [Maven Site Plugin 3][56]                               | [The Apache Software License, Version 2.0][3] |

## Go-client

### Compile Dependencies

| Dependency                         | License   |
| ---------------------------------- | --------- |
| github.com/exasol/exasol-driver-go | [MIT][58] |

### Test Dependencies

| Dependency                  | License   |
| --------------------------- | --------- |
| github.com/antchfx/xmlquery | [MIT][59] |
| github.com/stretchr/testify | [MIT][60] |

[17]: http://www.eclipse.org/legal/epl-v20.html
[3]: http://www.apache.org/licenses/LICENSE-2.0.txt
[32]: https://maven.apache.org/surefire/maven-surefire-plugin/
[13]: http://www.apache.org/licenses/LICENSE-2.0.html
[48]: http://maven.apache.org/plugins/maven-clean-plugin/
[10]: https://github.com/mockito/mockito
[41]: https://opensource.org/licenses/MIT
[42]: https://maven.apache.org/surefire/maven-failsafe-plugin/
[24]: https://www.mojohaus.org/flatten-maven-plugin/
[26]: http://www.mojohaus.org/exec-maven-plugin
[14]: https://commons.apache.org/proper/commons-codec/
[34]: http://www.mojohaus.org/versions-maven-plugin/
[9]: http://opensource.org/licenses/BSD-3-Clause
[20]: https://maven.apache.org/plugins/maven-compiler-plugin/
[60]: https://github.com/stretchr/testify/blob/v1.8.0/LICENSE
[0]: https://github.com/exasol/exasol-test-setup-abstraction-java/
[45]: https://www.eclipse.org/legal/epl-2.0/
[19]: http://www.gnu.org/licenses/lgpl.txt
[44]: https://www.jacoco.org/jacoco/trunk/doc/maven.html
[12]: http://code.google.com/p/rest-assured
[11]: https://github.com/mockito/mockito/blob/main/LICENSE
[30]: http://zlika.github.io/reproducible-build-maven-plugin
[1]: https://github.com/exasol/exasol-test-setup-abstraction-java/blob/main/LICENSE
[58]: https://github.com/exasol/exasol-driver-go/blob/v0.4.3/LICENSE
[15]: https://www.apache.org/licenses/LICENSE-2.0.txt
[18]: http://sonarsource.github.io/sonar-scanner-maven/
[22]: https://maven.apache.org/enforcer/maven-enforcer-plugin/
[5]: https://www.eclipse.org/legal/epl-v20.html
[52]: http://maven.apache.org/plugins/maven-install-plugin/
[59]: https://github.com/antchfx/xmlquery/blob/v1.3.11/LICENSE
[4]: https://junit.org/junit5/
[28]: https://sonatype.github.io/ossindex-maven/maven-plugin/
[16]: https://github.com/itsallcode/junit5-system-extensions
[2]: https://javalin.io/
[8]: http://hamcrest.org/JavaHamcrest/
[54]: http://maven.apache.org/plugins/maven-deploy-plugin/
[56]: http://maven.apache.org/plugins/maven-site-plugin/
[50]: http://maven.apache.org/plugins/maven-resources-plugin/
[40]: https://github.com/exasol/artifact-reference-checker-maven-plugin
[46]: https://github.com/exasol/error-code-crawler-maven-plugin
[38]: https://maven.apache.org/plugins/maven-jar-plugin/
[36]: https://maven.apache.org/plugins/maven-assembly-plugin/
