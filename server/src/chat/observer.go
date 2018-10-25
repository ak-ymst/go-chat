package chat

import (
    "net"
    "fmt"
)


type Observer struct {
    Senders []Sender
    Subject <-chan Notification
}

func (observer Observer) WaitNotice() {
    notice := <-observer.Subject

    switch notice.Type {
    case Message:
        for i := range observer.Senders {
            observer.Senders[i].SendMessage(notice.Message);
        }

        break;

    case Join:
        observer.Senders = appendSender(notice.ClientId, notice.Connection, observer.Senders);

        fmt.Printf("Client %d join, now menber count is %d\n", notice.ClientId, len(observer.Senders));
        break;

    case Defect:
        observer.Senders = removeSender(notice.ClientId, observer.Senders);

        fmt.Printf("Client %d defect, now menber count is %d\n", notice.ClientId, len(observer.Senders));
        break;

    default:

    }

    observer.WaitNotice();
}

func appendSender(senderId int, connection net.Conn, senders []Sender) []Sender {
    return append(senders, Sender{ Id: senderId, Connection: connection})
}

func removeSender(senderId int, senders []Sender) []Sender {
    var find = -1;

    for i := range senders {
        if (senders[i].Id == senderId) {
            find = i;
            break;
        }
    }

    if (find == -1) {
        return senders;
    }

    return append(senders[:find], senders[find+1:]...);
}

