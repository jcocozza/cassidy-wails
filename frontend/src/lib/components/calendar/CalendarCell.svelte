<script lang="ts">
    import { IsToday, ParseDateYYYYMMDD } from "$lib/model/date";
    import type { model } from "$lib/wailsjs/go/models";
    import NewActivityModal from "../activity/NewActivityModal.svelte";
    import ActivityList from "../activity/ActivityList.svelte";
    import { createEventDispatcher } from "svelte";

    export let activity_list: model.ActivityList;
    export let num_cycle_days: number;
    export let usr: model.User;
    export let equipment_choices: model.Equipment[];
    export let activity_type_list: model.ActivityTypeWithSubtypes[]
    export let display_completion: boolean;
    export let today: HTMLDivElement;
    export let is_hovering: boolean = false;

    const dispatch = createEventDispatcher()

</script>

<td class="cal" style="padding: 0;" on:mouseenter={() => {is_hovering = true}} on:mouseleave={() => {is_hovering = false}}>
    <div class="col">
        <div class="row">
            <div class="col">
                {#if num_cycle_days != 7}
                    {ParseDateYYYYMMDD(activity_list.date)}
                {:else}
                    {ParseDateYYYYMMDD(activity_list.date)}
                {/if}
            </div>
            <div class="col">
                <NewActivityModal
                    bind:usr={usr}
                    bind:equipment_choices={equipment_choices}
                    bind:date={activity_list.date}
                    bind:activity_list={activity_list.activity_list}
                    bind:activity_type_list={activity_type_list}
                    bind:is_hovering={is_hovering}
                    on:new={() => {dispatch("update")}}
                />
            </div>
        </div>

        <div class="row">
            <ActivityList
                bind:activity_list={activity_list}
                bind:date={activity_list.date}
                bind:activity_type_list={activity_type_list}
                bind:display_completion={display_completion}
                bind:equipment_choices={equipment_choices}
                on:new={() => {dispatch("update")}}
            />
        </div>
    </div>
    <!-- If it is today, assign an invisible div so that we can return to it later -->
    {#if IsToday(activity_list.date.toString())}
        <div bind:this={today}></div>
    {/if}
</td>

<style>
    .cal {
      border: 1px solid lightgray;
      text-align: center;
      vertical-align: top;
    }
</style>