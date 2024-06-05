<script lang="ts">
    import { alerts } from "../../../stores/alert";
    import Alert from "./Alert.svelte";
</script>

<div class="alert-container container">
    {#each $alerts as alert}
        <Alert bind:alert={alert}/>
    {/each}
</div>