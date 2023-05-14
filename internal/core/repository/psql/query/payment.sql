-- name: PaymentGetAll :many
SELECT *
FROM payments
WHERE deleted_at IS NULL
    AND organization_id = sqlc.arg('organization_id')
    AND CASE
        WHEN sqlc.arg('status')::INTEGER != 0 THEN STATUS = sqlc.arg('status')
        ELSE TRUE
    END
    AND CASE
        WHEN sqlc.arg('type')::INTEGER != 0 THEN TYPE = sqlc.arg('type')
        ELSE TRUE
    END
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: PaymentGetAllCount :one
SELECT COUNT(*)
FROM payments
WHERE deleted_at IS NULL
    AND organization_id = sqlc.arg('organization_id')
    AND CASE
        WHEN sqlc.arg('status')::INTEGER != 0 THEN STATUS = sqlc.arg('status')
        ELSE TRUE
    END
    AND CASE
        WHEN sqlc.arg('type')::INTEGER != 0 THEN TYPE = sqlc.arg('type')
        ELSE TRUE
    END;

-- name: PaymentGetOne :one
SELECT *
FROM payments
WHERE id = sqlc.arg('id');

-- name: PaymentInsertOne :one
INSERT INTO payments (
        organization_id,
        amount,
        requisites,
        STATUS,
        TYPE
    )
VALUES (
        sqlc.arg('organization_id'),
        sqlc.arg('amount'),
        sqlc.arg('requisites'),
        sqlc.arg('status'),
        sqlc.arg('type')
    )
RETURNING id;