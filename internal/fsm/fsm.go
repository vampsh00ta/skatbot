package fsm

import (
	"fmt"
	"github.com/go-telegram/bot/models"
	"sync"
)

type fsm struct {
	mu   *sync.Mutex
	data map[int64]*User
}

type Fsm interface {
	SetKeyboard(userId int64, keyboard *models.InlineKeyboardMarkup)
	InitKeyboard(userId int64, keyboard *models.InlineKeyboardMarkup)

	GetKeyboard(userId int64) *models.InlineKeyboardMarkup
	ClearKeyboard(userId int64)

	AddData(userId int64, val interface{})
	GetData(userId int64) interface{}
	DeleteData(userId int64)
}

type User struct {
	currVal  interface{}
	keyboard []*models.InlineKeyboardMarkup
}

func New() Fsm {
	return &fsm{
		&sync.Mutex{},
		make(map[int64]*User),
	}
}
func (fsm fsm) InitKeyboard(userId int64, keyboard *models.InlineKeyboardMarkup) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	data, ok := fsm.data[userId]

	if !ok {
		fsm.data[userId] = &User{keyboard: []*models.InlineKeyboardMarkup{keyboard}}
		return
	}
	(*data).keyboard = []*models.InlineKeyboardMarkup{keyboard}

}

func (fsm fsm) ClearKeyboard(userId int64) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	data, ok := fsm.data[userId]
	if !ok {
		return
	}

	data.keyboard = []*models.InlineKeyboardMarkup{}
	return
}

func (fsm fsm) SetKeyboard(userId int64, keyboard *models.InlineKeyboardMarkup) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	data, ok := fsm.data[userId]

	if !ok {
		fsm.data[userId] = &User{keyboard: []*models.InlineKeyboardMarkup{}}
		return
	}
	(*data).keyboard = append((*data).keyboard, keyboard)
	return
}
func (fsm fsm) GetKeyboard(userId int64) *models.InlineKeyboardMarkup {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	data, ok := fsm.data[userId]
	if !ok {
		return nil
	}

	if len((data).keyboard) > 0 {
		(*data).keyboard = (*data).keyboard[0 : len(data.keyboard)-1]
	}
	step := len(data.keyboard)
	return data.keyboard[step-1]
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
	fmt.Println()
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
