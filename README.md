# chrono-server



[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com) 
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Chrono is a simple, in-memory, ephemeral key-value store written in Go. 
It's designed for simplicity, making it a perfect solution for managing temporary state, sessions, or short-lived data.

---

## Features

* **High Performance:** Being entirely in-memory, all operations are extremely fast.
* **Ephemeral by Design:** Keys are set with a Time-To-Live (TTL) and are automatically deleted upon expiration.
* **Simple Text Protocol:** Interact with the server using a simple, human-readable command protocol.

---

## Getting Started

Follow these instructions to get the Chrono server running on your local machine.

### Prerequisites

You need to have **Go** (version 1.20 or later) installed on your system.

### Building from Source

1.  Clone the repository:
    ```bash
    git clone https://github.com/alexs60/chrono-server.git

    cd chrono-server
    ```

2.  Build the executable:
    ```bash
    go build -o chrono-server ./cmd/chrono-server
    ```
    This will create a single binary named `chrono-server` in the root

### Running the Server

Simply execute the compiled binary:

```
./chrono-server
```

You should see a log message indicating the server has started: Chrono server started on port 8181


### 4. Usage

Usage
You can connect to the server using a simple TCP client like netcat or telnet.

Connecting
Open a new terminal window and run:


```
nc localhost 8181
```
#### Commands API
Once connected, you can issue the following commands:

```
SET key "value" ttl
```
Sets a key to hold a string value, with an expiration time in seconds.

key: The key to store the data under.

value: The string value to store.

ttl: The time-to-live in seconds. A TTL of 0 means the key will never expire.

Example:

> SET user:1:session "some-session-token" 3600

> OK

> GET key

Retrieves the value of a key. If the key does not exist or has expired, it returns (nil).

key: The key to retrieve.

Example:

> GET user:1:session
some-session-token

> GET non_existent_key
(nil)

### 5. License

This project is licensed under the MIT License - see the LICENSE file for details.