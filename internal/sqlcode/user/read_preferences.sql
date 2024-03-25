SELECT user.cycle_start, user.cycle_days, user.initial_start_date
FROM user
WHERE user.uuid = ?