-- name: LicenseGetAll :many
SELECT *
FROM license
WHERE TRUE
    AND CASE
        WHEN sqlc.arg('search')::VARCHAR != '' THEN document_number ILIKE '%' || sqlc.arg('search') || '%'
        OR organization_name ILIKE '%' || sqlc.arg('search') || '%'
        OR stir_number ILIKE '%' || sqlc.arg('search') || '%'
        OR reestr_number ILIKE '%' || sqlc.arg('search') || '%'
        ELSE TRUE
    END
    AND CASE
        WHEN sqlc.arg('license_type')::INTEGER != 0 THEN license_type = sqlc.arg('license_type')
        ELSE TRUE
    END
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: LicenseGetAllCount :one
SELECT COUNT(*)
FROM license
WHERE TRUE
    AND CASE
        WHEN sqlc.arg('search')::VARCHAR != '' THEN document_number ILIKE '%' || sqlc.arg('search') || '%'
        OR organization_name ILIKE '%' || sqlc.arg('search') || '%'
        OR stir_number ILIKE '%' || sqlc.arg('search') || '%'
        OR reestr_number ILIKE '%' || sqlc.arg('search') || '%'
        ELSE TRUE
    END
    AND CASE
        WHEN sqlc.arg('license_type')::INTEGER != 0 THEN license_type = sqlc.arg('license_type')
        ELSE TRUE
    END;

-- name: LicenseGetOne :one
SELECT *
FROM license
WHERE id = sqlc.arg('id');

-- name: LicenseInsertOne :one
INSERT INTO license (
        document_number,
        granted_date,
        lifetime,
        organization_name,
        stir_number,
        reestr_number,
        work_category,
        doc_file,
        license_type
    )
VALUES (
        sqlc.arg('document_number'),
        sqlc.arg('granted_date'),
        sqlc.arg('lifetime'),
        sqlc.arg('organization_name'),
        sqlc.arg('stir_number'),
        sqlc.arg('reestr_number'),
        sqlc.arg('work_category'),
        sqlc.arg('doc_file'),
        sqlc.arg('license_type')
    )
RETURNING id;