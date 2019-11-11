package Track

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/markelog/pento/back/database/models"
)

// Track type
type Track struct {
	db    *gorm.DB
	Model *models.Track
}

// New Track
func New(db *gorm.DB) *Track {
	return &Track{
		db: db,
	}
}

// CreateArgs are create arguments for track type
type CreateArgs struct {
	Email  string     `json:"email"`
	Active bool       `json:"active"`
	Name   string     `json:"name"`
	Start  *time.Time `json:"start"`
	Stop   *time.Time `json:"stop"`
}

// Create Track
func (track *Track) Create(args *CreateArgs) error {
	var (
		err error
		tx  = track.db.Begin()
	)

	if args.Active == true {
		err = track.start(tx, args)
	} else {
		err = track.stop(tx, args)
	}

	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (track *Track) start(tx *gorm.DB, args *CreateArgs) error {
	user := &models.User{
		Email:  args.Email,
		Active: args.Active,
	}

	err := tx.Where(models.User{
		Email:  args.Email,
		Active: args.Active,
	}).FirstOrCreate(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Create(&models.Track{
		UserID: user.ID,
		Name:   args.Name,
		Start:  args.Start,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (track *Track) stop(tx *gorm.DB, args *CreateArgs) error {
	user := &models.User{
		Email:  args.Email,
		Active: args.Active,
	}

	err := tx.Model(&user).Update("active", false).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	data := &models.Track{
		UserID: user.ID,
		Stop:   nil,
	}

	err = tx.Model(&data).Update("stop", args.Stop).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// ListValue result value for List() method
type ListValue struct {
	Name  string     `json:"name"`
	Start *time.Time `json:"start"`
	Stop  *time.Time `json:"stop"`
}

// List Tracks
func (track *Track) List(email string) ([]ListValue, error) {
	var (
		tracks []models.Track
		result []ListValue
	)

	user := track.db.Table("users").Select("id").Where(
		"email = ?", email,
	).QueryExpr()

	err := track.db.Select("name, start, stop").Where(
		"user_id = (?)",
		user,
	).Find(&tracks).Error

	if err != nil {
		return nil, err
	}

	for _, track := range tracks {
		result = append(result, ListValue{
			Start: track.Start,
			Stop:  track.Stop,
			Name:  track.Name,
		})
	}

	return result, nil
}
