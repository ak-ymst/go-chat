package chat

import (
    "net"
)

type Sender struct {
    Id int
    Connection net.Conn
}

func (sender Sender) SendMessage(message string) {
    var buf = []byte(message);

    _, error := sender.Connection.Write(buf)
    if error != nil {
        panic(error);
    }
}
