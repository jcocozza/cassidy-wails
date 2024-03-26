<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import BaseActivityModal from "./BaseActivityModal.svelte";
    import type { model } from "../../wailsjs/go/models";

    export let equipment_choices: model.Equipment[];
    export let activity_type_list: model.ActivityTypeWithSubtypes[];
    export let total_num_date_activities: number;
    export let is_hovering: boolean = false;
    export let activity: model.Activity;

    const dispatch = createEventDispatcher()
</script>

<BaseActivityModal
    bind:equipment_choices={equipment_choices}
    bind:activity_type_list={activity_type_list}
    bind:is_hovering={is_hovering}
    bind:activity={activity}
    bind:total_num_date_activities={total_num_date_activities}
    edit_type={"update"}
    on:update={() => {dispatch("update")}}
/>
