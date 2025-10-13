# Homebox Go SDK

A Go SDK for the [Homebox](https://github.com/sysadminsmedia/homebox) inventory management API, automatically generated from the OpenAPI specification.

## Installation

```bash
go get github.com/jeffresc/go-homebox
```

## Usage
For usage examples, see the [example/](example/) directory.

## Development

### Prerequisites

- Go 1.22 or higher
- Make

### Regenerating the SDK

The SDK is automatically generated from the OpenAPI specification. To regenerate it:

```bash
make generate
```

This will:
1. Download the latest OpenAPI spec from the URL specified in `.openapi-url`
2. Generate the Go client code using `oapi-codegen`
3. Place the generated code in the `client/` directory

### Updating the OpenAPI Spec URL

Edit the `.openapi-url` file to point to a different version or location of the OpenAPI specification. The format is:

```
https://raw.githubusercontent.com/sysadminsmedia/homebox/refs/heads/main/docs/en/api/openapi-3.0.json
```

To use a specific version/tag instead of `main`, change `main` to the desired tag:

```
https://raw.githubusercontent.com/sysadminsmedia/homebox/refs/heads/v0.10.0/docs/en/api/openapi-3.0.json
```

After updating the URL, run `make generate` to regenerate the SDK.

### Makefile Commands

- `make help` - Show available commands
- `make generate` - Generate client from OpenAPI spec
- `make clean` - Clean generated files
- `make test` - Run tests
- `make deps` - Install/update dependencies

## Contributing

Contributions are welcome! Please note that the code in the `client/` directory is auto-generated and should not be edited directly. Instead, submit issues or PRs to improve the generation process or documentation.

## License

MIT
