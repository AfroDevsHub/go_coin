package constants

type Occupation int

const (
	SOFTWARE_ENGINEER Occupation = iota
	HARDWARE_ENGINEER
	ELECTRICAL_ENGINEER
	MECHANICAL_ENGINEER
	CIVIL_ENGINEER
	BIOMEDICAL_ENGINEER
	DOCTOR
	NURSE
	TEACHER
	PROFESSOR
	ARTIST
	MUSICIAN
	WRITER
	ACCOUNTANT
	LAWYER
	POLICE_OFFICER
	FIREFIGHTER
	CHEF
	ARCHITECT
	SCIENTIST
	STUDENT
	RETIREE
	ENTREPRENEUR
	ATHLETE
	JOURNALIST
	DESIGNER
	PHARMACIST
	SOCIAL_WORKER
	OTHER_OCCUPATION
)

var occupations = map[Occupation]string{
	SOFTWARE_ENGINEER:   "Software Engineer",
	HARDWARE_ENGINEER:   "Hardware Engineer",
	ELECTRICAL_ENGINEER: "Electrical Engineer",
	MECHANICAL_ENGINEER: "Mechanical Engineer",
	CIVIL_ENGINEER:      "Civil engineer",
	BIOMEDICAL_ENGINEER: "Biomedical Engineer",
	DOCTOR:              "Doctor",
	NURSE:               "Nurse",
	TEACHER:             "Teacher",
	PROFESSOR:           "Professor",
	ARTIST:              "Artist",
	MUSICIAN:            "Musician",
	WRITER:              "Writer",
	ACCOUNTANT:          "Accountant",
	LAWYER:              "Lawyer",
	POLICE_OFFICER:      "Police Officer",
	FIREFIGHTER:         "Firefighter",
	CHEF:                "Chef",
	ARCHITECT:           "Architect",
	SCIENTIST:           "Scientist",
	STUDENT:             "Student",
	RETIREE:             "Retiree",
	ENTREPRENEUR:        "Entrepreneur",
	ATHLETE:             "Athlete",
	JOURNALIST:          "Journalist",
	DESIGNER:            "Designer",
	PHARMACIST:          "Pharmacist",
	SOCIAL_WORKER:       "Social Worker",
	OTHER_OCCUPATION:    "Other",
}
