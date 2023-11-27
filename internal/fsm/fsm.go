package fsm

import (
	"github.com/go-telegram/bot/models"
	"sync"
)

type fsm struct {
	mu   *sync.Mutex
	data map[int64]*User
}

type Fsm interface {
	SetKeyboard(userId int64, keyboard *models.ReplyMarkup)
	GetKeyboard(userId int64) *models.ReplyMarkup
	AddData(userId int64, val interface{})
	GetData(userId int64) interface{}
	DeleteData(userId int64)
}

type User struct {
	currVal  interface{}
	keyboard *models.ReplyMarkup
}

func New() Fsm {
	return &fsm{
		&sync.Mutex{},
		make(map[int64]*User),
	}
}
func (fsm fsm) SetKeyboard(userId int64, keyboard *models.ReplyMarkup) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	data, ok := fsm.data[userId]
	if !ok {
		fsm.data[userId] = &User{keyboard: keyboard}
		return
	}
	data.keyboard = keyboard
	return
}
func (fsm fsm) GetKeyboard(userId int64) *models.ReplyMarkup {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	data, ok := fsm.data[userId]
	if !ok {
		return nil
	}

	return data.keyboard
}
func (fsm fsm) AddData(userId int64, val interface{}) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	data, ok := fsm.data[userId]
	if !ok {
		fsm.data[userId] = &User{currVal: val}
		return
	}
	(*data).currVal = val
	return
}
func (fsm fsm) GetData(userId int64) interface{} {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	val, ok := fsm.data[userId]
	if !ok {
		return nil
	}
	return val.currVal
}
func (fsm fsm) DeleteData(userId int64) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	data, ok := fsm.data[userId]
	if !ok {
		return
	}
	data.currVal = nil
}
