package config

type (
  JsonUser struct {
    ID        uint `gorm:"primary_key" json:"id"`

    Name      string  `gorm:"type:varchar(64);not null" json:"name"`
    Email     string  `gorm:"type:varchar(100);unique_index;not null" json:"email"`
  }

  JsonRecord struct {
    ID        uint `gorm:"primary_key" json:"id"`

    Title     string  `gorm:"type:varchar(100)" form:"title" json:"title"`
    Content   string  `gorm:"type:varchar(2048)" form:"content" json:"content"`
    Type      string  `gorm:"type:varchar(20)" json:"type"`
    IconUrl   string  `gorm:"type:varchar(256)" form:"icon_url" json:"icon_url"`
  }

  ParamPageable struct {
    // Page  uint `json:"page" form:"page" query:"page" validate:"min=1"`
    Page  uint `query:"page" validate:"min=1"`
    Limit uint `query:"limit" validate:"min=1"`
  }

  ParamLogin struct {
    Email     string `form:"email" validate:"min=6;required"`
    Password  string `form:"password" validate:"min=6;reuiqred"`
  }
)
