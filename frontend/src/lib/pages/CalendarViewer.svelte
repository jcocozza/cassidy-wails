<script lang="ts">
    import Calendar from "$lib/components/calendar/Calendar.svelte";
    import { onMount } from "svelte";
    import type { model } from "../wailsjs/go/models";
    import { List } from "../wailsjs/go/controllers/EquipmentHandler";

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
            bind:equipment_choices={equipment_choices}
        />
    </div>
</div>