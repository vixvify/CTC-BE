package infra

import (
	"server/internal/models"
	"server/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamRepoGorm struct {
	db *gorm.DB
}

func NewTeamRepoGorm(db *gorm.DB) repository.TeamRepository {
	return &TeamRepoGorm{db: db}
}

func (r *TeamRepoGorm) WithTx(tx *gorm.DB) repository.TeamRepository {
	return &TeamRepoGorm{db: tx}
}

func (r *TeamRepoGorm) FindAll() ([]models.Team, error) {
	var teams []models.Team
	err := r.db.Find(&teams).Error
	return teams, err
}

func (r *TeamRepoGorm) FindByID(id uuid.UUID) (models.Team, error) {
	var team models.Team
	err := r.db.Where("id = ?", id).First(&team).Error
	return team, err
}

func (r *TeamRepoGorm) Create(team models.Team) (models.Team, error) {
	err := r.db.Create(&team).Error
	return team, err
}

func (r *TeamRepoGorm) Update(id uuid.UUID, data models.Team) (models.Team, error) {
	if err := r.db.
		Model(&models.Team{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"teamname": data.Teamname,
			"school":   data.School,
			"call_1":   data.Call_1,
			"call_2":   data.Call_2,
			"relation": data.Relation,
			"name_1":   data.Name_1,
			"name_2":   data.Name_2,
			"name_3":   data.Name_3,
			"name_4":   data.Name_4,
		}).
		Error; err != nil {
		return models.Team{}, err
	}

	var team models.Team
	if err := r.db.First(&team, "id = ?", id).Error; err != nil {
		return models.Team{}, err
	}

	return team, nil
}
