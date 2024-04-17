<script lang="ts">
    import EditActivityModal from "./EditActivityModal.svelte";

    import trashbin from '$lib/static/trash-bin-trash-svgrepo-com.svg?raw'
    import { ConvertDuration } from "../../model/date";
    import { createEventDispatcher } from "svelte";

    import { DeleteActivity } from '../../wailsjs/go/controllers/ActivityHandler'
    import type { model } from "../../wailsjs/go/models";

    export let activity: model.Activity;
    export let activity_type_list: model.ActivityTypeWithSubtypes[]
    export let equipment_choices: model.Equipment[];
    export let display_completion: boolean = false;
    export let total_num_date_activities: number;

    let visibility = "hidden"

    let is_hovering = false;
    function handleMouseOver() {
        is_hovering = true;
        visibility = "visible"
    }
    function handleMouseLeave() {
        is_hovering = false;
        visibility = "hidden"
    }

    async function deleteActivity(act: model.Activity) {
        console.log("Activity::: deleting:", act)
        await DeleteActivity(act.uuid)
        dispatch("delete")
    }

    const dispatch = createEventDispatcher();
</script>

{#if activity && activity.activity_type && activity.completed && activity.planned}
    <div class="activity-card" style="position: relative;">
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <!-- svelte-ignore a11y-mouse-events-have-key-events -->
        <div class="card" on:mouseover={handleMouseOver} on:mouseleave={handleMouseLeave}>
            {#if display_completion}
                <div style="position: absolute; top: 0; left: 0px; width: 5px; height: 100%; background-color: {activity.color}; z-index: 1;border-radius: 2.5px;"></div>
            {/if}
            <div class="card-body" style="margin-left: 5px; margin-right: 5px;">
                <div class="btn-group" style="float: right;">
                    <EditActivityModal
                        bind:activity={activity}
                        bind:equipment_choices={equipment_choices}
                        bind:is_hovering={is_hovering}
                        bind:activity_type_list={activity_type_list}
                        bind:total_num_date_activities={total_num_date_activities}
                        on:update={() => {dispatch("update")}}
                        />
                        <!--{#if is_hovering}-->
                            <button class="btn btn-primary btn-sm" type="button" style:visibility={visibility} on:click={async () => {await deleteActivity(activity)}}>{@html trashbin}</button>
                        <!--{/if}-->
                </div>

                <div>
                    <strong>{activity.activity_type.name}</strong>
                    {#if activity.is_race}
                        🏁
                    {/if}

                    <br>
                    <!-- {#if activity.name} -->
                    <div class="text-container">
                        <span class="text-container"><i>{activity.name}</i></span>
                    </div>
                        <br>
                    <!-- {/if} -->

                    {#if activity.description}
                        <div class="text-container">
                            <span class="text-secondary">{activity.description}</span>
                            <br>
                            <br>
                        </div>
                    {/if}
                </div>
                <!-- If there is completed, just display that. If there is planned, just display that. Otherwise don't display anything. -->
                {#if activity.completed.distance?.length != 0 || activity.completed.duration != 0 || activity.completed.vertical?.length != 0}
                    <div class="row text-nowrap">
                        <div class="col">
                            {activity.completed.distance?.length} {activity.completed.distance?.unit}
                        </div>
                        <div class="col">
                            {activity.completed.vertical?.length} {activity.completed.vertical?.unit}
                        </div>
                    </div>
                    <div class="row">
                        <div class="col">
                            <span>{activity.completed.pace}</span>
                            <br>
                            {ConvertDuration(activity.completed.duration)}
                        </div>
                    </div>
                    <br>
                {:else if activity.planned.distance?.length != 0 || activity.planned.duration != 0 || activity.planned.vertical?.length != 0}
                    <div class="row text-secondary text-nowrap">
                        <div class="col">
                            {activity.planned.distance?.length} {activity.planned.distance?.unit}
                        </div>
                        <div class="col">
                            {activity.planned.vertical?.length} {activity.planned.vertical?.unit}
                        </div>
                    </div>
                    <div class="row text-secondary">
                        <div class="col">
                            <span>{activity.planned.pace}</span>
                            <br>
                            {ConvertDuration(activity.planned.duration)}
                        </div>
                    </div>
                    <br>
                {/if}

                {#if activity.type_subtype_list !== null }
                    <div class="container justify-content-md-center">
                        {#each activity.type_subtype_list as subtype}
                            <span style="margin-right: 1px; margin-bottom: 1px;" class="badge rounded-pill text-bg-secondary tag">{subtype.activity_subtype.name}</span>
                        {/each}
                    </div>
                {/if}
            </div>
        </div>
    </div>
{/if}


<style>
    .card-body {
        padding: 0; /* Remove padding from the card body */
        line-height: 1.2em;
        text-overflow: ellipsis;
        font-size: 11px;
        text-align: left;
    }
    .text-container {
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }
</style>
