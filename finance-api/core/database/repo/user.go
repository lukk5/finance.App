package repo

import (
	"context"
	"finance-api-v1/core/database"
	"finance-api-v1/core/entities"
	"gorm.io/gorm"
)

type UserRepo interface {
	Save(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, id uint, user *entity.User) error
	Get(ctx context.Context, id uint) (*entity.User, error)
}

type userDB struct {
	db *gorm.DB
}

func (udb *userDB) Get(ctx context.Context, id uint) (*entity.User, error) {
	db := database.FromContext(ctx, udb.db)

	var user entity.User

	if err := db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		// failed
		if database.IsRecordNotFoundErr(err) {
			return nil, database.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (udb *userDB) Save(ctx context.Context, user *entity.User) error {

	db := database.FromContext(ctx, udb.db)

	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		// TODO: log
		if database.IsKeyConflictErr(err) {
			return database.ErrKeyConflict
		}
		return err
	}
	return nil
}

func (udb *userDB) Update(ctx context.Context, id uint, user *entity.User) error {

	db := database.FromContext(ctx, udb.db)

	fields := make(map[string]interface{})

	if user.UserName != "" {
		fields["userName"] = user.UserName
	}

	chain := db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).UpdateColumns(fields)

	if chain.Error != nil {
		return chain.Error
	}
	if chain.RowsAffected == 0 {
		return database.ErrNotFound
	}
	return nil
}

// NewUserRepo creates a new account db with given db
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userDB{
		db: db,
	}
}
