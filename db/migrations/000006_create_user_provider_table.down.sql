-- DROP TABLE IF EXISTS user_provider;

-- Safe rollback with data validation
-- Rollback: create_user_providers

-- Check if table has data before dropping (optional safety check)
DO $$
DECLARE
    record_count INTEGER;
BEGIN
    -- Check if table exists first
    IF EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'user_providers') THEN
        -- Count records
        SELECT COUNT(*) INTO record_count FROM user_providers;
        
        -- Log warning if data exists
        IF record_count > 0 THEN
            RAISE WARNING 'Rolling back user_providers table with % records. Data will be lost!', record_count;
            -- Optionally uncomment the line below to prevent accidental data loss
            -- RAISE EXCEPTION 'Rollback aborted to prevent data loss. Use force rollback if intentional.';
        END IF;
        
        -- Drop triggers
        -- DROP TRIGGER IF EXISTS update_user_providers_updated_at ON user_providers;
        
        -- Drop function (be careful, other tables might use it)
        -- Only drop if no other triggers are using it
        -- IF NOT EXISTS (
        --     SELECT 1 FROM information_schema.triggers 
        --     WHERE action_statement LIKE '%update_updated_at_column%'
        --     AND trigger_name != 'update_user_providers_updated_at'
        -- ) THEN
        --     DROP FUNCTION IF EXISTS update_updated_at_column();
        -- END IF;
        
        -- Drop indexes (optional, will be dropped with table anyway)
        DROP INDEX IF EXISTS idx_user_providers_last_login;
        DROP INDEX IF EXISTS idx_user_providers_active;
        DROP INDEX IF EXISTS idx_user_providers_provider_id;
        DROP INDEX IF EXISTS idx_user_providers_user_id;
        
        -- Drop table
        DROP TABLE user_providers;
        
        RAISE NOTICE 'Successfully rolled back user_providers table';
    ELSE
        RAISE NOTICE 'Table user_providers does not exist, rollback skipped';
    END IF;
END $$;