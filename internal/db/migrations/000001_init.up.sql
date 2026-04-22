CREATE TABLE users (
  id UUID PRIMARY KEY,
  full_name VARCHAR(100) NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

-- password_hash (KHÔNG lưu password plain text)
-- email UNIQUE
CREATE TABLE accounts (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  account_number VARCHAR(20) UNIQUE NOT NULL,
  balance BIGINT NOT NULL DEFAULT 0,
  currency VARCHAR(10) NOT NULL DEFAULT 'VND',
  status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE',
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_accounts_user FOREIGN KEY (user_id) REFERENCES users (id)
);


-- BIGINT cho balance → tránh float
-- status: ACTIVE / BLOCKED / CLOSED
CREATE TABLE transactions (
  id UUID PRIMARY KEY,
  account_id UUID NOT NULL,
  amount BIGINT NOT NULL,
  type VARCHAR(20) NOT NULL, -- DEPOSIT / WITHDRAW / TRANSFER
  reference_id UUID,
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_transactions_account FOREIGN KEY (account_id) REFERENCES accounts (id)
);

-- Mọi thay đổi tiền đều phải có transaction
-- entries dùng để:
-- Ghi mỗi lần tiền tăng/giảm
-- Audit (check lại lịch sử)
-- Rebuild balance nếu cần
-- hống lỗi logic, double spending
CREATE TABLE entries (
  id UUID PRIMARY KEY,
  account_id UUID NOT NULL,
  amount BIGINT NOT NULL,
  transaction_id UUID,
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_entries_account FOREIGN KEY (account_id) REFERENCES accounts (id),
  CONSTRAINT fk_entries_transaction FOREIGN KEY (transaction_id) REFERENCES transactions (id)
);

-- Tách transfer và transaction → dễ audit
CREATE TABLE transfers (
  id UUID PRIMARY KEY,
  from_account_id UUID NOT NULL,
  to_account_id UUID NOT NULL,
  amount BIGINT NOT NULL,
  status VARCHAR(20) DEFAULT 'SUCCESS',
  created_at TIMESTAMP DEFAULT now(),
  CONSTRAINT fk_transfer_from FOREIGN KEY (from_account_id) REFERENCES accounts (id),
  CONSTRAINT fk_transfer_to FOREIGN KEY (to_account_id) REFERENCES accounts (id)
);

CREATE INDEX idx_user_email ON users (email);
CREATE INDEX idx_user_fullname ON users (full_name);

CREATE INDEX idx_accounts_user_id ON accounts (user_id);
CREATE INDEX idx_transactions_account_id ON transactions (account_id);
CREATE INDEX idx_transfers_from_account_id ON transfers (from_account_id);
CREATE INDEX idx_transfers_to_account_id ON transfers (to_account_id);
CREATE INDEX idx_entries_account_id ON entries (account_id);
CREATE INDEX idx_entries_transaction_id ON entries (transaction_id);

--users
--  └── accounts
--        └── entries   <<< CORE LEDGER
--        └── transactions
--        └── transfers