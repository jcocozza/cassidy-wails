<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { GetPreviousCycleRange, type DateObject, GetNextCycleRange, GetActivityTypes, type CycleSummary, GetCycleSummary } from "../../../func";
    import { CreateObserver } from "../../../../functions/infiniteScroll";
    import ActivityList from "../activity/ActivityList.svelte";
    import type { ActivityTypeWithSubtypes } from "../../../types";

    import { overrideItemIdKeyNameBeforeInitialisingDndZones } from "svelte-dnd-action";
    overrideItemIdKeyNameBeforeInitialisingDndZones("uuid");


    export let activity_type_list: ActivityTypeWithSubtypes[] = [];
    export let calendar_data: CyclePackage[];

    type CyclePackage = {
        cycle_dates: DateObject[];
        cycle_summary: CycleSummary;
    }

    const number_cycles_per_call: number = 10;
    const number_cycles_removed_per_call: number = 5;
    const total_number_cycles: number = 100;

    let top_elm: HTMLDivElement;
    let bottom_elm: HTMLDivElement;
    let scroller: HTMLDivElement;

    let top_observer: IntersectionObserver | null = null;
    let is_at_bottom: IntersectionObserver | null = null;

    let is_at_top: boolean = false;
    let is_at_bottom: boolean = false;

    let is_calling: Boolean = false;

    async function getPreviousCycles(number_cycles: number) {
        if (!is_calling) {
            is_calling = true;
            is_at_top = true;
            try {
                let start_cycle = calendar_data[0].cycle_dates[0].date;
                let end_cycle = calendar_data[0].cycle_dates[calendar_data[0].cycle_dates.length - 1].date;

                const row_height = 200;
                await GetPreviousCycleRange(start_cycle, end_cycle, number_cycles).then((data) => {
                    if (calendar_data.length > total_number_cycles) {
                        calendar_data = calendar_data.slice(0, -number_cycles_removed_per_call)
                    }
                    let cycle_package_list: CyclePackage[] = [];
                    data.forEach(async (date_obj_list) => {
                        let start = date_obj_list[0].date
                        let end = date_obj_list[date_obj_list.length - 1].date
                        let cycle_summary = await GetCycleSummary(start, end);
                        let cycle_package: CyclePackage = {cycle_dates: date_obj_list, cycle_summary: cycle_summary}
                        cycle_package_list = [...cycle_package_list, cycle_package]
                    })
                    calendar_data = [...cycle_package_list.reverse(), ...calendar_data]
                })
                scroller.scrollTo(0, (row_height * number_cycles))
            } catch (error) {
                console.error("failed to get previous cycles!!!:" + error)
            } finally {
                is_at_top = false;
                is_calling = false;
            }
        }
    }
    async function getNextCycles(number_cycles: number) {
        if (!is_calling) {
            is_calling = true;
            is_at_bottom = true;
            try {
                let start_cycle = calendar_data[0].cycle_dates[0].date;
                let end_cycle = calendar_data[0].cycle_dates[calendar_data[0].cycle_dates.length - 1].date;

                await GetNextCycleRange(start_cycle, end_cycle, number_cycles).then((data) => {
                    console.log(data)
                    if (calendar_data.length > total_number_cycles) {
                        calendar_data = calendar_data.slice(number_cycles_removed_per_call)
                    }
                    let cycle_package_list: CyclePackage[] = [];
                    data.forEach(async (date_obj_list) => {
                        let start = date_obj_list[0].date
                        let end = date_obj_list[date_obj_list.length - 1].date
                        let cycle_summary = await GetCycleSummary(start, end);
                        let cycle_package: CyclePackage = {cycle_dates: date_obj_list, cycle_summary: cycle_summary}
                        cycle_package_list = [...cycle_package_list, cycle_package]
                    })

                    calendar_data = [...calendar_data, ...cycle_package_list]
                })
            } catch (error) {
                console.error("failed to get next cycles!!" + error)
            } finally {
                is_at_bottom = false;
                is_calling = false;
            }
        }
    }

    $: {
        if (top_elm) {
            top_observer = CreateObserver(top_elm, async () => { await getPreviousCycles(number_cycles_per_call); });
        }
        if (bottom_elm) {
            bottom_observer = CreateObserver(bottom_elm, async () => { await getNextCycles(number_cycles_per_call); });
        }
    }

    onMount(async () => {
        if (top_elm) {
            top_observer = CreateObserver(top_elm, async () => { await getPreviousCycles(number_cycles_per_call); });
        }
        if (bottom_elm) {
            bottom_observer = CreateObserver(bottom_elm, async () => { await getNextCycles(number_cycles_per_call); });
        }
        activity_type_list = await GetActivityTypes();
    })
    onDestroy(() => {
        // Cleanup observers to avoid memory leaks
        if (top_observer) top_observer.disconnect();
        if (bottom_observer) bottom_observer.disconnect();
    });
</script>

{#if calendar_data.length != 0}
<div class="calendar container">
    <div class="overflow-scroll" style="max-height:100vh" bind:this={scroller}> <!--bind:this={scroller}> -->
        {#if is_at_top === false}
            <div bind:this={top_elm} class="spinner-border" role="status"></div>
        {/if}
        <table>
            <thead>
                <tr>
                    {#each Array(calendar_data[0].cycle_dates.length).fill(null) as _, index}
                        <th> {calendar_data[0].cycle_dates[index].day_of_week}</th>
                    {/each}
                    <th>Summary</th>
                </tr>
            </thead>
            <tbody>
                {#each calendar_data as cycle_date_list, row_index}
                    <tr>
                        {#each Array(cycle_date_list.cycle_dates.length).fill(null) as _, column_index}
                            <td>
                                {cycle_date_list.cycle_dates[column_index].date}
                                <ActivityList bind:date={cycle_date_list.cycle_dates[column_index].date} bind:activity_type_list={activity_type_list}/>
                            </td>
                        {/each}
                        <td>
                            {cycle_date_list.cycle_summary.planned_distance}
                            {cycle_date_list.cycle_summary.planned_duration}
                            {cycle_date_list.cycle_summary.planned_vertical}
                        </td>
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
    th {
      border: 1px solid black;
      padding: 8px 12px;
      text-align: center;
      position: sticky;
      top: 0;
      z-index: 2;
      background-color: #f2f2f2;
    }
    td {
      border: 1px solid black;
      padding: 8px 12px;
      text-align: center;
      vertical-align: top;
    }
    .calendar table {
        table-layout: fixed;
        width: 100%;
    }
</style>