package service

import (
    "mockexample/log"
)

// Service is just an example of a service you might create
// that takes a logger as a param or option... or something.
type Service struct{
    logger log.Logger
}

// Typically this logger would be optional but for purposes of
// this example we'll just make it the only param.
func NewService(logger log.Logger) *Service {
    return &Service{
        logger: logger,
    }
}

func (s *Service) IsThingEmpty(thing string) bool {
    s.logger.Debug("checking if '" + thing + "' is empty")
    if thing == "" {
        return true
    }

    return false
}