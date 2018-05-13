package config

import "time"

type (
  JsonUser struct {
    ID        uint `gorm:"primary_key" json:"id" valid:"-"`

    Name      string  `gorm:"type:varchar(64);not null" json:"name" form:"name" valid:"required"`
    Email     string  `gorm:"type:varchar(100);unique_index;not null" json:"email" form:"email" valid:"required,email"`
  }
  JsonUserPut struct {
    UpdatedAt time.Time

    Name      string  `gorm:"type:varchar(64);not null" json:"name" form:"name" valid:"optional"`
    Password  string  `gorm:"type:varchar(64);not null" form:"password" valid:"optional,length(6|20)"`
  }

  JsonRecord struct {
    ID        uint `gorm:"primary_key" json:"id" valid:"-"`

    Title     string  `gorm:"type:varchar(100)" form:"title" json:"title" valid:"required"`
    Content   string  `gorm:"type:varchar(2048)" form:"content" json:"content" valid:"-"`
    Type      string  `gorm:"type:varchar(20)" json:"type" valid:"required,in(PASSWORD|TEXT|KEY)"`
    IconUrl   string  `gorm:"type:varchar(256)" form:"icon_url" json:"icon_url" valid:"-"`
  }

  JsonLogin struct {
    Email     string `form:"email" valid:"required,email"`
    Password  string `form:"password" valid:"required,length(6|20)"`
  }

  ParamPageable struct {
    // Page  uint `json:"page" form:"page" query:"page"`
    Page  uint `query:"page" valid:"required,matches(^[1-9][0-9]*$)"`
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
