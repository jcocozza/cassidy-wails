UPDATE persisted_user_login
SET user_uuid = ?
WHERE id = 0; -- id will always be 0
