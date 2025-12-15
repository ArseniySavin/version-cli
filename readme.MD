# version-cli

[RU](./docs/ru.MD)

**Description**
- **Project:** Command-line utility for managing build versions according to [SemVer](https://semver.org/).
- **Purpose:** Parses and increments versions, generates tags, and can push a variable value to GitLab via the API.

**Key Features**
- **SemVer:** Support for the Major.Minor.Patch format, `-build` metadata, and `+revision`.
- **Increments:** Patch increment (`--patch`) and release/minor increment (`--release`).
- **Build metadata:** Support for `--build`, `--revision`, and automatic revision generation with `--auto-revision`.
- **Tag output:** The `--tag` flag returns the tag representation of the version.
- **Push:** Sends the value to the GitLab Variable API via the `push` command.

**SemVer Specification**
- Official specification: https://semver.org/

**Installation**
- Build from source:

```
CGO_ENABLED=0 go build -o version-cli .
```

**Usage Examples**
- Show version from `--value`:

```
./version-cli --value "v1.2.3" version
> v1.2.3
```

- Get tag:

```
./version-cli --value "v1.2.3+20200101" version --tag
> v1.2.3.20200101
```

- Increment minor | major version:

```
./version-cli --value "v1.2.3" version --patch
> v1.2.4

./version-cli --value "v1.2.3" version --release
> v1.3.0
```

- Auto revision:

```
./version-cli --value "v1.2.3" version --release --build DEV --auto-revision
> v1.3.0-DEV+8968705834406
```

- Push value to the GitLab Variable API:

```
./version-cli push --url "https://gitlab.com/api/v4/projects/123/variables/CI_VERSION" --private-token $ACCESS_TOKEN
```

**Main Flags and Parameters**
- **`--value`**: Version value (required). Examples: `v1.2.3`, `1.2.3+12345`.
- **`--release`**: Increment minor version (reset patch).
- **`--patch`**: Increment patch version.
- **`--build`**: Build metadata name (e.g. `Dev`, `Alpha`).
- **`--revision`**: Explicitly set revision. If empty and `--auto-revision` is specified, the revision will be generated automatically.
- **`--auto-revision`**: Generates a revision based on the current time + crc32.
- **`--tag`**: Output tag format (dots instead of `-`/`+` in metadata).
- **`--url`**: (for `push`) Full GitLab Variable API URL.
- **`--private-token`** / **`--job-token`**: Authentication tokens for `push`.
- **`--debug`**: Enable verbose debug output.

```
./version-cli --value v1.0.2-Alpha.1 [--debug]
version [--release --patch --build=value --revision=value --auto-revision --tag] \
push --url "https://gitlab.kz/api/v4/projects/321/variables/COUNTER" --private-token LTxRf3GYQdC-ASfjzhzb [--tag]
```

**How `version-cli` Builds Version Strings**
- Format without metadata: `vMAJOR.MINOR.PATCH`.
- With `build` and/or `revision`:
  - `vMAJOR.MINOR.PATCH-build` (if only `build` is specified).
- With `--tag`, a similar string is produced, but metadata is separated by dots (e.g. `v1.2.3.Dev.12345`).

**CI Examples (GitLab)**
- To manage versions, create a `CI_VERSION` variable in the repository and pass it through the pipeline, using the `push` command.

- Pipeline example:

```
tag-release-job:
  stage: tag-release
  <<: *job-release
  script:
    - git remote set-url origin "https://oauth2:$ACCESS_TOKEN@$CI_SERVER_HOST/$CI_PROJECT_PATH.git"
    - git tag -a $TAG_VERSION -m "$CI_COMMIT_BRANCH | $CI_COMMIT_TAG | $CI_COMMIT_MESSAGE"
    - git push origin --tags
    - version-cli --value=$CI_VERSION version --patch push --url "$CI_API_V4_URL/projects/$CI_PROJECT_ID/variables/CI_VERSION" --private-token $ACCESS_TOKEN
```