package chat

import (
    "net"
)

type NotificationType int

const (
    Message NotificationType = iota
    Join
    Defect
)

type Notification struct {
    Type NotificationType
    ClientId int
    Connection net.Conn
    Message string
}
