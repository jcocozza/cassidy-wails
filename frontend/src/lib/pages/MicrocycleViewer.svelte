<script lang="ts">
    import Microcycle from "$lib/components/calendar/Microcycle.svelte";
    import { onMount } from "svelte";
    import { type NextPrevious } from "../model/date";
    import { goto } from "$app/navigation";
    import { EmptyTotalsDifferences, EmptyTotal,} from "../model/microcycle";
    import Bar from "$lib/components/charts/Bar.svelte";
    import NCycleLineChart from "$lib/components/charts/NCycleLineChart.svelte";
    import { model } from "../wailsjs/go/models";
    import { GetMicrocycle } from "../wailsjs/go/controllers/MicrocycleHandler";
    import { GetNextPrevious } from '../wailsjs/go/controllers/MiscHandler'
    import { List } from "../wailsjs/go/controllers/EquipmentHandler";
    import { ListActivityTypes } from "../wailsjs/go/controllers/ActivityTypeHandler";
    import { LoadUser } from "$lib/wailsjs/go/main/App";

    export let start_date: string;
    export let end_date: string;
    let usr: model.User;

    let equipment_choices: model.Equipment[] = [];

    let microcycle: model.Microcycle;
    let activity_type_list: model.ActivityTypeWithSubtypes[] = [];
    let next_previous: NextPrevious;

    let is_loading: boolean = false;

    function loadPage(url: string) {
        goto(url);
    }

    async function updateMicrocycle() {
        microcycle = await GetMicrocycle(start_date, end_date)
    }

    $: (async () => {
        microcycle = await GetMicrocycle(start_date, end_date)
        next_previous = await GetNextPrevious(start_date, end_date);
    })();

    onMount(async () => {
        usr = await LoadUser()
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

        is_loading = true;
        microcycle = await GetMicrocycle(start_date, end_date)
        activity_type_list = await ListActivityTypes()
        equipment_choices = await List()
        next_previous = await GetNextPrevious(start_date, end_date);
        is_loading = false;
    });
</script>

<div class="container microcycle-viewer">
    {#if is_loading && microcycle}
        loading...
    {:else}
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
            <button class="btn btn-primary" on:click={() => {loadPage("/microcycle/"+next_previous.previous_start_date+"/"+next_previous.previous_end_date)}}> Previous </button>
            <button class="btn btn-primary" on:click={() => {loadPage("/microcycle/"+next_previous.next_start_date +"/"+next_previous.next_end_date)}}> Next </button>
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
    {/if}
</div>