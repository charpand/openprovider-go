# OpenProvider Terraform Provider

Terraform provider for managing Openprovider domains.

## Requirements

- Terraform >= 1.3

## Usage

```hcl
terraform {
  required_providers {
    openprovider = {
      source  = "charpand/openprovider"
      version = ">= 0.1.0"
    }
  }
}

provider "openprovider" {
  username = var.openprovider_username
  password = var.openprovider_password
}
```

## Documentation

Registry docs are generated from templates and examples:

- Templates live in `templates/`.
- Examples live in `examples/`.
- Generated docs live in `docs/`.

To regenerate docs, install `tfplugindocs` and run:

```bash
./scripts/docs
```

## Development

```bash
./scripts/format
./scripts/lint
./scripts/test
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
