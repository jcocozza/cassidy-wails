<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import { default_activity_interface, type ActivityInterface, NewActivity } from "../../../model/activity";
    import type { ActivityType, ActivityTypeSubtype, ActivityTypeWithSubtypes } from "../../../model/activity_type";
    import type { User } from "../../../model/user";
    import { get } from "svelte/store";
    import { page } from "$app/stores";

    export let activity_type_list: ActivityTypeWithSubtypes[];
    export let date: string;
    export let activity_list: ActivityInterface[];

    export let is_hovering: boolean = false;

    let usr: User = get(page).data.session.user

    let form_hidden: boolean = true;

    //let is_hidden: boolean = false;
    let planned_shown: string = "planned";
    let equipment_shown: boolean = false;
    let input_new_activity = default_activity_interface(date, activity_list.length + 1, usr)
    let selected_values: ActivityTypeSubtype[] = [];
    let current_activity_type: ActivityType;
    let _new_activity: ActivityInterface;
    let duration: string = "";
    let duration_completed: string = "";

    function toggleHidden() {
        form_hidden = !form_hidden;
    };

    function filterActivityTypeList(activity_type_id: number): ActivityTypeWithSubtypes {
        let l = activity_type_list.filter(item => item.activity_type.id === input_new_activity.activity_type.id)
        return l[0]
    }

    function setActivityType() {
        selected_values = [];
        let l = activity_type_list.filter(item => item.activity_type.id === input_new_activity.activity_type.id);
        current_activity_type = l[0].activity_type;
        if (l.length > 0) {
            current_activity_type = l[0].activity_type;
            console.log("CHANGING TYPE:", current_activity_type);

            input_new_activity.activity_type.id = current_activity_type.id;
            input_new_activity.activity_type.name = current_activity_type.name;
        } else {
            console.error("No matching activity type found");
        }
    };

    function handleCheckboxChange(event: Event & { currentTarget: EventTarget & HTMLInputElement }, typeSubtype: ActivityTypeSubtype) {
        console.log("CHECK EVENT:", event.currentTarget.checked)
        if (event.currentTarget.checked) {
            // Add the value to the selectedValues array if the checkbox is checked
            selected_values = [...selected_values, typeSubtype];
            input_new_activity.type_subtype_list = selected_values;
        } else {
            // Remove the value from the selectedValues array if the checkbox is unchecked
            console.log("SELECTED VALUES:", selected_values)
            console.log("TYPESUBTYPE:",typeSubtype)
            selected_values = selected_values.filter(value => value.activity_subtype.id != typeSubtype.activity_subtype.id);
            console.log("remvoing subtypes:", selected_values)
            input_new_activity.type_subtype_list = selected_values;
        }
    };

    const dispatch = createEventDispatcher();

    async function addActivity() {
        await NewActivity(input_new_activity, date, usr)
        .then((data) => {
            if (data === null) {
                return
            } else {
                dispatch("new")
                _new_activity = data
                //let act = new Activity(_new_activity)
                activity_list = [...activity_list, _new_activity]
            }
        });
    };

    function validateDuration(dur: string): number {
        const durationPattern = /^[0-9]{1,2}:[0-5][0-9]:[0-5][0-9]$/;
        if (durationPattern.test(dur)) {
            alert("Duration is valid!");
            let [hours, minutes, seconds] = dur.split(':').map(Number);

            let tot_seconds = (hours * 60 * 60) + (minutes * 60) + seconds
            return tot_seconds
        } else {
            alert("Invalid duration format. Please use HH:MM:SS");
            return -1
        }
    };

    function validateForm() {
        if (input_new_activity.activity_type.id == -1) {
            alert("Must choose an activity type.");
            return
        }
        if (duration !== ""){
            let durationSeconds = validateDuration(duration)
            if (durationSeconds == -1) {
                return
            }
            input_new_activity.planned.duration = durationSeconds
        }
        if (duration_completed !== ""){
            let durationSeconds = validateDuration(duration_completed)
            if (durationSeconds == -1) {
                return
            }
            input_new_activity.completed.duration = durationSeconds
        }
        addActivity();
        toggleHidden();
    };
</script>

<div class="NewActivityModal">
    {#if is_hovering}
        <button class="btn btn-primary btn-sm" on:click={toggleHidden}>&#8853;</button>
    {/if}
    {#if !form_hidden}
    <div class="modal" id="newActivityModal" tabindex="-1" role="dialog" aria-labelledby="newActivityModalLabel" aria-hidden={form_hidden}>
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="exampleModalLabel">New Activity</h5>
                    <button type="button" class="close btn btn-primary" aria-label="Close" on:click={toggleHidden}>
                        <span aria-hidden="true">&times;</span>
                    </button>
                </div>
                <div class="modal-body">
                    <!-- NEW ACTIVITY FORM-->
                    <form>
                        <div class="form-group">
                            <label for="newActivityType">Activity Type:</label>
                            <select class="form-select" bind:value={input_new_activity.activity_type.id} on:change={setActivityType} required>
                                <option disabled selected>Select an activity type...</option>
                                {#each activity_type_list as actType, index}
                                    <option value={actType.activity_type.id}>{actType.activity_type.name}</option>
                                {/each}
                            </select>
                            {#if input_new_activity.activity_type.id !== -1 && current_activity_type}
                                {#each filterActivityTypeList(input_new_activity.activity_type.id).subtype_list as subType, index}
                                    {#if subType}
                                        <div class="form-check">
                                            <label class="form-check-label" for="subtype-{subType.id}">{subType.name}</label>
                                            <input class="form-check-input" type="checkbox" value={subType} id="subtype-{subType.id}" on:change={(event) => handleCheckboxChange(event, {activity_subtype: subType, activity_type: current_activity_type, activity_uuid: "", id: -1})}>
                                        </div>
                                    {/if}
                                {/each}
                            {/if}
                        </div>

                        <div class="form-group">
                            <label for="newActivityOrder">Order:</label>
                            <select class="form-select" bind:value={input_new_activity.order}>
                                {#if activity_list !== null && activity_list.length > 0}
                                    {#each activity_list as act, index}
                                        <option value={index+1}>{index + 1}</option>
                                    {/each}
                                    <option value={activity_list.length + 1} selected>{activity_list.length + 1}</option>
                                {:else}
                                    <option value={1}>1</option>
                                {/if}
                            </select>
                        </div>

                        <div class="form-group">
                            <label for="newActivityDate">Date:</label>
                            <input type="date" class="form-control" id="newActivityDate" bind:value={input_new_activity.date}>
                        </div>

                        <div class="form-group">
                            <label for="newActivityName">Name:</label>
                            <input type="text" class="form-control" id="newActivityName" bind:value={input_new_activity.name}>
                        </div>

                        <div class="form-group">
                            <label for="newActivityDescription">Description:</label>
                            <textarea class="form-control" id="newActivityDescription" bind:value={input_new_activity.description}></textarea>
                        </div>

                        <div class="form-group">
                            <label for="newActivityNotes">Notes:</label>
                            <textarea class="form-control" id="newActivityNotes" bind:value={input_new_activity.notes}></textarea>
                        </div>

                        <div class="btn-group" role="group">
                            <button class="btn btn-outline-primary" on:click={() => {planned_shown = "planned"}}   class:active={planned_shown === "planned"}  >Planned</button>
                            <button class="btn btn-outline-primary" on:click={() => {planned_shown = "completed"}} class:active={planned_shown === "completed"}>Completed</button>
                          </div>


                        {#if planned_shown === "planned"}
                            <!-- Planned Stuff -->
                            <div class="form-group">
                                <div class="row">
                                    <div class="col">
                                        <div class="row">
                                            <div class="input-group">
                                                <label for="editedPlannedDistance">Distance:</label>
                                                <input type="number" step="0.01" class="form-control" id="editedPlannedDistance" bind:value={input_new_activity.planned.distance.length}>
                                                <select id="activitylenunits" bind:value={input_new_activity.planned.distance.unit}>
                                                    <option value="m">m</option>
                                                    <option value="yd">yd</option>
                                                    <option value="mi">mi</option>
                                                    <option value="km">km</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="row">
                                            <div class="input-group">
                                                <label for="editedPlannedVertical">Vertical:</label>
                                                <input type="number" step="0.01" class="form-control" id="editedPlannedVertical" bind:value={input_new_activity.planned.vertical.length}>

                                                <select bind:value={input_new_activity.planned.vertical.unit}>
                                                    <option value="ft">ft</option>
                                                    <option value="m">m</option>
                                                </select>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="col">
                                        <label for="editedPlannedDuration">Duration:</label>
                                        <input type=text class = "form-control" id="editedPlannedDuration" bind:value={duration} placeholder="hh:mm:ss">
                                    </div>
                                </div>
                            </div>
                        {:else if planned_shown === "completed"}
                            <!-- Completed Stuff -->
                            <div class="form-group">
                                <div class="row">
                                    <div class="col">
                                        <div class="row">
                                            <div class="input-group">
                                                <label for="editedCompletedDistance">Distance:</label>
                                                <input type="number" step="0.01" class="form-control" id="editedCompletedDistance" bind:value={input_new_activity.completed.distance.length} placeholder="distance">
                                                <select id="activitylenunits" bind:value={input_new_activity.completed.distance.unit}>
                                                    <option value="m">m</option>
                                                    <option value="yd">yd</option>
                                                    <option value="mi">mi</option>
                                                    <option value="km">km</option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="row">
                                            <div class="input-group">
                                                <label for="editedCompletedVertical">Vertical:</label>
                                                <input type="number" step="0.01" class="form-control" id="editedCompletedVertical" bind:value={input_new_activity.completed.vertical.length}>

                                                <select bind:value={input_new_activity.completed.vertical.unit}>
                                                    <option value="ft">ft</option>
                                                    <option value="m">m</option>
                                                </select>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="col">
                                        <label for="editedCompletedDuration">Duration:</label>
                                        <input type=text class = "form-control" id="editedCompletedDuration" bind:value={duration_completed} placeholder="hh:mm:ss">
                                    </div>
                                </div>
                            </div>
                        {/if}

                        <!-- Equipment stuff -->
                        <div class="form-group">
                            <div class="form-check form-switch">
                                <label class="form-check-label" for="equipmentToggle">Equipment</label>
                                <input class="form-check-input" type="checkbox" id="equipmentToggle" bind:checked={equipment_shown}>
                            </div>

                            {#if equipment_shown}
                                <p>Sorry! Not implemented yet. For now, create an activity and add equipment then.</p>
                            {/if}
                        </div>
                    </form>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" on:click={toggleHidden}>Close</button>
                    <button type="button" class="btn btn-primary" on:click={validateForm}>Save changes</button>
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