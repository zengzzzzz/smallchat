# Go SmallChat Server

A simple chat server implemented in Go, allowing multiple clients to connect, exchange messages, and set nicknames. This reposity reference from [smallchat](https://github.com/antirez/smallchat).

## Table of Contents

- [Go SmallChat Server](#go-smallchat-server)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Usage](#usage)
  - [Application Example](#application-example)
  - [License](#license)

## Features

- Multiple clients can connect to the server.
- Clients can set their nicknames using the `/nick` command.
- Messages sent by clients are broadcasted to all connected clients.
- Supports messages with newlines.

## Prerequisites

- [Go](https://golang.org/dl/) (Go 1.21 or higher recommended)

## Usage


1. Change to the project directory:

   ```sh
   cd go-chat-server
   ```

2. Build the server:

   ```sh
   go build chat_server.go
   ```

3. Run the server:

   ```sh
   ./main
   ```

4. Open multiple terminal windows and use telnet or netcat to connect to the server:

   ```sh
   nc localhost 7711
   ```

5. Set your nickname using the `/nick` command:

   ```sh
   /nick MyNickname
   ```

6. Start sending messages, including newlines:

   ```sh
   Hello, everyone!
   This is a multiline message.
   ```

## Application Example

```sh
go run chat_server.go
```

1. Start the server, and it will begin listening for client connections.

2. Open multiple terminal windows and connect to the server using telnet or netcat.

3. Set nicknames for each client using the `/nick` command.

4. Clients can send messages, and the server will broadcast these messages to all other clients.

   Example:

   ```
   arduinoCopy code
   [Client 1] /nick Alice
   [Client 2] /nick Bob
   [Client 1] Hello, Bob!
   [Client 2] Hi, Alice!
   ```

## License

This project is licensed under the  Apache-2.0 License.