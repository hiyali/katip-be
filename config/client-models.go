package config

type (
  JsonUser struct {
    ID        uint `gorm:"primary_key" json:"id" valid:"-"`

    Name      string  `gorm:"type:varchar(64);not null" json:"name" valid:"required"`
    Email     string  `gorm:"type:varchar(100);unique_index;not null" json:"email" valid:"required,email"`
    AvatarUrl string  `gorm:"type:varchar(256)" json:"avatar_url" valid:"optional,url"`
  }
  JsonUserPost struct {
    Name      string  `gorm:"type:varchar(64);not null" json:"name" valid:"optional"`
    Email     string  `gorm:"type:varchar(100);unique_index;not null" json:"email" valid:"required,email"`
    Password  string  `gorm:"type:varchar(64);not null" json:"password" valid:"optional,length(6|20)"`
  }
  JsonUserPut struct {
    Name      string  `gorm:"type:varchar(64);not null" json:"name" valid:"optional"`
    AvatarUrl string  `gorm:"type:varchar(256)" json:"avatar_url" valid:"optional,url"`
  }
  JsonUserChangePassword struct {
    Password	string  `gorm:"type:varchar(64);not null" valid:"required,length(6|20)"`
    NewPassword string  `gorm:"type:varchar(64);not null" json:"new_password" valid:"required,length(6|20)"`
  }

  JsonRecord struct {
    ID        uint `gorm:"primary_key" json:"id" valid:"-"`

    Title     string  `gorm:"type:varchar(100)" json:"title" valid:"required"`
    Content   string  `gorm:"type:varchar(2048)" json:"content" valid:"-"`
    Type      string  `gorm:"type:varchar(20)" json:"type" valid:"required,in(PASSWORD|TEXT|KEY)"`
    IconUrl   string  `gorm:"type:varchar(256)" json:"icon_url" valid:"optional,url"`
  }

  JsonLogin struct {
    Email     string `valid:"required,email"`
    Password  string `valid:"required,length(6|20)"`
  }

  ParamPageable struct {
    // Page  uint `json:"page" form:"page" query:"page"`
    Skip  uint `query:"skip" valid:"optional,matches(^[0-9]*$)"`
    Limit uint `query:"limit" valid:"optional,matches(^[1-9][0-9]*$)"`
  }

  JsonValidationError struct {
    Name      string  `json:"name"`
    Err       string  `json:"err"`
    Validator string  `json:"validator"`
  }
)

func (e *JsonValidationError) Error() string {
  return e.Name + ": " + e.Err
}
