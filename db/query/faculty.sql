-- name: CreateFaculty :one
INSERT INTO faculties (
  name
) VALUES (
  $1
) RETURNING *;

-- name: ListFacultyBySchool :many
SELECT F.* FROM faculties F
JOIN school_faculty_associations SFA ON SFA.faculty_id = F.id
WHERE SFA.school_id = $1;

-- name: RandomFacultyID :one
SELECT id FROM faculties
ORDER BY RANDOM()
LIMIT 1;
