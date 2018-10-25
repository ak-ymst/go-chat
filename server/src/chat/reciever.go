package chat

import (
    "net"
)

type Reciever struct {
    Id int
    Connection net.Conn
    Observer chan<- Notification
}

func (reciever Reciever) Start() {
    reciever.Observer <- Notification{ Type: Join, ClientId: reciever.Id, Connection:reciever.Connection }

    reciever.WaitMessage();
}

func (reciever Reciever) WaitMessage() {
    var buf = make([]byte, 1024);

    n, error := reciever.Connection.Read(buf);
    if (error != nil) {
        reciever.Observer <- Notification{ Type: Defect, ClientId: reciever.Id}
        return;
    }

    reciever.Observer <- Notification{ Type: Message, ClientId: reciever.Id, Message: string(buf[:n])}

    reciever.WaitMessage();
}
