<script lang="ts">
    import { onMount } from "svelte";
    import Chart from "chart.js/auto";
    import { ConvertDuration } from "../../model/date";
    import type { model } from "../../../../wailsjs/go/models";

    export let microcycle: model.Microcycle;
    $: start_date = microcycle.start_date;
    $: end_date = microcycle.end_date;
    $: date_list = createLabels(microcycle.cycle_activities)

    let chartType: string = "distance"
    let yAxisTitle: string = "distance(" + microcycle.summary.totals.total_completed_distance.unit + ")"
    let ctx;
    let canvas: { getContext: (arg0: string) => any; };
    let chart: Chart

    function createLabels(cycle_activities: model.ActivityList[]): string[] {
        let date_labels: string[] = []
        cycle_activities.forEach((act_list) => {
            date_labels = [...date_labels, act_list.date_object.date]
        })
        return date_labels
    }

    // Function to append or create a key with an array
    function appendOrCreateKey(dictionary: { [key: string]: any[] }, key: string, value: any): void {
        if (dictionary.hasOwnProperty(key)) {
            // Key exists, append to the existing array
            dictionary[key].push(value);
        } else {
            // Key doesn't exist, create a new key with an array
            dictionary[key] = value;
        }
    }

    function MakeDatasets(mc: model.Microcycle, type: string) {
        let act_type_dict: {[key: string]: number[]} = {};
        mc.summary.totals_by_activity_type_and_date.forEach((t) => {
            const stackData: number[] = Array(date_list.length).fill(0);  // Initialize with zeros for all labels
            appendOrCreateKey(act_type_dict, t.activity_type.name, stackData)
        })
        mc.summary.totals_by_activity_type_and_date.forEach((t) => {
            const index = date_list.indexOf(t.date.date);
            let act_type = t.activity_type.name;
            let lst =  act_type_dict[act_type];
            if (type == "distance") {
                lst[index] = t.total_completed_distance.length;
            } else if (type == "duration") {
                lst[index] = t.total_completed_duration;
            } else if (type == "vertical") {
                lst[index] = t.total_completed_vertical.length;
            }
        })

        let ds: any[] = [];
        for (let act_type in act_type_dict){
            let d = {
                label: act_type,
                data: act_type_dict[act_type]
            }
            ds = [...ds, d]
        }
        return ds
    }

    function updateChart(type: string) {
        if (chart) {
            chart.data.labels = createLabels(microcycle.cycle_activities)
            chart.data.datasets = MakeDatasets(microcycle, type);
            //let yAxis = chart.options?.scales?.y;

            if (type == "distance") {
                yAxisTitle = "distance (" + microcycle.summary.totals.total_completed_distance.unit + ")"
                // the error is not a problem here?
                chart.options.scales.y.title = { display: true, text: yAxisTitle }
                //chart.options.scales.suggestedMax = microcycle.summary.totals.total_completed_distance.length + 5
            } else if (type == "vertical") {
                yAxisTitle = "vertical (" + microcycle.summary.totals.total_completed_vertical.unit + ")"
                // the error is not a problem here?
                chart.options.scales.y.title = { display: true, text: yAxisTitle }
                //chart.options.scales.suggestedMax = microcycle.summary.totals.total_completed_vertical.length + 10
            } else {
                yAxisTitle = ""
                // the error is not a problem here?
                chart.options.scales.y.title = { display: true, text: yAxisTitle }
                //chart.options.scales.suggestedMax = microcycle.summary.totals.total_completed_duration + 10
            }
            chart.update()
        }
    }

    const initChart = () => {
        let date_labels: string[] = date_list;
        ctx = canvas.getContext('2d')

        let datasets = MakeDatasets(microcycle, chartType)
        chart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: date_labels,
                datasets: datasets
            },
            options: {
                scales: {
                    x: { stacked: true },
                    y: {
                        stacked: true,
                        ticks: {
                            // convert the ticks to string duration if we are dealing with duration
                            callback: (value) => {
                                if (chartType == "duration") {
                                    return ConvertDuration(Number(value));
                                } else {
                                    return value;
                                }
                            }
                        },
                        title: {
                            display: true,
                            text: yAxisTitle
                        },
                    }
                },
                plugins: {
                    tooltip: {
                        callbacks: {
                            // convert the tooltip labels to
                            label: (item) => {
                                if (chartType == "duration") {
                                    const value = item.raw;
                                    return ConvertDuration(Number(value));
                                }
                            }
                        }
                    }
                }
            },
        })

    }

    $: {
        console.log("PLOT:: start_date changed:", start_date);
        console.log("PLOT:: end_date changed:", end_date);
        console.log("PLOT::: microcycle changed:", microcycle);
        updateChart(chartType)
    }

    function handleRadioChange(event: Event) {
        const target = event.target as HTMLInputElement;
        chartType = target.value;
    }

    onMount(() => {
        initChart()
    })

</script>

<div class="completed-bar">

    <canvas bind:this={canvas}></canvas>

    <div class="btn-group d-flex justify-content-center" role="group" aria-label="Basic radio toggle button group">
        <input type="radio" class="btn-check" name="btnradio-bar" id="btnradio1-bar" autocomplete="off" value="distance" checked on:click={() => {chartType = "distance"}}>
        <label class="btn btn-outline-primary" for="btnradio1-bar">Distance</label>

        <input type="radio" class="btn-check" name="btnradio-bar" id="btnradio2-bar" autocomplete="off" value="duration" on:click={() => {chartType = "duration"}}>
        <label class="btn btn-outline-primary" for="btnradio2-bar">Duration</label>

        <input type="radio" class="btn-check" name="btnradio-bar" id="btnradio3-bar" autocomplete="off" value="vertical" on:click={() => {chartType = "vertical"}}>
        <label class="btn btn-outline-primary" for="btnradio3-bar">Vertical</label>
    </div>

</div>
