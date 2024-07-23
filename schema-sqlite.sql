-- Active: 1721762008159@@127.0.0.1@3306@kplus
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    phone TEXT NOT NULL UNIQUE,
    email TEXT UNIQUE,
    role TEXT NOT NULL DEFAULT 'user',
    phone_verified BOOLEAN NOT NULL DEFAULT 0,
    email_verified BOOLEAN NOT NULL DEFAULT 0,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_details (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    identity_number TEXT NOT NULL UNIQUE,
    full_name TEXT NOT NULL,
    legal_name TEXT NOT NULL,
    place_of_birth TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    salary DECIMAL NOT NULL,
    selfie TEXT NOT NULL,
    selfie_with_national_id TEXT NOT NULL,
    national_id_image TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS loan_limits (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    limit DECIMAL NOT NULL, tenor INTEGER NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    contract_number TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    otr DECIMAL NOT NULL,
    fee DECIMAL NOT NULL,
    installment DECIMAL NOT NULL,
    interest DECIMAL NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    asset_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS installments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    transaction_id INTEGER NOT NULL,
    installment DECIMAL NOT NULL,
    due_date DATE NOT NULL,
    paid_date DATE,
    period INTEGER NOT NULL,
    status TEXT NOT NULL DEFAULT 'unpaid',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transactions (id) ON DELETE CASCADE
);