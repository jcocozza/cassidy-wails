<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import type { model } from "../../wailsjs/go/models";
    import { NewEquipment } from "$lib/model/equipment";
    import { CreateEquipment } from "../../wailsjs/go/controllers/EquipmentHandler";

    export let usr: model.User
    export let equipment_type_list: model.EquipmentType[];
    let new_equipment: model.Equipment = NewEquipment(usr)
    let is_shown: boolean = false;

    function toggleHidden() {
        is_shown = !is_shown;
    }

    const dispatch = createEventDispatcher();
    async function submit(new_equipment: model.Equipment) {
        await CreateEquipment(new_equipment)
        new_equipment = NewEquipment(usr)
        toggleHidden()
        dispatch('update')
    }
</script>

<div class="NewEquipmentModal">
    <button class="btn btn-primary" on:click={toggleHidden}>Create New</button>

    {#if is_shown}
        <div class="modal">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title"> Create New Equipment </h5>

                        <button type="button" class="close btn btn-primary" aria-label="Close" on:click={toggleHidden}>
                            <span aria-hidden="true">&times;</span>
                        </button>
                    </div>

                    <div class="modal-body">
                        <form on:submit={() => {submit(new_equipment)}}>
                            <div class="form-group">

                                <label for="name">Name:</label>
                                <input class="form-control" id="name" type="text" bind:value={new_equipment.name}>

                                <label for="brand">Brand:</label>
                                <input class="form-control" id="brand" type="text" bind:value={new_equipment.brand}>

                                <label for="model">Model:</label>
                                <input class="form-control" id="model" type="text" bind:value={new_equipment.model}>

                                <label for="size">Size:</label>
                                <input class="form-control" id="size" type="text" bind:value={new_equipment.size}>

                                <label for="cost">Cost:</label>
                                <input class="form-control" id="cost" type="number" step="0.01" bind:value={new_equipment.cost}>

                                <label for="equipment_type">Type:</label>
                                <select class="equipment-type-dropdown" id="equipment_type" bind:value={new_equipment.equipment_type} required>
                                    <option value={null} selected disabled>Select equipment type</option>
                                    {#each equipment_type_list as equipment_type}
                                        <option value={equipment_type}>{equipment_type.name}</option>
                                    {/each}
                                </select>
                                <br>

                                <div class="input-group w-50">
                                    <label for="mileage">Initial mileage: </label>
                                    <input class="form-control" type="number" step="0.01" bind:value={new_equipment.mileage.length}>
                                    <select id="activitylenunits" bind:value={new_equipment.mileage.unit}>
                                        <option value="m">m</option>
                                        <option value="yd">yd</option>
                                        <option value="mi">mi</option>
                                        <option value="km">km</option>
                                    </select>
                                </div>

                                <label for="purchase_date">Purchase Date:</label>
                                <input class="form-control" id="purchase_date" type="date" bind:value={new_equipment.purchase_date}>

                                <label for="notes">Notes:</label>
                                <input class="form-control" id="notes" type="text" bind:value={new_equipment.notes}>

                                <label for="retired">Retire</label>
                                <input class="form-control" id="retired" type="checkbox" bind:checked={new_equipment.is_retired}>
                            </div>

                            <button class="btn btn-primary" type="submit">Submit</button>
                            <button class="btn btn-secondary" on:click={toggleHidden}>Cancel</button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    {/if}
</div>

<style>
    .modal {
        display: block;
    }
</style>