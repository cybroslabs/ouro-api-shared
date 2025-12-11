# ouro-api-shared

Shared API definitions and generated code for the Ouro Platform gRPC services.

## Overview

This repository contains Protocol Buffer definitions and generated code for the Ouro Platform's gRPC API. It serves as a shared library for both backend services and client applications.

## Features

- Protocol Buffer definitions for all Ouro Platform services
- Generated Go code (gRPC, GraphQL)
- Generated TypeScript/JavaScript code (Connect-RPC)
- Comprehensive API documentation

## Project Structure

```
.
├── proto/              # Protocol Buffer definitions
│   ├── acquisition/    # Device data acquisition services
│   ├── common/        # Common types and models
│   ├── cronjobs/      # Scheduled job management
│   ├── crypto/        # Cryptography services
│   ├── messaging/     # Message queue services
│   ├── services/      # Core API services
│   └── system/        # System configuration
├── gen/               # Generated code output
│   └── markdown/      # Generated API documentation
├── graph/             # GraphQL schema and resolvers
└── src/               # Additional source files

```

## Requirements

- Go 1.25 or later
- Node.js (for TypeScript generation)
- Buf CLI (for Protocol Buffer compilation)

## Installation

### As a Go Module

```bash
go get github.com/cybroslabs/ouro-api-shared
```

### As an npm Package

```bash
npm install @bufbuild/protobuf
```

## Code Generation

This project uses [Buf](https://buf.build/) for managing Protocol Buffer compilation.

### Generate All Code

```bash
make
```

This will generate:
- Go gRPC service definitions
- Go GraphQL bindings
- TypeScript/JavaScript Connect-RPC clients
- Markdown documentation

### Individual Generation Commands

See the [Makefile](Makefile) for available targets.

## API Services

The platform includes the following service categories:

- **Acquisition** - Device data collection and management
- **Configuration** - System and device configuration
- **Devices** - Device lifecycle management
- **Firmware** - Firmware image management and updates
- **Messaging** - Event streaming and message queues
- **Cryptography** - Secret management and encryption
- **User Management** - Authentication and authorization
- **System** - Platform-wide system services

For detailed API documentation, see the [generated documentation](gen/markdown/index.md).

## Development

### Prerequisites

Install dependencies:

```bash
go mod download
npm install
```

### Making Changes

1. Update Protocol Buffer definitions in `proto/`
2. Regenerate code: `make`
3. Test changes in dependent projects
4. Commit both `.proto` files and generated code

## Documentation

- [API Reference](gen/markdown/index.md) - Complete API documentation
- [Buf Schema Registry](https://buf.build) - Protocol Buffer schema management

## License

Copyright 2024 dDash s.r.o. All rights reserved.

This code is proprietary and may not be redistributed, modified, or published.
