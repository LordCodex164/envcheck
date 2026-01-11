#envcheck

Your Go CLI Tool thaat validates environment variables aganist a schema

## Features

- ✅ Validate required environment variables
- ✅ Type checking (string, int, bool, url)
- ✅ Default values for optional variables
- ✅ Strict mode to catch unexpected variables
- ✅ Helpful error messages with examples
- ✅ Config file support (.envcheck.yaml)
- ✅ Perfect for CI/CD pipelines

## Installation

### Quick Install (macOS/Linux)
```bash
curl -sSL https://raw.githubusercontent.com/yourusername/Lordcodex164/main/install.sh | bash
```

### Manual Installation

Download the latest binary for your platform from [Releases](https://github.com/Lordcodex164/envcheck/releases):

- **macOS (Intel)**: `envcheck-darwin-amd64`
- **macOS (Apple Silicon)**: `envcheck-darwin-arm64`
- **Linux**: `envcheck-linux-amd64`
- **Windows**: `envcheck-windows-amd64.exe`
```bash
# Example for macOS
curl -LO https://github.com/Lordcodex164/envcheck/releases/latest/download/envcheck-darwin-arm64
chmod +x envcheck-darwin-arm64
sudo mv envcheck-darwin-arm64 /usr/local/bin/envcheck
```

### Build from Source
```bash
git clone https://github.com/Lordcodex164/envcheck.git
cd envcheck
go build -o envcheck .
sudo mv envcheck /usr/local/bin/
```

## Quick Start

### 1. Create a schema
```bash
envcheck create
```

This generates `schema.yaml`:
```yaml
variables:
  - name: APP_DATABASE_URL
    required: true
    type: string
    description: PostgreSQL connection string
  - name: APP_PORT
    required: false
    type: int
    default: "8080"
    description: Server port
  - name: APP_DEBUG
    required: false
    type: bool
    default: "false"
    description: Enable debug mode
```

### 2. Edit the schema to match your app
```bash
vim schema.yaml
```

### 3. Validate your environment
```bash
# Set your variables
export APP_DATABASE_URL="postgres://localhost/mydb"
export APP_PORT="3000"

# Validate
envcheck validate
```

Output:
```
✓ APP_DATABASE_URL
✓ APP_PORT
✓ APP_DEBUG

✓ All environment variables validated successfully
```

## Usage

### Basic Commands
```bash
# Create a schema template
envcheck create

# Validate current environment
envcheck validate

# Validate with custom schema
envcheck validate --schema prod-schema.yaml

# Strict mode: fail on unexpected APP_* variables
envcheck validate --strict

# Verbose output
envcheck validate --verbose

# Create config file
envcheck init
```

### Config File Support

Create `.envcheck.yaml` in your project root:
```yaml
schema: schema.yaml
strict: true
verbose: false
```

Now just run:
```bash
envcheck validate
```

### CI/CD Integration

#### GitHub Actions
```yaml
name: Validate Environment
on: [push]

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Install envcheck
        run: |
          curl -LO https://github.com/Lordcodex164/envcheck/releases/latest/download/envcheck-linux-amd64
          chmod +x envcheck-linux-amd64
          sudo mv envcheck-linux-amd64 /usr/local/bin/envcheck
      
      - name: Validate environment
        env:
          APP_DATABASE_URL: ${{ secrets.DATABASE_URL }}
          APP_PORT: "8080"
        run: envcheck validate --strict
```

## Schema Reference

### Variable Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | ✓ | Environment variable name |
| `required` | bool | ✓ | Whether the variable must be set |
| `type` | string | ✓ | Variable type: `string`, `int`, `bool`, `url` |
| `default` | string | - | Default value if not set |
| `description` | string | - | Human-readable description |

### Example Schema
```yaml
variables:
  - name: DATABASE_URL
    required: true
    type: url
    description: PostgreSQL connection string
    
  - name: REDIS_URL
    required: true
    type: url
    description: Redis connection string
    
  - name: PORT
    required: false
    type: int
    default: "8080"
    description: HTTP server port
    
  - name: LOG_LEVEL
    required: false
    type: string
    default: "info"
    description: Logging level (debug, info, warn, error)
    
  - name: FEATURE_FLAG_BETA
    required: false
    type: bool
    default: "false"
    description: Enable beta features
```

## Error Messages

envcheck provides helpful error messages:
```
❌ Invalid variable: APP_PORT
   Expected type: int
   Got: "not-a-number"
   
   This variable must be a valid integer.
   
   Example: export APP_PORT=8080
```

## Exit Codes

- `0` - All validations passed
- `1` - Validation failed

Perfect for scripts:
```bash
if envcheck validate; then
  echo "Environment valid, starting app..."
  ./myapp
else
  echo "Environment validation failed"
  exit 1
fi
```

## Development

### Prerequisites

- Go 1.21 or higher

### Build
```bash
# Simple build
go build -o envcheck .

# Build for all platforms
./build-all.sh
```

### Run Tests
```bash
go test ./... -v
```

## Contributing

Contributions welcome! Please open an issue or PR.

## License

MIT License - see [LICENSE](LICENSE) file

## Author

([@Lordcodex_](https://x.com/Lordcodex_))

---

**Built with ❤️ using Go and Cobra**