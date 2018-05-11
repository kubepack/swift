---
title: Changelog | Swift
description: Changelog
menu:
  product_swift_0.8.1:
    identifier: changelog-swift
    name: Changelog
    parent: welcome
    weight: 10
product_name: swift
menu_name: product_swift_0.8.1
section_menu_id: welcome
url: /products/swift/0.8.1/welcome/changelog/
aliases:
  - /products/swift/0.8.1/CHANGELOG/
---

# Change Log

## [Unreleased](https://github.com/appscode/swift/tree/HEAD)

[Full Changelog](https://github.com/appscode/swift/compare/0.8.0...HEAD)

**Closed issues:**

- Improve docs [\#120](https://github.com/appscode/swift/issues/120)

**Merged pull requests:**

- Fix chart flags [\#130](https://github.com/appscode/swift/pull/130) ([tamalsaha](https://github.com/tamalsaha))
- Rename analytics to enable-analytics [\#127](https://github.com/appscode/swift/pull/127) ([diptadas](https://github.com/diptadas))
- Document how to deploy in custom namespace [\#126](https://github.com/appscode/swift/pull/126) ([tamalsaha](https://github.com/tamalsaha))
- Print swift installer namespace [\#125](https://github.com/appscode/swift/pull/125) ([tamalsaha](https://github.com/tamalsaha))

## [0.8.0](https://github.com/appscode/swift/tree/0.8.0) (2018-04-28)
[Full Changelog](https://github.com/appscode/swift/compare/0.7.3...0.8.0)

**Merged pull requests:**

- Fix installer [\#124](https://github.com/appscode/swift/pull/124) ([tamalsaha](https://github.com/tamalsaha))
- Prepare docs for 0.8.0 [\#123](https://github.com/appscode/swift/pull/123) ([tamalsaha](https://github.com/tamalsaha))
- Vendor Helm 2.9.0 [\#122](https://github.com/appscode/swift/pull/122) ([tamalsaha](https://github.com/tamalsaha))

## [0.7.3](https://github.com/appscode/swift/tree/0.7.3) (2018-04-21)
[Full Changelog](https://github.com/appscode/swift/compare/0.7.2...0.7.3)

## [0.7.2](https://github.com/appscode/swift/tree/0.7.2) (2018-04-21)
[Full Changelog](https://github.com/appscode/swift/compare/0.7.1...0.7.2)

## [0.7.1](https://github.com/appscode/swift/tree/0.7.1) (2018-04-21)
[Full Changelog](https://github.com/appscode/swift/compare/0.7.0...0.7.1)

**Fixed bugs:**

- Update flags for tiller [\#110](https://github.com/appscode/swift/pull/110) ([tamalsaha](https://github.com/tamalsaha))

**Closed issues:**

- support for multi-namespace chart installation [\#118](https://github.com/appscode/swift/issues/118)
- Failed to connect to tiller [\#115](https://github.com/appscode/swift/issues/115)
- unable to setup tiller-tls on 0.7.2 [\#109](https://github.com/appscode/swift/issues/109)

**Merged pull requests:**

- Update chart repository location [\#121](https://github.com/appscode/swift/pull/121) ([tamalsaha](https://github.com/tamalsaha))
- Support installing from local installer scripts [\#119](https://github.com/appscode/swift/pull/119) ([tamalsaha](https://github.com/tamalsaha))
- Move trap command to before onessl download steps [\#117](https://github.com/appscode/swift/pull/117) ([tamalsaha](https://github.com/tamalsaha))
- Update install instruction in RBAC doc [\#116](https://github.com/appscode/swift/pull/116) ([tamalsaha](https://github.com/tamalsaha))
- Add travis yaml [\#114](https://github.com/appscode/swift/pull/114) ([tahsinrahman](https://github.com/tahsinrahman))
- Updated chart version [\#113](https://github.com/appscode/swift/pull/113) ([diptadas](https://github.com/diptadas))
- Prepare docs for 0.7.3 [\#112](https://github.com/appscode/swift/pull/112) ([tamalsaha](https://github.com/tamalsaha))
- Search for tiller in self namespace before searching cluster. [\#111](https://github.com/appscode/swift/pull/111) ([tamalsaha](https://github.com/tamalsaha))
- Make it clear that installer is a single command [\#108](https://github.com/appscode/swift/pull/108) ([tamalsaha](https://github.com/tamalsaha))
- Fix installer [\#107](https://github.com/appscode/swift/pull/107) ([tamalsaha](https://github.com/tamalsaha))
- Update chart to match RBAC best practices for charts [\#106](https://github.com/appscode/swift/pull/106) ([tamalsaha](https://github.com/tamalsaha))
- Update chart version [\#105](https://github.com/appscode/swift/pull/105) ([tamalsaha](https://github.com/tamalsaha))
- Add OWNERS file for chart [\#104](https://github.com/appscode/swift/pull/104) ([tamalsaha](https://github.com/tamalsaha))
- Prepare docs for 0.7.2 release [\#103](https://github.com/appscode/swift/pull/103) ([tamalsaha](https://github.com/tamalsaha))
- Use glog middleware for logging [\#102](https://github.com/appscode/swift/pull/102) ([tamalsaha](https://github.com/tamalsaha))
- Rename validator method [\#101](https://github.com/appscode/swift/pull/101) ([tamalsaha](https://github.com/tamalsaha))
- Use appscode/grpc-go-addons [\#100](https://github.com/appscode/swift/pull/100) ([tamalsaha](https://github.com/tamalsaha))
- Use github.com/pkg/errors [\#99](https://github.com/appscode/swift/pull/99) ([tamalsaha](https://github.com/tamalsaha))
- Update changelog for 0.7.1 [\#98](https://github.com/appscode/swift/pull/98) ([tamalsaha](https://github.com/tamalsaha))
- Add --tiller-timeout flag with default 5 min deadline [\#97](https://github.com/appscode/swift/pull/97) ([tamalsaha](https://github.com/tamalsaha))
- Rename factory package to connectors [\#96](https://github.com/appscode/swift/pull/96) ([tamalsaha](https://github.com/tamalsaha))

## [0.7.0](https://github.com/appscode/swift/tree/0.7.0) (2018-02-14)
[Full Changelog](https://github.com/appscode/swift/compare/0.6.0...0.7.0)

**Closed issues:**

- Support basic auth for Chart repository [\#88](https://github.com/appscode/swift/issues/88)
- Get release history seems not work well [\#87](https://github.com/appscode/swift/issues/87)
- SSL support [\#86](https://github.com/appscode/swift/issues/86)
- Feature Request: Add ability to pass values.yaml into install and upgrade endpoints [\#76](https://github.com/appscode/swift/issues/76)
- Document how to use tiller-js-client [\#16](https://github.com/appscode/swift/issues/16)

**Merged pull requests:**

- Update installer script [\#95](https://github.com/appscode/swift/pull/95) ([tamalsaha](https://github.com/tamalsaha))
- Prepare docs for 0.7.0 release [\#94](https://github.com/appscode/swift/pull/94) ([tamalsaha](https://github.com/tamalsaha))
- Document Swift SSL options [\#93](https://github.com/appscode/swift/pull/93) ([tamalsaha](https://github.com/tamalsaha))
- Support self-signed ca certificate for Tiller [\#92](https://github.com/appscode/swift/pull/92) ([tamalsaha](https://github.com/tamalsaha))
- Do not write SETTINGS in response to ACKs in Cmux [\#91](https://github.com/appscode/swift/pull/91) ([tamalsaha](https://github.com/tamalsaha))
- Support SSL for chart repository [\#90](https://github.com/appscode/swift/pull/90) ([tamalsaha](https://github.com/tamalsaha))
- Pass username/password in chart URL as basic auth header [\#89](https://github.com/appscode/swift/pull/89) ([tamalsaha](https://github.com/tamalsaha))
- Update grpc-go to v1.9.2 [\#85](https://github.com/appscode/swift/pull/85) ([tamalsaha](https://github.com/tamalsaha))

## [0.6.0](https://github.com/appscode/swift/tree/0.6.0) (2018-01-24)
[Full Changelog](https://github.com/appscode/swift/compare/0.5.2...0.6.0)

**Merged pull requests:**

- Prepare docs for 0.6.0 [\#84](https://github.com/appscode/swift/pull/84) ([tamalsaha](https://github.com/tamalsaha))
- Revendor to Helm 2.8 [\#83](https://github.com/appscode/swift/pull/83) ([tamalsaha](https://github.com/tamalsaha))

## [0.5.2](https://github.com/appscode/swift/tree/0.5.2) (2018-01-06)
[Full Changelog](https://github.com/appscode/swift/compare/0.5.1...0.5.2)

**Fixed bugs:**

- Clone tunnel after api call [\#79](https://github.com/appscode/swift/pull/79) ([tamalsaha](https://github.com/tamalsaha))

**Closed issues:**

- Can not lookup tiller host [\#74](https://github.com/appscode/swift/issues/74)
- How can I install a release by specifying a local chart folder ? [\#67](https://github.com/appscode/swift/issues/67)
- Response Object for API Error [\#31](https://github.com/appscode/swift/issues/31)

**Merged pull requests:**

- Add changelog for 0.5.2 [\#82](https://github.com/appscode/swift/pull/82) ([tamalsaha](https://github.com/tamalsaha))
- Add front matter for docs [\#81](https://github.com/appscode/swift/pull/81) ([sajibcse68](https://github.com/sajibcse68))
- Fix analytics client id detection [\#80](https://github.com/appscode/swift/pull/80) ([tamalsaha](https://github.com/tamalsaha))
- Use tunnel tools from kutil [\#78](https://github.com/appscode/swift/pull/78) ([tamalsaha](https://github.com/tamalsaha))
- Set ClientID for analytics [\#77](https://github.com/appscode/swift/pull/77) ([tamalsaha](https://github.com/tamalsaha))
- Change left\_menu -\> menu\_name [\#73](https://github.com/appscode/swift/pull/73) ([sajibcse68](https://github.com/sajibcse68))
- Add aliases for README files [\#72](https://github.com/appscode/swift/pull/72) ([sajibcse68](https://github.com/sajibcse68))

## [0.5.1](https://github.com/appscode/swift/tree/0.5.1) (2017-11-28)
[Full Changelog](https://github.com/appscode/swift/compare/0.5.0...0.5.1)

**Closed issues:**

- Can not pass a range value as '--set' argument with swift proxy? [\#70](https://github.com/appscode/swift/issues/70)
- Connection leak? [\#66](https://github.com/appscode/swift/issues/66)
- Wrong yamls file addresses in install.md [\#62](https://github.com/appscode/swift/issues/62)
- Swift error: Failed to extract ServerMetadata from context [\#51](https://github.com/appscode/swift/issues/51)

**Merged pull requests:**

- Prepare docs for 0.5.1 release [\#71](https://github.com/appscode/swift/pull/71) ([tamalsaha](https://github.com/tamalsaha))
- Add front matter for swift cli [\#69](https://github.com/appscode/swift/pull/69) ([tamalsaha](https://github.com/tamalsaha))
- Close connection after usage. [\#68](https://github.com/appscode/swift/pull/68) ([tamalsaha](https://github.com/tamalsaha))
- Add Front matter of docs [\#65](https://github.com/appscode/swift/pull/65) ([sajibcse68](https://github.com/sajibcse68))
- Make chart namespaced [\#64](https://github.com/appscode/swift/pull/64) ([tamalsaha](https://github.com/tamalsaha))

## [0.5.0](https://github.com/appscode/swift/tree/0.5.0) (2017-10-30)
[Full Changelog](https://github.com/appscode/swift/compare/0.4.0...0.5.0)

**Merged pull requests:**

- Add changelog for 0.5.0 [\#63](https://github.com/appscode/swift/pull/63) ([tamalsaha](https://github.com/tamalsaha))
- Support listing releases from all namespaces [\#61](https://github.com/appscode/swift/pull/61) ([tamalsaha](https://github.com/tamalsaha))
- Add screenshot of release list [\#60](https://github.com/appscode/swift/pull/60) ([tamalsaha](https://github.com/tamalsaha))
- Add tutorial for RBAC enabled cluster [\#59](https://github.com/appscode/swift/pull/59) ([tamalsaha](https://github.com/tamalsaha))
- Fix service port in installer yamls [\#58](https://github.com/appscode/swift/pull/58) ([tamalsaha](https://github.com/tamalsaha))

## [0.4.0](https://github.com/appscode/swift/tree/0.4.0) (2017-10-24)
[Full Changelog](https://github.com/appscode/swift/compare/0.4.0-rc.0...0.4.0)

**Merged pull requests:**

- Add changelog for 0.4.0 [\#57](https://github.com/appscode/swift/pull/57) ([tamalsaha](https://github.com/tamalsaha))

## [0.4.0-rc.0](https://github.com/appscode/swift/tree/0.4.0-rc.0) (2017-10-16)
[Full Changelog](https://github.com/appscode/swift/compare/0.3.1...0.4.0-rc.0)

**Closed issues:**

- Network is unreachable [\#44](https://github.com/appscode/swift/issues/44)

**Merged pull requests:**

- Fix build [\#56](https://github.com/appscode/swift/pull/56) ([tamalsaha](https://github.com/tamalsaha))
- Truncate Chart helper values to 63 chars [\#55](https://github.com/appscode/swift/pull/55) ([tamalsaha](https://github.com/tamalsaha))
- Prepare 0.4.0-rc.0 release [\#54](https://github.com/appscode/swift/pull/54) ([tamalsaha](https://github.com/tamalsaha))
- Use client-go 5.x [\#53](https://github.com/appscode/swift/pull/53) ([tamalsaha](https://github.com/tamalsaha))
- Use log pkg from appscode/go [\#52](https://github.com/appscode/swift/pull/52) ([tamalsaha](https://github.com/tamalsaha))
- Add chart [\#50](https://github.com/appscode/swift/pull/50) ([tamalsaha](https://github.com/tamalsaha))

## [0.3.1](https://github.com/appscode/swift/tree/0.3.1) (2017-09-21)
[Full Changelog](https://github.com/appscode/swift/compare/0.3.0...0.3.1)

**Fixed bugs:**

- Set service account name in RBAC mode. [\#48](https://github.com/appscode/swift/pull/48) ([tamalsaha](https://github.com/tamalsaha))
- Update installer RBAC to support listing services [\#47](https://github.com/appscode/swift/pull/47) ([tamalsaha](https://github.com/tamalsaha))

**Merged pull requests:**

- Prepare docs for 0.3.1 release. [\#49](https://github.com/appscode/swift/pull/49) ([tamalsaha](https://github.com/tamalsaha))

## [0.3.0](https://github.com/appscode/swift/tree/0.3.0) (2017-09-21)
[Full Changelog](https://github.com/appscode/swift/compare/0.2.0...0.3.0)

**Merged pull requests:**

- Add changelog for 0.3.0 [\#46](https://github.com/appscode/swift/pull/46) ([tamalsaha](https://github.com/tamalsaha))
- Check for returned pods or services before connecting. [\#45](https://github.com/appscode/swift/pull/45) ([tamalsaha](https://github.com/tamalsaha))
- Use client-go 4.0.0 [\#43](https://github.com/appscode/swift/pull/43) ([tamalsaha](https://github.com/tamalsaha))
- Fix command in Developer-guide [\#42](https://github.com/appscode/swift/pull/42) ([the-redback](https://github.com/the-redback))
- Move analytics to common GA project [\#41](https://github.com/appscode/swift/pull/41) ([tamalsaha](https://github.com/tamalsaha))

## [0.2.0](https://github.com/appscode/swift/tree/0.2.0) (2017-09-08)
[Full Changelog](https://github.com/appscode/swift/compare/0.2.0-rc.0...0.2.0)

**Fixed bugs:**

- Remove streaming ListReleases [\#24](https://github.com/appscode/swift/issues/24)

**Closed issues:**

- can not create pvc [\#40](https://github.com/appscode/swift/issues/40)
- Add endpoint to get values from release [\#23](https://github.com/appscode/swift/issues/23)
- Mount an EmptyDir as scratch volume [\#35](https://github.com/appscode/swift/issues/35)

## [0.2.0-rc.0](https://github.com/appscode/swift/tree/0.2.0-rc.0) (2017-08-31)
[Full Changelog](https://github.com/appscode/swift/compare/0.1.0...0.2.0-rc.0)

**Fixed bugs:**

- Purge api not working? [\#28](https://github.com/appscode/swift/issues/28)

**Closed issues:**

- List all releases [\#32](https://github.com/appscode/swift/issues/32)
- Support Dry run  [\#30](https://github.com/appscode/swift/issues/30)
- Support YAML or JSON in values [\#29](https://github.com/appscode/swift/issues/29)
- Document Helm version dependency [\#17](https://github.com/appscode/swift/issues/17)
- build on osx [\#9](https://github.com/appscode/swift/issues/9)

**Merged pull requests:**

- Add changelog [\#39](https://github.com/appscode/swift/pull/39) ([diptadas](https://github.com/diptadas))
- Update readme [\#38](https://github.com/appscode/swift/pull/38) ([diptadas](https://github.com/diptadas))
- Prepare docs for 0.2.0-rc.0 release [\#37](https://github.com/appscode/swift/pull/37) ([tamalsaha](https://github.com/tamalsaha))
- Make chart-dir temporary & mount tmp dir [\#36](https://github.com/appscode/swift/pull/36) ([diptadas](https://github.com/diptadas))
- Format release content raw [\#34](https://github.com/appscode/swift/pull/34) ([diptadas](https://github.com/diptadas))
- List all releases [\#33](https://github.com/appscode/swift/pull/33) ([diptadas](https://github.com/diptadas))
- Remove ListReleases streaming api from readme. [\#27](https://github.com/appscode/swift/pull/27) ([tamalsaha](https://github.com/tamalsaha))
- Remove ListReleases api [\#26](https://github.com/appscode/swift/pull/26) ([tamalsaha](https://github.com/tamalsaha))
- Remove ListReleases api [\#25](https://github.com/appscode/swift/pull/25) ([tamalsaha](https://github.com/tamalsaha))
- Rename wheel [\#22](https://github.com/appscode/swift/pull/22) ([tamalsaha](https://github.com/tamalsaha))
- \[js client\] Fix main exports [\#21](https://github.com/appscode/swift/pull/21) ([tamalsaha](https://github.com/tamalsaha))
- Document Helm version [\#18](https://github.com/appscode/swift/pull/18) ([tamalsaha](https://github.com/tamalsaha))

## [0.1.0](https://github.com/appscode/swift/tree/0.1.0) (2017-08-06)
**Merged pull requests:**

- Document connectors [\#15](https://github.com/appscode/swift/pull/15) ([tamalsaha](https://github.com/tamalsaha))
- Update tiller-js-client [\#14](https://github.com/appscode/swift/pull/14) ([tamalsaha](https://github.com/tamalsaha))
- Update docs [\#13](https://github.com/appscode/swift/pull/13) ([tamalsaha](https://github.com/tamalsaha))
- Stop running proto generator builddeps everytime. [\#12](https://github.com/appscode/swift/pull/12) ([tamalsaha](https://github.com/tamalsaha))
- Add test chart [\#11](https://github.com/appscode/swift/pull/11) ([tamalsaha](https://github.com/tamalsaha))
- Fix builddeps.sh script [\#10](https://github.com/appscode/swift/pull/10) ([tamalsaha](https://github.com/tamalsaha))
- Search all namespaces for Tiller [\#8](https://github.com/appscode/swift/pull/8) ([tamalsaha](https://github.com/tamalsaha))
- Support icluster, direct and kubeconfig connector for Tiller server [\#7](https://github.com/appscode/swift/pull/7) ([tamalsaha](https://github.com/tamalsaha))
- Setup server stuff [\#6](https://github.com/appscode/swift/pull/6) ([tamalsaha](https://github.com/tamalsaha))
- Generate tiller-js-client [\#5](https://github.com/appscode/swift/pull/5) ([tamalsaha](https://github.com/tamalsaha))
- Add developer guide [\#4](https://github.com/appscode/swift/pull/4) ([tamalsaha](https://github.com/tamalsaha))
- Apply various cleanup [\#3](https://github.com/appscode/swift/pull/3) ([tamalsaha](https://github.com/tamalsaha))
- Implement Tiller release service [\#2](https://github.com/appscode/swift/pull/2) ([diptadas](https://github.com/diptadas))
- Add docker build scripts [\#1](https://github.com/appscode/swift/pull/1) ([tamalsaha](https://github.com/tamalsaha))



\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*