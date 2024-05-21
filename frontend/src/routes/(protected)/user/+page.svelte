<script lang="ts">
    import UserSettings from '$lib/components/user/Settings.svelte'
    import { GetStravaToken } from '$lib/wailsjs/go/controllers/UserHandler';
    import { LoadUser } from '$lib/wailsjs/go/main/App';
    import type { model, oauth2 } from '$lib/wailsjs/go/models';
    import { onMount } from 'svelte';

    let usr: model.User
    let strava_token: oauth2.Token;
    onMount(async () => {
        usr = await LoadUser()
        strava_token = await GetStravaToken(usr)
    })
</script>

<UserSettings bind:usr={usr} bind:existing_strava_token={strava_token}/>