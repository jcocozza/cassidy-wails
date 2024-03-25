<script lang="ts">
    import Chart from "chart.js/auto";
    import { afterUpdate, onMount } from "svelte";
    import { ConvertDuration } from "../../model/date";
    import { model } from "../../../../wailsjs/go/models";
    import { GetMicrocycleCurrentDates } from "../../../../wailsjs/go/controllers/UserHandler";
    import { GetNCycleSummary } from '../../../../wailsjs/go/controllers/MiscHandler'

    export let start_date: string = "";
    export let end_date: string = "";
    export let usr: model.User;

    let chart_type: string = "distance"
    let ctx;
    let canvas: { getContext: (arg0: string) => any; };
    let chart: Chart;
    let n_cycle_summary: model.NCycleSummary;
    let yAxisTitle = "";

    function updateChart(type: string) {
        chart.data.labels = n_cycle_summary.start_date_list.map(d => d.date)
        if (type == "distance") {
            chart.data.datasets = [
                {
                    label: "Completed Distance",
                    data: n_cycle_summary.completed_distances.map(dist => dist.length), // hacky list comprehension in a language that doesn't have it
                    fill: {
                        target: '1',
                        above: 'green',   // Area will be red above the origin
                        below: 'red'    // And blue below the origin
                    }
                },
                {
                    label: "Planned Distance",
                    data: n_cycle_summary.planned_distances.map(dist => dist.length),
                },
            ]
            yAxisTitle = "distance (" + n_cycle_summary.completed_distances[0].unit + ")";
            // the error is not a problem here?
            chart.options.scales.y.title = { display: true, text: yAxisTitle }
        } else if (type == "duration") {
            chart.data.datasets = [
                {
                    label: "Completed Duration",
                    data: n_cycle_summary.completed_durations,
                    fill: {
                        target: '1',
                        above: 'green',   // Area will be red above the origin
                        below: 'red'    // And blue below the origin
                    }
                },
                {
                    label: "Planned Duration",
                    data: n_cycle_summary.planned_durations
                }
            ]
            yAxisTitle = ""
            // the error is not a problem here?
            chart.options.scales.y.title = { display: true, text: yAxisTitle }
        } else if (type == "vertical") {
            chart.data.datasets = [
                {
                    label: "Completed Vertical",
                    data: n_cycle_summary.completed_verticals.map(vert => vert.length),
                    fill: {
                        target: '1',
                        above: 'green',   // Area will be red above the origin
                        below: 'red'    // And blue below the origin
                    }
                },
                {
                    label: "Planned Vertical",
                    data: n_cycle_summary.planned_verticals.map(vert => vert.length),
                }
            ]
            yAxisTitle = "vertical (" + n_cycle_summary.completed_verticals[0].unit + ")";
            // the error is not a problem here?
            chart.options.scales.y.title = { display: true, text: yAxisTitle }
        }
        chart.update()
    }

    const initChart = () => {
        ctx = canvas.getContext('2d')
        yAxisTitle = "distance(" + n_cycle_summary.completed_distances[0].unit + ")";
        chart = new Chart(ctx, {
            type: 'line',
            data: {
                labels: n_cycle_summary.start_date_list.map(d => d.date),
                datasets: [
                    {
                        label: "Completed Distance: " + n_cycle_summary.completed_distances[0].unit,
                        data: n_cycle_summary.completed_distances.map(dist => dist.length), // hacky list comprehension in a language that doesn't have it
                        fill: {
                            target: '1',
                            above: 'green',   // Area will be red above the origin
                            below: 'red'    // And blue below the origin
                        }
                    },
                    {
                        label: "Planned Distance: " + n_cycle_summary.planned_distances[0].unit,
                        data: n_cycle_summary.planned_distances.map(dist => dist.length)
                    }
                ]
            },
            options: {
                scales: {
                    y: {
                        suggestedMin: 0,
                        beginAtZero: true,
                        ticks: {
                            // convert the ticks to string duration if we are dealing with duration
                            callback: (value) => {
                                if (chart_type == "duration") {
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
                                if (chart_type == "duration") {
                                    const value = item.raw;
                                    return ConvertDuration(Number(value));
                                }
                            }
                        }
                    }
                }
            }
        })
    }

    onMount(async () => {
        if (start_date == "" || end_date == "") {
            let d = await GetMicrocycleCurrentDates(usr)
            n_cycle_summary = await GetNCycleSummary(d.start_date, d.end_date, usr);
        } else {
            n_cycle_summary = await GetNCycleSummary(start_date, end_date, usr)
        }
        initChart()
    })

    afterUpdate(() => {
        // console.log("PLOT:: start_date changed:", start_date);
        // console.log("PLOT:: end_date changed:", end_date);

        (async () => {
        try {
            n_cycle_summary = await GetNCycleSummary(start_date, end_date, usr);
            if (chart) {
                updateChart(chart_type);
            }
        } catch (error) {
            console.error("Error fetching data:", error);
        }
        })();
    });

    /*
   $: {
       console.log("PLOT:: start_date changed:", start_date);
       console.log("PLOT:: end_date changed:", end_date);
       async () => {
           n_cycle_summary = await GetNCycleSummary(start_date, end_date)
           if (chart) {
               updateChart(chart_type)
           }
       }
  }
  */
</script>

<div class="ncyclelinechart">
        <canvas bind:this={canvas}></canvas>

        <div class="btn-group d-flex justify-content-center" role="group" aria-label="Basic radio toggle button group">
            <input type="radio" class="btn-check" name="btnradio" id="btnradio1" autocomplete="off" value="distance" checked on:click={() => {chart_type = "distance"}}>
            <label class="btn btn-outline-primary" for="btnradio1">Distance</label>

            <input type="radio" class="btn-check" name="btnradio" id="btnradio2" autocomplete="off" value="duration" on:click={() => {chart_type = "duration"}}>
            <label class="btn btn-outline-primary" for="btnradio2">Duration</label>

            <input type="radio" class="btn-check" name="btnradio" id="btnradio3" autocomplete="off" value="vertical" on:click={() => {chart_type = "vertical"}}>
            <label class="btn btn-outline-primary" for="btnradio3">Vertical</label>
        </div>
</div>