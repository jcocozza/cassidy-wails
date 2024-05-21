<script lang="ts">
    import { goto } from "$app/navigation";
    import Redirect from "$lib/components/Redirect.svelte";
    import { GetMicrocycleCurrentDates } from "$lib/wailsjs/go/controllers/UserHandler";
    import { HasUser, Logout } from "$lib/wailsjs/go/main/App";
    import { controllers } from "$lib/wailsjs/go/models";
    import { onMount } from "svelte";

    let dates: controllers.MCCurrentDate = { start_date: new Date(), end_date: new Date() }; // Initialize with empty values
    let has_usr: boolean = false;

    async function logout() {
        await Logout()
        goto("/auth/login")
    }

    onMount(async () => {
        has_usr = await HasUser();
        if (has_usr) {
            dates = await GetMicrocycleCurrentDates();
        }
    })
</script>

<Redirect bind:passcondition={has_usr} logout={false}>
    {#if dates.start_date && dates.end_date}
        <ul class="nav nav-tabs">
            <!-- <li class="nav-item">
                <a class="nav-link" aria-current="page" href="/home">
                    Cassidy
                </a>
            </li>
            -->
            <li class="nav-item">
                <a class="nav-link" href={`/calendar`}>
                    calendar
                </a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href={`/microcycle`}>
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
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <!-- svelte-ignore a11y-no-static-element-interactions -->
                <!-- svelte-ignore a11y-missing-attribute -->
                <a class="nav-link" on:click={async () => {await logout()}}>
                    logout
                </a>
            </li>
        </ul>

        <slot />
    {:else}
        loading...
    {/if}
</Redirect>