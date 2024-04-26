<script lang="ts">
    import { EmptyLength } from "$lib/model/distance";
    import { model } from "../../wailsjs/go/models";
    import BaseActivityModal from "./BaseActivityModal.svelte";
    import { createEventDispatcher } from "svelte";

    export let usr: model.User;
    export let equipment_choices: model.Equipment[];
    export let activity_type_list: model.ActivityTypeWithSubtypes[];
    export let date: Date;
    export let activity_list: model.Activity[];
    export let is_hovering: boolean = false;

    let a = {activity_type: {id: -1, name: ""},
		date: new Date(date),
		description: "",
		equipment_list: [],
		name: "",
		notes: "",
		order: activity_list.length + 1,
		planned: {activity_uuid: "", distance: EmptyLength(false, usr), duration: 0, vertical: EmptyLength(true, usr), pace: ""},
		completed: {activity_uuid: "", distance: EmptyLength(false, usr), duration: 0, vertical: EmptyLength(true, usr), pace: ""},
		type_subtype_list: [],
		uuid: "",
		color: "",
		is_race: false,
		num_strides: 0
    }
    let new_activity = new model.Activity(a)

    const dispatch = createEventDispatcher()
</script>

<BaseActivityModal
    bind:equipment_choices={equipment_choices}
    bind:activity_type_list={activity_type_list}
    bind:is_hovering={is_hovering}
    bind:activity={new_activity}
    bind:total_num_date_activities={activity_list.length}
    edit_type={"create"}
    on:new={() => {dispatch("new")}}
/>
