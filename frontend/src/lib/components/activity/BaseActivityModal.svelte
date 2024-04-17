<script lang="ts">
    import { createEventDispatcher } from "svelte";

    import edit from '$lib/static/edit-3-svgrepo-com.svg?raw'
    import { ValidateDuration } from "$lib/model/duration";
    import { FormatDurationSimple } from "$lib/model/date";
    import type { model } from "../../wailsjs/go/models";
    import { EmptyActivityEquipment, HandleActivityEquipmentList } from "$lib/model/equipment";

    import { CreateActivity, UpdateActivity } from '../../wailsjs/go/controllers/ActivityHandler'
    import { DeleteActivityEquipment } from "../../wailsjs/go/controllers/EquipmentHandler";

    export let equipment_choices: model.Equipment[];
    export let activity_type_list: model.ActivityTypeWithSubtypes[];
    export let activity: model.Activity;
    export let is_hovering: boolean = false;
    export let total_num_date_activities: number = 0;
    export let edit_type: string; // either "update" or "create"

    let is_hidden: boolean = false;
    let edited_activity: model.Activity = activity;
    let planned_shown: string = "planned";
    let new_activity_equipment_list: model.ActivityEquipment[] = [];
    let ae_to_delete: model.ActivityEquipment[] = [];
    let subtype_list: model.ActivitySubtype[] = getSubtypeList();
    let subtype_bool_list: boolean[] = []; // this list corresponds to subtype_list and tells you if the activity has that subtype or not
    let activity_type_id: number = activity.activity_type!.id;
    let duration_planned: string = FormatDurationSimple(activity.planned!.duration);
    let duration_completed: string = FormatDurationSimple(activity.completed!.duration);

    let form_error: string = ""

    function getSubtypeList(): model.ActivitySubtype[] {
        if (edited_activity.activity_type?.id == -1) {
            return []
        }
        let l = activity_type_list.filter(item => item.activity_type?.id === edited_activity.activity_type?.id)
        return l[0].subtype_list
    };
    function toggleHidden() {
        is_hidden = !is_hidden
    };

    // check if an equipment is already assigned to an activity equipment object
    function equipmentIsUsed(equipment: model.Equipment, ae_list: model.ActivityEquipment[]): boolean {
        if (ae_list.length == 0) {
            return false
        }

        const found = ae_list.find(item => item.equipment?.id === equipment.id)

        const isInArray = found !== undefined;
        return isInArray
    }
    // add a new activity equipment object
    function addNewAE() {
        new_activity_equipment_list = [...new_activity_equipment_list, EmptyActivityEquipment(activity.uuid, activity.planned!.distance!.unit)]
    }
    // delete activity equipment
    function deleteAE(ae: model.ActivityEquipment) {
        ae_to_delete = [...ae_to_delete, ae]

        // filter out the deleted ae from the lists
        new_activity_equipment_list = new_activity_equipment_list.filter((item) => item.id !== ae.id)
        activity.equipment_list = activity.equipment_list.filter((item) => item.id !== ae.id)
    }
    function handleCheckboxChange(event: Event & { currentTarget: EventTarget & HTMLInputElement }, typeSubtype: model.ActivityTypeSubtype) {
        console.log("CHECK EVENT:", event.currentTarget.checked)
        if (event.currentTarget.checked) {
            // Add the value to the selectedValues array if the checkbox is checked
            edited_activity.type_subtype_list = [...edited_activity.type_subtype_list, typeSubtype];

            let idx = subtype_list.findIndex(subtype => subtype.id === typeSubtype.activity_subtype.id);
        } else {
            // Remove the value from the selectedValues array if the checkbox is unchecked
            edited_activity.type_subtype_list = edited_activity.type_subtype_list.filter(value => value.activity_subtype.id !== typeSubtype.activity_subtype.id);
        }
    };
    function setActivityType(id: number) {
        let l = activity_type_list.filter(item => item.activity_type?.id == id);
        edited_activity.activity_type = l[0].activity_type;
        subtype_list = getSubtypeList();
    };
    function subtypeIsInActivity(subType: model.ActivitySubtype): boolean {
        return edited_activity.type_subtype_list.some(sub_type => sub_type.activity_subtype.name === subType.name);
    }


    const dispatch = createEventDispatcher();
    function handleSubmit() {
        form_error = ""
        let planned_duration = ValidateDuration(duration_planned)
        let completed_duration = ValidateDuration(duration_completed)

        /*if (typeof planned_duration === "string" || typeof completed_duration === "string") {
            throw new Error("duration validation failure!")
        }*/

        if (typeof planned_duration === "string") {
            planned_shown = "planned"
            form_error = "planned duration is malformed"
            return
        } else if (typeof completed_duration === "string") {
            planned_shown = "completed"
            form_error = "completed duration is malformed"
            return
        }
        is_hidden = false;
        edited_activity.planned!.duration = planned_duration;
        edited_activity.completed!.duration = completed_duration;

        activity = edited_activity
        if (edit_type == "create") {
            (async () => {
                await CreateActivity(activity).then((data) => {
                    if (data !== null) {
                        activity = data
                    }
                })

                activity.equipment_list.forEach((ae) => {
                    ae.activity_uuid = activity.uuid
                })
                new_activity_equipment_list.forEach((ae) => {
                    ae.activity_uuid = activity.uuid
                })

                await HandleActivityEquipmentList(activity.equipment_list)
                await HandleActivityEquipmentList(new_activity_equipment_list)
                ae_to_delete.forEach((ae) => {
                    DeleteActivityEquipment(ae.id)
                })

                console.log("BaseActivityModal::: new activity:", activity)
                dispatch("new")
            })();
        } else if (edit_type == "update") {
            (async () => {
                await UpdateActivity(activity)

                await HandleActivityEquipmentList(activity.equipment_list)
                await HandleActivityEquipmentList(new_activity_equipment_list)

                ae_to_delete.forEach((ae) => {
                    DeleteActivityEquipment(ae.id)
                })

                console.log("BaseActivityModal::: update activity:", activity)
                dispatch("update")
            })();
        }
    }

    $: { // recalculate subtype_bool_list whenever subtype_list changes
        subtype_bool_list = subtype_list.map(subType => subtypeIsInActivity(subType));
    }
</script>

<div class="edit-activity-modal">
    {#if is_hovering}
        {#if edit_type == "create"}
            <button class="btn btn-primary btn-sm" type="button" on:click={toggleHidden}>&#8853;</button>
        {:else if edit_type == "update"}
            <button class="btn btn-primary btn-sm" type="button" on:click={toggleHidden}>{@html edit}</button>
        {/if}
    {/if}
    {#if is_hidden}
        <div class="modal" id="activityModal" tabindex="-1" role="dialog" aria-labelledby="activityModalLabel" aria-hidden={is_hidden}>
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title" id="exampleModalLabel">Edit {edited_activity.name}</h5>
                        <button type="button" class="close btn btn-primary" aria-label="Close" on:click={toggleHidden}>
                            <span aria-hidden="true">&times;</span>
                        </button>
                        </div>
                        <div class="modal-body">
                            <form on:submit={handleSubmit}>
                                <div class="row">
                                    <div class="form-group">
                                        <label for="newActivityType">Activity Type:</label>
                                        <select class="form-select" bind:value={activity_type_id} on:change={() => {setActivityType(activity_type_id)}} required>
                                            <option disabled selected>Select an activity type...</option>
                                            {#each activity_type_list as actType}
                                                <option value={actType.activity_type?.id}>{actType.activity_type?.name}</option>
                                            {/each}
                                        </select>
                                        {#if edited_activity.activity_type?.id !== -1}
                                            <div class="row">
                                                {#each subtype_list as subType, idx}
                                                    <div class="col-md-2">
                                                        {#if subType}
                                                            <div class="form-check">
                                                                <label class="form-check-label" for="subtype-{subType.id}">{subType.name}</label>
                                                                <input class="form-check-input" type="checkbox" id="subtype-{subType.id}" checked={subtype_bool_list[idx]} on:change={(event) => handleCheckboxChange(event, {activity_subtype: subType, activity_type: edited_activity.activity_type, activity_uuid: edited_activity.uuid, id: -1})}>
                                                            </div>
                                                        {/if}
                                                    </div>
                                                {/each}
                                            </div>
                                        {/if}
                                    </div>

                                    <div class="form-group">
                                        <label for="newActivityOrder">Order:</label>
                                        <select class="form-select" bind:value={edited_activity.order}>
                                            {#if total_num_date_activities > 0}
                                                {#each Array(total_num_date_activities).fill(null).map((_, index) => index) as act, index}
                                                    <option value={index+1}>{index + 1}</option>
                                                {/each}
                                                <option value={total_num_date_activities + 1} selected>{total_num_date_activities + 1}</option>
                                            {:else}
                                                <option value={1}>1</option>
                                            {/if}
                                        </select>
                                    </div>
                                </div>

                                <div class="row">
                                    <div class="col">
                                        <div class="form-group">
                                            <label for="name">Name:</label>
                                            <input type="text" class="form-control" id="name" bind:value={edited_activity.name}>
                                        </div>

                                        <div class="form-floating">
                                            <label for="description">Description</label>
                                            <textarea class="form-control tall-input" placeholder="activity description" id="description" bind:value={edited_activity.description}></textarea>
                                        </div>
                                    </div>
                                    <div class="col">
                                        <br>
                                        <div class="form-floating">
                                            <label for="notes">Notes</label>
                                            <textarea class="form-control tall-input" placeholder="Notes on the activity..." id="notes" bind:value={edited_activity.notes}></textarea>
                                        </div>

                                        <strong>Equipment</strong>
                                        <div class="form-group">
                                            {#each activity.equipment_list as ae}
                                                {#if ae.assigned_mileage}
                                                    <div class="input-group w-60">
                                                        <span>{ae.equipment?.name}</span>
                                                        <input class="form-control" type="number" step="0.01" bind:value={ae.assigned_mileage.length}>
                                                        <select id="activitylenunits" bind:value={ae.assigned_mileage.unit}>
                                                            <option value="m">m</option>
                                                            <option value="yd">yd</option>
                                                            <option value="mi">mi</option>
                                                            <option value="km">km</option>
                                                        </select>

                                                        <button type="button" class="btn btn-primary" on:click={() => {deleteAE(ae)}}>Delete</button>
                                                    </div>
                                                {/if}
                                            {/each}
                                        </div>

                                        {#each new_activity_equipment_list as nae}
                                            {#if nae.assigned_mileage}
                                                <div class="input-group">
                                                    <select bind:value={nae.equipment}>
                                                        <option value={null}>Select equipment...</option>
                                                        {#each equipment_choices as equipment}
                                                            {#if equipmentIsUsed(equipment, [...new_activity_equipment_list, ...activity.equipment_list])}
                                                                    <option value={equipment} disabled>{equipment.name}</option>
                                                            {:else}
                                                                <option value={equipment}>{equipment.name}</option>
                                                            {/if}
                                                        {/each}
                                                    </select>

                                                    <input class="form-control" type="number" step="0.01" bind:value={nae.assigned_mileage.length}>
                                                    <select id="activitylenunits" bind:value={nae.assigned_mileage.unit}>
                                                        <option value="m">m</option>
                                                        <option value="yd">yd</option>
                                                        <option value="mi">mi</option>
                                                        <option value="km">km</option>
                                                    </select>

                                                    <button type="button" class="btn btn-danger" on:click={() => {deleteAE(nae)}}>Delete</button>
                                                </div>
                                            {/if}
                                        {/each}

                                        <button type="button" class="btn btn-primary" on:click={addNewAE}>Add Activity Equipment</button>
                                    </div>
                                </div>
                                <br><br>
                                <div class="row">
                                    <ul class="nav nav-tabs">
                                        <!-- svelte-ignore a11y-click-events-have-key-events -->
                                        <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
                                        <li class="nav-item nav-link" class:active={planned_shown === "planned"} on:click={() => planned_shown = "planned"}>Planned</li>
                                        <!-- svelte-ignore a11y-click-events-have-key-events -->
                                        <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
                                        <li class="nav-item nav-link" class:active={planned_shown === "completed"} on:click={() => {planned_shown = "completed"}}>Completed</li>
                                    </ul>

                                    <div class="container">
                                        <!-- Planned Stuff -->
                                        {#if planned_shown === "planned"}
                                            {#if edited_activity.planned}
                                                <div class="row">
                                                    <div class="col">
                                                        <div class="input-group w-20">
                                                            <label for="editedPlannedDistance">Distance:</label>
                                                            <input type="number" step="0.01" class="form-control" id="editedPlannedDistance" bind:value={edited_activity.planned.distance.length}>
                                                            <select id="activitylenunits" bind:value={edited_activity.planned.distance.unit}>
                                                                <option value="m">m</option>
                                                                <option value="yd">yd</option>
                                                                <option value="mi">mi</option>
                                                                <option value="km">km</option>
                                                            </select>
                                                        </div>
                                                    </div>
                                                    <div class="col">
                                                        <div class="input-group w-20">
                                                            <label for="editedPlannedVertical">Vertical:</label>
                                                            <input type="number" step="0.01" class="form-control" id="editedPlannedVertical" bind:value={edited_activity.planned.vertical.length}>
                                                            <select bind:value={edited_activity.planned.vertical.unit}>
                                                                <option value="ft">ft</option>
                                                                <option value="m">m</option>
                                                            </select>
                                                        </div>
                                                    </div>
                                                    <div class="col">
                                                        <div class="input-group w-20">
                                                            <label for="editedPlannedDuration">Duration:</label>
                                                            <input type="text" class="form-control" id="editedPlannedDuration" bind:value={duration_planned} placeholder="hh:mm:ss" pattern="[0-9][0-9]:[0-5][0-9]:[0-5][0-9]">
                                                        </div>
                                                    </div>
                                                </div>
                                            {/if}
                                        {:else if planned_shown === "completed"}
                                            <!-- Completed Stuff -->
                                            {#if edited_activity.completed}
                                                <div class="row">
                                                    <div class="col">
                                                        <div class="input-group w-20">
                                                            <label for="editedCompletedDistance">Distance:</label>
                                                            <input type="number" step="0.01" class="form-control" id="editedCompletedDistance" bind:value={edited_activity.completed.distance.length}>
                                                            <select id="activitylenunits" bind:value={edited_activity.completed.distance.unit}>
                                                                <option value="m">m</option>
                                                                <option value="yd">yd</option>
                                                                <option value="mi">mi</option>
                                                                <option value="km">km</option>
                                                            </select>
                                                        </div>
                                                    </div>
                                                    <div class="col">
                                                        <div class="input-group w-20">
                                                            <label for="editedCompletedVertical">Vertical:</label>
                                                            <input type="number" step="0.01" class="form-control" id="editedCompletedVertical" bind:value={edited_activity.completed.vertical.length}>
                                                            <select bind:value={edited_activity.completed.vertical.unit}>
                                                                <option value="ft">ft</option>
                                                                <option value="m">m</option>
                                                            </select>
                                                        </div>
                                                    </div>
                                                    <div class="col">
                                                        <div class="input-group w-20">
                                                            <label for="editedCompletedDuration">Duration:</label>
                                                            <input type=text class="form-control" id="editedCompletedDuration" bind:value={duration_completed} placeholder="hh:mm:ss" pattern="[0-9][0-9]:[0-5][0-9]:[0-5][0-9]">
                                                        </div>
                                                    </div>
                                                </div>
                                            {/if}
                                        {/if}
                                    </div>
                                </div>

                                <div class="row">
                                    <label for="race">Race:</label>
                                    <input id="race" type="checkbox" bind:checked={edited_activity.is_race}>

                                    <label for="strides">Strides</label>
                                    <input id="strides" type="number" bind:value={edited_activity.num_strides}>
                                </div>

                                <div class="modal-footer">

                                    {#if form_error}
                                        <p>{form_error}</p>
                                    {/if}

                                    <button type="button" class="btn btn-secondary" on:click={toggleHidden}>Close</button>
                                    <input type="submit" class="btn btn-primary">
                                </div>
                            </form>
                        </div>
                </div>
            </div>
        </div>
    {/if}
</div>

<style>
    /* Ensure we can see the modal*/
    .modal {
        display: block;
    }
    /* allow the modal to grow */
    .modal-dialog {
        width: 80%;
        max-width: 800px;
    }
    .tall-input {
        min-height: 200px;
    }

</style>