SELECT user.uuid, user.username, user.password, user.units, user.cycle_start, user.cycle_days, user.initial_start_date
FROM user
WHERE username = ?;