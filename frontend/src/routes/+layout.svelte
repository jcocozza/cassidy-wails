<script lang="ts">
    import { onMount } from "svelte";
    import { GetMicrocycleCurrentDates } from "$lib/wailsjs/go/controllers/UserHandler";
    import { controllers } from "$lib/wailsjs/go/models";

    let dates: controllers.MCCurrentDate = { start_date: '', end_date: '' }; // Initialize with empty values

    onMount(async () => {
        dates = await GetMicrocycleCurrentDates();
    })
</script>

{#if dates.start_date && dates.end_date}
<ul class="nav nav-tabs">
    <li class="nav-item">
        <a class="nav-link" aria-current="page" href="/home">
            Cassidy
        </a>
    </li>
    <li class="nav-item">
        <a class="nav-link" href={`/calendar`}>
            calendar
        </a>
    </li>
    <li class="nav-item">
        <a class="nav-link" href={`/microcycle`}>
            microcycle
        </a>
    </li>
    <li class="nav-item">
        <a class="nav-link" href="/equipment">
            equipment
        </a>
    </li>
    <li class="nav-item">
        <a class="nav-link" href="/user">
            user
        </a>
    </li>
    <li class="nav-item">
        <a class="nav-link" href="/auth/logout">
            logout
        </a>
    </li>
</ul>
<slot />

{:else}
    loading...
{/if}