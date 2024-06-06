<script lang="ts">
    import { goto } from "$app/navigation";
    import { HasUser } from "$lib/wailsjs/go/main/App";
    import { onMount } from "svelte";

    onMount(async () => {
        if (await HasUser()) {
            goto("/microcycle")
        } else {
            goto("/auth/login")
        }
    })
</script>