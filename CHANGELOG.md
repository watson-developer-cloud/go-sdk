# [2.0.0](https://github.com/watson-developer-cloud/go-sdk/compare/v1.7.0...v2.0.0) (2020-12-10)


* Release/2.0.0 (#90) ([4f091a2](https://github.com/watson-developer-cloud/go-sdk/commit/4f091a21576297153b17817971bc20a186d0efef)), closes [#90](https://github.com/watson-developer-cloud/go-sdk/issues/90)


### BREAKING CHANGES

* version is now a string pointer

* feat(assistantv2): update models with latest definitions

* feat(comparecomplyv1): update models with latest definitions
* version is now a string pointer

* feat(discoveryv1): update models with latest definitions
* version is now a string pointer

* feat(discoveryv2): update models with latest definitions
* version is now a string pointer

* feat(languagetranslatorv3): update models with latest definitions

* feat(naturallanguageclassifierv1): update models with latest definitions
* version is now a string pointer

* feat(naturallanguageunderstandingv1): update models with latest definitions

* feat(personalityinsightsv3): update models with latest definitions
* version is now a string pointer

* feat(speechtotextv1): update models with latest definitions

* feat(texttospeechv1): update models with latest definitions
* VoiceModel is now renamed to CustomModel for all related models and operations

* feat(toneanalyzerv3): update models with latest definitions
* version is now a string pointer

* feat(visualrecognitionv3): update models with latest definitions
* version is now a string pointer

* feat(visualrecognitionv4): update models with latest definitions
* version is now a string pointer

* docs(examples): update examples

* build(update modules): update go modules
* the go sdk now requires go-sdk-core/v4 as a dependency

* chore(personalityinsightsv3): log deprecation warning

* fix(hand edit): make SpeechRecognitionAlternative utilize generic interface{}

* chore(copyright dates): fix copyright dates

* fix(hand edit): remove trainingStatus from CreateCollection and UpdateCollection

these are output parameters and should not be generated in the options model

* chore(deprecation): add deprecation message for visual recognition

* chore(deps): update strfmt

* test(integration tests): update tests for improvedNameFormatting constants

* fix(hand edit): unmarshal key_as_string for timestamp keys

* style(format): run gofmt

# [1.7.0](https://github.com/watson-developer-cloud/go-sdk/compare/v1.6.0...v1.7.0) (2020-08-26)


### Features

* **assistantv2:** add ListLogs and DeleteUserData operations ([ae4ecaf](https://github.com/watson-developer-cloud/go-sdk/commit/ae4ecafcb597fd646e4355a1496f685285276d7e))
* **default url:** update default url for services ([a63b102](https://github.com/watson-developer-cloud/go-sdk/commit/a63b102243708da5d024e6f4a30106457a73a1cb))
* **discoveryv2:** add Collection and Enrichment operations ([b6cea64](https://github.com/watson-developer-cloud/go-sdk/commit/b6cea64e29d4e1612ff5ceb1ebdc7aa37e19d53c))
* **languagetranslatorv3:** add ListLanguages operation ([fec7d7d](https://github.com/watson-developer-cloud/go-sdk/commit/fec7d7dcd73a95d24c81520e4b99d4975abd7664))
* **speechtotextv1:** update language constants for custom model operations ([27c2c66](https://github.com/watson-developer-cloud/go-sdk/commit/27c2c66f73070c64bb649fe12cdc2afd098016ff))
* **texttospeechv1:** update voice constants ([5eceaa0](https://github.com/watson-developer-cloud/go-sdk/commit/5eceaa0f9cdd62db9cb057791198047fb1f0147b))
* **visualrecognitionv4:** update service url and tests ([6d411e7](https://github.com/watson-developer-cloud/go-sdk/commit/6d411e7e83bdd4af70cc82094b3f670dc89805be))


### Reverts

* **visualrecognitionv4:** readd count to the metadata object ([38eed59](https://github.com/watson-developer-cloud/go-sdk/commit/38eed599b88376e1bb86be7b38e7a90eaaff7101))

# [1.6.0](https://github.com/watson-developer-cloud/go-sdk/compare/v1.5.0...v1.6.0) (2020-06-04)


### Features

* **assistantv1:** add generated code ([c66028b](https://github.com/watson-developer-cloud/go-sdk/commit/c66028b586368d28e0d83e21bf2d02f9118b49ed))
* **assistantv2:** add MessageStateless operation ([0ec503a](https://github.com/watson-developer-cloud/go-sdk/commit/0ec503ad1b86368f0b7b13879ff5b4412b846b05))
* **naturallanguageunderstandingv1:** generate new model definition ([85b7cdb](https://github.com/watson-developer-cloud/go-sdk/commit/85b7cdb37ce9ebf0e703360d5a74f221c3c9d1b6))
* **visualrecognitionv4:** generated code ([685fdfb](https://github.com/watson-developer-cloud/go-sdk/commit/685fdfb8f4a2aad1aaa981200c34bdaef8f87dba))

# [1.5.0](https://github.com/watson-developer-cloud/go-sdk/compare/v1.4.1...v1.5.0) (2020-04-24)


### Features

* **assistantv1:** re-generate service ([d17439a](https://github.com/watson-developer-cloud/go-sdk/commit/d17439a132ccaa18f2f7bb7cb447b36aa5cc6eab))
* **assistantv2:** re-generate service ([660c05a](https://github.com/watson-developer-cloud/go-sdk/commit/660c05a734cb11282c5c2a730a82f4898e2cbbec))
* **comparecomplyv1:** re-generate service ([dbb8cf6](https://github.com/watson-developer-cloud/go-sdk/commit/dbb8cf6ca0bec65399530f3c7dafa8217096ecb2))
* **discoveryv1:** re-generate service ([1c7cfdf](https://github.com/watson-developer-cloud/go-sdk/commit/1c7cfdfc8de3b6a14518ccacc61a98f19bf31b02))
* **discoveryv2:** re-generate service ([b912d2b](https://github.com/watson-developer-cloud/go-sdk/commit/b912d2bc0009c3972cf140d603a45cb9ec7893a3))
* **languagetranslatorv3:** re-generate service ([c36f8da](https://github.com/watson-developer-cloud/go-sdk/commit/c36f8daf017e57bf359e534f007bbe1e6541ceef))
* **naturallanguageunderstandingv1:** re-generate service ([24804cb](https://github.com/watson-developer-cloud/go-sdk/commit/24804cbb0d4152399823052d1df94b0cf4cbe343))
* **speechtotextv1:** re-generate service ([4475b24](https://github.com/watson-developer-cloud/go-sdk/commit/4475b249ea7dbc542e17cde90d366e52063d5083))
* **texttospeechv1:** re-generate service ([fe1231c](https://github.com/watson-developer-cloud/go-sdk/commit/fe1231cf4b44019c7f9b23917047e3f79378e169))
* **visualrecognitionv4:** re-generate service ([6e42648](https://github.com/watson-developer-cloud/go-sdk/commit/6e42648954c00c0442a757611a0184c8302c4869))

## [1.4.1](https://github.com/watson-developer-cloud/go-sdk/compare/v1.4.0...v1.4.1) (2020-02-18)


### Bug Fixes

* use go-sdk-core v2.1.1 ([d250f84](https://github.com/watson-developer-cloud/go-sdk/commit/d250f8480c2de3084f0c950c8cf7cc0fe25aba14))

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
