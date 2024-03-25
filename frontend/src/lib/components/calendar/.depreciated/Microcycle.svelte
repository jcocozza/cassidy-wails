<script lang="ts">
    import ActivityList from "$lib/components/activity/ActivityList.svelte";
    import { overrideItemIdKeyNameBeforeInitialisingDndZones } from "svelte-dnd-action";
    import type { ActivityTypeWithSubtypes } from "../../../model/activity_type";
    import { GetMicrocycle, type MicrocycleInterface } from "../../../model/microcycle";
    import { createEventDispatcher } from "svelte";
    import { ConvertDuration } from "$lib/model/date";

    overrideItemIdKeyNameBeforeInitialisingDndZones("uuid");

    export let activity_type_list: ActivityTypeWithSubtypes[];
    export let microcycle: MicrocycleInterface;

    const dispatch = createEventDispatcher();
    function dispatchUpdate() {
        console.log("Microcycle::: change")
        dispatch("update")
    }

</script>
{#if microcycle}
<div class="microcycle container">
    <div class="overflow-scroll" style="max-height:50vh">
        <table class="table">
            <thead>
                <tr>
                    {#each microcycle.cycle_activities as cycle, index}
                        <th> {cycle.date_object.day_of_week} </th>
                    {/each}
                    <th>Summary</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    {#each microcycle.cycle_activities as act_list, column_index}
                        <td>
                            <ActivityList
                                bind:activity_list={act_list}
                                bind:date={act_list.date_object.date}
                                bind:activity_type_list={activity_type_list}
                                on:change={dispatchUpdate}
                            />
                        </td>
                    {/each}
                    <td style="padding: 0;">
                        <ul class="list-group list-group-flush striped-list summary">
                            {#each microcycle.summary.totals_by_activity_type as act_type_total}
                                <!-- This is a hacky way to check for which activities types are in the current microcyle. This way we can distinguish between totals and planning better -->
                                {#if microcycle.summary.totals_by_activity_type_and_date.some(atd => atd.activity_type.id === act_type_total.activity_type.id)}
                                    <li class="list-group-item">
                                        <div class="row">
                                            <div class="col">
                                                <!-- The type -->
                                                <div class="col-md-auto d-flex align-items-center">
                                                    <strong>{act_type_total.activity_type.name}</strong>
                                                </div>
                                                <table>
                                                    <tbody>
                                                        <!-- Distance -->
                                                        <tr>
                                                            <td>{act_type_total.total_completed_distance.length} {act_type_total.total_completed_distance.unit}</td>
                                                            <td class="text-secondary">{act_type_total.total_planned_distance.length} {act_type_total.total_planned_distance.unit}</td>
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
                                                            <td>{act_type_total.total_completed_vertical.length} {act_type_total.total_completed_vertical.unit}</td>
                                                            <td class="text-secondary">{act_type_total.total_planned_vertical.length} {act_type_total.total_planned_vertical.unit}</td>
                                                        </tr>
                                                    </tbody>
                                                </table>
                                            </div>
                                        </div>
                                    </li>
                                {/if}
                            {/each}
                            <br>
                            <li class="list-group-item">
                                <div class="row">
                                    <div class="col">
                                        <!-- microcycle totals -->
                                        <div class="col-md-auto d-flex align-items-center">
                                            <strong>Totals</strong>
                                        </div>
                                        <table>
                                            <tbody>
                                                <!-- Distance -->
                                                <tr>
                                                    <td>{microcycle.summary.totals.total_planned_distance.length} {microcycle.summary.totals.total_planned_distance.unit}</td>
                                                    <td class="text-secondary">{microcycle.summary.totals.total_planned_distance.length} {microcycle.summary.totals.total_planned_distance.unit}</td>
                                                </tr>
                                                <!-- Duration -->
                                                <tr>
                                                    <td>{ConvertDuration(microcycle.summary.totals.total_completed_duration)}</td>
                                                    <td class="text-secondary">{ConvertDuration(microcycle.summary.totals.total_planned_duration)}</td>
                                                </tr>
                                                <!-- Vertical -->
                                                <tr>
                                                    <td>{microcycle.summary.totals.total_completed_vertical.length} {microcycle.summary.totals.total_completed_vertical.unit}</td>
                                                    <td class="text-secondary">{microcycle.summary.totals.total_planned_vertical.length} {microcycle.summary.totals.total_planned_vertical.unit}</td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </div>
                                </div>
                            </li>
                        </ul>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>
{/if}
<style>
    .microcycle table {
        table-layout: fixed;
        width: 100%;
    }
    .microcycle th {
      border: 1px solid lightgray;
      text-align: center;
      position: sticky;
      top: 0;
      z-index: 2;
      background-color: #f2f2f2;
    }
    .microcycle td {
      border: 1px solid lightgray;
      text-align: center;
      vertical-align: top;
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
    ul.striped-list > li:nth-of-type(even) {
        background-color: #e9e9f9 ;
    }
    ul.striped-list > li {
        border-bottom: 1px solid rgb(221,221,221);
        /*padding: 6px;*/
    }
    ul.striped-list > li:last-child {
        border-bottom: none;
        background-color: #e9e9f9 ;
    }
</style>