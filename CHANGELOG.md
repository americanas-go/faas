# Changelog
All notable changes to this project will be documented in this file.

## [unreleased]

### Testing

- Add github workflows ([afd5847](https://github.com/americanas-go/faas/commit/afd58472959ce2960b2a21a3a41e34b4e92d9cae))

## [v1.0.0-beta.5](https://github.com/americanas-go/faas/compare/v1.0.0-beta.4...v1.0.0-beta.5) - 2021-09-29

[719c5a4](https://github.com/americanas-go/faas/commit/719c5a4e9c684b2f43564836a0300ec8a8e2f5ac)...[757d52b](https://github.com/americanas-go/faas/commit/757d52b06f4a9215b6d16b85a7dc711d487e2c0a)

### Bug Fixes

- Fix repository config ([757d52b](https://github.com/americanas-go/faas/commit/757d52b06f4a9215b6d16b85a7dc711d487e2c0a))

## [v1.0.0-beta.4](https://github.com/americanas-go/faas/compare/v1.0.0-beta.3...v1.0.0-beta.4) - 2021-09-29

[575c529](https://github.com/americanas-go/faas/commit/575c529d20cc3da422aa905f950a6fb0e73d0b68)...[719c5a4](https://github.com/americanas-go/faas/commit/719c5a4e9c684b2f43564836a0300ec8a8e2f5ac)

### Bug Fixes

- Fix plugins config path ([719c5a4](https://github.com/americanas-go/faas/commit/719c5a4e9c684b2f43564836a0300ec8a8e2f5ac))

## [v1.0.0-beta.3](https://github.com/americanas-go/faas/compare/v1.0.0-beta.2...v1.0.0-beta.3) - 2021-09-29

[cc98d81](https://github.com/americanas-go/faas/commit/cc98d81438fe97f4b4e20c26eab8dbb5468ce1ac)...[575c529](https://github.com/americanas-go/faas/commit/575c529d20cc3da422aa905f950a6fb0e73d0b68)

### Bug Fixes

- Fixes cloudevents helper ([575c529](https://github.com/americanas-go/faas/commit/575c529d20cc3da422aa905f950a6fb0e73d0b68))

### Features

- Add userIdentity at Record ([02e82a7](https://github.com/americanas-go/faas/commit/02e82a7afd3ce00787271e070098d2fd8223e781))

## [v1.0.0-beta.2](https://github.com/americanas-go/faas/compare/v1.0.0-beta.1...v1.0.0-beta.2) - 2021-08-25

[ff2e5e1](https://github.com/americanas-go/faas/commit/ff2e5e10baf293028d71994ae79e8847aacdcd6f)...[cc98d81](https://github.com/americanas-go/faas/commit/cc98d81438fe97f4b4e20c26eab8dbb5468ce1ac)

### Features

- Adding s3 event  to cloudevent converter (wip) ([26115fb](https://github.com/americanas-go/faas/commit/26115fbb89de4e1983e03b7d8728246e580ee5f9))
- Adding s3 and dynamodb lambda triggers support ([43de409](https://github.com/americanas-go/faas/commit/43de409c94b6a046d46b253da62c3198b3f53355))

### Refactor

- Simplifying event conversion ([be041c4](https://github.com/americanas-go/faas/commit/be041c4fa714c3d3772802b6b31d25152d918d62))

## [v1.0.0-beta.1](https://github.com/americanas-go/faas/compare/v1.0.0-alpha-8...v1.0.0-beta.1) - 2021-06-28

[15d9155](https://github.com/americanas-go/faas/commit/15d9155f725fe6f340767b48e35aa428c202c6aa)...[ff2e5e1](https://github.com/americanas-go/faas/commit/ff2e5e10baf293028d71994ae79e8847aacdcd6f)

### Bug Fixes

- Sqs queue url ([1e0ae9a](https://github.com/americanas-go/faas/commit/1e0ae9a1266fca286fdde17eb9a74ed5fb2f5c90))

### Documentation

- Add godocs (wip) ([403ffe4](https://github.com/americanas-go/faas/commit/403ffe4679ee902b9ac81ae381694decd2914915))
- Add godoc for cloudevents package ([2a37974](https://github.com/americanas-go/faas/commit/2a379747a357750ec9f52a262a1b3a446423e3a7))
- Adds godocs for cmd ([f964de4](https://github.com/americanas-go/faas/commit/f964de42ee26d1415992bebe191563a8015c7462))
- Adds godocs for fx ([233444f](https://github.com/americanas-go/faas/commit/233444f90ab1a00c7621b2ca26a7a737f0a7ae38))
- Add godoc for datastore package ([785d2e5](https://github.com/americanas-go/faas/commit/785d2e5b83fbc01542a9d4c3e0b1ed8904eb1b98))
- Adds godocs for lambda ([4bdf473](https://github.com/americanas-go/faas/commit/4bdf4732b7fde08ad3059c73af431926f989030e))
- Adds godocs for nats ([4758a5d](https://github.com/americanas-go/faas/commit/4758a5d49517d51f8544c5dca0d7df9eac7fff07))

### Testing

- Adds wrapper provider event ([73a02b8](https://github.com/americanas-go/faas/commit/73a02b8e86c95c796cb3e038af820d3c15a16d75))
- Adds json bytes util func ([4c1c8d4](https://github.com/americanas-go/faas/commit/4c1c8d4aa1dfad4c7412c19532730960011c39fc))
- Adds nats helper ([90857bb](https://github.com/americanas-go/faas/commit/90857bb134f381170d32c4e5a5468c5c1e3581fd))
- Adds lambda helper ([18710f6](https://github.com/americanas-go/faas/commit/18710f674db57152928bdb5315ad652cb87e9f15))
- Adds cloudevents handler ([f2417de](https://github.com/americanas-go/faas/commit/f2417deeb3d77311bd920c677e733d2d3d50c78f))
- Adds event module ([caf2bad](https://github.com/americanas-go/faas/commit/caf2bad1c546587d9a3f41c2e123df9da6a7be2a))

## [v1.0.0-alpha-8](https://github.com/americanas-go/faas/compare/v1.0.0-alpha-7...v1.0.0-alpha-8) - 2021-06-15

[6509cb3](https://github.com/americanas-go/faas/commit/6509cb31ec113272b1d39704356b06699e9379ad)...[15d9155](https://github.com/americanas-go/faas/commit/15d9155f725fe6f340767b48e35aa428c202c6aa)

### Documentation

- Adds to package ([d27d18f](https://github.com/americanas-go/faas/commit/d27d18fd6e1e9ea8d10921db70404d5941980bfe))
- Fixing README.md ([345afc8](https://github.com/americanas-go/faas/commit/345afc87b7ef299fc4e4ca54ebd18ed31825889e))

### Features

- Returns error in fx.Run ([e2877e5](https://github.com/americanas-go/faas/commit/e2877e549e87295e61b5cefbcbd9e698cf59e4e7))

### Testing

- Removes unnecessary connection ([2a65f43](https://github.com/americanas-go/faas/commit/2a65f4376e8e79a687b1363e7d323b11a9a0233c))

## [v1.0.0-alpha-7](https://github.com/americanas-go/faas/compare/v1.0.0-alpha-6...v1.0.0-alpha-7) - 2021-06-07

[a08fdcc](https://github.com/americanas-go/faas/commit/a08fdcc2ac5f08357f4c3c2b4108b92f9104cc25)...[6509cb3](https://github.com/americanas-go/faas/commit/6509cb31ec113272b1d39704356b06699e9379ad)

### Bug Fixes

- Di for aws kinesis, sqs, sns ([c6c7e3f](https://github.com/americanas-go/faas/commit/c6c7e3f27e5b9b02ea3e8570c346e1c780b4691e))

## [v1.0.0-alpha-6](https://github.com/americanas-go/faas/compare/v1.0.0-alpha-5...v1.0.0-alpha-6) - 2021-06-04

[55e93bf](https://github.com/americanas-go/faas/commit/55e93bf2928c45a13021dc4da64a30cd6baf2dab)...[a08fdcc](https://github.com/americanas-go/faas/commit/a08fdcc2ac5f08357f4c3c2b4108b92f9104cc25)

### Miscellaneous Tasks

- Upgrades ignite to v1.0.0-alpha-16 ([33495cb](https://github.com/americanas-go/faas/commit/33495cb725a491241c2c03702cd05c3a2b181948))

## [v1.0.0-alpha-5](https://github.com/americanas-go/faas/compare/v1.0.0-alpha-4...v1.0.0-alpha-5) - 2021-06-04

[f1d3f1e](https://github.com/americanas-go/faas/commit/f1d3f1ef0335fc013e5a8971870f9501878998d7)...[55e93bf](https://github.com/americanas-go/faas/commit/55e93bf2928c45a13021dc4da64a30cd6baf2dab)

### Features

- Adding fx modules to sqs, sns and kinesis events ([fde8930](https://github.com/americanas-go/faas/commit/fde8930cea242daadd7113ba62656633c0cbeb61))

## [v1.0.0-alpha-4](https://github.com/americanas-go/faas/compare/v1.0.0-alpha-3...v1.0.0-alpha-4) - 2021-05-26

[d81dc02](https://github.com/americanas-go/faas/commit/d81dc023b81f06b04d6a9f2d12eec4ac6f180a3a)...[f1d3f1e](https://github.com/americanas-go/faas/commit/f1d3f1ef0335fc013e5a8971870f9501878998d7)

### Bug Fixes

- Fixes examples ([0d0d75b](https://github.com/americanas-go/faas/commit/0d0d75bbf7e38dc3e195923099ec283f08a33fc0))

### Features

- Adds fx modules for events ([9d23397](https://github.com/americanas-go/faas/commit/9d2339778821628f56b6161a21408995d8f2edf1))

## [v1.0.0-alpha-2](https://github.com/americanas-go/faas/compare/v1.0.0-alpha-1...v1.0.0-alpha-2) - 2021-05-14

[938b107](https://github.com/americanas-go/faas/commit/938b107538e34731b1eabce39b644f1d5bffa2e5)...[55b50e5](https://github.com/americanas-go/faas/commit/55b50e5891d5e08dd30430c4b48c1766f4486838)

### Bug Fixes

- Fixes configs and example ([55b50e5](https://github.com/americanas-go/faas/commit/55b50e5891d5e08dd30430c4b48c1766f4486838))

<!-- generated by git-cliff -->
