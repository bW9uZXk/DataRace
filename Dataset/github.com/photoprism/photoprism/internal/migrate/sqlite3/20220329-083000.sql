UPDATE files SET media_id = CASE WHEN photo_id IS NOT NULL AND file_missing = 0 AND deleted_at IS NULL THEN ((10000000000 - photo_id) || '-' || (1 + file_sidecar - file_primary) || '-' || file_uid) END WHERE 1;