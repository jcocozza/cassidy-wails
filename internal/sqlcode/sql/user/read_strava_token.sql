SELECT access_token, token_type, refresh_token, expiry
FROM strava_token
WHERE user_uuid = ?