SELECT activity_equipment.id, activity_equipment.activity_uuid, activity_equipment.assigned_mileage, activity_equipment.assigned_mileage_unit,
    equipment.id, equipment.user_uuid, equipment.name, equipment.brand, equipment.model, equipment.cost, equipment.size, equipment.purchase_date, equipment.notes, equipment.mileage, equipment.mileage_unit,
    equipment_type.id, equipment_type.name
FROM activity_equipment
INNER JOIN equipment ON equipment.id = activity_equipment.equipment_id
INNER JOIN equipment_type ON equipment_type.id = equipment.equipment_type_id
WHERE activity_equipment.activity_uuid IN (%s);