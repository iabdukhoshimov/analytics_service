-- name: UserInsertOne :one
INSERT INTO users (
        first_name,
        second_name,
        last_name,
        phone_number,
        email,
        region_id,
        profile_picture,
        hashed_password
    )
VALUES (
        sqlc.arg('first_name'),
        sqlc.arg('second_name'),
        sqlc.arg('last_name'),
        sqlc.arg('phone_number'),
        sqlc.arg('email'),
        sqlc.arg('region_id'),
        sqlc.arg('profile_picture'),
        sqlc.arg('hashed_password')
    )
RETURNING id;

-- name: UserGetOne :one
SELECT *
FROM users
WHERE id = sqlc.arg('id');

-- name: UserGetOneByEmail :one
SELECT *
FROM users
WHERE email = sqlc.arg('email');

-- name: UserDeleteOne :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = sqlc.arg('id');

-- name: UserUpdateOne :exec
UPDATE users
SET first_name = COALESCE(sqlc.narg(first_name), first_name),
    second_name = COALESCE(sqlc.narg(second_name), second_name),
    last_name = COALESCE(sqlc.narg(last_name), last_name),
    phone_number = COALESCE(sqlc.narg(phone_number), phone_number),
    email = COALESCE(sqlc.narg(email), email),
    region_id = COALESCE(sqlc.narg(region_id), region_id),
    profile_picture = COALESCE(sqlc.narg(profile_picture), profile_picture),
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    updated_at = NOW()
WHERE id = sqlc.arg('id');

-- name: UserGetAll :many
SELECT *
FROM users
WHERE deleted_at IS NULL
    AND CASE
        WHEN sqlc.arg('search')::VARCHAR != '' THEN first_name ILIKE '%' || sqlc.arg('search') || '%'
        OR second_name ILIKE '%' || sqlc.arg('search') || '%'
        OR last_name ILIKE '%' || sqlc.arg('search') || '%'
        OR phone_number ILIKE '%' || sqlc.arg('search') || '%'
        OR email ILIKE '%' || sqlc.arg('search') || '%'
        ELSE TRUE
    END
    AND CASE
        WHEN sqlc.arg('region_id')::INTEGER != 0 THEN region_id = sqlc.arg('region_id')
        ELSE TRUE
    END
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: UserGetAllCount :one
SELECT COUNT(*)
FROM users
WHERE deleted_at IS NULL
    AND CASE
        WHEN sqlc.arg('search')::VARCHAR != '' THEN first_name ILIKE '%' || sqlc.arg('search') || '%'
        OR second_name ILIKE '%' || sqlc.arg('search') || '%'
        OR last_name ILIKE '%' || sqlc.arg('search') || '%'
        OR phone_number ILIKE '%' || sqlc.arg('search') || '%'
        OR email ILIKE '%' || sqlc.arg('search') || '%'
        ELSE TRUE
    END
    AND CASE
        WHEN sqlc.arg('region_id')::INTEGER != 0 THEN region_id = sqlc.arg('region_id')
        ELSE TRUE
    END;