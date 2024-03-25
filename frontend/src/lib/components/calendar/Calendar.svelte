<script lang="ts">
    import { onMount } from 'svelte';

    import { overrideItemIdKeyNameBeforeInitialisingDndZones } from "svelte-dnd-action";
    overrideItemIdKeyNameBeforeInitialisingDndZones("uuid");

    import { FormatPercent } from '$lib/misc/percent';
    import ActivityList from '../activity/ActivityList.svelte';
    import { CreateObserver } from '../../../functions/infiniteScroll';
    import { ConvertDuration, IsToday } from '../../model/date';
    import type { model } from '../../../../wailsjs/go/models';
    import { GetMicrocycle, GetCalendar, GetNextNMicrocycles, GetPreviousNMicrocycles } from '../../../../wailsjs/go/controllers/MicrocycleHandler'
    import { ListActivityTypes } from '../../../../wailsjs/go/controllers/ActivityTypeHandler'

    export let initial_start_date: string;
    export let initial_end_date: string;
    export let equipment_choices: model.Equipment[];
    export let usr: model.User;

    let microcycle_list: model.Microcycle[];
    let activity_type_list: model.ActivityTypeWithSubtypes[] = [];

    let top_elm: HTMLDivElement;
    let bottom_elm: HTMLDivElement;
    let scroller: HTMLDivElement;
    let top_observer: IntersectionObserver | null = null;
    let bottom_observer: IntersectionObserver | null = null;
    let is_at_top: boolean = false;
    let is_at_bottom: boolean = false;
    let is_loading = false;

    let summary_col_is_visible: boolean = true;
    let planning_col_is_visible: boolean = true;
    let weighted_totals: boolean = false;
    let display_completion: boolean = true;
    let today: HTMLDivElement;

    const number_new_cycles_per_pull: number = 15;
    const max_number_cycles: number = 100;

    async function replaceMicrocycle(microcycle: model.Microcycle, index: number) {
        try {
            const updatedMicrocycle = await GetMicrocycle(microcycle.start_date, microcycle.end_date, usr);
            microcycle_list = [...microcycle_list.slice(0, index), updatedMicrocycle, ...microcycle_list.slice(index + 1)];
        } catch (error) {
            console.log(error)
        }
    }
    // scroll to the today div
    function ScrollToToday() {
        today.scrollIntoView({ behavior: 'smooth', block: 'end', inline: 'nearest' });
    }

    onMount(async () => {
        activity_type_list = await ListActivityTypes()
        microcycle_list = await GetCalendar(initial_start_date, initial_end_date, usr);
    })

    $: {
        if (top_elm) {
            is_at_top = true;

            top_observer = CreateObserver(top_elm, async () => {
                let start_date = microcycle_list[0].start_date
                let end_date = microcycle_list[0].end_date
                //console.log("Adding data at top")
                if (!is_loading) {
                    is_loading = true;

                    let before_scroll_height = scroller.scrollHeight;
                    await GetPreviousNMicrocycles(start_date, end_date, number_new_cycles_per_pull, usr).then(async (data) => {
                        microcycle_list = [...data, ...microcycle_list]
                        console.log(microcycle_list)
                        await new Promise(r => setTimeout(r, 100)); // the browers needs just a fraction of a second to catch up
                        let scroll_change = scroller.scrollHeight - before_scroll_height
                        //scroller.scrollTo(0,scroll_change);
                        scroller.scrollTo({
                            top: scroll_change,
                            behavior: "instant"
                            });
                        is_loading = false;
                    })
                }
            });

            is_at_top = false;
        }
        if (bottom_elm) {
            is_at_bottom = true;

            bottom_observer = CreateObserver(bottom_elm, async () => {
                //console.log("Adding data at bottom")
                let start_date = microcycle_list[microcycle_list.length - 1].start_date;
                let end_date = microcycle_list[microcycle_list.length - 1].end_date;

                if (!is_loading) {
                    is_loading = true;
                    await GetNextNMicrocycles(start_date, end_date, number_new_cycles_per_pull, usr).then((data) => {
                        microcycle_list = [...microcycle_list, ...data];
                        is_loading = false;
                    })
                }
            })
            is_at_bottom = false;
        }
    }
</script>

{#if microcycle_list}
    <div class="calendar container">
        <div class="row">
            <div class="col flex-grow-1">
                {usr.username}
            </div>
            <div class="col">
                <button class="btn btn-primary btn-sm" on:click={ScrollToToday}>Today</button>
            </div>
            <div class="col small-text d-flex justify-content-end">
                <div class="btn-group btn-group-vertical btn-group-sm">
                    <div class="btn-group btn-group-sm">
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
                        <button class="btn btn-primary btn-sm" on:click={() => {planning_col_is_visible = !planning_col_is_visible}}>
                            {#if planning_col_is_visible}
                                &gt; Plan
                            {:else}
                                &lt; Plan
                            {/if}
                        </button>
                    </div>
                    {#if planning_col_is_visible}
                        <div class="form-check form-switch">
                            <input class="form-check-input" type="checkbox" role="switch" id="flexSwitchCheckChecked2" bind:checked={weighted_totals}>
                            <label class="form-check-label" for="flexSwitchCheckChecked2">weighted totals (experimental)</label>
                        </div>
                    {:else}
                        <br>
                    {/if}
                </div>
            </div>
        </div>
        <div class="overflow-scroll" style="max-height:90vh" bind:this={scroller}>
            {#if is_at_top === false}
                <div bind:this={top_elm} role="status"></div>
            {/if}
            <table class="table table-striped">
                <thead>
                    {#if microcycle_list[0].cycle_activities.length == 7}
                        {#each microcycle_list[0].cycle_activities as cycle}
                            <th class="bg-body-secondary"> {cycle.date_object.day_of_week} </th>
                        {/each}
                    {:else}
                        {#each microcycle_list[0].cycle_activities as cycle}
                            <th class="bg-body-secondary"></th>
                        {/each}
                    {/if}

                    {#if summary_col_is_visible}
                        <th class="bg-body-secondary">
                            Summary
                            <table class="summary">
                                <tr>
                                    <td>Completed</td>
                                    <td>Planned</td>
                                </tr>
                            </table>
                        </th>
                    {/if}
                    {#if planning_col_is_visible}
                        <th class="bg-body-secondary">
                            Planning
                            <table class="summary">
                                <tr>
                                    <td>Realized <br> (%)</td>
                                    <td>Actual <br> (%)</td>
                                </tr>
                            </table>
                        </th>
                    {/if}
                </thead>
                <tbody>
                    {#each microcycle_list as cycle, index}
                        <tr>
                            {#each cycle.cycle_activities as activity_list}
                                <td style="padding: 0;">
                                    {#if microcycle_list[0].cycle_activities.length != 7}
                                        {activity_list.date_object.day_of_week}
                                    {/if}
                                    <ActivityList
                                        bind:activity_list={activity_list}
                                        bind:date={activity_list.date_object.date}
                                        bind:activity_type_list={activity_type_list}
                                        bind:display_completion={display_completion}
                                        bind:equipment_choices={equipment_choices}
                                        on:change={async () => {
                                            console.log('Calendar::: Change requested.');
                                            await replaceMicrocycle(cycle, index);
                                        }} />
                                        <!-- If it is today, assign an invisible div so that we can return to it later -->
                                        {#if IsToday(activity_list.date_object.date)}
                                            <div bind:this={today}></div>
                                        {/if}
                                </td>
                            {/each}
                            {#if summary_col_is_visible}
                                <td style="padding: 0;">
                                    <ul class="list-group list-group-flush striped-list summary">
                                        {#each cycle.summary.totals_by_activity_type as act_type_total}
                                            <!-- This is a hacky way to check for which activities types are in the current microcyle. This way we can distinguish between totals and planning better -->
                                            {#if cycle.summary.totals_by_activity_type_and_date.some(atd => atd.activity_type.id === act_type_total.activity_type.id)}
                                                <li class="list-group-item">
                                                    <div class="row">
                                                        <div class="col">
                                                            <!-- The type -->
                                                            <div class="col-md-auto d-flex align-items-center">
                                                                <strong>{act_type_total.activity_type.name}</strong>
                                                            </div>
                                                            <!-- The margin left is to ensure that the the pacing doesn't run away -->
                                                            <!-- This is kinda a hack and I'm looking for an alternative -->
                                                            <div style="margin-left: -10px;">
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
                                                                <td>{cycle.summary.totals.total_planned_distance.length} {cycle.summary.totals.total_planned_distance.unit}</td>
                                                                <td class="text-secondary">{cycle.summary.totals.total_planned_distance.length} {cycle.summary.totals.total_planned_distance.unit}</td>
                                                            </tr>
                                                            <!-- Duration -->
                                                            <tr>
                                                                <td>{ConvertDuration(cycle.summary.totals.total_completed_duration)}</td>
                                                                <td class="text-secondary">{ConvertDuration(cycle.summary.totals.total_planned_duration)}</td>
                                                            </tr>
                                                            <!-- Vertical -->
                                                            <tr>
                                                                <td>{cycle.summary.totals.total_completed_vertical.length} {cycle.summary.totals.total_completed_vertical.unit}</td>
                                                                <td class="text-secondary">{cycle.summary.totals.total_planned_vertical.length} {cycle.summary.totals.total_planned_vertical.unit}</td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </div>
                                            </div>
                                        </li>
                                    </ul>
                                </td>
                            {/if}
                            {#if planning_col_is_visible}
                                <td style="padding: 0;">
                                    <ul class="list-group list-group-flush striped-list summary">
                                        {#each cycle.summary.totals_by_activity_type_differences as act_type_difference}
                                            <!-- {#if cycle.summary.totals_by_activity_type.some(tat => tat.activity_type.id == act_type_difference.activity_type.id) } -->
                                                <li class="list-group-item">
                                                    <div class="row">
                                                        <div class="col">
                                                            <div class="col-md-auto d-flex align-items-center">
                                                                <strong>{act_type_difference.activity_type.name} Changes</strong>
                                                            </div>
                                                            <table>
                                                                <tbody>
                                                                    <!--
                                                                    <tr>
                                                                        <td>Realized</td>
                                                                        <td class="text-secondary">Planned</td>
                                                                    </tr>
                                                                    -->
                                                                    <!-- Distance -->
                                                                    <tr>
                                                                        <td>
                                                                            {act_type_difference.difference_completed_distance.length} {act_type_difference.difference_completed_distance.unit}
                                                                            <br>
                                                                            ({FormatPercent(act_type_difference.percent_difference_completed_distance)})
                                                                        </td>
                                                                        <td class="text-secondary">
                                                                            {act_type_difference.difference_planned_distance.length} {act_type_difference.difference_planned_distance.unit}
                                                                            <br>
                                                                            ({FormatPercent(act_type_difference.percent_difference_planned_distance)})
                                                                        </td>
                                                                    </tr>
                                                                    <!-- Duration -->
                                                                    <tr>
                                                                        <td>
                                                                            {ConvertDuration(act_type_difference.difference_completed_duration)}
                                                                            <br>
                                                                            ({FormatPercent(act_type_difference.percent_difference_completed_duration)})
                                                                        </td>
                                                                        <td class="text-secondary">
                                                                            {ConvertDuration(act_type_difference.difference_planned_duration)}
                                                                            <br>
                                                                            ({FormatPercent(act_type_difference.percent_difference_planned_duration)})
                                                                        </td>
                                                                    </tr>
                                                                    <!-- Vertical -->
                                                                    <tr>
                                                                        <td>
                                                                            {act_type_difference.difference_completed_vertical.length} {act_type_difference.difference_completed_vertical.unit}
                                                                            <br>
                                                                            ({FormatPercent(act_type_difference.percent_difference_completed_vertical)})
                                                                        </td>
                                                                        <td class="text-secondary">
                                                                            {act_type_difference.difference_planned_vertical.length} {act_type_difference.difference_planned_vertical.unit}
                                                                            <br>
                                                                            ({FormatPercent(act_type_difference.percent_difference_planned_vertical)})
                                                                        </td>
                                                                    </tr>
                                                                </tbody>
                                                            </table>
                                                        </div>
                                                    </div>
                                                </li>
                                            <!-- {/if} -->
                                        {/each}

                                        <br>
                                        <li class="list-group-item">
                                            <div class="row">
                                                <div class="col">
                                                    <div class="col-md-auto d-flex align-items-center">
                                                        <strong>Total Changes</strong>
                                                    </div>
                                                    <table>
                                                        <tbody>
                                                            <!--
                                                            <tr>
                                                                <td>Realized</td>
                                                                <td class="text-secondary">Planned</td>
                                                            </tr>
                                                            -->
                                                            <!-- Distance -->
                                                            <tr>
                                                                <td>
                                                                    {cycle.summary.totals_differences.difference_completed_distance.length} {cycle.summary.totals_differences.difference_completed_distance.unit}
                                                                    <br>
                                                                    {#if weighted_totals}
                                                                        ({FormatPercent(cycle.summary.weighted_totals_differences.percent_difference_completed_distance)})
                                                                    {:else}
                                                                        ({FormatPercent(cycle.summary.totals_differences.percent_difference_completed_distance)})
                                                                    {/if}
                                                                </td>
                                                                <td class="text-secondary text-center">
                                                                    {cycle.summary.totals_differences.difference_planned_distance.length} {cycle.summary.totals_differences.difference_planned_distance.unit}
                                                                    <br>
                                                                    {#if weighted_totals}
                                                                        ({FormatPercent(cycle.summary.weighted_totals_differences.percent_difference_planned_distance)})
                                                                    {:else}
                                                                        ({FormatPercent(cycle.summary.totals_differences.percent_difference_planned_distance)})
                                                                    {/if}
                                                                </td>
                                                            </tr>
                                                            <!-- Duration -->
                                                            <tr>
                                                                <td class="text-center">
                                                                    {ConvertDuration(cycle.summary.totals_differences.difference_completed_duration)}
                                                                    <br>
                                                                    {#if weighted_totals}
                                                                        ({FormatPercent(cycle.summary.weighted_totals_differences.percent_difference_completed_duration)})
                                                                    {:else}
                                                                        ({FormatPercent(cycle.summary.totals_differences.percent_difference_completed_duration)})
                                                                    {/if}
                                                                </td>
                                                               <td class="text-secondary text-center">
                                                                    {ConvertDuration(cycle.summary.totals.total_planned_duration)}
                                                                    <br>
                                                                    {#if weighted_totals}
                                                                        ({FormatPercent(cycle.summary.weighted_totals_differences.percent_difference_planned_duration)})
                                                                    {:else}
                                                                        ({FormatPercent(cycle.summary.totals_differences.percent_difference_planned_duration)})
                                                                    {/if}
                                                                </td>
                                                            </tr>
                                                            <!-- Vertical -->
                                                            <tr>
                                                                <td>
                                                                    {cycle.summary.totals.total_completed_vertical.length} {cycle.summary.totals.total_completed_vertical.unit}
                                                                    <br>
                                                                    {#if weighted_totals}
                                                                        ({FormatPercent(cycle.summary.weighted_totals_differences.percent_difference_completed_vertical)})
                                                                    {:else}
                                                                        ({FormatPercent(cycle.summary.totals_differences.percent_difference_completed_vertical)})
                                                                    {/if}
                                                                </td>
                                                                <td class="text-secondary">
                                                                    {cycle.summary.totals.total_planned_vertical.length} {cycle.summary.totals.total_planned_vertical.unit}
                                                                    <br>
                                                                    {#if weighted_totals}
                                                                        ({FormatPercent(cycle.summary.weighted_totals_differences.percent_difference_planned_vertical)})
                                                                    {:else}
                                                                        ({FormatPercent(cycle.summary.totals_differences.percent_difference_planned_vertical)})
                                                                    {/if}
                                                                </td>
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
                    {/each}
                </tbody>
            </table>
            {#if is_at_bottom === false}
                <div bind:this={bottom_elm} class="spinner-border" role="status"></div>
            {/if}
        </div>
    </div>
{/if}

<style>
    .calendar th {
      border: 1px solid lightgray;
      text-align: center;
      position: sticky;
      top: 0;
      z-index: 2;
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