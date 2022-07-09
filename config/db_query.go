package config

// Write SQL query here
const (
	// Advisor queries
	SELECT_ALL_ADVISORS = `select * from advisors`
	SELECT_ADVISOR      = `select * from advisors where $1=$2`
	INSERT_ADVISOR      = `insert into advisors(name,email) values ($1,$2)`
	UPDATE_ADVISOR      = `update advisors set $3=$4 where $1=$2`

	// Department queries
	SELECT_ALL_DEPARTMENTS = `select * from departments`
	SELECT_DEPARTMENTS     = `select * from departments where $1=$2`
	INSERT_DEPARTMENTS     = `insert into departments(name,school) values ($1,$2)`
	UPDATE_DEPARTMENTS     = `update advisors set $3=$4 where $1=$2`
)
