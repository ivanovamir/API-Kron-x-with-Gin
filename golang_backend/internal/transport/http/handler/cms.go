package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kron-x/internal/dto"
	"net/http"
)

// @Summary Get all sliders
// @Tags cms
// @Description Get all sliders
// @ID get_all_sliders
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.Slider}
// @Failure 404 {object} dto.ErrorResponse
// @Router /cms/sliders [get]
func (h *Handler) GetAllSliders(c *gin.Context) {

	allSlidersDTO, err := h.service.GetAllSliders()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": allSlidersDTO,
		})
	}

}

// @Summary Get company about
// @Tags cms
// @Description Get company about
// @ID get_all_about
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.About}
// @Failure 404 {object} dto.ErrorResponse
// @Router /cms/about [get]
func (h *Handler) GetAbout(c *gin.Context) {
	aboutDTO, err := h.service.GetAbout()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": aboutDTO,
		})
	}
}

// @Summary Get company requisites
// @Tags cms
// @Description Get company requisites
// @ID get_all_requisites
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.Requisites}
// @Failure 404 {object} dto.ErrorResponse
// @Router /cms/req [get]
func (h *Handler) GetRequisite(c *gin.Context) {
	reqDTO, err := h.service.GetRequisite()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": reqDTO,
		})
	}
}

// @Summary Get all head offices
// @Tags cms
// @Description Get all head offices
// @ID get_all_head offices
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.HeadOffice}
// @Failure 404 {object} dto.ErrorResponse
// @Router /cms/headoffice [get]
func (h *Handler) GetAllHeadOffice(c *gin.Context) {
	allHeadOfficesDTO, err := h.service.GetAllHeadOffice()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": allHeadOfficesDTO,
		})
	}
}

// @Summary Get all features
// @Tags cms
// @Description Get all features
// @ID get_all_features
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.Feature}
// @Failure 404 {object} dto.ErrorResponse
// @Router /cms/features [get]
func (h *Handler) GetAllFeature(c *gin.Context) {
	allFeaturesDTO, err := h.service.GetAllFeature()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": allFeaturesDTO,
		})
	}
}

// @Summary Get all services
// @Tags cms
// @Description Get all services
// @ID get_all_services
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.Service}
// @Failure 404 {object} dto.ErrorResponse
// @Router /cms/service [get]
func (h *Handler) GetAllService(c *gin.Context) {
	allServicesDTO, err := h.service.GetAllService()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": allServicesDTO,
		})
	}
}

// @Summary Get service detail
// @Tags cms
// @Description Get service detail
// @ID get_detail_service
// @Accept json
// @Produce json
// @Success 200 {object} dto.DefaultData{data=[]dto.Service}
// @Failure 404 {object} dto.ErrorResponse
// @Param service_id query string false "Required service id"
// @Router /cms/service_detail [get]
func (h *Handler) GetServiceById(c *gin.Context) {
	serviceId := c.Query("service_id")
	serviceDTO, err := h.service.GetServiceById(serviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": []string{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": serviceDTO,
		})
	}
}

// @Summary Send feedback call request
// @Tags cms
// @Description Send feedback call request
// @ID send_feedback_call
// @Accept json
// @Produce json
// @Param input body dto.FeedbackCall true "new feed back request info"
// @Failure 404 {object} dto.ErrorResponse
// @Router /cms/feedback_call [post]
func (h *Handler) FeedbackCall(c *gin.Context) {
	var feedbackCall *dto.FeedbackCall
	if err := c.ShouldBindJSON(&feedbackCall); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong json input",
		})
	} else {
		err := h.service.FeedbackCall(feedbackCall)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprint(err),
			})
		}
	}
}

// @Summary Send feedback selection request
// @Tags cms
// @Description Send feedback selection request
// @ID send_feedback_selection
// @Accept json
// @Produce json
// @Param input body dto.FeedbackSelection true "new feed back details selection request info"
// @Failure 404 {object} dto.ErrorResponse
// @Router /cms/feedback_selection [post]
func (h *Handler) FeedbackSelection(c *gin.Context) {
	var feedbackSelection *dto.FeedbackSelection
	if err := c.ShouldBindJSON(&feedbackSelection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong json input",
		})
	} else {
		err := h.service.FeedbackSelection(feedbackSelection)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprint(err),
			})
		}
	}
}

// @Summary Send support request
// @Tags cms
// @Description Send support request
// @ID send_support
// @Accept json
// @Produce json
// @Param input body dto.Support true "new support request info"
// @Failure 404 {object} dto.ErrorResponse
// @Router /cms/support [post]
func (h *Handler) Support(c *gin.Context) {
	var support *dto.Support
	if err := c.ShouldBindJSON(&support); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong json input",
		})
	} else {
		err := h.service.Support(support)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprint(err),
			})
		}
	}
}
