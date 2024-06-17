# OpenAPI Documentation Action

Action for generating documentation using the Web Component
[elements](https://github.com/stoplightio/elements) from
[stoplightio](https://github.com/stoplightio/).

## Usage

### Example workflow

```yaml
name: Documentation
on: 
    push:
        branches: ["master"]

jobs:
    doc:
        name: Documentation
        runs-on: ubuntu-latest
        steps:
            - name: Checkout
              uses: actions/checkout@v4

            - name: Build
              uses: alexandrelamberty/openapi-stoplight-doc-action@v1.0.0
              with:
                  title: My API Documentation
                  file: ./api-spec.yaml
                  directory: ./docs

            - name: Publish
              uses: peaceiris/actions-gh-pages@v3
              with:
                  github_token: ${{ secrets.GITHUB_TOKEN }}
                  publish_dir: ./docs
```

### Inputs

- **title** (optional): The title of the generated documentation. Defaults to
  `API Documentation`.
- **file** (optional): The path to your OpenAPI specification file in YAML
  format. Defaults to `api.yaml`.
- **directory** (optional): The directory where the documentation will be
  generated. Defaults to the `root` of your repository.
