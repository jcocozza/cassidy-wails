<script lang="ts">
    import { GetUser } from "$lib/wailsjs/go/main/App";
    import { onMount } from "svelte";
    import { GetMicrocycleCurrentDates } from "../../../wailsjs/go/controllers/UserHandler";
    let d: any;

    onMount(() => {
        GetUser().then((usr) => {
            let d = GetMicrocycleCurrentDates(usr)
        })
    })
</script>


{#await d}
    Loading...
{:then dish}
        <ul class="nav nav-tabs">
            <li class="nav-item">
                <a class="nav-link" aria-current="page" href="/home">
                    Cassidy
                </a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/calendar/{dish.start_date}/{dish.end_date}">
                    calendar
                </a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/microcycle/{dish.start_date}/{dish.end_date}">
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
                <a class="nav-link" href="/auth/logout">
                    logout
                </a>
            </li>
        </ul>

        <slot />
{/await}


