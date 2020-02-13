# [1.4.0](https://github.com/watson-developer-cloud/go-sdk/compare/v1.3.2...v1.4.0) (2020-02-13)


### Features

* **assistantv1:** New param `IncludeAudit` following options ([23c7652](https://github.com/watson-developer-cloud/go-sdk/commit/23c765204e81cdfcf3a36073b6615aa0e8c44244))
* **assistantv1:** New params `Interpretation` and `Role` in `RuntimeEntity` ([9d375d7](https://github.com/watson-developer-cloud/go-sdk/commit/9d375d7550b555242b17fbaf945220245a628716))
* **assistantv2:** New params `Interpretation`, `Alternatives` and `Role` in `RuntimeEntity` ([e82ca09](https://github.com/watson-developer-cloud/go-sdk/commit/e82ca093e8285801b2066fe7e89459a7cd5a5147))
* **assistantv2:** New params `Locale` and `ReferenceTime` in `MessageContextGlobalSystem` ([bcd5e9a](https://github.com/watson-developer-cloud/go-sdk/commit/bcd5e9a65c6b402aef5f2b00afd7e2f5403c3eff))
* **vr4:** New object operations ([3f87c56](https://github.com/watson-developer-cloud/go-sdk/commit/3f87c560e1f8b04f95d1bcae9769789ed6c0e45a))

## [1.3.2](https://github.com/watson-developer-cloud/go-sdk/compare/v1.3.1...v1.3.2) (2020-01-24)


### Bug Fixes

* **discoveryv1:** Add custom UnmarshalJSON for AggregationResult ([4d41191](https://github.com/watson-developer-cloud/go-sdk/commit/4d411912d27fb4568be692349829a1170c133271))

## [1.3.1](https://github.com/watson-developer-cloud/go-sdk/compare/v1.3.0...v1.3.1) (2020-01-17)


### Bug Fixes

* **nlu:** Add Deprecated param `Model` back in `CategoriesOptions` ([6022eec](https://github.com/watson-developer-cloud/go-sdk/commit/6022eec2da270423ffd81bde9cbbf1d54f2f5ae9))

# [1.3.0](https://github.com/watson-developer-cloud/go-sdk/compare/v1.2.0...v1.3.0) (2020-01-17)


### Features

* **core:** Update core to v2.1.0 ([7609848](https://github.com/watson-developer-cloud/go-sdk/commit/7609848452cda4ad0d6e4da79f112d152580bd65))
* **speech to text:** New params `EndOfPhraseSilenceTime` and `SplitTranscriptAtPhraseEnd ` in `RecognizeOptions` and `CreateJobOptions` ([af0a3c3](https://github.com/watson-developer-cloud/go-sdk/commit/af0a3c3cdda026f7ffb215e1ca17701f8c76dd1c))

# [1.2.0](https://github.com/watson-developer-cloud/go-sdk/compare/v1.1.1...v1.2.0) (2019-11-27)


### Features

* **assistantv1:** New param `NewDisambiguationOptOut` in `UpdateDialogNodeOptions ` ([1e55910](https://github.com/watson-developer-cloud/go-sdk/commit/1e559109b298c91ff2c44074ba3489b7349c7ba3))
* **assistantv1:** New param `OffTopic` in `WorkspaceSystemSettings` ([7acadea](https://github.com/watson-developer-cloud/go-sdk/commit/7acadea9c21d58c2cda0db6bb7d533ca927c2487))
* **assistantv1:** New param `Webhooks` in `CreateWorkspaceOptions` and `UpdateWorkspaceOptions` ([e29359d](https://github.com/watson-developer-cloud/go-sdk/commit/e29359d6cc96c4ea5ba47a649b8b4ebcefc301d8))
* **assistantv1:** New properties `randomize`, `max_suggestions` and `suggestion_text_policy` in  `WorkspaceSystemSettingsDisambiguation` ([d141ee3](https://github.com/watson-developer-cloud/go-sdk/commit/d141ee3daca3edf4d74b95a7435bcb67ce5f108a))
* **asssistantv1:** New param `DisambiguationOptOut` in `CreateDialogNodeOptions` ([80b940d](https://github.com/watson-developer-cloud/go-sdk/commit/80b940d4a3bb38ed9cb254f7f4df94e497c84556))
* **discoveryv1:**  Property `title` not present in `QueryResult` and `QueryNoticesResult` ([a1ab2b3](https://github.com/watson-developer-cloud/go-sdk/commit/a1ab2b3288105de2a98aa8ce6417491122fb63fe))
* **discoveryv2:** New service discovery v2 for CP4D ([9139c6d](https://github.com/watson-developer-cloud/go-sdk/commit/9139c6dea422d69542e0432719c8fc74f721701d))
* **visual recognition v4:** Add thumbnail support ([252562f](https://github.com/watson-developer-cloud/go-sdk/commit/252562fdfb1290ebf3fa30ce72b5f3370142debe))
* **visual recognition v4:** New meethod `GetTrainingUsage()` ([c65c43d](https://github.com/watson-developer-cloud/go-sdk/commit/c65c43da3fb447242e93300af09fd99faeba15df))
* **visual recognition v4:** Properties `width` and `height` in ImageDimensions are optional ([224ce24](https://github.com/watson-developer-cloud/go-sdk/commit/224ce245356e8281473314a42f000c59946b79a4))
* **visual recognition v4:** Updates to `ImageDetails` and `Image` model ([221adc8](https://github.com/watson-developer-cloud/go-sdk/commit/221adc8a3f55e4fd14619aa2698a94f9f0697912))

## [1.1.1](https://github.com/watson-developer-cloud/go-sdk/compare/v1.1.0...v1.1.1) (2019-11-25)


### Bug Fixes

* **go module:** Regenerate go module with correct path ([8b3be06](https://github.com/watson-developer-cloud/go-sdk/commit/8b3be061d44cfc4e9c7dc3dd7a9a89f218e930ed))

# [1.1.0](https://github.com/watson-developer-cloud/go-sdk/compare/v1.0.0...v1.1.0) (2019-11-18)


### Bug Fixes

* **go module:** Move from dep to go module ([5c2d42a](https://github.com/watson-developer-cloud/go-sdk/commit/5c2d42ae3e70e5411e19d9a0c1ea26fccc7ef17c))
* **semantic release:** Update files for semantic release ([405833d](https://github.com/watson-developer-cloud/go-sdk/commit/405833d514a9bf602e49964ca90797c6b6210b69))


### Features

* **core:** Update core version ([e633856](https://github.com/watson-developer-cloud/go-sdk/commit/e633856bf5de23a153e4ee22169b1fd7bfd27461))
