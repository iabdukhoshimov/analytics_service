INSERT INTO regions(name)
VALUES ('Tashkent'),
    ('Andijan'),
    ('Bukhara'),
    ('Fergana'),
    ('Jizzakh'),
    ('Namangan'),
    ('Navoiy'),
    ('Qashqadaryo'),
    ('Samarqand'),
    ('Sirdaryo'),
    ('Surxondaryo'),
    ('Tashkent'),
    ('Xorazm'),
    ('Karakalpakstan');

-- insert classificators group
INSERT INTO classificator_group(name)
VALUES ('general');

-- insert dynamic category
INSERT INTO dynamic_category(name)
VALUES ('СЕС'),
    ('Пожарная безопасность'),
    ('Кадастр');

-- insert status
INSERT INTO STATUS(name)
VALUES ('in-active'),
    ('in-progress'),
    ('rejected'),
    ('cancelled'),
    ('approved'),
    ('completed');

-- insert license type
INSERT INTO license_type(name)
VALUES ('permission'),
    ('license');

-- insert payment type
INSERT INTO payment_types(name)
VALUES ('cash'),
    ('payme'),
    ('click'),
    ('p2p'),
    ('uzum'),
    ('terminal');