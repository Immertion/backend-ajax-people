package user

type User struct {
	Id             int    `json:"-" db:"id"`
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`
	Password       string `json:"password"`
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
	AvatarPath     string `json:"avatar_path"`
	IsModerated    bool   `json:"is_moderated"`
}
