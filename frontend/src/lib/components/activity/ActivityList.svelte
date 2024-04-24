<script lang="ts">
    import ActivityComp from "$lib/components/activity/Activity.svelte";
    import { createEventDispatcher } from "svelte";
    import { dndzone, type DndEvent } from "svelte-dnd-action";
    import type { model } from "../../wailsjs/go/models";
    import { UpdateActivity } from "../../wailsjs/go/controllers/ActivityHandler";

    export let date: Date;
    export let activity_type_list: model.ActivityTypeWithSubtypes[] = [];
    export let activity_list: model.ActivityList;
    export let display_completion: boolean;
    export let equipment_choices: model.Equipment[] = [];

    // Handle drag-and-drop consideration
    // This function is called when an item is dragged over another
    function handleDndConsider(e: CustomEvent<DndEvent<model.Activity>>) {
        const sortedActivities: model.Activity[] = e.detail.items;
            sortedActivities.forEach((activity, index) => { // Update the order for each activity based on its position in the sorted array
                activity.order = index + 1; // Assuming order start from 1
        });
        activity_list.activity_list = sortedActivities; // Update the activities array with the sorted activities
    }

    // Handle drag-and-drop finalization
    // This function is called when an item is dropped
    async function handleDndFinalize(e: CustomEvent<DndEvent<model.Activity>>) {
        const sortedActivities: model.Activity[] = e.detail.items;
            sortedActivities.forEach((activity, index) => { // Update the order for each activity based on its position in the sorted array
                    activity.order = index + 1; // Assuming order start from 1
                    activity.date = date;
        });

        activity_list.activity_list = sortedActivities; // Update the activities array with the sorted activities

        await Promise.all(activity_list.activity_list.map(async (act) => {
            await UpdateActivity(act)
        }));

        dispatchChange()
    };

    const dispatch = createEventDispatcher();
    function dispatchChange() {
        dispatch("change")
        console.log("CHANGE TO ACTIVITY LIST DISPATCHED")
    }

    let is_hovering = false;
    function handleMouseOver() {
        if (!is_hovering) {
            is_hovering = true;
        }
    }
    function handleMouseLeave() {
        if (is_hovering) {
        is_hovering = false;
    }
    }

</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="activity-list" on:mouseenter={handleMouseOver} on:mouseleave={handleMouseLeave}>
    {#await activity_list.activity_list}
        loading...
    {:then}
        {#if activity_list.activity_list}
            <div class="min-size-drop-zone" use:dndzone="{{ items: activity_list.activity_list }}" on:consider="{handleDndConsider}"  on:finalize="{handleDndFinalize}">
                {#each activity_list.activity_list as act (act.uuid)}
                    <ActivityComp
                        bind:activity={ act }
                        bind:equipment_choices={equipment_choices}
                        bind:display_completion={display_completion}
                        bind:activity_type_list={activity_type_list}
                        bind:total_num_date_activities={activity_list.activity_list.length}
                        on:delete={dispatchChange}
                        on:update={dispatchChange}
                    />
                {/each}
            </div>
        {/if}
    {/await}
</div>


<style>
    /* Add any styling for the activity list container here */
    .activity-list {
      /* Example styles */
      margin-top: 20px;
      min-height: 100px;
      position: relative;
    }
    .min-size-drop-zone {
        min-height: 100px; /* Set the minimum height */
        /* min-width: 200px; Set the minimum width */
        /* border: 2px dashed #ccc; /* Optional: Add a dashed border for better visualization */
        padding: 10px; /* Optional: Add padding to the drop zone */
    }
</style>