<script lang="ts">
    import { GetMicrocycleCurrentDates } from "$lib/wailsjs/go/controllers/UserHandler";
    import { controllers } from "$lib/wailsjs/go/models";
    import { onMount } from "svelte";

    let d: controllers.MCCurrentDate = { start_date: '', end_date: '' }; // Initialize with empty values

    async function GetMCD() {
        try {
            const dates = await GetMicrocycleCurrentDates();
            console.log(dates.start_date);
            console.log(dates.end_date);
            d = dates; // Update d with the retrieved dates
        } catch (error) {
            console.error("Error fetching microcycle current dates:", error);
        }
    }

    onMount(async () => {
        await GetMCD();
    });

</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

{#if d.start_date && d.end_date}
    <p>Start Date: {d.start_date}</p>
    <p>End Date: {d.end_date}</p>
{:else}
    <p>Loading...</p>
{/if}
