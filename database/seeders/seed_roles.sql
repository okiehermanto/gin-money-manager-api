INSERT INTO roles (id, label, name)
VALUES
    (gen_random_uuid(), 'Admin', 'admin'),
    (gen_random_uuid(), 'User', 'user')
ON CONFLICT (id) DO NOTHING;