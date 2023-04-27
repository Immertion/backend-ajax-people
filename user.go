package user

import "errors"

type User struct {
	Id             int    `json:"-" db:"id"`
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Age            int    `json:"age"`
	Mail           string `json:"mail"`
	StatusUser     string `json:"statusUser"`
	EducationLevel string `json:"educationLevel"`
	StudyProgramId int    `json:"studyProgramId"`
	SchoolId       int    `json:"schoolId"`
	AdmissionYear  string `json:"admissionYear"`
	GraduationYear string `json:"graduationYear"`
	IsAdmin        bool   `json:"isAdmin"`
	IsVerificated  bool   `json:"isVerificated"`
	IsVisible      bool   `json:"isVisible"`
	AvatarPath     string `json:"avatarPath"`
	IsModerated    bool   `json:"isModerated"`
}

type UpdateUserInput struct {
	FirstName      *string `json:"firstName"`
	LastName       *string `json:"lastName"`
	Age            *int    `json:"age"`
	StatusUser     *string `json:"statusUser"`
	EducationLevel *string `json:"educationLevel"`
	StudyProgramId *int    `json:"studyProgramId"`
	SchoolId       *int    `json:"schoolId"`
	AdmissionYear  *string `json:"admissionYear"`
	GraduationYear *string `json:"graduationYear"`
	AvatarPath     *string `json:"avatarPath"`
}

func (i UpdateUserInput) Validate() error {
	if i.FirstName == nil && i.LastName == nil && i.Age == nil && i.StatusUser == nil &&
		i.EducationLevel == nil && i.StudyProgramId == nil && i.SchoolId == nil &&
		i.AdmissionYear == nil && i.GraduationYear == nil && i.AvatarPath == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type Faculty struct {
	Id    int    `json:"-" db:"id"`
	Title string `json:"title" binding:"required"`
}
