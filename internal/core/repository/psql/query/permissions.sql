-- name: PermissionGetOneByRoleAndPath :one
SELECT roles.name AS role_name,
    paths.path AS path,
    permissions.can_insert AS can_insert,
    permissions.can_update AS can_update,
    permissions.can_delete AS can_delete,
    permissions.can_read AS can_read
FROM permissions
    JOIN roles ON roles.id = permissions.role_id
    JOIN paths ON paths.id = permissions.path_id
WHERE roles.id = sqlc.arg('role_id')
    AND paths.path = sqlc.arg('path')
LIMIT 1;