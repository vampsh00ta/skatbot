package user

import "errors"

type User map[int64]int

func (user User) SetSubjectId(userId int64, subjectId int) {
	user[userId] = subjectId
}

func (user User) GetSubjectId(userId int64) (int, error) {
	subjectId, ok := user[userId]
	if !ok {
		return 0, errors.New("no such user")
	}
	return subjectId, nil
}
