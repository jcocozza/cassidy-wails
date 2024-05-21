-- change the uuid of an activity
UPDATE activity
SET uuid = ?
WHERE uuid = ?;

UPDATE planned
SET activity_uuid = ?
WHERE activity_uuid = ?;

UPDATE completed
SET activity_uuid = ?
WHERE activity_uuid = ?;
