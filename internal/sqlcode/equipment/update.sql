UPDATE equipment
SET name = ?, brand = ?, model = ?, cost = ?, size = ?, purchase_date = ?, notes = ?, mileage = ?, mileage_unit = ?, is_retired = ?
WHERE equipment.id = ?;