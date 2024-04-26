<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import type { model } from "../../wailsjs/go/models";
    import { DeleteEquipment, UpdateEquipment } from "../../wailsjs/go/controllers/EquipmentHandler";


    export let equipment: model.Equipment;
    export let is_hovering: boolean = false;

    let is_shown: boolean = false;
    let is_editing: boolean = false;
    function toggleHidden() {
        is_shown = !is_shown;
    }

    function toggleEdit() {
        is_editing = !is_editing;
    }

    const dispatch = createEventDispatcher();
    async function deleteEquipment(equipment: model.Equipment) {
        console.log("EquipmentModal::: deleting:", equipment)
        await DeleteEquipment(equipment.id)
        dispatch("delete")
    }
</script>

<div class="EquipmentModal">
    {#if is_hovering}
        <button class="btn btn-primary" on:click={toggleHidden}>View/Edit</button>
    {/if}

    {#if is_shown}
        {#if equipment && equipment.mileage && equipment.equipment_type}
            <div class="modal">
                <div class="modal-dialog">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h5 class="modal-title"> Equipment: {equipment.name} </h5>

                            <div class="form-check form-switch">
                                <input class="form-check-input" type="checkbox" role="switch" id="flexSwitchCheckChecked" bind:checked={is_editing}>
                                <label class="form-check-label" for="flexSwitchCheckChecked">Toggle Editing</label>
                            </div>
                            <button type="button" class="close btn btn-primary" aria-label="Close" on:click={toggleHidden}>
                                <span aria-hidden="true">&times;</span>
                            </button>
                        </div>
                        <div class="modal-body">
                            <form on:submit={() => {UpdateEquipment(equipment.id, equipment); toggleEdit();}}>
                                <div class="form-group">
                                    <label for="equipment_brand">Brand:</label>
                                    <input class="form-control" id="equipment_brand" type="text" bind:value={equipment.brand} aria-label="Disabled input example" disabled={!is_editing} readonly={!is_editing}>

                                    <label for="equipment_model">Model:</label>
                                    <input class="form-control" id="equipment_model" type="text" bind:value={equipment.model} aria-label="Disabled input example" disabled={!is_editing} readonly={!is_editing}>

                                    <label for="equipment_size">Size:</label>
                                    <input class="form-control" id="equipment_size" bind:value={equipment.size} aria-label="Disabled input example" disabled={!is_editing} readonly={!is_editing}>

                                    <label for="equipment_cost">Cost:</label>
                                    <input class="form-control" id="equipment_cost" type="number" step="0.01" bind:value={equipment.cost} aria-label="Disabled input example" disabled={!is_editing} readonly={!is_editing}>

                                    <label for="equipment_type_name">Type:</label>
                                    <input class="form-control" id="equipment_type_name" type="text" bind:value={equipment.equipment_type.name} aria-label="Disabled input example" disabled readonly>

                                    <div class="input-group w-50">
                                        <label for="mileage">Mileage: </label>
                                        <input class="form-control" type="number" step="0.01" bind:value={equipment.mileage.length} disabled={!is_editing} readonly={!is_editing}>
                                        <select id="activitylenunits" bind:value={equipment.mileage.unit} disabled={!is_editing}>
                                            <option value="m">m</option>
                                            <option value="yd">yd</option>
                                            <option value="mi">mi</option>
                                            <option value="km">km</option>
                                        </select>
                                    </div>

                                    <label for="equipment_purchase_date">Date Purchased:</label>
                                    <input class="form-control" id="equipment_purchase_date" type="date" bind:value={equipment.purchase_date} aria-label="Disabled input example" disabled={!is_editing} readonly={!is_editing}>

                                    <label for="equipment_notes">Notes:</label>
                                    <input class="form-control" id="equipment_notes" type="text" bind:value={equipment.notes} aria-label="Disabled input example" disabled={!is_editing} readonly={!is_editing}>

                                    <!--
                                    <div class="form-check">
                                        <input class="form-check-input" type="checkbox" id="equipment-{equipment.id}-is-retired" bind:checked={equipment.is_retired}>
                                        <label class="form-check-label text-danger" for="equipment-{equipment.id}-is-retired">Retire</label>
                                    </div>
                                    -->
                                </div>

                                {#if is_editing}
                                    <input type="submit">

                                    <button class="btn btn-danger" on:click={() => {deleteEquipment(equipment); toggleEdit(); toggleHidden();}}>Delete</button>
                                {/if}
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        {/if}
    {/if}
</div>

<style>
    .modal {
        display: block;
    }
</style>