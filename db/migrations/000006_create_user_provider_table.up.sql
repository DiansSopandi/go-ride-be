CREATE TABLE IF NOT EXISTS user_providers (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL CHECK (provider IN ('local', 'google', 'facebook', 'github', 'apple')),
    provider_id VARCHAR(255) NOT NULL,
    provider_email VARCHAR(255), -- Email dari provider (bisa berbeda dengan user.email)
    provider_data JSONB, -- Data tambahan dari provider (avatar, name, etc)
    access_token_hash VARCHAR(255), -- Hash dari access token (jangan simpan plain text)
    refresh_token_hash VARCHAR(255), -- Hash dari refresh token
    token_expires_at TIMESTAMP, -- Kapan token expired
    is_active BOOLEAN DEFAULT true, -- Apakah provider account masih aktif
    last_login_at TIMESTAMP, -- Kapan terakhir login dengan provider ini
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Constraints
    CONSTRAINT uq_user_providers_user_provider UNIQUE (user_id, provider),
    CONSTRAINT uq_user_providers_provider_id UNIQUE (provider, provider_id)
);

-- Indexes untuk performance
CREATE INDEX idx_user_providers_user_id ON user_providers(user_id);
CREATE INDEX idx_user_providers_provider_id ON user_providers(provider, provider_id);
CREATE INDEX idx_user_providers_active ON user_providers(is_active) WHERE is_active = true;
CREATE INDEX idx_user_providers_last_login ON user_providers(last_login_at DESC);

-- Trigger untuk auto-update updated_at
-- CREATE OR REPLACE FUNCTION update_updated_at_column()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     NEW.updated_at = CURRENT_TIMESTAMP;
--     RETURN NEW;
-- END;
-- $$ language 'plpgsql';

-- CREATE TRIGGER update_user_providers_updated_at
--     BEFORE UPDATE ON user_providers
--     FOR EACH ROW
--     EXECUTE FUNCTION update_updated_at_column();