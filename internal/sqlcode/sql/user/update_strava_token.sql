UPDATE strava_token
SET access_token = ?, token_type = ?, refresh_token = ?, expiry = ?
WHERE user_uuid = ?;