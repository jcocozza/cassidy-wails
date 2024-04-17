<script lang="ts">
    import CalendarViewer from "$lib/pages/CalendarViewer.svelte";
    import { GetMicrocycleCurrentDates } from "$lib/wailsjs/go/controllers/UserHandler";
    import { LoadUser } from '$lib/wailsjs/go/main/App'
    import type { model } from "$lib/wailsjs/go/models";
    import { onMount } from "svelte";

    let start_date: Date;
    let end_date: Date;
    let usr: model.User;

    onMount(async () => {
        usr = await LoadUser()
        let d = await GetMicrocycleCurrentDates()
        start_date = d.start_date
        end_date = d.end_date
    })
</script>

{#if start_date && end_date && usr}
    <CalendarViewer bind:usr={usr} bind:initial_start_date={start_date} bind:initial_end_date={end_date} />
{:else}
    <p>Loading CalendarViewer...</p>
{/if}
