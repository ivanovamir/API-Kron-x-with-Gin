package service

import (
	"kron-x/internal/dto"
	"kron-x/internal/repository"
	"kron-x/pkg/email"
)

type CMSService struct {
	repo repository.CMS
}

func NewCMSService(repo repository.CMS) *CMSService {
	return &CMSService{
		repo: repo,
	}
}

func (s *CMSService) GetAllSliders() ([]*dto.Slider, error) {
	return s.repo.GetAllSliders()
}

func (s *CMSService) GetAbout() ([]*dto.About, error) {
	return s.repo.GetAbout()
}

func (s *CMSService) GetRequisite() ([]*dto.Requisites, error) {
	return s.repo.GetRequisite()
}

func (s *CMSService) GetAllHeadOffice() ([]*dto.HeadOffice, error) {
	return s.repo.GetAllHeadOffice()
}

func (s *CMSService) GetAllFeature() ([]*dto.Feature, error) {
	return s.repo.GetAllFeature()
}

func (s *CMSService) GetAllService() ([]*dto.Service, error) {
	return s.repo.GetAllService()
}

func (s *CMSService) GetServiceById(id string) ([]*dto.Service, error) {
	return s.repo.GetServiceById(id)
}

func (s *CMSService) FeedbackCall(feedbackCallData *dto.FeedbackCall) error {
	if err := s.repo.FeedbackCall(feedbackCallData); err != nil {
		return err
	} else {
		newEmail := email.NewEmail()
		if err := newEmail.SendFeedbackCall(feedbackCallData); err != nil {
			return err
		} else {
			return nil
		}
	}

}

func (s *CMSService) FeedbackSelection(feedbackSelectionData *dto.FeedbackSelection) error {
	if err := s.repo.FeedbackSelection(feedbackSelectionData); err != nil {
		return err
	} else {
		newEmail := email.NewEmail()
		if err := newEmail.SendFeedbackSelection(feedbackSelectionData); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func (s *CMSService) Support(supportData *dto.Support) error {
	if err := s.repo.Support(supportData); err != nil {
		return err
	} else {
		newEmail := email.NewEmail()
		if err := newEmail.SendSupport(supportData); err != nil {
			return err
		} else {
			return nil
		}
	}
}
