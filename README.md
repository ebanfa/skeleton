# Skeleton Project

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [System API](#system-api)
  - [Key Concepts](#key-concepts)
  - [Usage Examples](#usage-examples)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Development](#development)
  - [Building](#building)
  - [Running](#running)
  - [Testing](#testing)
  - [Generating Mocks](#generating-mocks)
- [Contributing](#contributing)
- [License](#license)

## Overview

Skeleton is a library for building applications with components, services, operations, and system management functionalities. It provides a foundation for creating scalable and extensible software systems.

## Features

- Modular component-based architecture
- Service management (start, stop, restart)
- Operation execution framework
- Plugin system
- Integrated event bus
- Database abstraction layer
- Logging system

## System API

The System API is the core of the Skeleton project, providing a framework for building applications. It can be used to define components, manage services, execute operations, and control the system's lifecycle.

### Key Concepts

#### Components
Components are the building blocks of the application, encapsulating specific functionalities. Each component has a unique identifier (`ID`), name (`Name`), type (`Type`), and description (`Description`).

#### Services
Services are components that can be started and stopped. They implement the `StartableInterface` and provide functionality to the system during runtime.

#### Operations
Operations are components that represent units of work that can be executed within the system. They implement the `OperationInterface` and can perform various tasks based on input parameters.

#### System
The `SystemInterface` represents the core execution environment for components in the application, providing functionalities such as system initialization, configuration management, component registration, service control, and operation execution.

### Usage Examples

#### Component Creation and Registration
```go
// Create an instance of the ExampleComponentFactory
factory := ExampleComponentFactory{}

// Create a new ExampleComponent using the factory
component, err := factory.CreateComponent()
if err != nil {
    logger.log(LevelError, "Error creating component:", err)
}

// Initialize the system
sys := NewSystem()

// Register a component factory
err := sys.ComponentRegistry().RegisterComponent("example", ExampleComponent{})
if err != nil {
    logger.log(LevelError, err)
}
```

#### Service Management
```go
// Start a service
err := sys.StartService("exampleService", ctx)
if err != nil {
    logger.log(LevelError, err)
}

// Stop a service
err := sys.StopService("exampleService", ctx)
if err != nil {
    logger.log(LevelError, err)
}

// Restart a service
err := sys.RestartService("exampleService", ctx)
if err != nil {
    logger.log(LevelError, err)
}
```

#### Operation Execution
```go
// Execute an operation
output, err := sys.ExecuteOperation(ctx, "exampleOperation", &OperationInput{Data: inputData})
if err != nil {
    logger.log(LevelError, err)
}
```

## Getting Started

### Prerequisites

- Go (version X.X or higher)
- Make

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/skeleton.git
   cd skeleton
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

## Usage

For more detailed usage instructions and API documentation, please refer to the [README.md](docs/README.md) file.

## Project Structure

```
.
├── common/         # Common utilities and interfaces
├── component/      # Component-related code
├── db/             # Database interfaces and implementations
├── mocks/          # Auto-generated mocks for testing
├── plugin/         # Plugin system implementation
├── store/          # Data store implementations
├── system/         # Core system functionality
├── types/          # Type definitions
├── Makefile        # Build and development commands
└── README.md       # This file
```

## Development

### Building

To build the project:

```
make build
```

This will compile the project and create an executable named `skeleton` in the `bin/` directory.

### Running

To run the project:

```
make run
```

### Testing

To run tests:

```
make go-test
```

To run tests with coverage:

```
make go-test-with-cover
```

This will generate a coverage report and open it in your default web browser.

### Generating Mocks

To generate mock implementations for interfaces:

```
make generate-mocks
```

## Contributing

Contributions to the Skeleton project are welcome! If you have any suggestions, improvements, or bug fixes, please feel free to open an issue or submit a pull request on the GitHub repository.

## License

This project is licensed under the [Apache License](LICENSE).