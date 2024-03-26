<script lang="ts">
    import { page } from "$app/stores";
    import { GetMicrocycleCurrentDates } from "$lib/wailsjs/go/controllers/UserHandler";
    import MicrocycleViewer from "$lib/pages/MicrocycleViewer.svelte";
    import { onMount } from "svelte";
    import type { model } from "$lib/wailsjs/go/models";
    $: start_date = $page.params.startdate;
    $: end_date = $page.params.enddate;

    async function currentMC() {
        let d = await GetMicrocycleCurrentDates()
        start_date = d.start_date
        end_date = d.end_date
    }

    onMount(async () => {
        await currentMC()
    })
</script>
<MicrocycleViewer
    bind:start_date={start_date}
    bind:end_date={end_date}>
</MicrocycleViewer>