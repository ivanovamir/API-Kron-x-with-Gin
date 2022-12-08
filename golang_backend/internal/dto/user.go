package dto

type RegisterUser struct {
	Inn      string `json:"inn" binding:"required"`
	Email    string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

type UserAuth struct {
	Email    string `json:"mail"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Code     string `json:"code"`
}

type UserInformation struct {
	Email string `json:"mail" binding:"required"`
}

type CodeGenerator struct {
	Phone string `json:"phone"`
	Email string `json:"mail"`
}

type SendOrderEmailUr struct {
	Address    string
	OrderId    string
	Email      string
	Inn        string
	Comment    string
	LogoImage  string
	CartImage  string
	CheckImage string
	Product    []*OrderProduct
	TotalPrice string
}

type SendOrderEmailPhyz struct {
	Address    string
	OrderId    string
	Phone      string
	Comment    string
	LogoImage  string
	CartImage  string
	CheckImage string
	Product    []*OrderProduct
	TotalPrice string
}
