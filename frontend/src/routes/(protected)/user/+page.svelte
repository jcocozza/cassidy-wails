<script lang="ts">
    import UserSettings from '$lib/components/user/Settings.svelte'
    import type { model } from '../../../../wailsjs/go/models';
    export let usr: model.User;
</script>

<UserSettings bind:usr={usr}/>