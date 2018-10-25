package main

import (
    "fmt"
    "net"
    "./chat"
)

func main() {

    listener, error := net.Listen("tcp", "localhost:10000");

    if error != nil {
        panic(error);
    }

    fmt.Println("Server running at localhost:10000");

    var channel = make(chan chat.Notification);
    var observer chat.Observer = chat.Observer{ Senders: make([]chat.Sender, 0, 5), Subject: channel};
    go observer.WaitNotice();
    waitClient(listener, 0, observer, channel);

}

func waitClient(listener net.Listener, sequence int, observer chat.Observer, channel chan chat.Notification) {
    connection, error := listener.Accept();

    if error != nil {
        panic(error);
    }

    var reciever chat.Reciever = chat.Reciever{ Id: sequence, Connection: connection, Observer: channel};
    go reciever.Start();

    waitClient(listener, sequence + 1, observer, channel);
}
