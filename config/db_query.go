package config

// Write SQL query here
const (
	// Advisor queries
	SELECT_ALL_ADVISORS = `select * from advisors`
	SELECT_ADVISOR      = `select * from advisors where $1=$2`
	INSERT_ADVISOR      = `insert into advisors(name,email) values ($1,$2)`
	UPDATE_ADVISOR      = `update advisors set $3=$4 where $1=$2`
)
