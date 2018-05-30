package config

type (
	JsonUser struct {
		ID        uint   `json:"id" valid:"-"`
		Name      string `json:"name" valid:"required"`
		Email     string `json:"email" valid:"required,email"`
		AvatarUrl string `json:"avatar_url" valid:"optional,url"`
	}
	JsonUserPost struct {
		Name     string `json:"name" valid:"optional"`
		Email    string `json:"email" valid:"required,email"`
		Password string `json:"password" valid:"optional,length(6|20)"`
	}
	JsonUserPut struct {
		Name      string `json:"name" valid:"optional"`
		AvatarUrl string `json:"avatar_url" valid:"optional,url"`
	}
	JsonUserChangePassword struct {
		Password    string `json:"password" valid:"required,length(6|20)"`
		NewPassword string `json:"new_password" valid:"required,length(6|20)"`
	}
	JsonUserResetPassword struct {
		Token       string `json:"token" valid:"required,length(40)"`
		NewPassword string `json:"new_password" valid:"required,length(6|20)"`
	}

	JsonRecord struct {
		ID      uint   `json:"id" valid:"-"`
		Title   string `json:"title" valid:"required"`
		Content string `json:"content" valid:"-"`
		Type    string `json:"type" valid:"required,in(PASSWORD|TEXT|KEY)"`
		IconUrl string `json:"icon_url" valid:"optional,url"`
	}

	JsonLogin struct {
		Email    string `valid:"required,email"`
		Password string `valid:"required,length(6|20)"`
	}

	ParamPageable struct {
		// Page  uint `json:"page" form:"page" query:"page"`
		Skip  uint `query:"skip" valid:"optional,matches(^[0-9]*$)"`
		Limit uint `query:"limit" valid:"optional,matches(^[1-9][0-9]*$)"`
	}

	JsonValidationError struct {
		Name      string `json:"name"`
		Err       string `json:"err"`
		Validator string `json:"validator"`
	}
)

func (e *JsonValidationError) Error() string {
	return e.Name + ": " + e.Err
}
