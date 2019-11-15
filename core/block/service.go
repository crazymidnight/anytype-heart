package block

import (
	"errors"
	"log"
	"sync"

	"github.com/anytypeio/go-anytype-middleware/core/anytype"
	"github.com/anytypeio/go-anytype-middleware/pb"
)

var (
	ErrBlockNotFound    = errors.New("block not found")
	ErrBlockAlreadyOpen = errors.New("block already open")
)

type Service interface {
	OpenBlock(id string) error
	CloseBlock(id string) error
	Close() error
}

func NewService(accountId string, lib anytype.Anytype, sendEvent func(event *pb.Event)) Service {
	return &service{
		accountId:   accountId,
		anytype:     lib,
		sendEvent:   sendEvent,
		smartBlocks: make(map[string]smartBlock),
	}
}

type service struct {
	anytype     anytype.Anytype
	accountId   string
	sendEvent   func(event *pb.Event)
	smartBlocks map[string]smartBlock
	m           sync.RWMutex
}

func (s *service) OpenBlock(id string) (err error) {
	s.m.Lock()
	defer s.m.Unlock()
	if _, ok := s.smartBlocks[id]; ok {
		return ErrBlockAlreadyOpen
	}
	sb, err := openSmartBlock(s, id)
	if err != nil {
		return
	}
	s.smartBlocks[id] = sb
	return nil
}

func (s *service) CloseBlock(id string) (err error) {
	s.m.Lock()
	defer s.m.Unlock()
	if sb, ok := s.smartBlocks[id]; ok {
		delete(s.smartBlocks, id)
		return sb.Close()
	}
	return ErrBlockNotFound
}

func (s *service) CreateBlock(req pb.RpcBlockCreateRequest) (string, error) {
	s.m.RLock()
	defer s.m.RUnlock()
	if sb, ok := s.smartBlocks[req.ContextId]; ok {
		return sb.Create(req)
	}
	return "", ErrBlockNotFound
}

func (s *service) Close() error {
	s.m.Lock()
	defer s.m.Unlock()
	for _, sb := range s.smartBlocks {
		if err := sb.Close(); err != nil {
			log.Printf("block[%s] close error: %v", sb.GetId(), err)
		}
	}
	return nil
}
