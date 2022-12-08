package repository

import (
	"fmt"
	"gorm.io/gorm"
	"kron-x/internal/dto"
	"kron-x/internal/models"
	"kron-x/pkg"
)

type CMSRepository struct {
	db *gorm.DB
}

func NewCMSRepository(db *gorm.DB) *CMSRepository {
	return &CMSRepository{
		db: db,
	}
}

func (r *CMSRepository) GetAllSliders() ([]*dto.Slider, error) {
	var sliders []*models.Slider
	if err := r.db.Find(&sliders).Error; err != nil {
		return nil, err
	} else {
		if len(sliders) == 0 {
			return []*dto.Slider{}, nil
		} else {
			var slidersDTO []*dto.Slider
			fmt.Println()
			for x := range sliders {

				result := &dto.Slider{
					Id:        fmt.Sprint(sliders[x].Id),
					MainText:  sliders[x].MainText,
					UpperText: sliders[x].UpperText,
					DownText:  sliders[x].DownText,
					Link:      sliders[x].Link,
					Image:     pkg.PhotoChecker(sliders[x].Image),
				}
				slidersDTO = append(slidersDTO, result)
			}
			return slidersDTO, nil
		}
	}
}

func (r *CMSRepository) GetAbout() ([]*dto.About, error) {
	var about *models.About
	if err := r.db.First(&about).Error; err != nil {
		return nil, err
	} else {
		if about == nil {
			return []*dto.About{}, nil
		} else {
			var aboutDTO []*dto.About
			result := &dto.About{
				MainImage:      pkg.PhotoChecker(about.MainImage),
				SecondaryImage: pkg.PhotoChecker(about.SecondaryImage),
				Image1:         pkg.PhotoChecker(about.Image1),
				Image2:         pkg.PhotoChecker(about.Image2),
				Image3:         pkg.PhotoChecker(about.Image3),
				Image4:         pkg.PhotoChecker(about.Image4),
				Image5:         pkg.PhotoChecker(about.Image5),
				Image6:         pkg.PhotoChecker(about.Image6),
				Image7:         pkg.PhotoChecker(about.Image7),
			}
			aboutDTO = append(aboutDTO, result)
			return aboutDTO, err
		}
	}
}

func (r *CMSRepository) GetRequisite() ([]*dto.Requisites, error) {
	var req *models.Requisite
	if err := r.db.First(&req).Error; err != nil {
		return nil, err
	} else {
		if req == nil {
			return []*dto.Requisites{}, nil
		} else {
			var reqDTO []*dto.Requisites
			result := &dto.Requisites{
				CompanyName: req.CompanyName,
				Address:     req.Address,
				Inn:         req.Inn,
				Kpp:         req.Kpp,
				Ogrn:        req.Ogrn,
				Rs:          req.Rs,
				Bank:        req.Bank,
				Ks:          req.Ks,
				Bik:         req.Bik,
			}
			reqDTO = append(reqDTO, result)
			return reqDTO, nil
		}
	}
}

func (r *CMSRepository) GetAllHeadOffice() ([]*dto.HeadOffice, error) {
	var headOffices []*models.HeadOffice
	if err := r.db.Find(&headOffices).Error; err != nil {
		return nil, err
	} else {
		if len(headOffices) == 0 {
			return []*dto.HeadOffice{}, nil
		} else {
			var headofficesDTO []*dto.HeadOffice
			for x := range headOffices {
				result := &dto.HeadOffice{
					Id:       fmt.Sprint(headOffices[x].Id),
					Title:    headOffices[x].Title,
					Phone:    headOffices[x].Phone,
					Mail:     headOffices[x].Mail,
					Schedule: headOffices[x].Schedule,
					Address:  headOffices[x].Address,
				}
				headofficesDTO = append(headofficesDTO, result)
			}
			return headofficesDTO, nil
		}
	}
}

func (r *CMSRepository) GetAllFeature() ([]*dto.Feature, error) {
	var features []*models.Feature
	if err := r.db.Find(&features).Error; err != nil {
		return nil, err

	} else {
		if len(features) == 0 {
			return []*dto.Feature{}, nil
		} else {
			var featuresDTO []*dto.Feature
			for x := range features {
				result := &dto.Feature{
					Id:     fmt.Sprint(features[x].Id),
					Header: features[x].Header,
					Body:   features[x].Body,
					Icon:   pkg.PhotoChecker(features[x].Icon),
				}
				featuresDTO = append(featuresDTO, result)
			}
			return featuresDTO, nil
		}
	}
}

func (r *CMSRepository) GetAllService() ([]*dto.Service, error) {
	var services []*models.Service
	if err := r.db.Find(&services).Error; err != nil {
		return nil, err
	} else {
		if len(services) == 0 {
			return []*dto.Service{}, nil
		} else {
			var servicesDTO []*dto.Service
			for x := range services {
				result := &dto.Service{
					Id:        fmt.Sprint(services[x].Id),
					Title:     services[x].Title,
					Body:      services[x].Body,
					Image:     pkg.PhotoChecker(services[x].Image),
					CreatedAt: services[x].CreatedAt,
				}
				servicesDTO = append(servicesDTO, result)
			}
			return servicesDTO, nil
		}
	}
}

func (r *CMSRepository) GetServiceById(id string) ([]*dto.Service, error) {
	var service *models.Service
	if err := r.db.Where("id=?", id).First(&service).Error; err != nil {
		return nil, err
	} else {
		var serviceDTO []*dto.Service

		result := &dto.Service{
			Id:        fmt.Sprint(service.Id),
			Title:     service.Title,
			Body:      service.Body,
			Image:     pkg.PhotoChecker(service.Image),
			CreatedAt: service.CreatedAt,
		}
		serviceDTO = append(serviceDTO, result)

		return serviceDTO, nil
	}
}

func (r *CMSRepository) FeedbackCall(feedbackCallData *dto.FeedbackCall) error {
	var feedBackCall *models.FeedbackCall

	feedBackCall = &models.FeedbackCall{
		Name:  feedbackCallData.Name,
		Phone: feedbackCallData.Phone,
	}
	if err := r.db.Create(&feedBackCall).Error; err != nil {
		return fmt.Errorf("error occured db")
	} else {
		return nil
	}
}

func (r *CMSRepository) FeedbackSelection(feedbackSelectionData *dto.FeedbackSelection) error {
	var feedBackSelection *models.FeedbackSelection

	feedBackSelection = &models.FeedbackSelection{
		Name:    feedbackSelectionData.Name,
		Phone:   feedbackSelectionData.Phone,
		Email:   feedbackSelectionData.Email,
		Comment: feedbackSelectionData.Comment,
	}
	if err := r.db.Create(&feedBackSelection).Error; err != nil {
		return fmt.Errorf("error occured db")
	} else {
		return nil
	}
}

func (r *CMSRepository) Support(supportData *dto.Support) error {
	var support *models.Support

	support = &models.Support{
		Name: supportData.Name,
		Body: supportData.Body,
	}
	if err := r.db.Create(&support).Error; err != nil {
		return fmt.Errorf("error occured db")
	} else {
		return nil
	}
}
