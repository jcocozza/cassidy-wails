<script lang="ts">
    import { page } from "$app/stores";
    import CalendarViewer from "$lib/pages/CalendarViewer.svelte";
    import { LoadUser } from '$lib/wailsjs/go/main/App'
    import type { model } from "$lib/wailsjs/go/models";
    import { onMount } from "svelte";

    $: start_date = $page.params.startdate;
    $: end_date = $page.params.enddate;

    let usr: model.User;
    onMount(async () => {
        usr = await LoadUser()
    })
</script>
<CalendarViewer bind:usr={usr} bind:initial_start_date={start_date} bind:initial_end_date={end_date} />