<script lang="ts">
    import { GetMicrocycleCurrentDates } from "$lib/wailsjs/go/controllers/UserHandler";
    import MicrocycleViewer from "$lib/pages/MicrocycleViewer.svelte";
    import { onMount } from "svelte";

    let start_date: string;
    let end_date: string;

    onMount(async () => {
        let d = await GetMicrocycleCurrentDates()
        start_date = d.start_date
        end_date = d.end_date
    })
</script>
<MicrocycleViewer
    bind:start_date={start_date}
    bind:end_date={end_date}>
</MicrocycleViewer>