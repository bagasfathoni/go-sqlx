package config

// Write SQL query here
const (
	// Advisor queries
	SELECT_ALL_ADVISORS = `select name, email from advisors`
	SELECT_ADVISOR      = `select name, email from advisors where $1=$2`
	INSERT_ADVISOR      = `insert into advisors(name,email) values ($1,$2)`
	UPDATE_ADVISOR      = `update advisors set $3=$4 where $1=$2`

	// Department queries
	SELECT_ALL_DEPARTMENTS = `select name, email from departments`
	SELECT_DEPARTMENTS     = `select name, email from departments where $1=$2`
	INSERT_DEPARTMENTS     = `insert into departments(name,school) values ($1,$2)`
	UPDATE_DEPARTMENTS     = `update advisors set $3=$4 where $1=$2`

	// Student queries
	SELECT_ALL_STUDENTS = `select students.name, students.email, students.gpa, students.credit,	advisors.name, departments.name from students JOIN advisors ON advisors.id = students.advisor JOIN departments ON departments.id = students.department;`
	SELECT_STUDENT      = `select students.name, students.email, students.gpa, students.credit,	advisors.name, departments.name from students JOIN advisors ON advisors.id = students.advisor JOIN departments ON departments.id = students.department where $1=$2`
	INSERT_STUDENT      = `insert into students(name,email,gpa,credit,department,advisor) values ($1,$2,$3,$4,$5,$6)`
	UPDATE_STUDENT      = `update students set $3=$4 where $1=$2 `
)
