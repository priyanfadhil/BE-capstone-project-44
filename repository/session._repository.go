package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type sessionRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *sessionRepositoryMysqlLayer) CreateSession(session model.Session) error {
	res := r.DB.Create(&session)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *sessionRepositoryMysqlLayer) GetAllSession() []model.Session {
	sessions := []model.Session{}
	r.DB.Find(&sessions)
	//r.DB.Model(&model.User{}).Association("rents").Find(&users)

	return sessions
}

func (r *sessionRepositoryMysqlLayer) GetOneSessionByID(id int) (session model.Session, err error) {
	res := r.DB.Where("id = ?", id).Find(&session)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *sessionRepositoryMysqlLayer) UpdateOneSessionByID(id int, session model.Session) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&session)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *sessionRepositoryMysqlLayer) DeleteSessionByID(id int) error {
	res := r.DB.Delete(&model.Session{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func SessionMysqlRepository(db *gorm.DB) domain.AdapterSessionRepository {
	return &sessionRepositoryMysqlLayer{
		DB: db,
	}
}
