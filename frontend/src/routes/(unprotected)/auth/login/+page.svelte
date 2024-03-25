<script lang="ts">
    import Login from "$lib/components/user/Login.svelte";

    import type { ActionData } from './$types';

	export let form: ActionData;
</script>

<Login bind:form={form}></Login>