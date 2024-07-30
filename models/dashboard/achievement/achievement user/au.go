package models

import (
	"time"

	models "github.com/SymbioSix/ProgressieAPI/models/auth"
)

type usr_user struct {
		`user_id`  uuid,
		`Rank_Id`  varchar,
		`email_verified` bool,
		`phone_verified` bool,
		`last_email_updated` timestamp,
		`total_courses_finished` int,
		`total_subcourses_finished` int,
		`total_point_achieved` int,
}

type csr_enrollment struct {
		`user_id` PK, FK, uuid,
		`course_id` PK, FK, varchar,
		`progress` float4,
}

type R