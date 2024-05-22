<script lang="ts">
    import { goto } from "$app/navigation";
    import { CreateUser, PersistUser } from "$lib/wailsjs/go/controllers/UserHandler";
    import { SetUser } from "$lib/wailsjs/go/main/App";
    import { model } from "../../wailsjs/go/models";

    let user_already_exits: boolean = false;
    let other_error: boolean = false;
    let new_user: model.User = new model.User({
        username: "",
        password: "",
        units: "imperial",
        cycle_start: "Monday",
        cycle_days: 7,
        initial_cycle_start: ""
    });

    let keep_logged_in: boolean = false;

    async function signup() {
        try {
            user_already_exits = false;
            other_error = false;
            let usr = await CreateUser(new_user)
            await SetUser(usr)
            if (keep_logged_in) {
				await PersistUser(usr)
			}
            goto("/microcycle")
        } catch (error) {
            if (error === "user already exists") {
                user_already_exits = true;
            } else {
                other_error = true;
                console.error(error)
            }
        }
    }

    function toLogin() {
        goto("/auth/login");
    };
</script>

<div class="signup container">
    <h3>Create a New Account</h3>

    <form>
        <h5>Basic Info:</h5>
        <input class="form-control" bind:value="{new_user.username}" type="email" name="username" placeholder="Email" required />

        {#if user_already_exits}
            <p style="color: red;">Username already exists</p>
        {/if}
        {#if other_error}
            <p style="color: red;">An unknown error occurred</p>
        {/if}
        <input class="form-control" bind:value="{new_user.password}" type="password" name="password" placeholder="Password" required />

        <h5>Personal Preferences:</h5>
        <p>(You can change these at any time)</p>

        <label for="units">Units:</label>
        <select id="units" class="form-control" bind:value={new_user.units} name="units" required>
            <option value="imperial">Imperial (mi)</option>
            <option value="metric">Metric (km)</option>
        </select>

        <label for="cycle_start">Cycle Start Day:</label>
        <select id="cycle_start" class="form-control" bind:value={new_user.cycle_start} name="cycle_start" required>
            <option value="Monday">Monday</option>
            <option value="Tuesday">Tuesday</option>
            <option value="Wednesday">Wednesday</option>
            <option value="Thursday">Thursday</option>
            <option value="Friday">Friday</option>
            <option value="Saturday">Saturday</option>
            <option value="Sunday">Sunday</option>
        </select>

        <p>
            Note: For cycles whose length is not divisble by 7 (i.e. 10), the start day of the week has no effect.
            Instead, you select a <i>start date</i> which will determine where the cycles are calculated from.
        </p>

        <label for="cycle_days">Number of days in a cycle:</label>
        <input id="cycle_days" class="form-control" bind:value="{new_user.cycle_days}" name="cycle_days" type="number" required>

        {#if new_user.cycle_days != 7}
            <label for="initial_cycle_start">Initial Cycle Start Date:</label>
            <input class="form-control" id="initial_cycle_start" name="initial_cycle_start" type="date" bind:value={new_user.initial_cycle_start} required>
        {/if}

        <div class="form-check form-switch">
			<input class="form-check-input" type="checkbox" role="switch" id="flexSwitchCheckDefault" bind:value={keep_logged_in}>
			<label class="form-check-label" for="flexSwitchCheckDefault">Keep Me Logged In</label>
		</div>

        <button class="btn btn-primary" on:click={async () => {await signup()}}>Sign Up</button>
    </form>
    <p>
    Already have an account?
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
    <strong class="pe-auto" on:click={toLogin}>Login</strong>
    </p>
</div>

