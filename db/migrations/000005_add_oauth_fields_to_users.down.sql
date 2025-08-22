-- Rollback migration
DROP INDEX IF EXISTS idx_users_provider_id;
ALTER TABLE users DROP CONSTRAINT IF EXISTS unique_provider_id;
ALTER TABLE users DROP COLUMN IF EXISTS picture;
ALTER TABLE users DROP COLUMN IF EXISTS provider_id;
ALTER TABLE users DROP COLUMN IF EXISTS provider;
ALTER TABLE users ALTER COLUMN password SET NOT NULL;