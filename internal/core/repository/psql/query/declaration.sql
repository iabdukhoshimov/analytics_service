-- name: DeclarationGetAll :many
SELECT d.*,
    o.name AS organization_name
FROM declaration d
    LEFT JOIN organization o ON o.id = d.organization_id
    AND o.deleted_at IS NULL
WHERE d.deleted_at IS NULL
    AND d.organization_id = sqlc.arg('organization_id')
    AND o.id IS NOT NULL
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: DeclarationGetAllCount :one
SELECT COUNT(*)
FROM declaration d
    LEFT JOIN organization o ON o.id = d.organization_id
    AND o.deleted_at IS NULL
WHERE d.deleted_at IS NULL
    AND d.organization_id = sqlc.arg('organization_id')
    AND o.id IS NOT NULL;

-- name: DeclarationGetOne :one
SELECT d.*,
    o.name AS organization_name
FROM declaration AS d
    LEFT JOIN organization o ON o.id = d.organization_id
WHERE d.id = sqlc.arg('id');

-- name: DeclarationInsertOne :one
INSERT INTO declaration (
        organization_id,
        danger_rate,
        reasons_of_danger,
        converage_of_the_danger_area,
        proof,
        location_info,
        residents_info,
        life_insurance,
        tech_document
    )
VALUES (
        sqlc.arg('organization_id'),
        sqlc.arg('danger_rate'),
        sqlc.arg('reasons_of_danger'),
        sqlc.arg('converage_of_the_danger_area'),
        sqlc.arg('proof'),
        sqlc.arg('location_info'),
        sqlc.arg('residents_info'),
        sqlc.arg('life_insurance'),
        sqlc.arg('tech_document')
    )
RETURNING id;