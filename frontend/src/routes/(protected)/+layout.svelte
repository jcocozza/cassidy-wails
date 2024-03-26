<script lang="ts">
    import { onMount } from "svelte";
    import { GetMicrocycleCurrentDates } from "../../lib/wailsjs/go/controllers/UserHandler";
    import { controllers } from "$lib/wailsjs/go/models";

    let dates: controllers.MCCurrentDate | null = null;

    onMount(async () => {
        dates = await GetMicrocycleCurrentDates();
    })
</script>

{#if dates}
<ul class="nav nav-tabs">
    <li class="nav-item">
        <a class="nav-link" aria-current="page" href="/home">
            Cassidy
        </a>
    </li>
    <li class="nav-item">
        <a class="nav-link" href={`/calendar/${encodeURIComponent(dates.start_date)}/${encodeURIComponent(dates.end_date)}`}>
            calendar
        </a>
    </li>
    <li class="nav-item">
        <a class="nav-link" href={`/microcycle/${encodeURIComponent(dates.start_date)}/${encodeURIComponent(dates.end_date)}`}>
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

{:else}
    This is the layout.
{/if}

<slot />