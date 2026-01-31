# bit2bin
[![Go Version](https://img.shields.io/github/go-mod/go-version/hairenjun/bit2bin)](https://github.com/hairenjun/bit2bin) [![Go Reference](https://pkg.go.dev/badge/hairenjun/bit2bin.svg)](https://pkg.go.dev/hairenjun/bit2bin) [![License](https://img.shields.io/github/license/hairenjun/bit2bin)](LICENSE)

A high-performance command-line utility written in Go that converts text files containing bit sequences (0s and 1s) into actual binary files.

## Features

- **Cross-Platform:** Supports Linux and Windows on both AMD64 and ARM64 (AArch64).
- **Statically Linked:** No external dependencies or runtime libraries required.
- **Automatic Padding:** If the input bit count is not a multiple of 8, the tool automatically pads the final byte with trailing zeros.
- **Efficient:** Uses buffered I/O to handle large bitstring files.

## Quick Start

```bash
# Create a simple input file with bit sequence
echo -n "0101010100110010" > input.txt

# Convert to binary
./bit2bin input.txt output.bin

# Output:
# Conversion complete.
# Successfully wrote 2 bytes to output.bin
```

## Installation

### Pre-compiled Binaries

Download the appropriate binary for your platform from the [Releases](https://github.com/hairenjun/bit2bin/releases) section:

- **Linux AMD64:** `bit2bin-linux-amd64`
- **Linux ARM64:** `bit2bin-linux-arm64`
- **Windows AMD64:** `bit2bin-windows-amd64.exe`
- **Windows ARM64:** `bit2bin-windows-arm64.exe`

After downloading, make the binary executable (Linux/macOS):

```bash
chmod +x bit2bin-linux-amd64
```

### Build from Source with Makefile

You need [Go](https://go.dev/) installed. Use the provided `Makefile` to compile for all platforms:

```bash
# Build for all platforms (Linux/Windows AMD64/ARM64)
make all

# Build for specific platform
make linux-amd64
make linux-arm64
make windows-amd64
make windows-arm64

# Clean build directory
make clean
```

Binaries will be placed in the `build/` directory.

### Build from Source with go install

```bash
# Install directly to your GOPATH/bin
go install github.com/hairenjun/bit2bin@latest

# Build for current platform
go build -o bit2bin main.go
```

## Usage

### Basic Syntax

```bash
bit2bin <input.txt> <output.bin>
```

- `input.txt`: Text file containing a sequence of 0s and 1s
- `output.bin`: Binary file to write the converted bytes

### Command-Line Options

The tool accepts two positional arguments:
1. Input file path (required)
2. Output file path (required)

No additional flags or options are available.

## Examples

### Basic Conversion

Convert a 16-bit sequence to a 2-byte binary file:

```bash
# Create input file with 16 bits (2 bytes)
echo -n "0101010100110010" > input.txt

# Convert to binary
./build/bit2bin-linux-amd64 input.txt output.bin

# Output:
# Conversion complete.
# Successfully wrote 2 bytes to output.bin
```

### Automatic Padding

When the input bit count is not a multiple of 8, the tool automatically pads the final byte with trailing zeros:

```bash
# Create input file with 5 bits (not a complete byte)
echo -n "10110" > input.txt

# Convert to binary
./build/bit2bin-linux-amd64 input.txt output.bin

# Output:
# Conversion complete.
# Note: Padded last byte with 3 zeros to complete the byte.
# Successfully wrote 1 byte to output.bin
```

In this example, the 5 bits `10110` are padded to form the complete byte `10110000`.

### Large File Processing

The tool efficiently handles large bitstring files using buffered I/O:

```bash
# Generate a large bit sequence (1 million bits)
python3 -c "print('01' * 500000)" > large_input.txt

# Convert to binary
./build/bit2bin-linux-amd64 large_input.txt large_output.bin

# Output:
# Conversion complete.
# Successfully wrote 125000 bytes to large_output.bin
```

## Input Format

The input file should contain only the characters `0` and `1`. Any other characters (including whitespace, newlines, or other symbols) are ignored during processing.

**Valid input examples:**
- `01010101` (8 bits)
- `1100110011110000` (16 bits)
- `10110` (5 bits, will be padded to 8)

**Invalid characters are silently skipped:**
- `0101 0101` (space is ignored)
- `0101\n0101` (newline is ignored)

## Output Format

The output file contains raw binary data. Each group of 8 bits from the input produces exactly 1 byte in the output. If the final group has fewer than 8 bits, it is padded with zeros to complete the byte.

## Contributing

Contributions are welcome! To contribute to bit2bin:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests to ensure everything works
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Development Setup

```bash
# Clone the repository
git clone https://github.com/hairenjun/bit2bin.git
cd bit2bin

# Build the project
make all

# Run the binary
./build/bit2bin-linux-amd64 input.txt output.bin
```

### Testing

To test your changes:

```bash
# Build for your platform
go build -o bit2bin main.go

# Test with sample input
echo -n "0101010100110010" > test_input.txt
./bit2bin test_input.txt test_output.bin

# Verify output
xxd test_output.bin
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Copyright (c) 2025 bit2bin contributors
