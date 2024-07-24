package psql

import (
	"fmt"
	"test/pkg/domain"
	"test/pkg/storage"

	//"go.uber.org/zap"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
type PSQL struct {
	config *Config
	logger *zap.Logger
	db     *gorm.DB
}
*/

/*
func NewPSQL(config *Config, logger *zap.Logger) (*PSQL, error) {

	db, err := gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&storage.Msg{}); err != nil {
		return nil, err
	}

	return &PSQL{
		config: config,
		logger: logger,
		db:     db,
	}, nil
}
*/

type PSQL struct {
	db *gorm.DB // Поле для хранения соединения с базой данных
}

func NewPSQL(db *gorm.DB) *PSQL {

	return &PSQL{db: db}

}

func (p *PSQL) CreateMsg(msg *domain.Msg) (*storage.Msg, error) {

	preMsg := storage.Msg{
		TimeStamp: msg.TimeStamp,
		Text:      msg.Text,
	}

	tx := p.db.Create(&preMsg)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &preMsg, nil
}

func (p *PSQL) DeleteMsg(id uint) error {
	var preMsg storage.Msg

	tx := p.db.Delete(&preMsg, id)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return fmt.Errorf("message with ID %d not found", id)
	}

	return nil
}

func (p *PSQL) UpdateMsg(id uint, newMsg *domain.Msg) (*storage.Msg, error) {

	var existingMsg storage.Msg

	tx := p.db.First(&existingMsg, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("message with ID %d not found", id)
		}
		return nil, tx.Error
	}

	existingMsg.TimeStamp = newMsg.TimeStamp
	existingMsg.Text = newMsg.Text

	tx = p.db.Save(&existingMsg)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingMsg, nil
}

func (p *PSQL) ReadMsg(id uint) (*storage.Msg, error) {

	var msg storage.Msg

	tx := p.db.First(&msg, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("message with ID %d not found", id)
		}
		return nil, tx.Error
	}

	return &msg, nil
}

func (p *PSQL) ReadAllMsgs() ([]storage.Msg, error) {
	var msgs []storage.Msg

	tx := p.db.Find(&msgs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return msgs, nil
}
