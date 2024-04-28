<script lang="ts">
    import compwstrava from '$lib/static/strava/strava_api_logos/compatible with strava/cptblWith_strava_light/api_logo_cptblWith_strava_horiz_light.svg?raw'
    import connectwstrava from '$lib/static/strava/connect_with_strava/btn_strava_connectwith_orange/btn_strava_connectwith_orange.svg?raw'

    import { oauth2, type model } from "../../wailsjs/go/models";
    import { CreateStravaToken, UpdateUser } from "$lib/wailsjs/go/controllers/UserHandler";
    import { BackfillData, GetNewData, OpenStravaAuth, RevokeAccess, StartListener } from '$lib/wailsjs/go/strava/Strava';
    import { GetMostRecentDate } from '$lib/wailsjs/go/controllers/ActivityHandler';

    export let usr: model.User;
    export let existing_strava_token: oauth2.Token | null;
    let is_editing = false;
    //let strava_token: oauth2.Token;
    let backfilling: boolean = false;
    let redirect_message: boolean = false;
    let error_message: string = ""

    let getting_new_data: boolean = false;

    let now = new Date()
    let month: string = ""
    let day: string = ""
    let year: number
    let dateString: string

    $: if (usr) {
        now = new Date(usr.initial_cycle_start)
        month = '' + (now.getMonth() + 1)
		day = '' + now.getDate()
		year = now.getFullYear()
    }

	$: if (month.length < 2)
        month = '0' + month;

    $: if (day.length < 2)
        day = '0' + day;

	$: dateString = [year, month, day].join('-');

    function toggleEdit() {
        is_editing = !is_editing
    };

    async function update() {
        //usr.initial_cycle_start = new Date(usr.initial_cycle_start)
        usr.initial_cycle_start = new Date(`${dateString}T00:00:00`);
        usr = await UpdateUser(usr)
        toggleEdit()
    }

    async function adf() {
        await OpenStravaAuth()
        redirect_message = true
        try {
            let strava_token = await StartListener()
            await CreateStravaToken(usr, strava_token)
            existing_strava_token = strava_token
        } catch (error) {
            error_message = String(error)
            redirect_message = false
            return
        }
        redirect_message = false
    }

    async function backfillStravaData() {
        backfilling = true;
        await BackfillData(usr)
        backfilling = false;
    }

    async function updateStravaData() {
        getting_new_data = true;
        let latest = await GetMostRecentDate()
        await GetNewData(usr, latest)
        getting_new_data = false;
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
                <input class="form-control" id="initial_cycle_start" name="initial_cycle_start" type="date" bind:value={dateString} disabled={!is_editing} readonly={!is_editing} required>
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

<div class="container">
    <div class="accordion" id="accordionExample">
        <div class="accordion-item">
          <h2 class="accordion-header">
            <button class="accordion-button btn" type="button" data-bs-toggle="collapse" data-bs-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
                {@html compwstrava}
            </button>
          </h2>
          <div id="collapseOne" class="accordion-collapse collapse show" data-bs-parent="#accordionExample">
            <div class="accordion-body">
                <div class="col">
                    {#if !existing_strava_token}
                        <div class="row">
                            <button class="btn btn-sm" type="button" on:click={adf}>{@html connectwstrava}</button>
                        </div>
                    {/if}
                    <div class="row">
                        {#if redirect_message}
                            <p>An authorization prompt should have opened in your browser, please check it and grant authorization.</p>
                        {/if}
                        {#if error_message}
                            <div class="container">
                                <div class="alert alert-danger" role="alert">
                                    {error_message}
                                </div>
                            </div>
                        {/if}
                    </div>
                </div>

                {#if existing_strava_token}
                    <button class="btn btn-primary" type="button" on:click={backfillStravaData} disabled={backfilling}>Backfill all strava data</button>
                    <button class="btn btn-primary" type="button" on:click={updateStravaData} disabled={getting_new_data}>Refresh Strava Data</button>
                    <button class="btn btn-primary" type="button" on:click={() => {RevokeAccess(usr); existing_strava_token = null}} disabled={getting_new_data}>Revoke Strava Access</button>
                    {#if backfilling}
                        <p>Currently backfilling strava data. This can take some time. Please be patient.</p>
                        <div class="spinner-border" role="status">
                            <span class="visually-hidden">Loading...</span>
                        </div>
                    {/if}
                    {#if getting_new_data}
                        <p>Currently getting new strava data. This can take some time. Please be patient.</p>
                        <div class="spinner-border" role="status">
                            <span class="visually-hidden">Loading...</span>
                        </div>
                    {/if}
                {/if}
            </div>
          </div>
      </div>
    </div>
</div>

{/if}