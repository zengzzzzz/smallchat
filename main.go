package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	MAX_CLIENTS  = 1000
	MAX_NICK_LEN = 32
	SERVER_PORT  = "7711"
)

type Client struct {
	conn net.Conn
	nick string
}

type ChatState struct {
	serverListener net.Listener
	clients        map[net.Conn]*Client
}

var Chat *ChatState

func createTCPServer(port string) (net.Listener, error) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func initChat() {
	Chat = &ChatState{
		clients: make(map[net.Conn]*Client),
	}
}

func handleClientConnection(conn net.Conn) {
	client := &Client{
		conn: conn,
	}
	Chat.clients[conn] = client

	welcomeMessage := "Welcome to Simple Chat! Use /nick <nickname> to set your nick.\n"
	conn.Write([]byte(welcomeMessage))

	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Disconnected client: %v\n", client.nick)
			delete(Chat.clients, conn)
			conn.Close()
			return
		}
		message := string(buf[:n])
		message = strings.TrimSpace(message)
		if strings.HasPrefix(message, "/nick") {
			newNick := strings.TrimPrefix(message, "/nick")
			client.nick = newNick
		} else {
			message = strings.Replace(message, "\n", "\n", -1)
			for otherConn, otherClient := range Chat.clients {
				if otherConn != conn {
					msgToSend := fmt.Sprintf("%s: %s", client.nick, message)
					_, err := otherConn.Write([]byte(msgToSend))
					if err != nil {
						fmt.Printf("Error sending message to client: %v\n", otherClient.nick)
					}
				}
			}
		}
	}
}

func main() {
	initChat()
	serverListener, err := createTCPServer(SERVER_PORT)
	if err != nil {
		fmt.Printf("Error creating server: %v\n", err)
		os.Exit(1)
	}
	defer serverListener.Close()
	for {
		conn, err := serverListener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			os.Exit(1)
		}
		go handleClientConnection(conn)
	}
}
