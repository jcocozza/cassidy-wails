<script lang="ts">
    import { page } from "$app/stores";
    import CalendarViewer from "$lib/pages/CalendarViewer.svelte";
    import type { model } from "../../../../../../wailsjs/go/models";

    export let usr: model.User;
    $: start_date = $page.params.startdate;
    $: end_date = $page.params.enddate;
</script>

<CalendarViewer bind:usr={usr} bind:initial_start_date={start_date} bind:initial_end_date={end_date} />