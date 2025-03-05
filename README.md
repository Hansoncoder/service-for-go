
# Go Server Demo Overview

This project is a simple server demo implemented in Go, showcasing a typical project structure and layered architecture design for building backend services with Go. The project includes configuration files, database operations, API layer, logging, service layer, etc., making it a great starting point for learning Go backend development.

## Project Directory Structure

### Root Directory
- **go.mod**: The Go module file that manages project dependencies.
- **go.sum**: The Go checksum file ensuring the integrity and consistency of dependencies.
- **logs**: Directory for storing log files, such as `2025-03-05.log`, to record runtime logs of the application.
- **tmp**: Temporary folder containing build logs and intermediate files, such as `build-errors.log` and the compiled `main` executable.

### `cmd` Directory
- **main.go**: The entry point of the server. It initializes the necessary configurations, database connections, API routes, and starts the HTTP server.

### `config` Directory
- **config.yaml**: Configuration file defining application settings (e.g., database connection information, server configurations, etc.). The file is loaded to dynamically configure the runtime environment of the application.

### `internal` Directory
This directory contains the core business logic of the server, divided into several sub-modules, each with its own responsibility.

- **api**: Defines HTTP routes and request handlers, responsible for mapping user requests to appropriate service logic.
- **config**: Handles loading and managing application configuration to ensure correct configuration information is passed to other modules.
- **database**: Manages database connections and operations, including connection pooling and transaction handling.
- **models**: Defines the data models used in the application, typically corresponding to database table structures.
- **repository**: Handles data persistence operations, abstracting the database interaction logic, and provides data access interfaces for the service layer.
- **service**: Contains the core business logic, such as user management, order processing, etc.
- **utils**: A collection of utility functions, including features like logging, error handling, encryption/decryption, etc.

### `pkg` Directory
- **errors**: Custom error types and related utilities for unified error management throughout the project.

### `gotest.sql` File
- A simple database initialization script that defines the database table structures. Running this SQL file initializes the database.

## How to Run

1. Clone the project:
   ```bash
   git clone https://your-repo-url
   cd your-repo-directory
   ```

2. Configure the database:
   Edit the `config/config.yaml` file with your database connection details.

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Start the server:
   ```bash
   go run cmd/main.go
   ```

The server will start and listen on the port specified in the configuration file.

## Summary
This Go Server Demo demonstrates a typical backend development structure with Go. With its clear layered architecture, it helps developers understand how to write scalable backend applications in Go.
