<script lang="ts">
    import SignUp from "$lib/components/user/SignUp.svelte";
    import type { ActionData } from "./$types";

    export let form: ActionData;
</script>

<SignUp bind:form={form}></SignUp>