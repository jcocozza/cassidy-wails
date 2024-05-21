<script lang="ts">
    // Bootstrap settings
    import 'bootstrap/dist/css/bootstrap.min.css';
    import 'bootstrap/dist/js/bootstrap.min.js';
    import '@popperjs/core/dist/umd/popper.min.js'
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';

    onMount(() => {
        goto("/auth/login")
    })
</script>

<slot />