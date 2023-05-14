-- name: OrganizationGetAll :many
SELECT *
FROM organization
WHERE TRUE
    AND CASE
        WHEN sqlc.arg('search')::VARCHAR != '' THEN name ILIKE '%' || sqlc.arg('search') || '%'
        OR full_name ILIKE '%' || sqlc.arg('search') || '%'
        OR phone_number ILIKE '%' || sqlc.arg('search') || '%'
        OR location ILIKE '%' || sqlc.arg('search') || '%'
        ELSE TRUE
    END
    AND CASE
        WHEN sqlc.arg('parent_organization')::VARCHAR != '' THEN parent_organization = sqlc.arg('parent_organization')
        ELSE TRUE
    END
    AND deleted_at IS NULL
ORDER BY name ASC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: OrganizationGetAllCount :one
SELECT COUNT(*)
FROM organization
WHERE TRUE
    AND CASE
        WHEN sqlc.arg('search')::VARCHAR != '' THEN name ILIKE '%' || sqlc.arg('search') || '%'
        OR full_name ILIKE '%' || sqlc.arg('search') || '%'
        OR phone_number ILIKE '%' || sqlc.arg('search') || '%'
        OR location ILIKE '%' || sqlc.arg('search') || '%'
        ELSE TRUE
    END
    AND CASE
        WHEN sqlc.arg('parent_organization')::VARCHAR != '' THEN parent_organization = sqlc.arg('parent_organization')
        ELSE TRUE
    END
    AND deleted_at IS NULL;

-- name: OrganizationGetOne :one
SELECT *
FROM organization
WHERE id = sqlc.arg('id');

-- name: OrganizationInsertOne :one
INSERT INTO organization (
        name,
        full_name,
        phone_number,
        parent_organization,
        location
    )
VALUES (
        sqlc.arg('name'),
        sqlc.arg('full_name'),
        sqlc.arg('phone_number'),
        sqlc.arg('parent_organization'),
        sqlc.arg('location')
    )
RETURNING id;
