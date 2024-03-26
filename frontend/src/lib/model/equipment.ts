import { model } from "../wailsjs/go/models";
import { EmptyLength, type Distance } from "./distance";

import { CreateActivityEquipment, UpdateActivityEquipment } from '../wailsjs/go/controllers/EquipmentHandler'

/**
 * Create an empty equipment
 * @param usr the user creating the equipment
 * @returns a new equipment object
 */
export function NewEquipment(usr: model.User): model.Equipment {
    let a = {
        brand: "",
        user_uuid: usr.uuid,
        cost: 0,
        equipment_type: {id: -1, name: ""},
        id: -1,
        mileage: EmptyLength(false, usr),
        model: "",
        name: "",
        notes: "",
        purchase_date: "",
        size: "",
        is_retired: false
    }
    let e = new model.Equipment(a)
    return e
}
export function EmptyActivityEquipment(uuid: string, unit: string): model.ActivityEquipment {
    let a = {
        activity_uuid: uuid,
        assigned_mileage: {unit: unit, length: 0},
        equipment: null,
        id: -1
    }
    return new model.ActivityEquipment(a)
}
/**
 * Update/create activity equipment based on wether it is new or not.
 * @param activity_equipment_list an activity equipment list for an activity
 */
export async function HandleActivityEquipmentList(activity_equipment_list: model.ActivityEquipment[]) {
    activity_equipment_list.forEach((ae, index) => {
        if (ae.id == -1) {
            // run create a new one
            CreateActivityEquipment(ae)
        } else {
            // run update
            UpdateActivityEquipment(ae.id, ae)
        }
    });
};