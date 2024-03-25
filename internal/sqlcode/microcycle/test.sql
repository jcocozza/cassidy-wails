WITH DateRanges AS (
    -- Define your date ranges here
    SELECT '2024-01-01' AS start_date, '2024-01-08' AS end_date
    UNION ALL
    SELECT '2024-01-09', '2024-01-16'
    UNION ALL
    SELECT '2024-01-17', '2024-01-24'
    UNION ALL
    SELECT '2024-01-25', '2024-02-01'
)
SELECT
    DateRanges.start_date,
    DateRanges.end_date,
    SUM(planned.distance) AS total_planned_distance
FROM
    DateRanges
LEFT JOIN planned ON planned.activity_uuid IN (
    SELECT activity_uuid
    FROM activity
    WHERE (date BETWEEN DateRanges.start_date AND DateRanges.end_date)
)
GROUP BY
    DateRanges.start_date, DateRanges.end_date;