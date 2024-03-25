<script lang="ts">
    import { onMount } from 'svelte';
    import { GetPlannedData, type PlannedApiData } from '../../../../func'
    import Chart from "chart.js/auto";
    import type { MicrocycleInterface } from "../../../../model/microcycle";
    import type { ActivityListInterface } from "../../../../model/activity_list";

    type dataset = {
        label: string;
        data: any[];
    }

    export let microcycle: MicrocycleInterface;
    $: start_date = microcycle.start_date;
    $: end_date = microcycle.end_date;
    $: date_list = createLabels(microcycle.cycle_activities)

    let is_loading: boolean = true;
    let planned_data: PlannedApiData;
    let datasets: dataset[];
    let current_view: number = 1;
    let ctx;
    let canvas: { getContext: (arg0: string) => any; };
    let chart: Chart

    function createLabels(cycle_activities: ActivityListInterface[]): string[] {
        let date_labels: string[] = []

        cycle_activities.forEach((act_list) => {
            date_labels = [...date_labels, act_list.date_object.date]
        })
        return date_labels
    }

    const triggerChartUpdate = async () => {
        datasets = [];
        if (chart) {
            is_loading = true;
            await GetPlannedData(start_date, end_date)
            .then((data): void => {
                planned_data = data;
                if (chart && data) {
                    updateChart(data)
                }
                is_loading = false;
            });
        }
    };

    $: {
        console.log("PLOT:: start_date changed:", start_date);
        console.log("PLOT:: end_date changed:", end_date);
        console.log("PLOT::: microcycle changed:", microcycle);
        triggerChartUpdate();
    }

    const initChart = () => {
        let date_labels: string[] = date_list
        console.log("DATA: ",datasets)
        ctx = canvas.getContext('2d');
        chart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: date_labels,
                datasets: []
            }
        });

        chart.data.datasets = [datasets[0]];
        chart.update();
    };

    function updateChart(data: PlannedApiData) {
        let ds = [{
                    label: 'Planned Distance',
                    data: data.data.planned_distance
                },
                {
                    label: 'Planned Duration',
                    data: data.data.planned_duration
                },
                {
                    label: 'Planned Vertical',
                    data: data.data.planned_vertical
                }];

        let date_labels: string[] = date_list

        chart.data.labels = date_labels;
        chart.data.datasets = [ds[current_view]];
        datasets = ds;
        chart.update();
    };

    function setData(id: number) {
        chart.data.datasets = [datasets[id]]
        chart.update()
        current_view = id;
    }

    onMount(async () => {
            console.log("LOADING BAR FOR DATE LIST:", date_list)
            await GetPlannedData(start_date, end_date).then((data) => {
                planned_data = data
                datasets = [{
                            label: 'Planned Distance',
                            data: planned_data.data.planned_distance
                        },
                        {
                            label: 'Planned Duration',
                            data: planned_data.data.planned_duration
                        },
                        {
                            label: 'Planned Vertical',
                            data: planned_data.data.planned_vertical
                    }];
            initChart();
        })
    });
</script>

<div class="planned-bar">
    <div class="btn-group" role="group" aria-label="Basic example">
        <button type="button" class="btn btn-secondary" on:click={() => setData(0)}>Distance</button>
        <button type="button" class="btn btn-secondary" on:click={() => setData(1)}>Duration</button>
        <button type="button" class="btn btn-secondary" on:click={() => setData(2)}>Vertical</button>
    </div>
    <canvas bind:this={canvas}></canvas>
</div>
