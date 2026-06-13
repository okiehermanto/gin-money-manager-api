CREATE TABLE todos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    todo TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
