<script lang="ts">
    import { onMount } from "svelte";

	let url: string = ""
	onMount(() => {
		url = window.location.href
	})
</script>


{url}