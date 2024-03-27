<script lang="ts">
    import Microcycle from "$lib/components/calendar/Microcycle.svelte";
    import { onMount } from "svelte";
    import { type NextPrevious } from "../model/date";
    import Bar from "$lib/components/charts/Bar.svelte";
    import NCycleLineChart from "$lib/components/charts/NCycleLineChart.svelte";
    import { model } from "../wailsjs/go/models";
    import { GetMicrocycle } from "../wailsjs/go/controllers/MicrocycleHandler";
    import { GetNextPrevious } from '../wailsjs/go/controllers/MiscHandler'
    import { List } from "../wailsjs/go/controllers/EquipmentHandler";
    import { ListActivityTypes } from "../wailsjs/go/controllers/ActivityTypeHandler";

    export let start_date: string;
    export let end_date: string;
    export let usr: model.User;

    let equipment_choices: model.Equipment[] = [];

    let microcycle: model.Microcycle;
    let activity_type_list: model.ActivityTypeWithSubtypes[] = [];
    let next_previous: NextPrevious;

    let is_loading: boolean = false;

    function loadPage(direction: string) {
        if (direction == "next") {
            start_date = next_previous.next_start_date;
            end_date = next_previous.next_end_date;
        } else if (direction == "previous") {
            start_date = next_previous.previous_start_date;
            end_date = next_previous.previous_end_date;
        }
    }

    async function updateMicrocycle() {
        microcycle = await GetMicrocycle(start_date, end_date)
    }

    $: (async () => {
        microcycle = await GetMicrocycle(start_date, end_date)
        next_previous = await GetNextPrevious(start_date, end_date);
    })();

    onMount(async () => {
        /*
        microcycle = new model.Microcycle({
            start_date: start_date,
            end_date:end_date,
            cycle_activities: [],
            summary: {
                totals: EmptyTotal(usr),
                totals_differences: EmptyTotalsDifferences(usr),
                previous_totals: EmptyTotal(usr),
                average_previous_totals: EmptyTotal(usr),
                weighted_totals_differences: EmptyTotalsDifferences(usr),
                totals_by_activity_type: [],
                totals_by_activity_type_differences: [],
                previous_totals_by_activity_type: [],
                totals_by_activity_type_and_date: []
            }
        });
        */
        is_loading = true;

        const [
        microcycleData,
        activityTypes,
        equipment,
        nextPrev
        ] = await Promise.all([
            GetMicrocycle(start_date, end_date),
            ListActivityTypes(),
            List(),
            GetNextPrevious(start_date, end_date)
        ]);

        microcycle = microcycleData;
        activity_type_list = activityTypes;
        equipment_choices = equipment;
        next_previous = nextPrev;
        is_loading = false;
    });
</script>


{#if microcycle && usr && !is_loading}
<div class="container microcycle-viewer">
    <div class="row">
        <Microcycle
            bind:usr={usr}
            bind:microcycle={microcycle}
            bind:activity_type_list={activity_type_list}
            bind:equipment_choices={equipment_choices}
            on:update={updateMicrocycle}
        />
    </div>

    <div class="d-flex justify-content-between" style="padding: 0%;">
        <button class="btn btn-primary" on:click={() => {loadPage("previous")}}> Previous </button>
        <button class="btn btn-primary" on:click={() => {loadPage("next")}}> Next </button>
    </div>
    <div class="row">
        {#if microcycle.cycle_activities.length > 0}
            <div class="col">
                <Bar bind:microcycle={microcycle}></Bar>
            </div>
        {/if}

        <div class="col">
            <div class="container">
                <NCycleLineChart
                    bind:start_date={microcycle.start_date}
                    bind:end_date={microcycle.end_date}
                />
            </div>
        </div>
    </div>
</div>
{:else}
    Loading mcviewer...
{/if}