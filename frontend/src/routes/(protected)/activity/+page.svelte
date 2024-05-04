<!--
    Because everything is statically generated, we can't use dynamic routing :(
    So to reach the activity page, you must include an activity uuid in the query args:
    ie. goto("activity?uuid=123456")
-->
<script lang="ts">
    import ActivityViewer from "$lib/components/activity/ActivityViewer.svelte";
    import { GetActivity } from "$lib/wailsjs/go/controllers/ActivityHandler";
    import type { model } from "$lib/wailsjs/go/models";
    import { onMount } from "svelte";

    let activity: model.Activity;
    let uuid: string | null
    onMount(async () => {
        const urlParams = new URLSearchParams(window.location.search);
        uuid = urlParams.get('uuid');
        if (uuid) {
            activity = await GetActivity(uuid)
        }
    })
</script>
{#if activity}
    <ActivityViewer bind:activity={activity}/>
{:else}
    <p>loading...</p>
{/if}