<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import EquipmentModal from "./EquipmentModal.svelte";
    import type { model } from "../../wailsjs/go/models";


    export let equipment: model.Equipment;
    let is_hovering: boolean = false;

    function handleMouseOver() {
        is_hovering = true;
    }
    function handleMouseLeave() {
        is_hovering = false;
    }
    const dispatch = createEventDispatcher();
</script>

<!-- svelte-ignore a11y-mouse-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="equipment-card card" on:mouseover={handleMouseOver} on:mouseleave={handleMouseLeave}>
    <div class="card-body">
        <h5> {equipment.name} | {equipment.brand}, {equipment.model} </h5>
        <h6> {equipment.equipment_type?.name }</h6>
        <p> { equipment.mileage?.length} {equipment.mileage?.unit} </p>

        <p> {equipment.notes} </p>
        <EquipmentModal
            bind:equipment={equipment}
            bind:is_hovering={is_hovering}
            on:delete={() => {dispatch('delete')}}
        />
    </div>
</div>