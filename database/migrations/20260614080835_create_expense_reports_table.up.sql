CREATE TABLE expense_reports (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    balance_amount BIGINT DEFAULT 0,
    balance_currency VARCHAR(5) DEFAULT 'IDR', 
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_expense_reports_user_id ON expense_reports (user_id);
CREATE INDEX idx_expense_reports_name ON expense_reports (name);
CREATE INDEX idx_expense_reports_type ON expense_reports (type);



CREATE TABLE expense_report_statements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    expense_report_id UUID NOT NULL, -- Fixed: Typo 'UUId' to 'UUID'
    author_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    attachment VARCHAR(255),
    amount_amount BIGINT DEFAULT 0,
    amount_currency VARCHAR(5) DEFAULT 'IDR',
    settled BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_expense_report FOREIGN KEY(expense_report_id) REFERENCES expense_reports(id) ON DELETE CASCADE,
    CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_statement_expense_report_id ON expense_report_statements (expense_report_id);
CREATE INDEX idx_statement_author_id ON expense_report_statements (author_id);
CREATE INDEX idx_statement_name ON expense_report_statements (name);