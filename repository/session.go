package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type SessionRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *SessionRepositoryMysqlLayer) CreateSession(session model.Session) error {
	res := r.DB.Create(&session)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *SessionRepositoryMysqlLayer) GetAllSessions() []model.Session {
	sessions := []model.Session{}
	r.DB.Find(&sessions)
	//r.DB.Model(&model.User{}).Association("rents").Find(&users)

	return sessions
}

func (r *SessionRepositoryMysqlLayer) GetOneSessionByID(id int) (session model.Session, err error) {
	res := r.DB.Where("id = ?", id).Find(&session)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *SessionRepositoryMysqlLayer) UpdateOneSessionByID(id int, session model.Session) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&session)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *SessionRepositoryMysqlLayer) DeleteSessionByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Session{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewSessionMysqlRepository(db *gorm.DB) domain.AdapterSessionRepository {
	return &SessionRepositoryMysqlLayer{
		DB: db,
	}
}
