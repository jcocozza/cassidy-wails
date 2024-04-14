<script lang="ts">
    import compwstrava from '$lib/static/strava/strava_api_logos/compatible with strava/cptblWith_strava_light/api_logo_cptblWith_strava_horiz_light.svg?raw'
    import connectwstrava from '$lib/static/strava/connect_with_strava/btn_strava_connectwith_orange/btn_strava_connectwith_orange.svg?raw'

    import type { model } from "../../wailsjs/go/models";
    import { UpdateUser } from "$lib/wailsjs/go/controllers/UserHandler";
    import { OpenStravaAuth, StartListener } from '$lib/wailsjs/go/strava/Strava';

    export let usr: model.User
    let is_editing = false;

    function toggleEdit() {
        is_editing = !is_editing
    };

    async function update() {
        usr = await UpdateUser(usr)
        toggleEdit()
    }

    async function adf() {
        StartListener()
        await OpenStravaAuth()
    }

</script>

{#if usr}
<div class="user-settings container">
    <form>
        <div class="form-group">
            <label for="unit_class">Unit System:</label>
            <select class="form-control" id="unit_class" name="unit_class" bind:value={usr.units} disabled={!is_editing} required>
                <option value="imperial"> Imperial (mile, ft, etc)</option>
                <option value="metric"> Metric (km, m, etc)</option>
            </select>


            <label for="cycle_start">Cycle Start Day:</label>
            <select class="form-control" id="cycle_start" name="cycle_start" bind:value={usr.cycle_start} disabled={!is_editing} required>
                <option value="Monday"> Monday </option>
                <option value="Tuesday"> Tuesday </option>
                <option value="Wednesday"> Wednesday </option>
                <option value="Thursday"> Thursday </option>
                <option value="Friday"> Friday </option>
                <option value="Saturday"> Saturday </option>
                <option value="Sunday"> Sunday </option>
            </select>

            <label for="cycle_days"> Number of days in a cycle:</label>
            <input class="form-control" id="cycle_days" name="cycle_days" type="number" bind:value={usr.cycle_days} disabled={!is_editing} readonly={!is_editing} required>

            {#if usr.cycle_days != 7}
                <label for="initial_cycle_start">Initial Cycle Start Date:</label>
                <input class="form-control" id="initial_cycle_start" name="initial_cycle_start" type="date" bind:value={usr.initial_cycle_start} disabled={!is_editing} readonly={!is_editing} required>
            {/if}

            <input type="hidden" id="uuid" name="uuid" bind:value={usr.uuid}>
        </div>

        {#if is_editing}
            <button class="btn btn-primary" on:click={async () => {await update()}}>Update</button>
        {/if}
    </form>
    {#if is_editing}
        <button class="btn btn-secondary" on:click={toggleEdit}>Cancel</button>
    {:else}
        <button class="btn btn-primary" on:click={toggleEdit}>Edit Settings</button>
    {/if}
</div>

<!-- TODO: password
<div class="password-edit">
    <div class="form-group">
        <label for="old_password">Password:</label>
        <input id="old_password" type="password" bind:value={usr.}>
    </div>
</div>
-->

{@html compwstrava}

<button class="btn btn-primary btn-sm" type="button" on:click={adf}>{@html connectwstrava}</button>
{/if}