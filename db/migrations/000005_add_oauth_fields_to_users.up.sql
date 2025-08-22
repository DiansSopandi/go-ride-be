-- Make password nullable for OAuth users
ALTER TABLE users ALTER COLUMN password DROP NOT NULL;

-- Add provider field (local, google, facebook, etc.)
ALTER TABLE users ADD COLUMN provider VARCHAR(50) NOT NULL DEFAULT 'local';

-- Add provider_id for OAuth providers
ALTER TABLE users ADD COLUMN provider_id VARCHAR(255) NULL;

-- Add profile picture URL
ALTER TABLE users ADD COLUMN picture TEXT NULL;

-- Add unique constraint for provider + provider_id combination
ALTER TABLE users ADD CONSTRAINT unique_provider_id 
UNIQUE (provider, provider_id);

-- Add index for faster OAuth lookups
CREATE INDEX idx_users_provider_id ON users (provider, provider_id);

-- Update existing users to have 'local' provider
UPDATE users SET provider = 'local' WHERE provider IS NULL;