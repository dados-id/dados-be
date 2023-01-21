-- name: CreateTag :one
INSERT INTO tags (
  name
) VALUES (
  $1
) RETURNING *;

-- name: RandomTag :one
SELECT name FROM tags
ORDER BY RANDOM()
LIMIT 1;

-- name: ListTagsByProfessorRatingId :many
SELECT T.name as tags
FROM professor_ratings PR
JOIN professor_rating_tags PRT ON PRT.professor_rating_id = PR.id
JOIN tags T ON PRT.tag_name = T.name
WHERE PR.id = $1;

-- name: ListTagsByProfessorRatingId :many
-- SELECT T.name as tags
-- FROM professor_ratings PR
-- JOIN professor_rating_tags PRT ON PRT.professor_rating_id = PR.id
-- JOIN tags T ON PRT.tag_name = T.name
-- JOIN users U ON U.id = PR.user_id
-- WHERE U.id = $1;
