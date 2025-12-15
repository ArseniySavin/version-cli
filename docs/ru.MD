# version-cli

**Описание**
- **Проект:**: Утилита командной строки для управления версией сборки согласно [SemVer](https://semver.org/).
- **Назначение:**: 
    - Помогает в DevOps процессе разработки.
    - Парсит и инкрементирует версии, формирует теги и может отправлять значение переменной в GitLab через API.

**Ключевые возможности**
- **SemVer:**: Поддержка формата Major.Minor.Patch, метаданных `-build` и `+revision`.
- **Инкременты:**: Увеличение патча (`--patch`) и релиза/минора (`--release`).
- **Метаданные сборки:**: Поддержка `--build`, `--revision` и генерации ревизии `--auto-revision`.
- **Tag output:**: Флаг `--tag` возвращает теговую форму версии.
- **Push:**: Отправка значения в GitLab Variable API через команду `push`.

**Ссылка на спецификацию SemVer**
- Официальная спецификация: https://semver.org/

**Установка**
- Сборка из исходников:

```
CGO_ENABLED=0 go build -o version-cli .
```

**Пример использования**
- Показать версию из `--value`:

```
./version-cli --value "v1.2.3" version
> v1.2.3
```
- Получить tag:

```
./version-cli --value "v1.2.3+20200101" version --tag
> v1.2.3.20200101

```
- Увеличить minor | major инкремента
```
./version-cli --value "v1.2.3" version --patch
> v1.2.4

./version-cli --value "v1.2.3" version --release
> v1.3.0
```
- Авто ревизия
```
./version-cli --value "v1.2.3" version --release --build DEV --auto-revision
> v1.3.0-DEV+8968705834406
```
- Отправить значение в GitLab Variable API (push):

```
./version-cli push --url "https://gitlab.com/api/v4/projects/123/variables/CI_VERSION" --private-token $ACCESS_TOKEN
```

**Основные флаги и параметры**
- **`--value`**: Значение версии (обязательно). Пример: `v1.2.3`, `1.2.3+12345`.
- **`--release`**: Увеличить minor (сбросить patch).
- **`--patch`**: Увеличить patch.
- **`--build`**: Имя build-метаданных (например `Dev`, `Alpha`).
- **`--revision`**: Явно указать ревизию. Если пусто и указан `--auto-revision`, ревизия сгенерируется автоматически.
- **`--auto-revision`**: Генерирует ревизию на основе текущего времени + crc32.
- **`--tag`**: Вывести tag-формат (точки вместо `-`/`+` в метаданных).
- **`--url`**: (для `push`) Полный URL GitLab Variable API.
- **`--private-token`** / **`--job-token`**: Токены для аутентификации при `push`.
- **`--debug`**: Включить подробный вывод для отладки.

```
./version-cli --value v1.0.2-Alpha.1 [--debug]
version [--release --patch --build=value --revision=value --auto-revision --tag]\
push --url "https://gitlab.kz/api/v4/projects/321/variables/COUNTER" --private-token LTxRf3GYQdC-ASfjzhzb [--tag]
```


**Как `version-cli` формирует строки**
- Формат без метаданных: `vMAJOR.MINOR.PATCH`.
- С `build` и/или `revision`:
	- `vMAJOR.MINOR.PATCH-build` (если только `build`);
- При `--tag` выводится похожая строка, но с точками для разделения метаданных (например `v1.2.3.Dev.12345`).

**Примеры CI (GitLab)**
- Для того что бы можно было вести версию, на репозитории заводим переменную `CI_VERSION` и в pipeline можно передавать `CI_VERSION` и использовать команду `push`.

- Пример из pipeline:

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