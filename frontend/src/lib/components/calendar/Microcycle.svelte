<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import { ConvertDuration, GetWeekday, ParseDateYYYYMMDD, } from '$lib/model/date';
    import { model } from '../../wailsjs/go/models';
    import ActivityList from '../activity/ActivityList.svelte';

    import { overrideItemIdKeyNameBeforeInitialisingDndZones } from "svelte-dnd-action";
    import NewActivityModal from '../activity/NewActivityModal.svelte';
    overrideItemIdKeyNameBeforeInitialisingDndZones("uuid");

    export let activity_type_list: model.ActivityTypeWithSubtypes[];
    export let microcycle: model.Microcycle;
    export let equipment_choices: model.Equipment[];
    export let usr: model.User;

    let display_completion: boolean = true
    let summary_col_is_visible: boolean = true

    const dispatch = createEventDispatcher();
    function dispatchUpdate() {
        console.log("Microcycle::: change")
        dispatch("update")
    }

    function gotoToday() {
        dispatch("today")
    }

</script>

{#if microcycle && microcycle.cycle_activities && microcycle.summary && microcycle.summary.totals}
    <div class="calendar container">

        <div class="row">

            <div class="col flex-grow-1">
                {usr.username}
            </div>

            <div class="col">
                <button class="btn btn-primary btn-sm" on:click={gotoToday}>Today</button>
            </div>
            <div class="col small-text d-flex justify-content-end">
                <div class="btn-group-vertical">
                    <div class="btn-group">
                        <select bind:value={display_completion} id="display_completion">
                            <option value={true}>Completion</option>
                            <option value={false}>None</option>
                        </select>
                        <button class="btn btn-primary btn-sm" on:click={() => {summary_col_is_visible = !summary_col_is_visible}}>
                            {#if summary_col_is_visible}
                            &gt; Totals
                            {:else}
                            &lt; Totals
                            {/if}
                        </button>
                    </div>
                </div>
            </div>

        </div>

        <div class="overflow-scroll" style="max-height:50vh">
            <table class="table">
                <thead>
                    {#each microcycle.cycle_activities as cycle}
                        <th>
                            <div class="row">
                                <div class="col">
                                    {#if microcycle.cycle_activities.length == 7}
                                        {GetWeekday(cycle.date)}
                                    {:else}
                                        {ParseDateYYYYMMDD(cycle.date)}
                                    {/if}
                                </div>
                                <div class="col">
                                    <NewActivityModal
                                        bind:usr={usr}
                                        bind:equipment_choices={equipment_choices}
                                        bind:date={cycle.date}
                                        bind:activity_list={cycle.activity_list}
                                        bind:activity_type_list={activity_type_list}
                                        is_hovering={true}
                                        on:new={dispatchUpdate}
                                    />
                                </div>
                            </div>
                        </th>
                    {/each}

                    {#if summary_col_is_visible}
                        <th>
                            Summary
                            <table class="summary">
                                <tr>
                                    <td>Completed</td>
                                    <td>Planned</td>
                                </tr>
                            </table>
                        </th>
                    {/if}
                </thead>
                <tbody>
                    <tr>
                        {#each microcycle.cycle_activities as activity_list}
                            <td>
                                <ActivityList
                                    bind:activity_list={activity_list}
                                    bind:date={activity_list.date}
                                    bind:activity_type_list={activity_type_list}
                                    bind:display_completion={display_completion}
                                    bind:equipment_choices={equipment_choices}
                                    on:change={dispatchUpdate}
                                />
                            </td>
                        {/each}
                        {#if summary_col_is_visible}
                            <td style="padding: 0;">
                                <ul class="list-group list-group-flush striped-list summary">
                                    {#each microcycle.summary.totals_by_activity_type as act_type_total}
                                        <!-- This is a hacky way to check for which activities types are in the current microcyle. This way we can distinguish between totals and planning better -->
                                        {#if microcycle.summary.totals_by_activity_type_and_date.some(atd => atd.activity_type?.id === act_type_total.activity_type?.id)}
                                            <li class="list-group-item">
                                                <div class="row">
                                                    <div class="col">
                                                        <!-- The type -->
                                                        <div class="col-md-auto d-flex align-items-center">
                                                            <strong>{act_type_total.activity_type?.name}</strong>
                                                        </div>
                                                        <div style="margin-left: -12px;">
                                                            <table>
                                                                <tbody>
                                                                    <!-- Distance -->
                                                                    <tr>
                                                                        <td>{act_type_total.total_completed_distance?.length} {act_type_total.total_completed_distance?.unit}</td>
                                                                        <td class="text-secondary">{act_type_total.total_planned_distance?.length} {act_type_total.total_planned_distance?.unit}</td>
                                                                    </tr>
                                                                    <!-- Duration -->
                                                                    <tr>
                                                                        <td>{ConvertDuration(act_type_total.total_completed_duration)}</td>
                                                                        <td class="text-secondary">{ConvertDuration(act_type_total.total_planned_duration)}</td>
                                                                    </tr>
                                                                    <!-- Pace -->
                                                                    <tr>
                                                                        <td>{act_type_total.completed_pace}</td>
                                                                        <td class="text-secondary">{act_type_total.planned_pace}</td>
                                                                    </tr>
                                                                    <!-- Vertical -->
                                                                    <tr>
                                                                        <td>{act_type_total.total_completed_vertical?.length} {act_type_total.total_completed_vertical?.unit}</td>
                                                                        <td class="text-secondary">{act_type_total.total_planned_vertical?.length} {act_type_total.total_planned_vertical?.unit}</td>
                                                                    </tr>
                                                                </tbody>
                                                            </table>
                                                        </div>
                                                    </div>
                                                </div>
                                            </li>
                                        {/if}
                                    {/each}
                                    <br>
                                    <li class="list-group-item">
                                        <div class="row">
                                            <div class="col">
                                                <!-- cycle totals -->
                                                <div class="col-md-auto d-flex align-items-center">
                                                    <strong>Totals</strong>
                                                </div>
                                                <table>
                                                    <tbody>
                                                        <!-- Distance -->
                                                        <tr>
                                                            <td>{microcycle.summary.totals.total_completed_distance?.length} {microcycle.summary.totals.total_completed_distance?.unit}</td>
                                                            <td class="text-secondary">{microcycle.summary.totals.total_planned_distance?.length} {microcycle.summary.totals.total_planned_distance?.unit}</td>
                                                        </tr>
                                                        <!-- Duration -->
                                                        <tr>
                                                            <td>{ConvertDuration(microcycle.summary.totals.total_completed_duration)}</td>
                                                            <td class="text-secondary">{ConvertDuration(microcycle.summary.totals.total_planned_duration)}</td>
                                                        </tr>
                                                        <!-- Vertical -->
                                                        <tr>
                                                            <td>{microcycle.summary.totals.total_completed_vertical?.length} {microcycle.summary.totals.total_completed_vertical?.unit}</td>
                                                            <td class="text-secondary">{microcycle.summary.totals.total_planned_vertical?.length} {microcycle.summary.totals.total_planned_vertical?.unit}</td>
                                                        </tr>
                                                    </tbody>
                                                </table>
                                            </div>
                                        </div>
                                    </li>
                                </ul>
                            </td>
                        {/if}
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
{/if}

<style>
    .calendar th {
      border: 1px solid lightgray;
      text-align: center;
      /*position: sticky;*/
      /*top: 0;*/
      /*z-index: 2;*/
    }
    .calendar td {
      border: 1px solid lightgray;
      text-align: center;
      vertical-align: top;
    }
    .calendar table {
        table-layout: fixed;
        width: 100%;
    }
    .small-text {
        font-size: 10px;
        line-height: 1.2em;
        text-align: left;
    }
    .summary {
        font-size: 10px;
        line-height: 1.2em;
        text-align: left;
    }
    .summary table {
        table-layout: auto;
    }
    .summary td {
        border: none;
        white-space: nowrap;
        border-top: 1px solid lightgray;
    }
</style>