SELECT equipment.id, equipment_type.id, equipment_type.name, equipment.name, equipment.brand, equipment.model, equipment.cost, equipment.size, equipment.purchase_date, equipment.notes, equipment.mileage, equipment.mileage_unit, equipment.is_retired
FROM equipment
INNER JOIN equipment_type ON equipment.equipment_type_id = equipment_type.id
WHERE equipment.user_uuid = ?