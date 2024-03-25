SELECT user.uuid, user.username, user.password, user.units, user.cycle_start, user.cycle_days
FROM user
WHERE username = ?;