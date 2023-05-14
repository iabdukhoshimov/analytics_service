-- insert public role
INSERT INTO roles (id, name)
VALUES (0, 'public');

-- insert roles
INSERT INTO roles (name)
VALUES ('user'),
    ('inspector'),
    ('moderator'),
    ('admin');

-- insert paths
INSERT INTO paths (path)
VALUES ('/v1/ping'),
    ('/v1/users'),
    ('/v1/login');

-- insert permissions
INSERT INTO permissions (
        role_id,
        path_id,
        can_insert,
        can_update,
        can_delete,
        can_read
    )
VALUES (1, 1, false, false, false, TRUE),
    (2, 1, false, false, false, TRUE),
    (3, 1, false, false, false, TRUE),
    (4, 1, false, false, false, TRUE),
    (0, 2, TRUE, TRUE, TRUE, TRUE),
    (1, 2, TRUE, TRUE, TRUE, TRUE),
    (2, 2, TRUE, TRUE, TRUE, TRUE),
    (3, 2, TRUE, TRUE, TRUE, TRUE),
    (4, 2, TRUE, TRUE, TRUE, TRUE),
    (0, 3, TRUE, TRUE, TRUE, TRUE),
    (1, 3, TRUE, TRUE, TRUE, TRUE),
    (2, 3, TRUE, TRUE, TRUE, TRUE),
    (3, 3, TRUE, TRUE, TRUE, TRUE),
    (4, 3, TRUE, TRUE, TRUE, TRUE);