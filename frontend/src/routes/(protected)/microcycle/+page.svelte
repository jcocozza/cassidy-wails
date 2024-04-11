<script lang="ts">
    import { GetMicrocycleCurrentDates } from "$lib/wailsjs/go/controllers/UserHandler";
    import MicrocycleViewer from "$lib/pages/MicrocycleViewer.svelte";
    import { onMount } from "svelte";
    import type { model } from "$lib/wailsjs/go/models";
    import { LoadUser } from "$lib/wailsjs/go/main/App";

    let start_date: string = "";
    let end_date: string = "";
    let usr: model.User;
    onMount(async () => {
        const [user, dates] = await Promise.all([LoadUser(), GetMicrocycleCurrentDates()]);
        usr = user;
        start_date = dates.start_date;
        end_date = dates.end_date;
    });

</script>

{#if start_date && end_date && usr}
    <MicrocycleViewer
        bind:usr={usr}
        bind:start_date={start_date}
        bind:end_date={end_date}>
    </MicrocycleViewer>
{:else}
    <p>Loading MicrocycleViewer...</p>
{/if}