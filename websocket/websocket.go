package websocket

import (
	"dodo-open-go/log"
	"dodo-open-go/tools"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
	"time"
)

// Client WebSocket client interface
type Client interface {
	Connect() error
	Listen() error
	Write(event *WSEventMessage) error
	Reconnect() error
	Close()
}

// New a WebSocket instance
func New(url string, options ...OptionHandler) (Client, error) {
	conf := &config{
		messageQueueSize: 10000,
	}

	for _, optionHandler := range options {
		if optionHandler == nil {
			return nil, errors.New("invalied OptionHandler (nil detected)")
		}
		if err := optionHandler(conf); err != nil {
			return nil, err
		}
	}

	c := &client{
		Url:             url,
		messageChan:     make(messageChan, conf.messageQueueSize),
		closeChan:       make(errorChan, 16),
		heartbeatTicker: time.NewTicker(time.Second * 25),
		isConnected:     false,
	}

	return c, nil
}

type (
	// messageChan message channel
	messageChan chan *WSEventMessage

	// errorChan error channel to handle errors
	errorChan chan error

	// client WebSocket client implement
	client struct {
		Url             string          // WebSocket connect url
		conn            *websocket.Conn // WebSocket connection
		messageChan     messageChan     // message channel
		closeChan       errorChan       // errors channel
		heartbeatTicker *time.Ticker    // ticker for heartbeat
		isConnected     bool            // connection status
	}

	config struct {
		messageQueueSize int
	}
)

func (c *client) Connect() error {
	if c.Url == "" {
		return errors.New("WebSocket URL invalid")
	}

	var err error
	c.conn, _, err = websocket.DefaultDialer.Dial(c.Url, nil)
	if err != nil {
		log.Errorf("connect error: %v", err)
		return err
	}

	return nil
}

// Listen message and handle it
func (c *client) Listen() error {
	defer c.Close()

	// read message
	go c.readMessage()
	// listen and handle message
	go c.listenMessageAndHandle()

	for {
		select {
		case err := <-c.closeChan:
			log.Errorf("stop listening, error: %v", err)
			if DefaultHandlers.ErrorHandler != nil {
				DefaultHandlers.ErrorHandler(err)
			}
			return err
		case <-c.heartbeatTicker.C:
			packet := &WSEventMessage{Type: HeartbeatType}
			_ = c.Write(packet)
		}
	}
}

// Write message
func (c *client) Write(event *WSEventMessage) error {
	m, _ := tools.JSON.Marshal(event)
	if err := c.conn.WriteMessage(websocket.TextMessage, m); err != nil {
		log.Errorf("write message failed cause: %v", err)
		c.closeChan <- err
		return err
	}
	return nil
}

// Reconnect to the WebSocket server
func (c *client) Reconnect() error {
	// TODO
	return nil
}

// Close connection and stop heartbeat ticker
func (c *client) Close() {
	if err := c.conn.Close(); err != nil {
		log.Errorf("close connection failed cause: %v", err)
	}
	c.heartbeatTicker.Stop()
	close(c.messageChan)
	close(c.closeChan)
}

// readMessage read message from connection
func (c *client) readMessage() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Errorf("read message error cause: %v", err)
			c.closeChan <- err
			return
		}
		event := &WSEventMessage{}
		if err = tools.JSON.Unmarshal(message, &event); err != nil {
			log.Errorf("json unmarshal failed cause: %v", err)
			continue
		}
		event.RawData = message
		c.messageChan <- event
	}
}

// listenMessageAndHandle listen and handle message from message channel
func (c *client) listenMessageAndHandle() {
	defer func() {
		if err := recover(); err != nil {
			c.closeChan <- fmt.Errorf("we got panic: %v", err)
		}
	}()

	for event := range c.messageChan {
		if event.Type == HeartbeatType {
			continue
		}

		if err := ParseDataAndHandle(event); err != nil {
			log.Errorf("try to parse and handle message failed cause: %v", err)
		}
	}
}
