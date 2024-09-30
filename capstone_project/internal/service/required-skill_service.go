package service

import (
	"capstone_project/internal/api/dto/game"
	"capstone_project/internal/models"
	"capstone_project/internal/repository"
)

type RequiredSkillService struct {
	repo *repository.RequiredSkillRepository
}

func NewRequiredSkillService(repo *repository.RequiredSkillRepository) *RequiredSkillService {
	return &RequiredSkillService{repo: repo}
}

func (s *RequiredSkillService) CreateRequiredSkill(createDTO game.RequiredSkillCreateDTO) (game.RequiredSkillDTO, error) {
	return game.RequiredSkillDTO{}, nil
}

func (s *RequiredSkillService) ListRequiredSkills() ([]game.RequiredSkillDTO, error) {
	skills, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var skillDTOs []game.RequiredSkillDTO
	for _, skill := range skills {
		skillDTOs = append(skillDTOs, game.RequiredSkillDTO{
			ID:   skill.ID,
			Name: skill.Name,
			Value: skill.Name,
			Label: skill.Name,
		})
	}

	return skillDTOs, nil
}

func (s *RequiredSkillService) CreateRequiredSkills(createDTO game.RequiredSkillCreateDTO) ([]game.RequiredSkillDTO, error) {
	var skillDTOs []game.RequiredSkillDTO

	for _, skillName := range createDTO.Skills {
		skill := &models.RequiredSkill{
			Name: skillName,
		}

		err := s.repo.Create(skill)
		if err != nil {
			return nil, err
		}

		skillDTOs = append(skillDTOs, game.RequiredSkillDTO{
			ID:   skill.ID,
			Name: skill.Name,
		})
	}

	return skillDTOs, nil
}
