-- Get the date of the most recent activity with completion data
SELECT MAX(DATE(a.date)) AS lastest_completed
FROM activity a
INNER JOIN completed c ON a.uuid = c.activity_uuid
WHERE (c.distance != 0 OR c.moving_duration != 0 OR c.elapsed_duration != 0 OR c.vertical != 0) AND a.user_uuid  = ?;
