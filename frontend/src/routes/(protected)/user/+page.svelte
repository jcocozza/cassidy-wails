<script lang="ts">
    import UserSettings from '$lib/components/user/Settings.svelte'
    import { LoadUser } from '$lib/wailsjs/go/main/App';
    import type { model } from '$lib/wailsjs/go/models';
    import { onMount } from 'svelte';

    let usr: model.User
    onMount(async () => {
        usr = await LoadUser()
    })
</script>

<UserSettings bind:usr={usr}/>