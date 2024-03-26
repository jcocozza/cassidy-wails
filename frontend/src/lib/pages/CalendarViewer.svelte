<script lang="ts">
    import Calendar from "$lib/components/calendar/Calendar.svelte";
    import { onMount } from "svelte";
    import type { model } from "../wailsjs/go/models";
    import { List } from "../wailsjs/go/controllers/EquipmentHandler";

    export let initial_start_date: string;
    export let initial_end_date: string;
    export let usr: model.User;
    let equipment_choices: model.Equipment[] = [];

    onMount(async () => {
        equipment_choices = await List()
    });
</script>

<div class="calendar-viewer">

    <div class="row">
        <Calendar
            bind:usr={usr}
            bind:initial_start_date={initial_start_date}
            bind:initial_end_date={initial_end_date}
            bind:equipment_choices={equipment_choices}
        />
    </div>
</div>