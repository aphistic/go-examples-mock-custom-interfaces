package main

import (
    "mockexample/log"
    "mockexample/service"
)

// See log/logger.go for the starting comments/description of this example.

func main() {
    l := log.NewStdoutLogger()
    s := service.NewService(l)

    if s.IsThingEmpty("not empty") {
        l.Info("Thing is empty\n")
    } else {
        l.Error("Thing is not empty\n")
    }
}